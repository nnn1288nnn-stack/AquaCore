#!/bin/bash

# 澎湖數位老船長 - API 测试命令
# Penghu Digital Captain - API Testing Commands

# 颜色定义
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# API 基础 URL
API_BASE="http://localhost:8080"

echo -e "${BLUE}╔════════════════════════════════════════════════════════════╗${NC}"
echo -e "${BLUE}║     澎湖數位老船長 - Agent API 测试命令                   ║${NC}"
echo -e "${BLUE}║   Penghu Digital Captain - Agent API Testing Commands    ║${NC}"
echo -e "${BLUE}╚════════════════════════════════════════════════════════════╝${NC}\n"

# 测试 1: 简单导航和提取
echo -e "${YELLOW}[测试 1] 导航到网站并提取数据${NC}"
echo -e "${BLUE}命令:${NC}"
echo 'curl -X POST http://localhost:8080/api/agent/navigate \'
echo '  -H "Content-Type: application/json" \'
echo '  -d '\''{'
echo '    "url": "https://example.com",'
echo '    "extractors": {'
echo '      "title": "h1",'
echo '      "paragraphs": "p"'
echo '    }'
echo '  }'\'

echo -e "\n${GREEN}完整命令:${NC}\n"

curl -X POST $API_BASE/api/agent/navigate \
  -H "Content-Type: application/json" \
  -d '{
    "url": "https://example.com",
    "extractors": {
      "title": "h1",
      "paragraphs": "p"
    }
  }' 2>/dev/null | jq . || echo "❌ 请确保 Go API 服务正在运行 (http://localhost:8080)"

echo -e "\n═══════════════════════════════════════════════════════════\n"

# 测试 2: 获取页面状态
echo -e "${YELLOW}[测试 2] 获取当前页面状态${NC}"
echo -e "${BLUE}命令:${NC}"
echo 'curl -X POST http://localhost:8080/api/agent/state \'
echo '  -H "Content-Type: application/json" \'
echo '  -d '\''{"url": "https://example.com"}'\'

echo -e "\n${GREEN}完整命令:${NC}\n"

curl -X POST $API_BASE/api/agent/state \
  -H "Content-Type: application/json" \
  -d '{"url": "https://example.com"}' 2>/dev/null | jq . || echo "❌ 请确保 Go API 服务正在运行"

echo -e "\n═══════════════════════════════════════════════════════════\n"

# 测试 3: 点击元素
echo -e "${YELLOW}[测试 3] 点击页面元素${NC}"
echo -e "${BLUE}命令:${NC}"
echo 'curl -X POST http://localhost:8080/api/agent/click \'
echo '  -H "Content-Type: application/json" \'
echo '  -d '\''{'
echo '    "url": "https://example.com",'
echo '    "element_index": 5'
echo '  }'\'

echo -e "\n${GREEN}完整命令:${NC}\n"

curl -X POST $API_BASE/api/agent/click \
  -H "Content-Type: application/json" \
  -d '{
    "url": "https://example.com",
    "element_index": 5
  }' 2>/dev/null | jq . || echo "❌ 请确保 Go API 服务正在运行"

echo -e "\n═══════════════════════════════════════════════════════════\n"

# 测试 4: 填充表单字段
echo -e "${YELLOW}[测试 4] 填充表单字段${NC}"
echo -e "${BLUE}命令:${NC}"
echo 'curl -X POST http://localhost:8080/api/agent/fill-form \'
echo '  -H "Content-Type: application/json" \'
echo '  -d '\''{'
echo '    "url": "https://example.com/form",'
echo '    "fields": {'
echo '      "1": "user@example.com",'
echo '      "2": "password123"'
echo '    }'
echo '  }'\'

echo -e "\n${GREEN}完整命令:${NC}\n"

curl -X POST $API_BASE/api/agent/fill-form \
  -H "Content-Type: application/json" \
  -d '{
    "url": "https://example.com/form",
    "fields": {
      "1": "user@example.com",
      "2": "password123"
    }
  }' 2>/dev/null | jq . || echo "❌ 请确保 Go API 服务正在运行"

echo -e "\n═══════════════════════════════════════════════════════════\n"

