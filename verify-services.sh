#!/bin/bash

# 澎湖數位老船長 - 服務驗證腳本
# Penghu Digital Captain - Service Verification Script

echo "🐠 澎湖數位老船長 - 服務驗證"
echo "================================="

# 顏色定義
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# 檢查 Docker 服務狀態
echo -e "\n${YELLOW}📋 檢查 Docker 容器狀態...${NC}"
if command -v docker &> /dev/null && command -v docker-compose &> /dev/null; then
    cd /home/ouo/project
    if docker-compose ps | grep -q "Up"; then
        echo -e "${GREEN}✅ Docker 服務運行中${NC}"
        docker-compose ps --format "table {{.Name}}\t{{.Status}}\t{{.Ports}}"
    else
        echo -e "${RED}❌ Docker 服務未運行${NC}"
        echo -e "${YELLOW}💡 請運行: docker-compose up -d${NC}"
        exit 1
    fi
else
    echo -e "${RED}❌ Docker 或 Docker Compose 未安裝${NC}"
    exit 1
fi

# 檢查 Nginx Web 服務器
echo -e "\n${YELLOW}🌐 檢查 Nginx Web 服務器...${NC}"
if curl -s --max-time 5 http://192.168.50.75/ > /dev/null; then
    echo -e "${GREEN}✅ Nginx Web 服務器可訪問${NC}"

    # 檢查健康狀態
    if curl -s http://192.168.50.75/health | grep -q "healthy"; then
        echo -e "${GREEN}✅ 健康檢查通過${NC}"
    else
        echo -e "${RED}❌ 健康檢查失敗${NC}"
    fi

    # 檢查代理功能 - Golang API
    if curl -s --max-time 5 http://192.168.50.75/api/golang/health > /dev/null; then
        echo -e "${GREEN}✅ Nginx 代理到 Golang API 正常${NC}"
    else
        echo -e "${RED}❌ Nginx 代理到 Golang API 失敗${NC}"
    fi

    # 檢查代理功能 - Python AI API
    if curl -s --max-time 5 http://192.168.50.75/api/ai/docs > /dev/null; then
        echo -e "${GREEN}✅ Nginx 代理到 Python AI API 正常${NC}"
    else
        echo -e "${RED}❌ Nginx 代理到 Python AI API 失敗${NC}"
    fi
else
    echo -e "${RED}❌ Nginx Web 服務器不可訪問${NC}"
    echo -e "${YELLOW}💡 請檢查 IP 位址 192.168.50.75 是否正確配置${NC}"
fi

# 檢查 API 服務
echo -e "\n${YELLOW}🔌 檢查 API 服務...${NC}"

# Golang API
if curl -s --max-time 5 http://localhost:8080/health > /dev/null; then
    echo -e "${GREEN}✅ Golang API 服務運行中${NC}"
else
    echo -e "${RED}❌ Golang API 服務不可訪問${NC}"
fi

# Python AI API
if curl -s --max-time 5 http://localhost:8000/docs > /dev/null; then
    echo -e "${GREEN}✅ Python AI API 服務運行中${NC}"
else
    echo -e "${RED}❌ Python AI API 服務不可訪問${NC}"
fi

# 檢查資料庫連線
echo -e "\n${YELLOW}🗄️ 檢查資料庫連線...${NC}"
if docker-compose exec -T mariadb mysqladmin ping -h localhost --silent; then
    echo -e "${GREEN}✅ MariaDB 資料庫連線正常${NC}"
else
    echo -e "${RED}❌ MariaDB 資料庫連線失敗${NC}"
fi

# 檢查網路連線
echo -e "\n${YELLOW}📡 檢查網路配置...${NC}"
if ip addr show | grep -q "192.168.50.75"; then
    echo -e "${GREEN}✅ IP 位址 192.168.50.75 已配置${NC}"
else
    echo -e "${YELLOW}⚠️ IP 位址 192.168.50.75 未配置${NC}"
fi

# 檢查端口
echo -e "\n${YELLOW}🔌 檢查端口狀態...${NC}"
ports=(80 443 3306 8080 8000)
for port in "${ports[@]}"; do
    if netstat -tln | grep -q ":$port "; then
        echo -e "${GREEN}✅ 端口 $port 正在監聽${NC}"
    else
        echo -e "${RED}❌ 端口 $port 未監聽${NC}"
    fi
done

echo -e "\n${GREEN}🎉 驗證完成！${NC}"
echo -e "${YELLOW}💡 如有問題，請參考 [DEPLOYMENT.md](DEPLOYMENT.md) 故障排除部分${NC}"