# 🚀 澎湖數位老船長 - 服務啟動與部署指南

**项目名称**: Penghu Digital Captain (澎湖數位老船長)  
**版本**: 1.0.0  
**最后更新**: 2026-04-07  
**主要目录**: `/home/ouo/project`

---

## 📋 概述

本指南描述如何啟動澎湖數位老船長系統的所有服務，包括：

- **MariaDB 資料庫** (端口 3306)
- **Golang API 服務** (端口 8080)
- **Python AI 服務** (端口 8000)
- **Nginx Web 伺服器** (192.168.50.75:80)

系統採用 Docker Compose 進行容器化部署，確保一致性和易於管理。

---

## 🛠️ 環境要求

### 系統要求
- **作業系統**: Linux (Ubuntu/Debian/CentOS)
- **Docker**: 20.10+
- **Docker Compose**: 2.0+
- **記憶體**: 至少 4GB RAM
- **磁碟空間**: 至少 5GB 可用空間

### 網路要求
- **IP 位址**: 192.168.50.75 (用於 Nginx)
- **開放端口**:
  - 80 (HTTP)
  - 443 (HTTPS, 可選)
  - 3306 (MariaDB)
  - 8080 (Golang API)
  - 8000 (Python AI)

### 依賴檢查
```bash
# 檢查 Docker 和 Docker Compose
docker --version
docker-compose --version

# 檢查網路介面
ip addr show | grep 192.168.50.75

# 檢查端口是否可用
netstat -tlnp | grep -E ':80|:443|:3306|:8080|:8000'
```

---

## 📁 專案結構

```
/home/ouo/project/
├── docker-compose.yml          # Docker Compose 配置
├── .env                        # 環境變數 (需要配置)
├── .env.example               # 環境變數範例
├── DEPLOYMENT.md              # 🚀 服務啟動與部署指南 (本文件)
├── QUICKSTART.md              # 快速啟動指南
├── TESTING.md                 # 測試清單
├── PROJECT_COMPLETION.md      # 專案完成報告
├── verify-services.sh         # 🔍 服務驗證腳本
├── test-nginx-proxy.sh       # 🌐 Nginx 代理測試腳本
├── validate-nginx-config.sh  # ⚙️ Nginx 配置驗證腳本
├── test-api.sh                # 🧪 API 測試腳本
├── golang-api/                # Golang API 服務
│   ├── Dockerfile
│   ├── main.go
│   └── ...
├── python-ai/                 # Python AI 服務
│   ├── Dockerfile
│   ├── main.py
│   └── ...
├── nginx/                     # Nginx 配置
│   ├── conf/
│   │   └── nginx.conf
│   └── html/
├── web/                       # 靜態網頁檔案
├── mysql-init/               # 資料庫初始化腳本
├── docs/                      # 文檔
└── README.md                  # 專案總覽
```

---

## ⚙️ 環境配置

### 1. 複製環境變數檔案
```bash
cd /home/ouo/project
cp .env.example .env
```

### 2. 編輯環境變數
```bash
nano .env
```

**重要變數配置**:
```bash
# 資料庫配置
DB_ROOT_PASSWORD=your_secure_root_password
DB_NAME=aquaculture
DB_USER=appuser
DB_PASSWORD=your_secure_app_password

# API 金鑰 (可選)
OPENAI_API_KEY=your_openai_key
GOOGLE_API_KEY=your_google_key

# Nginx 綁定 IP (已配置為 192.168.50.75)
```

### 3. 設定網路介面 (如果需要)
確保 192.168.50.75 IP 位址可用：
```bash
# 檢查當前 IP
ip addr show

# 如果需要新增 IP (範例)
sudo ip addr add 192.168.50.75/24 dev eth0
```

---

## 🚀 服務啟動流程

### 步驟 1: 進入專案目錄
```bash
cd /home/ouo/project
```

### 步驟 2: 建置並啟動所有服務
```bash
# 建置並啟動 (推薦用於首次運行)
docker-compose up --build -d

# 或僅啟動 (如果已建置過)
docker-compose up -d
```

**預期輸出**:
```
Creating network "project_app-network" called a project
Creating aquaculture_db ... done
Creating golang-api ... done
Creating python-ai ... done
Creating nginx-proxy ... done
```

### 步驟 3: 驗證 Nginx 配置
```bash
# 驗證 Nginx 配置語法和設置
./validate-nginx-config.sh
```

**配置驗證將檢查**:
- Nginx 容器運行狀態
- 配置檔案語法正確性
- 靜態文件目錄存在性
- Upstream 服務配置
- 代理路由配置
- 端口綁定設置
```bash
# 查看所有容器狀態
docker-compose ps

# 查看日誌
docker-compose logs -f
```