# 测试 5: 提交表单
echo -e "${YELLOW}[测试 5] 提交表单${NC}"
echo -e "${BLUE}命令:${NC}"
echo 'curl -X POST http://localhost:8080/api/agent/submit-form \'
echo '  -H "Content-Type: application/json" \'
echo '  -d '\''{'
echo '    "url": "https://example.com/form",'
echo '    "submit_button_index": 3,'
echo '    "wait_time": 2'
echo '  }'\'

echo -e "\n${GREEN}完整命令:${NC}\n"

curl -X POST $API_BASE/api/agent/submit-form \
  -H "Content-Type: application/json" \
  -d '{
    "url": "https://example.com/form",
    "submit_button_index": 3,
    "wait_time": 2
  }' 2>/dev/null | jq . || echo "❌ 请确保 Go API 服务正在运行"

echo -e "\n═══════════════════════════════════════════════════════════\n"

# 测试 6: 获取会话信息
echo -e "${YELLOW}[测试 6] 获取 Agent 会话信息${NC}"
echo -e "${BLUE}命令:${NC}"
echo 'curl -X GET http://localhost:8080/api/agent/session'

echo -e "\n${GREEN}完整命令:${NC}\n"

curl -X GET $API_BASE/api/agent/session \
  -H "Content-Type: application/json" 2>/dev/null | jq . || echo "❌ 请确保 Go API 服务正在运行"

echo -e "\n═══════════════════════════════════════════════════════════\n"

# 测试 7: Python FastAPI - 导航和提取
echo -e "${YELLOW}[测试 7] Python FastAPI - 导航和提取${NC}"
echo -e "${BLUE}命令:${NC}"
echo 'curl -X POST http://localhost:8000/api/agent/navigate \'
echo '  -H "Content-Type: application/json" \'
echo '  -d '\''{'
echo '    "url": "https://example.com",'
echo '    "extractors": {'
echo '      "title": "h1"'
echo '    }'
echo '  }'\'

echo -e "\n${GREEN}完整命令:${NC}\n"

curl -X POST http://localhost:8000/api/agent/navigate \
  -H "Content-Type: application/json" \
  -d '{
    "url": "https://example.com",
    "extractors": {
      "title": "h1"
    }
  }' 2>/dev/null | jq . || echo "❌ 请确保 Python API 服务正在运行 (http://localhost:8000)"

echo -e "\n═══════════════════════════════════════════════════════════\n"

# 测试 8: Python FastAPI - 表格提取
echo -e "${YELLOW}[测试 8] Python FastAPI - 表格提取${NC}"
echo -e "${BLUE}命令:${NC}"
echo 'curl -X POST http://localhost:8000/api/agent/extract-table \'
echo '  -H "Content-Type: application/json" \'
echo '  -d '\''{'
echo '    "url": "https://example.com/data",'
echo '    "selector": "table.data-table"'
echo '  }'\'

echo -e "\n${GREEN}完整命令:${NC}\n"

curl -X POST http://localhost:8000/api/agent/extract-table \
  -H "Content-Type: application/json" \
  -d '{
    "url": "https://example.com/data",
    "selector": "table.data-table"
  }' 2>/dev/null | jq . || echo "❌ 请确保 Python API 服务正在运行"

echo -e "\n═══════════════════════════════════════════════════════════\n"

# 测试 9: Python FastAPI - 获取会话
echo -e "${YELLOW}[测试 9] Python FastAPI - 获取 Agent 会话${NC}"
echo -e "${BLUE}命令:${NC}"
echo 'curl -X GET http://localhost:8000/api/agent/session'

echo -e "\n${GREEN}完整命令:${NC}\n"

curl -X GET http://localhost:8000/api/agent/session \
  -H "Content-Type: application/json" 2>/dev/null | jq . || echo "❌ 请确保 Python API 服务正在运行"

echo -e "\n═══════════════════════════════════════════════════════════\n"

echo -e "${GREEN}✅ API 测试脚本完成${NC}\n"

echo "提示:"
echo "  • Go API 服务: http://localhost:8080"
echo "  • Python API 服务: http://localhost:8000"
echo "  • 确保 OpenCli 已在系统中可用"
echo "  • 使用 'jq' 格式化 JSON 输出"
