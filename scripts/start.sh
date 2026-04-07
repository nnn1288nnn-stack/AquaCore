#!/bin/bash

# 澎湖數位老船長 - 啟動腳本
# 快速啟動所有微服務

set -e

echo "🚀 澎湖數位老船長 - 啟動微服務系統"
echo "=================================="

# 檢查 Docker
if ! command -v docker &> /dev/null; then
    echo "❌ 錯誤: Docker 未安裝"
    exit 1
fi

if ! command -v docker-compose &> /dev/null; then
    echo "❌ 錯誤: Docker Compose 未安裝"
    exit 1
fi

# 檢查 .env 文件
if [ ! -f .env ]; then
    echo "⚠️  警告: .env 文件不存在，使用默認配置"
    echo "   建議: cp .env.example .env 並編輯配置"
fi

echo ""
echo "📋 啟動步驟:"
echo "1️⃣  拉取/更新容器镜像..."
docker-compose pull

echo ""
echo "2️⃣  構建自定義鏡像..."
docker-compose build --no-cache

echo ""
echo "3️⃣  啟動所有服務..."
docker-compose up -d

echo ""
echo "⏳ 等待服務就緒 (30秒)..."
sleep 30

echo ""
echo "✅ 服務啟動狀態:"
docker-compose ps

echo ""
echo "📍 服務地址:"
echo "   🌐 網頁首頁: http://localhost"
echo "   🚀 Golang API: http://localhost:8080"
echo "   🤖 Python AI: http://localhost:8000"
echo "   🗄️  MariaDB: localhost:3306"
echo ""
echo "📖 查看日誌:"
echo "   docker-compose logs -f"
echo ""
echo "✨ 澎湖數位老船長已啟動！"
