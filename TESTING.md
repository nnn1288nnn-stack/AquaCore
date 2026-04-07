<!-- 澎湖數位老船長 - OpenCli Agent 整合测试清单 -->

# 🧪 澎湖數位老船長 - OpenCli Agent 整合测试清单

**项目**: Penghu Digital Captain (澎湖數位老船長)  
**模块**: Local OpenCli Agent Integration  
**版本**: 1.0.0  
**最后更新**: 2026-04-07  

---

## 📋 测试清单

### 第一阶段: 环境和依赖验证

#### 1.1 系统依赖检查
- [ ] OpenCli 命令行工具已安装
  ```bash
  which opencli
  opencli --version
  ```
- [ ] Go 1.21+ 已安装
  ```bash
  go version
  ```
- [ ] Python 3.10+ 已安装
  ```bash
  python --version
  ```
- [ ] 依赖包已安装 (Go)
  ```bash
  go list ./...
  ```
- [ ] 依赖包已安装 (Python)
  ```bash
  pip list | grep -E "fastapi|pydantic"
  ```

#### 1.2 OpenCli 功能验证
- [ ] OpenCli doctor 命令成功
  ```bash
  opencli doctor
  ```
- [ ] OpenCli 可以打开网页
  ```bash
  opencli open "https://example.com"
  ```
- [ ] OpenCli 可以获取页面状态
  ```bash
  opencli get-state
  ```

---

### 第二阶段: 代码编译和单元测试

#### 2.1 Go 代码编译
- [ ] `utils/opencli.go` 编译成功（无错误）
  ```bash
  cd golang-api
  go build ./utils
  ```
- [ ] `utils/agent.go` 编译成功
  ```bash
  go build ./utils
  ```
- [ ] `handlers/agent.go` 编译成功
  ```bash
  go build ./handlers
  ```
- [ ] 完整项目编译成功
  ```bash
  go build -o api .
  ```

#### 2.2 Go 单元测试
- [ ] `opencli_test.go` 通过
  ```bash
  go test ./utils -run TestOpenCli
  ```
- [ ] `agent_test.go` 通过
  ```bash
  go test ./utils -run TestAgent
  ```
- [ ] 所有测试通过
  ```bash
  go test ./...
  ```

#### 2.3 Python 代码验证
- [ ] 导入 `local_agent.py` 成功
  ```bash
  cd python-ai
  python -c "from utils.local_agent import LocalAgentService; print('OK')"
  ```
- [ ] `LocalAgentService` 类可以实例化
  ```python
  from utils.local_agent import LocalAgentService
  agent = LocalAgentService(debug=True)
  ```
- [ ] FastAPI 应用启动成功
  ```bash
  python -m uvicorn main:app --reload
  ```

#### 2.4 Python 单元测试
- [ ] `test_local_agent.py` 通过
  ```bash
  pytest test_local_agent.py -v
  ```

---

### 第三阶段: API 集成测试

#### 3.1 Go API 端点测试

**启动 Go API 服务:**
```bash
cd golang-api
go run main.go
# 或
./api
```

- [ ] `POST /api/agent/navigate` 端点可访问
  ```bash
  curl -X POST http://localhost:8080/api/agent/navigate \
    -H "Content-Type: application/json" \
    -d '{"url":"https://example.com","extractors":{"title":"h1"}}'
  ```

- [ ] `POST /api/agent/click` 端点可访问
  ```bash
  curl -X POST http://localhost:8080/api/agent/click \
    -H "Content-Type: application/json" \
    -d '{"url":"https://example.com","element_index":1}'
  ```

- [ ] `POST /api/agent/fill-form` 端点可访问
  ```bash
  curl -X POST http://localhost:8080/api/agent/fill-form \
    -H "Content-Type: application/json" \
    -d '{"url":"https://example.com/form","fields":{"1":"test"}}'
  ```

- [ ] `POST /api/agent/submit-form` 端点可访问
  ```bash
  curl -X POST http://localhost:8080/api/agent/submit-form \
    -H "Content-Type: application/json" \
    -d '{"url":"https://example.com","submit_button_index":3}'
  ```

- [ ] `GET /api/agent/session` 端点可访问
  ```bash
  curl http://localhost:8080/api/agent/session
  ```

#### 3.2 Python FastAPI 端点测试

**启动 Python API 服务:**
```bash
cd python-ai
python -m uvicorn main:app --reload --port 8000
```

- [ ] `POST /api/agent/navigate` 端点可访问
  ```bash
  curl -X POST http://localhost:8000/api/agent/navigate \
    -H "Content-Type: application/json" \
    -d '{"url":"https://example.com","extractors":{"title":"h1"}}'
  ```

- [ ] `POST /api/agent/form-submit` 端点可访问
  ```bash
  curl -X POST http://localhost:8000/api/agent/form-submit \
    -H "Content-Type: application/json" \
    -d '{"url":"https://example.com","fields":{"1":"test"}}'
  ```

- [ ] `POST /api/agent/extract-table` 端点可访问
  ```bash
  curl -X POST http://localhost:8000/api/agent/extract-table \
    -H "Content-Type: application/json" \
    -d '{"url":"https://example.com/data","selector":"table"}'
  ```

- [ ] `GET /api/agent/session` 端点可访问
  ```bash
  curl http://localhost:8000/api/agent/session
  ```

---

### 第四阶段: 功能测试

#### 4.1 点击操作测试
- [ ] 能够导航到网页
- [ ] 能够点击页面元素
- [ ] 点击后能够获取新页面状态
- [ ] 点击操作返回正确的 HTTP 状态码

