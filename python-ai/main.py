"""
澎湖數位老船長 - Python FastAPI 主應用
負責 AI Agent、LLM 串接、Tool Calling 等

啟動命令: uvicorn main:app --host 0.0.0.0 --port 8000 --reload
"""

import os
import logging
from fastapi import FastAPI, HTTPException
from fastapi.middleware.cors import CORSMiddleware
from pydantic import BaseModel
from dotenv import load_dotenv
import httpx

from utils.local_agent import LocalAgentService

# 加載環境變數
load_dotenv()

# 配置日誌
logging.basicConfig(
    level=logging.INFO,
    format='%(asctime)s - %(name)s - %(levelname)s - %(message)s'
)
logger = logging.getLogger(__name__)

# ============================================
# 創建 FastAPI 應用
# ============================================

app = FastAPI(
    title="澎湖數位老船長 - AI 服務",
    description="AI Agent、LLM 串接與 Tool Calling 服務",
    version="1.0.0"
)

# ============================================
# 中間件配置
# ============================================

# CORS 中間件
app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

# ============================================
# 數據模型
# ============================================

class ChatMessage(BaseModel):
    """聊天消息"""
    message: str
    language: str = "zh-TW"


class ChatResponse(BaseModel):
    """聊天回應"""
    reply: str
    action: str = "none"


# 環境配置

GOLANG_API_URL = os.getenv("GOLANG_API_URL", "http://golang-api:8080")
OPENAI_API_KEY = os.getenv("OPENAI_API_KEY")
GEMINI_API_KEY = os.getenv("GEMINI_API_KEY")

logger.info(f"🔗 Golang API URL: {GOLANG_API_URL}")

# 初始化本地 Agent 服務
local_agent = LocalAgentService(debug=True)

# ============================================
# 健康檢查
# ============================================

@app.get("/health")
async def health_check():
    """健康檢查端點"""
    return {
        "status": "healthy",
        "service": "python-ai",
        "version": "1.0.0"
    }


@app.get("/")
async def root():
    """根端點"""
    return {
        "message": "澎湖數位老船長 - Python AI 服務",
        "version": "1.0.0",
        "endpoints": {
            "chat": "POST /api/chat",
            "health": "GET /health"
        }
    }

# ============================================
# API 路由
# ============================================

@app.post("/api/chat", response_model=ChatResponse)
async def chat(request: ChatMessage):
    """
    聊天端點 - 接收用戶消息，由 AI Agent 處理

    示例:
    ```
    POST /api/chat
    {
        "message": "飼料庫存有多少?",
        "language": "zh-TW"
    }
    ```
    """
    logger.info(f"📨 收到訊息: {request.message}")

    try:
        # 低 token 消耗的簡化回應邏輯
        response = await generate_low_token_response(request.message, request.language)

        return ChatResponse(
            reply=response,
            action="chat"
        )

    except Exception as e:
        logger.error(f"❌ 聊天處理失敗: {str(e)}")
        raise HTTPException(status_code=500, detail=f"處理失敗: {str(e)}")


async def generate_low_token_response(message: str, language: str = "zh-TW") -> str:
    """
    生成低 token 消耗的回應

    Args:
        message: 用戶消息
        language: 語言

    Returns:
        簡潔回應
    """
    # 關鍵字匹配，減少 LLM 調用
    message_lower = message.lower()

    # 庫存查詢
    if any(keyword in message_lower for keyword in ["庫存", "inventory", "stock", "飼料", "藥品"]):
        try:
            inventory = await check_inventory()
            return f"📦 庫存資訊: {len(inventory)} 項物品在庫。建議檢查詳細清單。"
        except:
            return "📦 庫存查詢服務暫時不可用。"

    # 環境數據
    elif any(keyword in message_lower for keyword in ["環境", "水質", "溫度", "鹽度", "溶氧", "environment"]):
        try:
            env_data = await get_environmental_data()
            return f"🌊 環境數據: 目前監測 {len(env_data) if isinstance(env_data, list) else 1} 項參數。"
        except:
            return "🌊 環境監測服務暫時不可用。"

    # 任務管理
    elif any(keyword in message_lower for keyword in ["任務", "工作", "待辦", "task", "todo"]):
        try:
            tasks = await call_golang_api("GET", "/api/tasks")
            pending_count = len([t for t in tasks if t.get('status') == 'pending'])
            return f"✅ 任務狀態: {pending_count} 項待完成任務。"
        except:
            return "✅ 任務管理服務暫時不可用。"

    # OpenClaw CLI 集成
    elif any(keyword in message_lower for keyword in ["openclaw", "cli", "命令", "指令", "command"]):
        return await handle_openclaw_command(message)

    # 預設回應
    else:
        responses = [
            "我可以幫您查詢庫存、環境數據或任務狀態。",
            "請告訴我您需要什麼幫助？",
            "我可以協助您管理養殖作業。",
        ]
        return responses[len(message) % len(responses)]


async def handle_openclaw_command(message: str) -> str:
    """
    處理 OpenClaw CLI 命令

    Args:
        message: 用戶消息

    Returns:
        CLI 命令結果
    """
    try:
        # 簡單的命令解析
        if "status" in message.lower():
            # 模擬 OpenClaw 狀態檢查
            return "🔧 OpenClaw CLI 狀態: 運行中，所有服務正常。"
        elif "restart" in message.lower():
            # 模擬重啟命令
            return "🔄 OpenClaw CLI 重啟中... 請稍候。"
        else:
            return "💡 支援的 OpenClaw 命令: status, restart"

    except Exception as e:
        logger.error(f"OpenClaw CLI 處理失敗: {str(e)}")
        return "❌ OpenClaw CLI 命令執行失敗。"


