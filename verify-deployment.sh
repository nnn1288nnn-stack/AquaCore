#!/bin/bash

# 澎湖數位老船長 - 部署驗證腳本
# 驗證 Vue.js 前端與微服務架構的整合

echo "🐠 澎湖數位老船長 - 部署驗證"
echo "================================="

# 檢查 Docker 服務狀態
echo "📋 檢查服務狀態..."
docker compose ps

echo ""
echo "🌐 測試前端訪問..."
curl -s -o /dev/null -w "HTTP 狀態: %{http_code}\n" http://192.168.50.75

echo ""
echo "🔍 檢查前端內容..."
if curl -s http://192.168.50.75 | grep -q "澎湖數位老船長"; then
    echo "✅ 前端應用正常載入"
else
    echo "❌ 前端應用載入失敗"
fi

echo ""
echo "🔗 測試 API 代理..."
# 注意: 這裡的 API 路徑需要根據實際配置調整
echo "API 代理測試將在後端服務啟動後進行"

echo ""
echo "🎉 部署驗證完成！"
echo "訪問 http://192.168.50.75 查看應用"