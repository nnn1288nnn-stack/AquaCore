# OpenClaw + OpenCli + Go API 集成指南

**澎湖數位老船長** - 微服務 + AI Agent 集成

---

## 📋 概览

本项目整合了：
- **Go API**: 高性能的业务 API  
- **Python FastAPI**: 轻量级 LLM 和本地 Agent
- **OpenCli Agent**: 浏览器自动化（部署在本地）
- **OpenClaw Framework**: Agent 框架和扩展

### 架构流程

```
┌─────────────────────────────────────────────────┐
│              調用方 (Web/CLI/Agent)             │
└────────────────┬─────────────────────────────────┘
                 │ HTTP
        ┌────────▼──────────┐
        │  Nginx (Port 80)  │ ← 反向代理
        └────────┬──────────┘
                 │
      ┌──────────┴──────────┐
      ▼                     ▼
┌──────────────┐     ┌──────────────────┐
│  Go API      │     │ Python FastAPI   │
│  (8080)      │     │ (8000)           │
│  - CRUD      │     │ - 本地 Agent     │
│  - 業務邏輯   │◄────┤ - OpenCli CLI    │
└──────┬───────┘     │ - LLM 集成       │
       │             └──────┬───────────┘
       │                    │
       └────────────────────┤ 浏览器自动化
                            ▼
                    ┌─────────────────┐
                    │ OpenCli 瀏覽器  │
                    │ 自動化 (本地)   │
                    └─────────────────┘
```

---

## 🔧 Go API 集成

### 1. SDK 使用

**位置**: `golang-api/utils/opencli.go`

#### 初始化 OpenCli 客户端

```go
import "github.com/penghu-digital-captain/golang-api/utils"

func main() {
    client := utils.NewOpenCliClient(debug := true)
    
    // 检查连接
    status, err := client.Doctor()
    if err != nil {
        log.Fatal("OpenCli 连接失败:", err)
    }
    log.Println("状态:", status)
}
```

#### 基本操作

```go
// 打开 URL
client.Open("https://example.com")

// 获取页面状态
state, _ := client.GetState()
fmt.Printf("Title: %s\n", state.Title)
fmt.Printf("Elements: %+v\n", state.Elements)

// 点击元素
client.Click(5)

// 输入文本
client.Type(3, "hello world")

// 获取值
value, _ := client.GetValue(3)

// 滚动
client.Scroll("down", 1000)

// 等待
client.Wait("selector", ".loading")

// JavaScript 执行 (只读)
result, _ := client.Eval("(function(){ return document.title; })()")
```

### 2. Agent 服务

**位置**: `golang-api/utils/agent.go`

高级操作接口：

```go
agent := utils.NewAgentService(debug := true)
agent.Initialize()

// 导航并提取数据
data, err := agent.NavigateAndExtract(
    "https://example.com",
    map[string]string{
        "price": ".product-price",
        "title": ".product-title",
    },
)

// 点击并导航
result, err := agent.ClickAndNavigate(7)

// 填充表单
agent.FillForm(map[int]string{
    3: "john@example.com",
    4: "password123",
})

// 提交表单
agent.SubmitForm(5)

// 提取表格数据
tableData, err := agent.ExtractTableData("table.data-table")
```

---

## 🐍 Python Agent 集成

### 1. 本地 Agent 服务

**位置**: `python-ai/utils/local_agent.py`

#### 初始化

```python
from utils.local_agent import LocalAgentService

# 创建 Agent 实例
agent = LocalAgentService(debug=True)

# 验证 OpenCli 连接
# （初始化时自动验证）
```

#### 导航和提取

```python
# 导航到 URL 并提取数据
result = agent.navigate_and_extract(
    url="https://example.com",
    extractors={
        "title": "h1",
        "price": ".price"
    }
)

print(result["extracted"]["title"])
print(result["extracted"]["price"])
```

#### 表单操作

```python
# 填充并提交表单
result = agent.fill_and_submit_form(
    fields={
        3: "user@example.com",    # 字段索引 -> 值
        4: "password123"
    },
    submit_index=5  # 提交按钮的索引
)

if result["success"]:
    print("表单提交成功")
    print(result["message"])
```

#### 表格提取

```python
# 提取表格数据
result = agent.extract_table("table.data-table")

if result["success"]:
    print(f"提取了 {result['row_count']} 行数据")
    for row in result["data"]:
        print(row)
```

### 2. FastAPI 路由

现有的 Python API 新增以下路由：

```
POST   /api/agent/navigate          - 导航并提取
POST   /api/agent/form-submit       - 提交表单
POST   /api/agent/extract-table     - 提取表格
GET    /api/agent/session           - 获取会话信息
```

#### 使用示例

```bash
# 导航
curl -X POST http://localhost:8000/api/agent/navigate \
  -H "Content-Type: application/json" \
  -d '{
    "url": "https://example.com",
    "extractors": {
      "title": "h1",
      "price": ".price"
    }
  }'

# 提交表单
curl -X POST http://localhost:8000/api/agent/form-submit \
  -H "Content-Type: application/json" \
  -d '{
    "fields": {"3": "user@example.com", "4": "password"},
    "submit_button_index": 5
  }'
```

---

## 🚀 Go API 路由

新增 Agent 处理器端点 (`golang-api/handlers/agent.go`):