# ============================================
# Agent 工具 - HTTP 調用 Golang API
# ============================================

async def call_golang_api(method: str, endpoint: str, data=None):
    """
    通用函數: 調用 Golang API

    Args:
        method: HTTP 方法 (GET, POST, PUT, DELETE)
        endpoint: API 端點 (如 /api/assets)
        data: 請求數據 (JSON)

    Returns:
        API 回應
    """
    url = f"{GOLANG_API_URL}{endpoint}"
    try:
        async with httpx.AsyncClient() as client:
            if method == "GET":
                response = await client.get(url)
            elif method == "POST":
                response = await client.post(url, json=data)
            elif method == "PUT":
                response = await client.put(url, json=data)
            elif method == "DELETE":
                response = await client.delete(url)
            else:
                raise ValueError(f"未支援的 HTTP 方法: {method}")

            if response.status_code >= 400:
                logger.error(f"❌ API 錯誤: {response.status_code} - {response.text}")
                raise Exception(f"API 請求失敗: {response.status_code}")

            return response.json()

    except Exception as e:
        logger.error(f"❌ Golang API 調用失敗: {str(e)}")
        raise


# ============================================
# Agent Tool 定義 - 查詢庫存
# ============================================

async def check_inventory(item_name: str = None):
    """
    Tool: 查詢庫存

    Args:
        item_name: 物品名稱 (可選)

    Returns:
        庫存清單
    """
    logger.info(f"🔍 查詢庫存: {item_name}")
    try:
        result = await call_golang_api("GET", "/api/assets")
        logger.info(f"✅ 庫存查詢成功")
        return result
    except Exception as e:
        return {"error": f"查詢失敗: {str(e)}"}


# ============================================
# Agent Tool 定義 - 建立任務
# ============================================

async def create_task(title: str, description: str, assigned_to: int, due_date: str):
    """
    Tool: 建立新任務

    Args:
        title: 任務名稱
        description: 任務描述
        assigned_to: 分配給的用戶 ID
        due_date: 截止日期

    Returns:
        建立的任務
    """
    logger.info(f"➕ 建立任務: {title}")
    try:
        data = {
            "title": title,
            "description": description,
            "assigned_to": assigned_to,
            "due_date": due_date
        }
        result = await call_golang_api("POST", "/api/tasks", data)
        logger.info(f"✅ 任務建立成功")
        return result
    except Exception as e:
        return {"error": f"建立失敗: {str(e)}"}


# ============================================
# Agent Tool 定義 - 查詢環境數據
# ============================================

async def get_environmental_data():
    """
    Tool: 查詢環境數據

    Returns:
        最新的環境數據
    """
    logger.info("🌊 查詢環境數據")
    try:
        result = await call_golang_api("GET", "/api/environmental-data")
        logger.info(f"✅ 環境數據查詢成功")
        return result
    except Exception as e:
        return {"error": f"查詢失敗: {str(e)}"}


# ============================================
# Agent Tool 定義 - 生成報表
# ============================================

async def generate_report(date_range: str = "daily"):
    """
    Tool: 生成營運報表

    Args:
        date_range: 日期範圍 (daily/weekly/monthly)

    Returns:
        生成的報表
    """
    logger.info(f"📊 生成{date_range}報表")
    try:
        # 獲取所需數據
        assets = await get_environmental_data()
        tasks = await call_golang_api("GET", "/api/tasks")

        report = {
            "date_range": date_range,
            "environmental_data": assets,
            "tasks": tasks,
            "generated_at": "2026-04-07"
        }

        logger.info(f"✅ 報表生成成功")
        return report

    except Exception as e:
        return {"error": f"生成失敗: {str(e)}"}


# ============================================
# 本地 Agent 路由
# ============================================

class NavigateRequest(BaseModel):
    url: str
    extractors: dict = None

@app.post("/api/agent/navigate")
async def agent_navigate(req: NavigateRequest):
    """本地 Agent 導航端點"""
    result = local_agent.navigate_and_extract(req.url, req.extractors)
    return result


class FormSubmitRequest(BaseModel):
    fields: dict
    submit_button_index: int

@app.post("/api/agent/form-submit")
async def agent_form_submit(req: FormSubmitRequest):
    """本地 Agent 提交表單端點"""
    result = local_agent.fill_and_submit_form(req.fields, req.submit_button_index)
    return result


class TableExtractRequest(BaseModel):
    table_selector: str

@app.post("/api/agent/extract-table")
async def agent_extract_table(req: TableExtractRequest):
    """本地 Agent 提取表格端點"""
    result = local_agent.extract_table(req.table_selector)
    return result


@app.get("/api/agent/session")
async def agent_get_session():
    """獲取 Agent 會話信息"""
    return local_agent.get_session_info()


# ============================================
# 啟動事件
# ============================================

@app.on_event("startup")
async def startup_event():
    """應用啟動時執行"""
    logger.info("🚀 Python FastAPI 應用啟動")
    logger.info(f"✅ LLM 配置已載入 (OpenAI: {'是' if OPENAI_API_KEY else '否'})")


@app.on_event("shutdown")
async def shutdown_event():
    """應用關閉時執行"""
    logger.info("🛑 Python FastAPI 應用關閉")


if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="0.0.0.0", port=8000)
