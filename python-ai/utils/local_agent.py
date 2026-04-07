"""
澎湖數位老船長 - 本地 Agent 集成模块
與 OpenClaw Agent 和 OpenCli 集成
"""

import os
import json
import logging
import subprocess
from typing import Dict, List, Optional, Any
from dataclasses import dataclass, asdict


# 日志配置
logging.basicConfig(
    level=logging.INFO,
    format='[%(asctime)s][AGENT] %(levelname)s: %(message)s'
)
logger = logging.getLogger(__name__)


@dataclass
class PageElement:
    """页面元素数据类"""
    index: int
    tag: str
    text: str
    value: str = ""
    selector: str = ""
    type: str = ""


@dataclass
class PageState:
    """页面状态数据类"""
    url: str
    title: str
    elements: Dict[int, Dict[str, Any]]
    raw_html: str = ""


class OpenCliAgent:
    """OpenCli 代理 - 与 OpenCli CLI 交互"""
    
    def __init__(self, debug: bool = False):
        self.debug = debug
        self._verify_opencli()
    
    def _verify_opencli(self) -> bool:
        """验证 OpenCli 是否可用"""
        try:
            result = subprocess.run(
                ["opencli", "doctor"],
                capture_output=True,
                text=True,
                timeout=5
            )
            if result.returncode == 0:
                logger.info("✅ OpenCli 连接正常")
                return True
            else:
                logger.warning("⚠️  OpenCli 连接失败")
                return False
        except FileNotFoundError:
            logger.error("❌ opencli 命令未找到，请确保已安装")
            return False
        except Exception as e:
            logger.error(f"❌ OpenCli 验证失败: {e}")
            return False
    
    def execute_command(self, *args: str) -> str:
        """执行 opencli 命令"""
        try:
            cmd = ["opencli"] + list(args)
            if self.debug:
                logger.debug(f"执行命令: {' '.join(cmd)}")
            
            result = subprocess.run(
                cmd,
                capture_output=True,
                text=True,
                timeout=30
            )
            
            if result.returncode != 0:
                raise RuntimeError(f"命令失败: {result.stderr}")
            
            return result.stdout.strip()
        except Exception as e:
            logger.error(f"命令执行失败: {e}")
            raise
    
    def open(self, url: str) -> bool:
        """打开 URL"""
        try:
            self.execute_command("operate", "open", url)
            logger.info(f"📖 已打开 URL: {url}")
            return True
        except Exception as e:
            logger.error(f"打开 URL 失败: {e}")
            return False
    
    def get_state(self) -> Optional[PageState]:
        """获取页面状态"""
        try:
            output = self.execute_command("operate", "state")
            
            # 尝试解析 JSON
            try:
                data = json.loads(output)
                return PageState(
                    url=data.get("url", ""),
                    title=data.get("title", ""),
                    elements=data.get("elements", {}),
                    raw_html=data.get("rawHtml", "")
                )
            except json.JSONDecodeError:
                # 如果解析失败，返回原始 HTML
                return PageState(
                    url="",
                    title="",
                    elements={},
                    raw_html=output
                )
        except Exception as e:
            logger.error(f"获取页面状态失败: {e}")
            return None
    
    def click(self, element_index: int) -> bool:
        """点击元素"""
        try:
            self.execute_command("operate", "click", str(element_index))
            logger.info(f"👆 已点击元素 [{element_index}]")
            return True
        except Exception as e:
            logger.error(f"点击元素失败: {e}")
            return False
    
    def type_text(self, element_index: int, text: str) -> bool:
        """输入文本"""
        try:
            self.execute_command("operate", "type", str(element_index), text)
            logger.info(f"⌨️  已输入文本到元素 [{element_index}]: {text[:50]}...")
            return True
        except Exception as e:
            logger.error(f"输入文本失败: {e}")
            return False
    
    def select_option(self, element_index: int, value: str) -> bool:
        """选择下拉选项"""
        try:
            self.execute_command("operate", "select", str(element_index), value)
            logger.info(f"🔽 已选择选项 [{element_index}] = {value}")
            return True
        except Exception as e:
            logger.error(f"选择选项失败: {e}")
            return False
    
    def get_value(self, element_index: int) -> str:
        """获取元素值"""
        try:
            return self.execute_command("operate", "get", "value", str(element_index))
        except Exception as e:
            logger.error(f"获取元素值失败: {e}")
            return ""
    
    def get_text(self, element_index: int) -> str:
        """获取元素文本"""
        try:
            return self.execute_command("operate", "get", "text", str(element_index))
        except Exception as e:
            logger.error(f"获取元素文本失败: {e}")
            return ""
    
    def scroll(self, direction: str = "down", amount: int = 0) -> bool:
        """滚动页面"""
        try:
            args = ["operate", "scroll", direction]
            if amount > 0:
                args.extend(["--amount", str(amount)])
            self.execute_command(*args)
            logger.info(f"📜 已滚动页面: {direction} ({amount})")
            return True
        except Exception as e:
            logger.error(f"滚动页面失败: {e}")
            return False
    
    def wait(self, wait_type: str, value: str) -> bool:
        """等待条件满足"""
        try:
            self.execute_command("operate", "wait", wait_type, value)
            logger.info(f"⏳ 已等待条件: {wait_type}={value}")
            return True
        except Exception as e:
            logger.error(f"等待条件失败: {e}")
            return False
    
    def eval_js(self, js_code: str) -> str:
        """执行 JavaScript 代码 (读取专用)"""
        try:
            return self.execute_command("operate", "eval", js_code)
        except Exception as e:
            logger.error(f"执行 JS 失败: {e}")
            return ""
    
    def screenshot(self, filepath: str) -> bool:
        """保存屏幕截图"""
        try:
            self.execute_command("operate", "screenshot", filepath)
            logger.info(f"📸 已保存截图: {filepath}")
            return True
        except Exception as e:
            logger.error(f"保存截图失败: {e}")
            return False