**成功啟動的容器狀態**:
```
     Name                   Command               State                    Ports
------------------------------------------------------------------------------------------------
aquaculture_db     docker-entrypoint.sh mysqld      Up      0.0.0.0:3306->3306/tcp
golang-api         /app/main                        Up      0.0.0.0:8080->8080/tcp
nginx-proxy        /docker-entrypoint.sh nginx      Up      192.168.50.75:80->80/tcp, 192.168.50.75:443->443/tcp
python-ai          python -m uvicorn main:app ...   Up      0.0.0.0:8000->8000/tcp
```

### 步驟 4: 驗證服務可用性

#### 4.1 使用自動化驗證腳本
```bash
cd /home/ouo/project

# 運行服務驗證腳本
./verify-services.sh
```

#### 4.2 專門測試 Nginx 代理功能
```bash
# 運行 Nginx 代理測試腳本
./test-nginx-proxy.sh
```

**代理測試腳本將檢查**:
- Nginx 靜態文件服務
- API 代理到 Golang 服務
- API 代理到 Python AI 服務
- 代理標頭設置
- 路由配置正確性

#### 4.2 手動檢查 Nginx Web 伺服器
```bash
# 測試主頁面
curl -I http://192.168.50.75/

# 預期回應: HTTP/1.1 200 OK
```

#### 4.2 檢查健康狀態
```bash
# Nginx 健康檢查
curl http://192.168.50.75/health

# 預期回應: healthy
```

#### 4.3 檢查 API 服務
```bash
# Golang API 健康檢查
curl http://192.168.50.75/api/golang/health

# Python AI API 健康檢查
curl http://192.168.50.75/api/ai/docs
```

#### 4.4 檢查資料庫連線
```bash
# 從容器內測試
docker-compose exec golang-api nc -z mariadb 3306

# 或使用 MySQL 客戶端
mysql -h 127.0.0.1 -u appuser -p aquaculture
```

---

## 🔧 服務管理指令

### 停止所有服務
```bash
cd /home/ouo/project
docker-compose down
```

### 重新啟動特定服務
```bash
# 重新啟動 Nginx
docker-compose restart nginx

# 重新啟動 Golang API
docker-compose restart golang-api

# 重新啟動 Python AI
docker-compose restart python-ai
```

### 查看服務日誌
```bash
# 查看所有服務日誌
docker-compose logs -f

# 查看特定服務日誌
docker-compose logs -f golang-api
docker-compose logs -f python-ai
docker-compose logs -f nginx
```

### 進入容器進行除錯
```bash
# 進入 Golang API 容器
docker-compose exec golang-api sh

# 進入 Python AI 容器
docker-compose exec python-ai bash

# 進入 Nginx 容器
docker-compose exec nginx sh

# 進入資料庫容器
docker-compose exec mariadb bash
```

### 清理和重建
```bash
# 停止並刪除容器
docker-compose down

# 刪除映像和卷
docker-compose down --volumes --rmi all

# 重新建置
docker-compose up --build -d
```

---

## 🌐 服務端點說明

### Nginx Web 伺服器 (192.168.50.75:80)
- **主頁面**: http://192.168.50.75/
- **健康檢查**: http://192.168.50.75/health
- **靜態檔案**: http://192.168.50.75/ (從 `/web` 目錄提供)

### API 端點
- **Golang API**: http://192.168.50.75/api/golang/
  - 健康檢查: `/api/golang/health`
  - Agent 端點: `/api/golang/agent/*`
- **Python AI API**: http://192.168.50.75/api/ai/
  - 文檔: `/api/ai/docs`
  - Agent 端點: `/api/ai/agent/*`

### 直接訪問 (用於開發)
- **Golang API**: http://localhost:8080
- **Python AI**: http://localhost:8000
- **MariaDB**: localhost:3306

---

## 🧪 測試和驗證

### 運行自動化測試
```bash
cd /home/ouo/project

# 運行 API 測試腳本
bash test-api.sh

# 或手動測試
curl -X POST http://192.168.50.75/api/golang/agent/navigate \
  -H "Content-Type: application/json" \
  -d '{"url":"https://example.com"}'
```

### 運行程式示例
```bash
# Golang 示例
cd golang-api/examples
docker-compose exec golang-api go run main.go

# Python 示例
cd python-ai
docker-compose exec python-ai python examples.py
```

### 完整測試清單
參考 [TESTING.md](TESTING.md) 進行全面測試。

---

## 🚨 故障排除

### 常見問題

#### 問題 1: IP 位址 192.168.50.75 不可用
**解決方案**:
```bash
# 檢查網路介面
ip link show

# 新增 IP 位址 (如果需要)
sudo ip addr add 192.168.50.75/24 dev eth0

# 或修改 docker-compose.yml 使用不同 IP
# 將 "192.168.50.75:80:80" 改為 "80:80"
```

#### 問題 2: 端口 80 被佔用
**解決方案**:
```bash
# 檢查端口使用情況
sudo netstat -tlnp | grep :80

# 停止衝突服務
sudo systemctl stop apache2  # 或其他 web 服務

# 或使用不同端口
# 修改 docker-compose.yml 中的端口映射
```

