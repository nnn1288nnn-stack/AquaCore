<!-- 澎湖數位老船長 - OpenCli Agent 快速启动指南 -->

# 🚀 澎湖數位老船長 - 快速启动指南

**项目**: Penghu Digital Captain (澎湖數位老船長)  
**功能**: Agent-Based Aquaculture Management System  
**最后更新**: 2026-04-07  

---

## ⚡ 5 分钟快速启动

### 前置要求

```bash
# 检查 OpenCli 是否已安装
which opencli

# 检查 Go 版本
go version  # 需要 1.21+

# 检查 Python 版本
python --version  # 需要 3.10+
```

### 步骤 1: 启动 Go API 服务

```bash
cd /home/ouo/project/golang-api

# 编译并运行
go run main.go

# 或编译后运行
go build -o api . && ./api
```

✅ 输出应该显示:
```
Server running on :8080
Database connected
```

### 步骤 2: 在新终端启动 Python API 服务

```bash
cd /home/ouo/project/python-ai

# 安装依赖（如未安装）
pip install -r requirements.txt

# 运行服务
python -m uvicorn main:app --reload --port 8000
```

✅ 输出应该显示:
```
Uvicorn running on http://127.0.0.1:8000
```

### 步骤 3: 测试 Agent 功能

**选项 A: 使用 curl 测试**

```bash
# 导航到网站并提取信息
curl -X POST http://localhost:8080/api/agent/navigate \
  -H "Content-Type: application/json" \
  -d '{
    "url":"https://example.com",
    "extractors":{"title":"h1","description":"p"}
  }'
```

**选项 B: 运行自动化测试脚本**

```bash
# 进入项目目录
cd /home/ouo/project

# 运行测试脚本
bash test-api.sh
```

**选项 C: 运行示例程序**

Go 示例:
```bash
cd golang-api/examples
go run main.go  # 然后选择示例编号
```

Python 示例:
```bash
cd python-ai
python examples.py  # 然后选择示例编号
```

---

## 📊 系统架构概览

```
┌─────────────────────────────────────────────────────────────┐
│                     Client Applications                      │
│           (Web, Mobile, Desktop, CLI Tools)                  │
└──────────────────┬──────────────────────────────────────────┘
                   │
        ┌──────────┼──────────┐
        │          │          │
   ┌────▼──┐   ┌───▼───┐   ┌─▼─────┐
   │  Web  │   │ REST  │   │ WebSocket
   │Server │   │Client │   │Client
   └────┬──┘   └───┬───┘   └─┬─────┘
        │          │          │
   ┌────▼──────────▼──────────▼────┐
   │      API Gateway (Nginx)       │
   └────┬──────────────────────────┘
        │
   ┌────┴─────────────────────────┐
   │                               │
┌──▼─────────────┐      ┌──────────▼──┐
│   Go API Sv    │      │Python FastAPI
│  (port 8080)   │      │(port 8000)
│                │      │
│■ OpenCli SDK   │      │■ LocalAgent
│■ Agent Service │      │  Service
│■ HTTP Handler  │      │■ FastAPI Rts
└──┬─────────────┘      └──────────┬──┘
   │                               │
   └────────────┬──────────────────┘
                │
    ┌───────────▼───────────┐
    │   OpenCli CLI Tool    │
    │  (Browser Automation) │
    └───────────┬───────────┘
                │
    ┌───────────▼───────────┐
    │    Chrome Browser     │
    │  (OpenCli Target)     │
    └───────────┬───────────┘
                │
    ┌───────────▼───────────┐
    │  Target Websites      │
    │  (Form Processing)    │
    └───────────────────────┘
```

---

## 🎯 常用命令速查表

### Go API 相关

```bash
# 构建 Go API
cd golang-api
go build -o api .

# 运行 Go API 与调试信息
DEBUG=1 ./api

# 运行 Go 单元测试
go test ./...

# 查看 Go API 的可用端点
# 打开 http://localhost:8080/swagger 或查看 handlers/agent.go

# 使用 Go 编写代理脚本
cd examples
go run main.go
```

### Python API 相关

```bash
# 安装依赖
pip install -r requirements.txt

# 运行 Python API
python main.py

# 运行 Python API 带热重载
python -m uvicorn main:app --reload

# 查看 Python API 文档
# 打开 http://localhost:8000/docs

# 运行 Python 示例
python examples.py

# 与 Python Agent 交互（Python REPL）
python -c "from utils.local_agent import LocalAgentService; a=LocalAgentService()"
```

### 测试和验证

```bash
# 运行所有 API 测试
bash test-api.sh

# 测试单个 Go 端点
curl -X POST http://localhost:8080/api/agent/navigate \
  -H "Content-Type: application/json" \
  -d '{"url":"https://example.com"}'

# 测试单个 Python 端点
curl -X POST http://localhost:8000/api/agent/navigate \
  -H "Content-Type: application/json" \
  -d '{"url":"https://example.com"}'

# 检查 OpenCli 状态
opencli doctor
```

### Docker 相关

```bash
# 完整启动所有服务
docker-compose up -d

# 查看运行的容器
docker-compose ps

# 查看日志
docker-compose logs -f

# 停止所有服务
docker-compose down

# 重建镜像
docker-compose build --no-cache
```

---

## 🔍 测试场景示例

### 场景 1: 简单网页爬虫

```bash
# 导航到网页并提取信息
curl -X POST http://localhost:8080/api/agent/navigate \
  -H "Content-Type: application/json" \
  -d '{
    "url": "https://example.com",
    "extractors": {
      "title": "h1",
      "description": "meta[name=description]",
      "links": "a"
    }
  }' | jq .
```