class LocalAgentService:
    """本地 Agent 服务 - 高级操作接口"""
    
    def __init__(self, debug: bool = False):
        self.opencli = OpenCliAgent(debug=debug)
        self.cache = {}
        self.session_id = self._generate_session_id()
    
    @staticmethod
    def _generate_session_id() -> str:
        """生成会话 ID"""
        import uuid
        from datetime import datetime
        return f"{uuid.uuid4().hex[:8]}-{int(datetime.now().timestamp())}"
    
    def navigate_and_extract(
        self,
        url: str,
        extractors: Optional[Dict[str, str]] = None
    ) -> Dict[str, Any]:
        """导航并提取数据"""
        logger.info(f"🔍 导航到: {url}")
        
        if not self.opencli.open(url):
            return {"success": False, "error": "打开 URL 失败"}
        
        # 等待页面加载
        self.opencli.wait("time", "1")
        
        # 获取页面状态
        state = self.opencli.get_state()
        if not state:
            return {"success": False, "error": "获取页面状态失败"}
        
        result = {
            "success": True,
            "url": state.url,
            "title": state.title,
            "elements": state.elements,
        }
        
        # 执行提取器
        if extractors:
            extracted = {}
            for key, selector in extractors.items():
                js_code = f"""
                (function() {{
                    const el = document.querySelector('{selector}');
                    return el ? el.textContent : '';
                }})()
                """
                value = self.opencli.eval_js(js_code)
                extracted[key] = value
            result["extracted"] = extracted
        
        return result
    
    def fill_and_submit_form(self, fields: Dict[int, str], submit_index: int) -> Dict[str, Any]:
        """填充并提交表单"""
        logger.info(f"📝 填充表单 ({len(fields)} 个字段)")
        
        try:
            # 填充字段
            for index, value in fields.items():
                if not self.opencli.type_text(index, value):
                    return {"success": False, "error": f"填充字段 [{index}] 失败"}
            
            # 提交表单
            if not self.opencli.click(submit_index):
                return {"success": False, "error": "点击提交按钮失败"}
            
            # 等待响应
            self.opencli.wait("time", "2")
            
            # 获取结果页面状态
            state = self.opencli.get_state()
            
            return {
                "success": True,
                "message": "表单已成功提交",
                "page_state": asdict(state) if state else None
            }
        except Exception as e:
            logger.error(f"表单提交失败: {e}")
            return {"success": False, "error": str(e)}
    
    def extract_table(self, table_selector: str) -> Dict[str, Any]:
        """提取表格数据"""
        logger.info(f"📊 提取表格: {table_selector}")
        
        js_code = f"""
        (function() {{
            const rows = document.querySelectorAll('{table_selector} tbody tr');
            const data = [];
            rows.forEach(row => {{
                const cells = row.querySelectorAll('td');
                const rowData = {{}};
                cells.forEach((cell, idx) => {{
                    rowData['col_' + idx] = cell.textContent.trim();
                }});
                data.push(rowData);
            }});
            return JSON.stringify(data);
        }})()
        """
        
        try:
            result = self.opencli.eval_js(js_code)
            data = json.loads(result)
            return {
                "success": True,
                "data": data,
                "row_count": len(data)
            }
        except Exception as e:
            logger.error(f"提取表格失败: {e}")
            return {"success": False, "error": str(e)}
    
    def get_session_info(self) -> Dict[str, str]:
        """获取会话信息"""
        return {
            "session_id": self.session_id,
            "status": "initialized"
        }


# 使用示例
if __name__ == "__main__":
    # 创建 Agent 实例
    agent = LocalAgentService(debug=True)
    
    # 验证连接
    logger.info("开始测试 OpenCli Agent...")
    
    # 导航示例
    # result = agent.navigate_and_extract("https://example.com")
    # print(result)