#### 問題 3: 容器啟動失敗
**解決方案**:
```bash
# 查看詳細日誌
docker-compose logs

# 檢查資源使用情況
docker system df

# 清理未使用的資源
docker system prune -a
```

#### 問題 4: 資料庫連線失敗
**解決方案**:
```bash
# 檢查資料庫容器狀態
docker-compose ps mariadb

# 查看資料庫日誌
docker-compose logs mariadb

# 檢查環境變數
cat .env | grep DB_

# 手動測試連線
docker-compose exec mariadb mysql -u appuser -p aquaculture
```

#### 問題 5: API 服務無法訪問
**解決方案**:
```bash
# 檢查服務狀態
docker-compose ps

# 查看服務日誌
docker-compose logs golang-api
docker-compose logs python-ai

# 測試直接訪問
curl http://localhost:8080/health
curl http://localhost:8000/docs
```

#### 問題 6: Nginx 配置錯誤
**解決方案**:
```bash
# 驗證 Nginx 配置
./validate-nginx-config.sh

# 查看 Nginx 日誌
docker compose logs nginx

# 檢查配置檔案
docker compose exec nginx cat /etc/nginx/nginx.conf
```

#### 問題 7: API 代理失敗
**解決方案**:
```bash
# 測試代理功能
./test-nginx-proxy.sh

# 檢查後端服務
curl http://localhost:8080/health  # Golang API
curl http://localhost:8000/docs    # Python AI

# 檢查網路連線
docker compose exec nginx ping golang-api
docker compose exec nginx ping python-ai
```

#### 問題 8: 靜態文件無法訪問
**解決方案**:
```bash
# 檢查靜態文件目錄
docker compose exec nginx ls -la /usr/share/nginx/html/

# 檢查文件權限
docker compose exec nginx test -r /usr/share/nginx/html/index.html && echo "可讀" || echo "不可讀"

# 檢查 Nginx 訪問日誌
docker compose logs nginx | grep " 404 "
```

---

## 📊 監控和維護

### 資源監控
```bash
# 查看容器資源使用
docker stats

# 查看磁碟使用
docker system df

# 查看網路使用
docker network ls
```

### 日誌管理
```bash
# 查看所有日誌
docker-compose logs -f --tail=100

# 匯出日誌
docker-compose logs > logs_$(date +%Y%m%d_%H%M%S).txt
```

### 備份資料
```bash
# 備份資料庫
docker-compose exec mariadb mysqldump -u root -p aquaculture > backup_$(date +%Y%m%d).sql

# 備份環境配置
cp .env .env.backup
```

---

## 🔄 更新和升級

### 更新服務
```bash
# 拉取最新映像
docker-compose pull

# 重新建置
docker-compose up --build -d

# 檢查更新
docker-compose ps
```

### 滾動更新
```bash
# 零停機更新
docker-compose up -d --no-deps golang-api
docker-compose up -d --no-deps python-ai
docker-compose up -d --no-deps nginx
```

---

## 📞 支援和聯絡

如遇到問題，請：

1. 檢查 [故障排除](#故障排除) 部分
2. 查看詳細日誌: `docker-compose logs`
3. 參考 [TESTING.md](TESTING.md) 進行診斷
4. 查看 [QUICKSTART.md](QUICKSTART.md) 快速驗證

---

## ✅ 檢查清單

啟動前檢查:
- [ ] Docker 和 Docker Compose 已安裝
- [ ] IP 位址 192.168.50.75 可用
- [ ] 端口 80, 443, 3306, 8080, 8000 未被佔用
- [ ] `.env` 檔案已正確配置
- [ ] 專案目錄權限正確

啟動後驗證:
- [ ] 所有容器狀態為 "Up"
- [ ] Nginx 可在 http://192.168.50.75/ 訪問
- [ ] API 端點回應正常
- [ ] 資料庫連線正常
- [ ] 日誌無錯誤

---

**🎉 成功啟動後，您可以通過 http://192.168.50.75 訪問澎湖數位老船長系統！**

---

## 📋 快速檢查清單

### 啟動前檢查:
- [ ] Docker 和 Docker Compose 已安裝
- [ ] IP 位址 192.168.50.75 可用
- [ ] 端口 80, 443, 3306, 8080, 8000 未被佔用
- [ ] `.env` 檔案已正確配置
- [ ] 專案目錄權限正確

### 啟動後驗證:
- [ ] 運行 `./verify-services.sh` 檢查所有服務
- [ ] Nginx 可在 http://192.168.50.75/ 訪問
- [ ] API 端點回應正常
- [ ] 資料庫連線正常
- [ ] 日誌無錯誤

---

**文件位置**: `/home/ouo/project/DEPLOYMENT.md`  
**最後更新**: 2026-04-07  
**維護者**: 開發團隊