```
POST   /api/agent/navigate          - 导航页面
POST   /api/agent/click             - 点击元素
POST   /api/agent/fill-form         - 填充表单
POST   /api/agent/submit-form       - 提交表单
GET    /api/agent/page-state        - 获取页面状态
GET    /api/agent/session           - 获取会话信息
```

#### 使用示例

```bash
# 导航
curl -X POST http://localhost:8080/api/agent/navigate \
  -H "Content-Type: application/json" \
  -d '{
    "url": "https://example.com",
    "extractors": {"title": "h1"}
  }'

# 点击
curl -X POST http://localhost:8080/api/agent/click \
  -H "Content-Type: application/json" \
  -d '{"element_index": 5}'

# 获取页面状态
curl http://localhost:8080/api/agent/page-state
```

---

## 🔗 集成工作流程

### 场景 1: 数据采集

```
浏览器 (OpenCli) 代理
    ↓
打开网页 → 获取页面状态 → 定位元素
    ↓
Python Agent 使用 JavaScript 提取
    ↓
存储到 Go API (MariaDB)
```

**代码示例**:

```python
# Python 端
agent = LocalAgentService()

# 1. 导航
nav_result = agent.navigate_and_extract(
    "https://shop.example.com/products",
    extractors={
        "product_name": ".product-name",
        "price": ".product-price"
    }
)

# 2. 提取表格
table_result = agent.extract_table("table.products")

# 3. 发送到 Go API 存储
import requests
requests.post(
    "http://localhost:8080/api/assets",
    json={
        "name": nav_result["extracted"]["product_name"],
        "quantity": table_result["data"][0]["quantity"],
        "unit": "件"
    }
)
```

### 场景 2: 自动化表单

```
接收表单参数
    ↓
提取表单页面 (OpenCli)
    ↓
定位form字段 → 填充数据 → 提交
    ↓
处理响应 → 记录结果到数据库
```

**代码示例**:

```go
// Go 端
handler := handlers.NewAgentHandler(agentService)

// 支持的 API 端点
router.POST("/api/agent/fill-form", handler.FillForm)
router.POST("/api/agent/submit-form", handler.SubmitForm)

// 使用
client.Post("/api/agent/fill-form", map[string]interface{}{
    "fields": map[string]string{
        "3": "test@example.com",
        "4": "password123"
    }
})
```

### 场景 3: AI Agent 决策

```
接收自然语言指令 (LLM)
    ↓
解析意图 → 生成 Agent 工作流
    ↓
调用 OpenCli 执行浏览器自动化
    ↓
收集结果 → LLM 生成报告
    ↓
返回用户友好的响应
```

**代码示例** (Python):

```python
from langchain import OpenAI, Agent, Tool

# 定义工具
navigate_tool = Tool(
    name="navigate",
    func=lambda url: agent.navigate_and_extract(url),
    description="导航到 URL 并提取页面数据"
)

extract_tool = Tool(
    name="extract_table",
    func=agent.extract_table,
    description="从页面中提取表格数据"
)

# 创建 Agent
llm = OpenAI(model="gpt-4")
tools = [navigate_tool, extract_tool]

# Agent 会自动决定何时使用这些工具
```

---

##⚙️ 部署和配置

### 本地开发

```bash
# 1. 验证 OpenCli 安装
opencli doctor

# 2. 启动 Go API
cd golang-api
go run main.go

# 3. 启动 Python Agent
cd python-ai
uvicorn main:app --reload

# 4. 测试集成
curl http://localhost:8000/api/agent/session
```

### Docker 部署

两个服务都已包含 Dockerfile。ensure OpenCli 在主机上可用（不容器化）：

```bash
# Ｈost 上运行 OpenCli daemon
opencli daemon start

# 容器网络必须能访问主机 OpenCli
docker-compose up -d
```

---

## 📊 文件结构

```
golang-api/
├── utils/
│   ├── opencli.go         # OpenCli 客户端 SDK
│   └── agent.go           # 高级 Agent 服务
└── handlers/
    └── agent.go           # API 处理器

python-ai/
├── utils/
│   └── local_agent.py     # 本地 Agent 服务
├── main.py                # FastAPI + 新 Agent 路由
└── requirements.txt       # 更新的依赖
```

---

##🔐 安全建议

1. **OpenCli 连接**：
   - 仅在受信任的网络使用
   - 设置防火墙规则限制访问

2. **API 认证**：
   - 添加 JWT 或 API Key 认证
   - 实现请求速率限制

3. **敏感信息**：
   - 表单密码通过环境变量传入
   - 不要硬编码凭证

---

## 🚨 故障排除

### OpenCli 连接失败

```bash
# 1. 验证安装
which opencli
opencli doctor

# 2. 检查 daemon
opencli daemon status

# 3. 重启 daemon
opencli daemon stop
opencli daemon start
```

### Go API 编译错误

```bash
# 更新依赖
go mod tidy
go get -u ./...

# 编译
go build
```

### Python 导入错误（

```bash
# 安装依赖
pip install -r requirements.txt

# 检查路径
export PYTHONPATH=$PYTHONPATH:$(pwd)
```

---

## 📚 参考资源

- [OpenCli 文档](https://github.com/openclaw/opencli)
- [Golang Gin 框架](https://gin-gonic.com/)
- [FastAPI 文档](https://fastapi.tiangolo.com/)
- [LangChain 文档](https://python.langchain.com/)

---

**澎湖數位老船長** 🐠 | OpenClaw + Go + Python 集成 | 2026-04-07
