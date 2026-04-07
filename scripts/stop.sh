#!/bin/bash

# 澎湖數位老船長 - 停止腳本

set -e

echo "🛑 澎湖數位老船長 - 停止微服務系統"
echo "=================================="

echo ""
echo "1️⃣  停止所有容器..."
docker-compose down

echo ""
echo "✅ 所有服務已停止"
echo ""
echo "💾 數據已保存（MariaDB 數據卷未刪除）"
echo ""
echo "提示: 要完全清除數據和卷，執行:"
echo "   docker-compose down -v"
