# 部署指南 - 澎湖數位老船長

## 目錄

1. [環境要求](#環境要求)
2. [本地開發部署](#本地開發部署)
3. [生產環境部署](#生產環境部署)
4. [故障排除](#故障排除)
5. [監控與日誌](#監控與日誌)
6. [備份與恢復](#備份與恢復)

---

## 環境要求

### 硬體要求

#### 開發環境
- CPU: 2+ 核心
- RAM: 4GB+
- 磁碟: 10GB+

#### 生產環境
- CPU: 4+ 核心
- RAM: 8GB+
- 磁碟: 50GB+ (附加備份空間)

### 軟體要求

#### 必須安裝

```bash
# 檢查安裝
docker --version
docker-compose --version

# 最低版本
Docker:         20.10+
Docker Compose: 2.0+
```

#### 可選安裝

```bash
# 開發工具
Golang:         1.21+
Python:         3.10+
MySQL Client:   8.0+
Git:            2.30+
```

### 網路要求

| 服務 | 端口 | 協議 |
|------|------|------|
| Nginx | 80 | HTTP |
| Golang API | 8080 | HTTP |
| Python AI | 8000 | HTTP |
| MariaDB | 3306 | TCP |

---

## 本地開發部署

### 1. 項目準備

```bash
# 切換到項目目錄
cd ~/project

# 複製環境配置
cp .env.example .env

# 編輯 .env 文件 (可選)
vim .env

# 檢查目錄結構
tree -L 2
```

### 2. 啟動所有服務

#### 方式 A: 使用啟動腳本 (推薦)

```bash
# 給予執行權限
chmod +x scripts/start.sh

# 執行啟動腳本
./scripts/start.sh

# 輸出示例:
# 🚀 澎湖數位老船長 - 啟動微服務系統
# 📋 啟動步驟:
# 1️⃣  拉取/更新容器鏡像...
# 2️⃣  構建自定義鏡像...
# 3️⃣  啟動所有服務...
# ⏳ 等待服務就緒 (30秒)...
# ✅ 服務啟動狀態:
```

#### 方式 B: 使用 Docker Compose 命令

```bash
# 後台啟動
docker-compose up -d

# 查看日誌
docker-compose logs -f

# 檢查服務狀態
docker-compose ps
```

### 3. 驗證部署

**等待 30 秒讓所有服務就緒**

```bash
# 檢查容器狀態
docker-compose ps

# 預期輸出:
# NAME                COMMAND             STATUS              PORTS
# aquaculture_db      mysql ...           Up (healthy)        3306
# golang-api          ./app               Up (healthy)        8080
# python-ai           uvicorn ...         Up (healthy)        8000
# nginx               nginx ...           Up (healthy)        80
```

### 4. 測試服務

```bash
# 1. 測試 Nginx
curl http://localhost/health

# 2. 測試 Golang API
curl http://localhost:8080/health
# 回應: {"status":"healthy","service":"golang-api"}

# 3. 測試 Python AI
curl http://localhost:8000/health
# 回應: {"status":"healthy","service":"python-ai"}

# 4. 訪問網頁首頁
# 開啟瀏覽器: http://localhost

# 5. 測試查詢 API
curl http://localhost:8080/api/assets
```

### 5. 查看日誌

```bash
# 查看所有服務日誌
docker-compose logs -f

# 查看特定服務日誌
docker-compose logs -f golang-api
docker-compose logs -f python-ai
docker-compose logs -f mariadb
docker-compose logs -f nginx

# 查看最後 100 行日誌
docker-compose logs --tail=100 golang-api
```

### 6. 停止服務

```bash
# 停止容器 (保留數據)
./scripts/stop.sh
# 或
docker-compose stop

# 完全移除 (刪除容器，保留卷)
docker-compose down

# 完全清除 (刪除容器和數據卷)
docker-compose down -v
```

---

## 生產環境部署

### 1. 生產環境檢查清單

- [ ] 更新 .env 文件所有敏感資訊
- [ ] 更改 MariaDB root 密碼
- [ ] 配置 SSL/TLS 證書
- [ ] 設置防火牆規則
- [ ] 配置 DNS 記錄
- [ ] 備份計劃
- [ ] 監控告警設置
- [ ] 日誌收集配置

### 2. 環境變數配置

```bash
# 生產環境 .env 示例

# 強密碼
DB_ROOT_PASSWORD=complex-password-2026!@#$%
DB_PASSWORD=complex-app-password!@#

# 環境標記
GO_ENV=production
PYTHON_ENV=production
LOG_LEVEL=warn

# API 安全
OPENAI_API_KEY=sk-your-production-key
GEMINI_API_KEY=your-production-gemini-key

# LINE 整合 (如需要)
LINE_CHANNEL_ID=your-production-channel-id
LINE_CHANNEL_SECRET=your-production-secret
```

### 3. SSL/TLS 配置

#### 使用 Let's Encrypt (推薦免費)

```bash
# 安裝 Certbot
sudo apt-get install certbot python3-certbot-nginx

# 獲得証書
sudo certbot certonly --standalone -d your-domain.com

# 証書位置
/etc/letsencrypt/live/your-domain.com/
```

#### Nginx SSL 配置

更新 `nginx/conf/nginx.conf`:

```nginx
server {
    listen 443 ssl http2;
    server_name your-domain.com;

    ssl_certificate /etc/letsencrypt/live/your-domain.com/fullchain.pem;
    ssl_certificate_key /etc/letsencrypt/live/your-domain.com/privkey.pem;
    ssl_protocols TLSv1.2 TLSv1.3;
    ssl_ciphers HIGH:!aNULL:!MD5;
    
    # ... 其他配置 ...
}

# HTTP 重定向至 HTTPS
server {
    listen 80;
    server_name your-domain.com;
    return 301 https://$server_name$request_uri;
}
```

### 4. 資源限制

`docker-compose.yml` 添加資源限制:

```yaml
services:
  golang-api:
    # ...
    deploy:
      resources:
        limits:
          cpus: '2'
          memory: 512M
        reservations:
          cpus: '0.5'
          memory: 256M

  python-ai:
    deploy:
      resources:
        limits:
          cpus: '2'
          memory: 1G
        reservations:
          cpus: '1'
          memory: 512M

  mariadb:
    deploy:
      resources:
        limits:
          cpus: '4'
          memory: 2G
        reservations:
          cpus: '2'
          memory: 1G
```

### 5. 部署在雲端 (示例: AWS/GCP/Azure)

#### AWS EC2 部署

```bash
# 1. 啟動 Ubuntu 實例
# 選擇: Ubuntu 22.04 LTS, t3.medium 或更高

# 2. 安裝 Docker
sudo apt-get update
sudo apt-get install -y docker.io docker-compose

# 3. 克隆項目
git clone <your-repo-url> ~/project

# 4. 改變目錄並啟動
cd ~/project
sudo docker-compose up -d

# 5. 配置安全組
# 允許: 80 (HTTP), 443 (HTTPS)
```

---

## 故障排除

### 服務無法啟動

**症狀:** 容器立即退出

```bash
# 1. 檢查日誌
docker-compose logs golang-api
docker-compose logs python-ai

# 2. 常見原因
# - 端口被佔用
# - 資料庫連接失敗
# - 環境變數缺失

# 解決:
# 檢查端口
sudo lsof -i :8080
sudo lsof -i :8000

# 清理與重啟
docker-compose down -v
docker-compose up -d
```

### 資料庫連接失敗

```bash
# 檢查 MariaDB 狀態
docker-compose exec mariadb mysql -u root -p$DB_ROOT_PASSWORD -e "SHOW DATABASES;"

# 檢查連接
docker-compose exec golang-api nc -zv mariadb 3306

# 重新初始化
docker-compose exec mariadb mysql -u root -p$DB_ROOT_PASSWORD < mysql-init/init.sql
```

### API 無回應

```bash
# 檢查服務是否運行
docker-compose ps

# 檢查網絡
docker network ls
docker network inspect project_app-network

# 測試連通性
docker-compose exec golang-api curl localhost:8080/health
docker-compose exec python-ai curl localhost:8000/health
```

### 記憶體不足

```bash
# 檢查資源使用
docker stats

# 增加 Docker 內存限制
# 修改 docker-compose.yml 中的 memory 限制

# 或重啟 Docker daemon
sudo systemctl restart docker
```

---

## 監控與日誌

### 1. 即時監控

```bash
# 實時資源監控
docker stats

# 持續日誌跟蹤
docker-compose logs -f --tail=50

# 按時間戳顯示
docker-compose logs --timestamps
```

### 2. 日誌收集

#### 配置 ELK Stack (可選)

```bash
# 收集 Docker 日誌到 Elasticsearch
# 這需要額外的 Filebeat 配置
```

#### 簡單的日誌備份

```bash
# 定期備份日誌
docker-compose logs > logs/backup-$(date +%Y%m%d).log

# 清理古舊日誌
find logs/ -mtime +30 -delete
```

### 3. 應用指標

**Golang API:**
- 響應時間
- 請求計數
- 錯誤率

**Python AI:**
- LLM API 調用次數
- Average latency
- Token 使用量

---

## 備份與恢復

### 1. 資料庫備份

```bash
# 完整備份
docker-compose exec mariadb mysqldump -u root -p$DB_ROOT_PASSWORD \
  --all-databases > backup-full-$(date +%Y%m%d).sql

# 單一資料庫備份
docker-compose exec mariadb mysqldump -u root -p$DB_ROOT_PASSWORD \
  aquaculture_db > backup-$(date +%Y%m%d).sql

# 排程備份 (cron)
# 編輯: crontab -e
# 添加: 0 2 * * * cd ~/project && docker-compose exec mariadb mysqldump -u root -p$DB_ROOT_PASSWORD aquaculture_db > backup-$(date +\%Y\%m\%d).sql
```

### 2. 資料庫恢復

```bash
# 從備份恢復
docker-compose exec -T mariadb mysql -u root -p$DB_ROOT_PASSWORD \
  aquaculture_db < backup-20260407.sql

# 驗證恢復
docker-compose exec mariadb mysql -u root -p$DB_ROOT_PASSWORD \
  -e "SELECT COUNT(*) FROM aquaculture_db.users;"
```

### 3. 卷備份

```bash
# 備份 MariaDB 卷
docker run --rm -v project_mariadb_data:/dbdata \
  -v $(pwd):/backup ubuntu tar czf /backup/mariadb-backup.tar.gz /dbdata

# 恢復卷
docker run --rm -v project_mariadb_data:/dbdata \
  -v $(pwd):/backup ubuntu tar xzf /backup/mariadb-backup.tar.gz -C /
```

### 4. 完整系統備份

```bash
# 備份整個項目 (包括代碼和配置)
tar -czf project-backup-$(date +%Y%m%d).tar.gz \
  --exclude='mariadb_data' \
  --exclude='.git' \
  ~/project

# 備份到遠程伺服器
scp project-backup-*.tar.gz user@remote-server:/backups/
```

---

## 自動更新與補丁

### 1. 定期更新容器鏡像

```bash
# 檢查更新
docker-compose pull

# 重新構建與重啟
docker-compose up -d --build
```

### 2. 安全補丁

```bash
# 更新基礎鏡像
# 編輯 Dockerfile，更新 FROM 行到最新版本

# 重新構建
docker-compose build --no-cache
docker-compose up -d
```

---

## 性能調優

### 1. MariaDB 調優

```bash
# 連接到容器並配置
docker-compose exec mariadb mysql -u root -p$DB_ROOT_PASSWORD

# 執行優化命令
OPTIMIZE TABLE users;
OPTIMIZE TABLE assets;
OPTIMIZE TABLE tasks;
OPTIMIZE TABLE environmental_data;
```

### 2. 應用層調優

**Golang:**
- 增加 worker 並發數
- 實施連接池
- 啟用 HTTP/2

**Python:**
- 增加 Uvicorn workers
- 啟用異步處理
- 實施快取層

### 3. 基礎設施調優

- 啟用 gzip 壓縮
- CDN 快取
- 數據庫複製
- 讀寫分離

---

## 災難恢復計劃

### RTO/RPO 目標

- **RTO (恢復時間目標):** < 5 分鐘
- **RPO (恢復點目標):** < 1 小時

### 備份策略

| 類型 | 頻率 | 保留期 | 位置 |
|------|------|--------|------|
| 完整備份 | 每日 0:00 | 30 天 | 本地 + 雲端 |
| 增量備份 | 每 6 小時 | 7 天 | 本地 |
| 日誌備份 | 每小時 | 3 天 | 本地 |

### 測試計劃

- 每周進行恢復測試
- 月度完整系統演習
- 文檔更新

---

## 聯絡支援

有任何部署相關問題，請參考：

- 📖 [README.md](../README.md)
- 📚 [API 文檔](./API.md)
- 🏗️ [架構文檔](./ARCHITECTURE.md)
- 🐛 提交 Issue 到版本控制系統

---

**澎湖數位老船長** 🐠 | 部署指南 v1.0 | 最後更新: 2026-04-07