### 场景 2: 表单自动填充和提交

```bash
# 第 1 步: 导航到表单页面
curl -X POST http://localhost:8080/api/agent/navigate \
  -H "Content-Type: application/json" \
  -d '{"url": "https://example.com/contact-form"}' | jq .

# 第 2 步: 填充表单字段
curl -X POST http://localhost:8080/api/agent/fill-form \
  -H "Content-Type: application/json" \
  -d '{
    "url": "https://example.com/contact-form",
    "fields": {
      "1": "John Doe",
      "2": "john@example.com",
      "3": "Hello, this is a test message"
    }
  }' | jq .

# 第 3 步: 提交表单
curl -X POST http://localhost:8080/api/agent/submit-form \
  -H "Content-Type: application/json" \
  -d '{
    "url": "https://example.com/contact-form",
    "submit_button_index": 4,
    "wait_time": 2
  }' | jq .
```

### 场景 3: 数据表格提取

```bash
# 从网页提取表格数据
curl -X POST http://localhost:8000/api/agent/extract-table \
  -H "Content-Type: application/json" \
  -d '{
    "url": "https://example.com/data-table",
    "selector": "table.data-table"
  }' | jq .
```

### 场景 4: 多步工作流 (Python)

```python
from utils.local_agent import LocalAgentService

# 初始化 Agent
agent = LocalAgentService(debug=True)

# 步骤 1: 导航和提取
result = agent.navigate_and_extract(
    url="https://example.com/products",
    extractors={
        "product_names": ".product-name",
        "prices": ".product-price"
    }
)
print(f"提取了 {len(result['extracted'])} 个产品")

# 步骤 2: 点击并导航
agent.opencli.click(5)
agent.opencli.wait("time", "1")

# 步骤 3: 获取新页面状态
new_state = agent.opencli.get_state()
print(f"当前 URL: {new_state.url}")
```

---

## 🛠️ 故障排除

### 问题 1: "command not found: opencli"

**解决方案:**
```bash
# 检查 OpenCli 是否在 PATH 中
echo $PATH

# 或直接运行
/home/ouo/skills/opencli-operate/bin/opencli --version

# 添加到 PATH (在 ~/.bashrc 或 ~/.zshrc 中)
export PATH="/home/ouo/skills/opencli-operate/bin:$PATH"
```

### 问题 2: "Port 8080 already in use"

**解决方案:**
```bash
# 查找占用端口的进程
sudo lsof -i :8080

# 杀死进程（如果是旧的 API 实例）
kill -9 <PID>

# 或改变端口（在 main.go 中）
PORT=8081 go run main.go
```

### 问题 3: "Python module not found"

**解决方案:**
```bash
# 从项目根目录运行
cd /home/ouo/project

# 设置 PYTHONPATH
export PYTHONPATH="${PYTHONPATH}:/home/ouo/project/python-ai"

# 验证导入
python -c "from utils.local_agent import LocalAgentService"
```

### 问题 4: OpenCli 命令超时

**解决方案:**
```bash
# 检查 OpenCli 状态
opencli doctor

# 使用更长的超时时间
# 在代码中修改超时参数
# Go: executeCommand() 中的 time.Duration
# Python: execute_command() 中的 timeout 参数
```

### 问题 5: "connection refused" 错误

**解决方案:**
```bash
# 确保 API 服务已启动
ps aux | grep -E "api|uvicorn"

# 测试连接
curl http://localhost:8080/health || echo "Go API 未响应"
curl http://localhost:8000/docs || echo "Python API 未响应"

# 查看防火墙设置
sudo firewall-cmd --list-ports  # 如果使用 firewalld
```

---

## 📚 更多资源

| 资源 | 位置 | 描述 |
|------|------|------|
| 集成指南 | [docs/INTEGRATION.md](docs/INTEGRATION.md) | 详细的集成文档 |
| 架构文档 | [docs/ARCHITECTURE.md](docs/ARCHITECTURE.md) | 系统架构说明 |
| 测试清单 | [TESTING.md](TESTING.md) | 完整的测试计划 |
| Go 示例 | [golang-api/examples/main.go](golang-api/examples/main.go) | 8 个 Go 使用示例 |
| Python 示例 | [python-ai/examples.py](python-ai/examples.py) | 8 个 Python 使用示例 |
| API 测试脚本 | [test-api.sh](test-api.sh) | 自动化 curl 测试 |
| 项目结构 | [README.md](README.md) | 项目总体信息 |

---

## ✨ 下一步

1. ✅ **现在完成**: 系统已部署和配置
2. 🚀 **接下来**: [运行第一个测试](TESTING.md)
3. 📖 **然后**: 阅读[集成指南](docs/INTEGRATION.md)了解高级用法
4. 🔧 **最后**: 根据需要自定义集成

---

## 💡 关键概念

### OpenCli 工作流程
```
启动 Chrome → 打开 URL → 获取页面状态 → 执行操作 → 等待 → 获取结果
```

### Agent 工作流程
```
创建会话 → 初始化OpenCli → 执行命令 → 缓存结果 → 返回数据
```

### API 架构
```
REST 请求 → HTTP 端点 → Service 层 → OpenCli 操作 → 浏览器 → 返回数据
```

---

## 📞 获取帮助

- 🐛 **报告 Bug**: 检查 Agent 的输出日志
- ❓ **技术问题**: 查看 [INTEGRATION.md](docs/INTEGRATION.md) 中的故障排除部分
- 📧 **反馈**: 查看项目的 CONTRIBUTING.md

---

**祝您使用愉快！ 🎉**

