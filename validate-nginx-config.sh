#!/bin/bash

# 澎湖數位老船長 - Nginx 配置驗證腳本
# Penghu Digital Captain - Nginx Configuration Validation Script

echo "🐠 澎湖數位老船長 - Nginx 配置驗證"
echo "==================================="

# 顏色定義
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# 檢查 Nginx 容器是否運行
echo -e "\n${YELLOW}檢查 Nginx 容器狀態...${NC}"
if docker ps | grep -q nginx-proxy; then
    echo -e "${GREEN}✅ Nginx 容器正在運行${NC}"
else
    echo -e "${RED}❌ Nginx 容器未運行${NC}"
    echo -e "${YELLOW}請運行: docker compose up -d nginx${NC}"
    exit 1
fi

# 測試 Nginx 配置語法
echo -e "\n${YELLOW}測試 Nginx 配置語法...${NC}"
if docker compose exec -T nginx nginx -t 2>/dev/null; then
    echo -e "${GREEN}✅ Nginx 配置語法正確${NC}"
else
    echo -e "${RED}❌ Nginx 配置語法錯誤${NC}"
    echo -e "${YELLOW}查看詳細錯誤:${NC}"
    docker compose exec nginx nginx -t
    exit 1
fi

# 檢查配置文件
echo -e "\n${YELLOW}檢查配置文件...${NC}"
if docker compose exec -T nginx test -f /etc/nginx/nginx.conf; then
    echo -e "${GREEN}✅ Nginx 配置文件存在${NC}"
else
    echo -e "${RED}❌ Nginx 配置文件不存在${NC}"
    exit 1
fi

# 檢查靜態文件目錄
echo -e "\n${YELLOW}檢查靜態文件目錄...${NC}"
if docker compose exec -T nginx test -d /usr/share/nginx/html; then
    echo -e "${GREEN}✅ 靜態文件目錄存在${NC}"

    # 檢查 index.html
    if docker compose exec -T nginx test -f /usr/share/nginx/html/index.html; then
        echo -e "${GREEN}✅ index.html 文件存在${NC}"
    else
        echo -e "${RED}❌ index.html 文件不存在${NC}"
    fi
else
    echo -e "${RED}❌ 靜態文件目錄不存在${NC}"
fi

# 檢查 upstream 配置
echo -e "\n${YELLOW}檢查 Upstream 配置...${NC}"
config=$(docker compose exec -T nginx cat /etc/nginx/nginx.conf 2>/dev/null)

if echo "$config" | grep -q "upstream golang_api"; then
    echo -e "${GREEN}✅ Golang API upstream 配置存在${NC}"
else
    echo -e "${RED}❌ Golang API upstream 配置缺失${NC}"
fi

if echo "$config" | grep -q "upstream python_ai"; then
    echo -e "${GREEN}✅ Python AI upstream 配置存在${NC}"
else
    echo -e "${RED}❌ Python AI upstream 配置缺失${NC}"
fi

# 檢查代理配置
echo -e "\n${YELLOW}檢查代理配置...${NC}"
if echo "$config" | grep -q "location /api/golang/"; then
    echo -e "${GREEN}✅ Golang API 代理配置存在${NC}"
else
    echo -e "${RED}❌ Golang API 代理配置缺失${NC}"
fi

if echo "$config" | grep -q "location /api/ai/"; then
    echo -e "${GREEN}✅ Python AI 代理配置存在${NC}"
else
    echo -e "${RED}❌ Python AI 代理配置缺失${NC}"
fi

# 檢查端口綁定
echo -e "\n${YELLOW}檢查端口綁定...${NC}"
if docker port nginx-proxy 2>/dev/null | grep -q "192.168.50.75:80"; then
    echo -e "${GREEN}✅ 端口 80 正確綁定到 192.168.50.75${NC}"
else
    echo -e "${RED}❌ 端口 80 未正確綁定${NC}"
fi

# 顯示配置摘要
echo -e "\n${BLUE}📋 配置摘要${NC}"
echo "-----------------------------"
echo "Nginx 配置檔案: /etc/nginx/nginx.conf"
echo "靜態文件目錄: /usr/share/nginx/html"
echo "監聽端口: 80 (綁定到 192.168.50.75)"
echo "Upstream 服務:"
echo "  - golang_api: golang-api:8080"
echo "  - python_ai: python-ai:8000"
echo "代理路由:"
echo "  - /api/golang/ → golang_api"
echo "  - /api/ai/ → python_ai"
echo "  - /api/ → golang_api (默認)"

echo -e "\n${GREEN}🎉 Nginx 配置驗證完成！${NC}"

# 提供測試建議
echo -e "\n${BLUE}💡 下一步測試建議${NC}"
echo "-----------------------------"
echo "1. 運行代理測試: ./test-nginx-proxy.sh"
echo "2. 運行完整驗證: ./verify-services.sh"
echo "3. 訪問網頁: http://192.168.50.75/"