#!/bin/bash

# 澎湖數位老船長 - Nginx 代理測試腳本
# Penghu Digital Captain - Nginx Proxy Test Script

echo "🐠 澎湖數位老船長 - Nginx 代理測試"
echo "==================================="

# 顏色定義
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

NGINX_URL="http://192.168.50.75"

# 測試函數
test_endpoint() {
    local url=$1
    local description=$2
    local expected_status=${3:-200}

    echo -n "測試 $description: "
    response=$(curl -s -w "HTTPSTATUS:%{http_code}" "$url" 2>/dev/null)
    body=$(echo "$response" | sed 's/HTTPSTATUS.*//')
    status=$(echo "$response" | grep -o "HTTPSTATUS:[0-9]*" | cut -d: -f2)

    if [ "$status" = "$expected_status" ]; then
        echo -e "${GREEN}✅ 成功 (狀態: $status)${NC}"
        return 0
    else
        echo -e "${RED}❌ 失敗 (狀態: $status)${NC}"
        echo -e "${YELLOW}  URL: $url${NC}"
        return 1
    fi
}

echo -e "\n${BLUE}🔍 測試 Nginx 基本功能${NC}"
echo "-----------------------------"

# 測試靜態網頁
test_endpoint "$NGINX_URL/" "靜態網頁服務" 200

# 測試健康檢查
test_endpoint "$NGINX_URL/health" "健康檢查端點" 200

echo -e "\n${BLUE}🔗 測試 API 代理功能${NC}"
echo "-----------------------------"

# 測試 Golang API 代理
echo -e "\n${YELLOW}Golang API 代理測試:${NC}"
test_endpoint "$NGINX_URL/api/golang/health" "Golang API 健康檢查" || echo "  💡 請確保 Golang API 服務正在運行"

# 測試 Python AI API 代理
echo -e "\n${YELLOW}Python AI API 代理測試:${NC}"
test_endpoint "$NGINX_URL/api/ai/docs" "Python AI API 文檔" || echo "  💡 請確保 Python AI 服務正在運行"

# 測試默認 API 路由
echo -e "\n${YELLOW}默認 API 路由測試:${NC}"
test_endpoint "$NGINX_URL/api/health" "默認 API 路由" || echo "  💡 這將路由到 Golang API"

echo -e "\n${BLUE}📊 測試詳細信息${NC}"
echo "-----------------------------"

# 顯示 Nginx 配置摘要
echo -e "\n${YELLOW}Nginx 配置摘要:${NC}"
echo "  監聽地址: 192.168.50.75:80"
echo "  靜態文件: /usr/share/nginx/html"
echo "  Golang API: /api/golang/ → golang-api:8080"
echo "  Python AI: /api/ai/ → python-ai:8000"
echo "  默認 API: /api/ → golang-api:8080"

# 檢查代理標頭
echo -e "\n${YELLOW}檢查代理標頭:${NC}"
response=$(curl -s -I "$NGINX_URL/api/golang/health" 2>/dev/null | grep -E "(X-|Host:)" || echo "無代理標頭")
if [ "$response" != "無代理標頭" ]; then
    echo -e "${GREEN}✅ 代理標頭正確設置${NC}"
    echo "$response" | sed 's/^/  /'
else
    echo -e "${RED}❌ 代理標頭未設置${NC}"
fi

echo -e "\n${BLUE}🔧 故障排除提示${NC}"
echo "-----------------------------"
echo "如果代理測試失敗:"
echo "1. 確保所有 Docker 容器都在運行:"
echo "   docker compose ps"
echo ""
echo "2. 檢查容器日誌:"
echo "   docker compose logs nginx"
echo "   docker compose logs golang-api"
echo "   docker compose logs python-ai"
echo ""
echo "3. 測試直接訪問:"
echo "   curl http://localhost:8080/health  # Golang API"
echo "   curl http://localhost:8000/docs    # Python AI"
echo ""
echo "4. 檢查網路連線:"
echo "   docker compose exec nginx ping golang-api"
echo "   docker compose exec nginx ping python-ai"

echo -e "\n${GREEN}🎉 Nginx 代理測試完成！${NC}"