#### 4.2 表单填充测试
- [ ] 能够填充单个表单字段
- [ ] 能够填充多个表单字段
- [ ] 能够获取表单字段的当前值
- [ ] 能够清除表单字段

#### 4.3 数据提取测试
- [ ] 能够使用 CSS 选择器提取数据
- [ ] 能够提取表格数据
- [ ] 能够使用 JavaScript 执行代码
- [ ] 能够提取嵌套元素

#### 4.4 会话管理测试
- [ ] 每个会话有唯一的 ID
- [ ] 会话信息可以被检索
- [ ] 会话缓存正确工作
- [ ] 会话可以正确清除

---

### 第五阶段: Docker 集成测试

#### 5.1 Docker 构建
- [ ] Go API Docker 镜像构建成功
  ```bash
  docker build -f golang-api/Dockerfile -t penghu-api:latest .
  ```

- [ ] Python API Docker 镜像构建成功
  ```bash
  docker build -f python-ai/Dockerfile -t penghu-ai:latest .
  ```

#### 5.2 Docker Compose 运行
- [ ] 整个应用栈通过 docker-compose 启动
  ```bash
  docker-compose up -d
  ```

- [ ] 所有容器都在运行
  ```bash
  docker-compose ps
  ```

- [ ] Go API 服务可达
  ```bash
  curl http://localhost:8080/api/agent/session
  ```

- [ ] Python API 服务可达
  ```bash
  curl http://localhost:8000/api/agent/session
  ```

#### 5.3 Docker 数据持久化
- [ ] 容器停止后数据保持
- [ ] 容器重启后数据完整
- [ ] 日志正确输出

---

### 第六阶段: 性能和监控

#### 6.1 性能测试
- [ ] 单次导航操作 < 5 秒
- [ ] 表单填充 + 提交 < 3 秒
- [ ] 表格提取 < 10 秒
- [ ] 并发请求能够正确处理（10+ 并发）

#### 6.2 错误处理
- [ ] 网络错误被正确捕获
- [ ] 无效的 URL 返回适当的错误
- [ ] 导航失败返回清晰的错误消息
- [ ] 缺少依赖项显示明确的错误

#### 6.3 日志和监控
- [ ] 重要操作都被记录
- [ ] 错误被正确记录
- [ ] 日志级别可以配置
- [ ] 日志格式清晰易读

---

### 第七阶段: 文档验证

- [ ] `docs/INTEGRATION.md` 内容准确
  - [ ] 架构图正确
  - [ ] 所有代码示例都可运行
  - [ ] API 端点文档完整
  - [ ] 故障排除指南有帮助

- [ ] `README.md` 更新包含 Agent 信息
- [ ] 每个文件都有适当的注释
- [ ] 没有过时的文档

---

### 第八阶段: 安全性审查

#### 8.1 输入验证
- [ ] URL 参数被验证
- [ ] CSS 选择器被验证
- [ ] JavaScript 代码被隔离
- [ ] 表单数据被清理

#### 8.2 访问控制
- [ ] API 端点需要身份验证（如适用）
- [ ] 敏感操作被记录
- [ ] 没有敏感数据暴露在日志中

#### 8.3 依赖安全
- [ ] 所有依赖都是安全的（无已知漏洞）
- [ ] 依赖版本是最新的

---

### 第九阶段: 示例应用测试

#### 9.1 Go 示例
- [ ] `golang-api/examples/main.go` 编译成功
- [ ] 所有 8 个示例都能运行
- [ ] 示例输出清晰明确

#### 9.2 Python 示例
- [ ] `python-ai/examples.py` 可以执行
- [ ] 所有 8 个示例都能运行
- [ ] 示例演示了关键功能

#### 9.3 API 测试脚本
- [ ] `test-api.sh` 脚本可执行
- [ ] 所有 curl 命令都能运行
- [ ] 脚本输出组织良好

---

## ✅ 测试结果总结

| 阶段 | 测试数 | 通过 | 失败 | 状态 |
|------|--------|------|------|------|
| 1. 环境验证 | 10 | 0 | 0 | ⏳ |
| 2. 编译测试 | 12 | 0 | 0 | ⏳ |
| 3. API 集成 | 9 | 0 | 0 | ⏳ |
| 4. 功能测试 | 12 | 0 | 0 | ⏳ |
| 5. Docker 测试 | 8 | 0 | 0 | ⏳ |
| 6. 性能监控 | 10 | 0 | 0 | ⏳ |
| 7. 文档验证 | 5 | 0 | 0 | ⏳ |
| 8. 安全审查 | 7 | 0 | 0 | ⏳ |
| 9. 示例测试 | 6 | 0 | 0 | ⏳ |
| **总计** | **79** | **0** | **0** | **⏳** |

---

## 🐛 已知问题

（在此列出任何已知问题）

---

## 📝 测试人员信息

| 项 | 详情 |
|----|------|
| 测试人员 | TBD |
| 测试日期 | TBD |
| 环境 | Linux / macOS / Windows |
| 测试版本号 | 1.0.0 |

---

## 🔄 后续步骤

1. **第一次测试运行** - 完成所有清单项
2. **Bug 修复** - 记录并修复任何失败的测试
3. **性能优化** - 根据性能测试结果优化代码
4. **安全加固** - 实施所有建议的安全措施
5. **正式发布** - 完成所有测试后发布到生产环境

---

## 📞 联系和支持

如有问题，请参考：
- [INTEGRATION.md](docs/INTEGRATION.md) - 集成指南
- [README.md](README.md) - 项目概述
- [ARCHITECTURE.md](docs/ARCHITECTURE.md) - 系统架构
