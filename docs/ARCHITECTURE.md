# 系統架構 - 澎湖數位老船長

## 1. 架構概览

澎湖數位老船長採用**微服務架構 (Microservices Architecture)**，將系統拆分為多個獨立、可擴展的服務，各自負責特定的業務領域。

```
┌─────────────────────────────────────────────────────┐
│                前端層 (Frontend)                    │
│  ┌─────────────┐  ┌──────────────┐  ┌────────────┐ │
│  │ LINE LIFF   │  │ Web Browser  │  │ Mobile App │ │
│  └──────┬──────┘  └──────┬───────┘  └─────┬──────┘ │
│         └──────────────────┬────────────────┘       │
└─────────────────────────────┼──────────────────────┘
                              │ HTTP/HTTPS
                              ↓
          ┌───────────────────────────────────────┐
          │     API 網關層 - Nginx (Port 80)      │
          │  - 反向代理                           │
          │  - 靜態文件服務                       │
          │  - 負載平衡                           │
          └───────┬──────────────┬──────────────┘
                  │              │
          ┌───────▼─────┐  ┌────▼──────────┐
          │             │  │              │
┌─────────▼──────────┐  │  ┌──────────────▼───────────┐
│  Golang 微服務 A   │  │  │ Python 微服務 B          │
│  (Port 8080)       │  │  │ (Port 8000)              │
├────────────────────┤  │  ├──────────────────────────┤
│ • 業務 API         │  │  │ • AI Agent               │
│ • CRUD 操作        │  │  │ • LLM 串接               │
│ • 高併發處理       │  │  │ • Tool Calling           │
│ • 資料驗證         │  │  │ • 自然語言推理           │
└────┬───────────────┘  │  └─────────┬────────────────┘
     │                  │            │
     │ HTTP 連接        │ HTTP 連接  │
     │                  │            │
     └──────────┬───────┴────────────┘
                │
       ┌────────▼────────┐
       │  MariaDB 資料庫 │
       │  (Port 3306)    │
       │                 │
       │  ┌────────────┐ │
       │  │ Users      │ │
       │  │ Tasks      │ │
       │  │ Assets     │ │
       │  │ Env Data   │ │
       │  └────────────┘ │
       └─────────────────┘
```

## 2. 核心服務組件

### 2.1 Golang 微服務 (業務 API)

**職責:**
- 處理所有 CRUD 操作
- 資料驗證與業務邏輯
- 高併發請求處理
- 資料庫連接 (GORM ORM)

**技術棧:**
- Golang 1.21+
- Gin Web Framework
- GORM ORM
- MariaDB Driver

**特點:**
- ✅ 高性能 (C 級編譯性能)
- ✅ 低內存占用
- ✅ 天生併發支援 (Goroutines)
- ✅ 靜態類型安全

**API 端點:**
```
GET    /api/dashboard              - 儀表板
GET    /api/environmental-data     - 環境數據列表
POST   /api/environmental-data     - 記錄環境數據
GET    /api/assets                 - 資產列表
POST   /api/assets                 - 新增資產
PUT    /api/assets/:id             - 更新資產
DELETE /api/assets/:id             - 刪除資產
GET    /api/tasks                  - 任務列表
POST   /api/tasks                  - 新增任務
PUT    /api/tasks/:id              - 更新任務
DELETE /api/tasks/:id              - 刪除任務
```

### 2.2 Python 微服務 (AI & Agent)

**職責:**
- LLM 模型串接 (OpenAI GPT-4o / Gemini)
- AI Agent 邏輯
- Tool Calling (函數調用)
- 自然語言處理

**技術棧:**
- Python 3.10+
- FastAPI Web Framework
- LangChain Agent Framework
- OpenAI SDK / Google Gemini API

**特點:**
- ✅ 豐富的 AI 生態 (LangChain, LLaMA, 等)
- ✅ 快速開發 (動態類型)
- ✅ 大量預構建庫
- ✅ 易於原型開發

**API 端點:**
```
POST   /api/chat                   - 聊天端點 (AI Agent)
POST   /api/agent/check-inventory  - 工具:查詢庫存
POST   /api/agent/create-task      - 工具:建立任務
POST   /api/agent/generate-report  - 工具:生成報表
```

### 2.3 MariaDB 資料庫

**職責:**
- 持久化數據存儲
- 交易支援
- 性能優化

**技術棧:**
- MariaDB 10.11 (MySQL 兼容)
- InnoDB 存儲引擎
- UTF-8MB4 编码（完整中文支援）

**資料表:**
- `users` - 用戶信息
- `environmental_data` - 環境監測數據
- `assets` - 庫存資產
- `tasks` - 任務管理
- `operation_logs` - 操作日誌

### 2.4 Nginx 網頁伺服器

**職責:**
- 反向代理
- 靜態文件服務
- 負載平衡
- SSL/TLS 終止

**配置:**
```nginx
# 代理 Golang API
location /api/golang { proxy_pass http://golang-api:8080; }

# 代理 Python AI
location /api/ai { proxy_pass http://python-ai:8000; }

# 靜態文件
location / { root /usr/share/nginx/html; }
```

## 3. 數據流向

### 3.1 基本操作流程

```
使用者 (前端)
    ↓
[ Nginx 反向代理 ]
    ├─ /api/assets → Golang API
    ├─ /api/tasks → Golang API  
    ├─ /api/chat → Python AI
    └─ / → 靜態文件
    ↓
後端服務
    ├─ Golang API (業務邏輯)
    └─ Python AI (分析與推理)
    ↓
[ 資料庫 ]
    ├─ 讀取/寫入
    └─ 返回結果
    ↓
迴應至使用者
```

### 3.2 AI Agent 工作流程

```
用戶訊息 ("飼料庫存有多少?")
    ↓
[ Python FastAPI ]
    ↓
[ LangChain Agent ]
    ├─ LLM 理解意圖 (GPT-4o / Gemini)
    └─ 判斷所需 Tool
    ↓
[ Tool Calling ]
    ├─ Tool: check_inventory
    └─ HTTP → Golang API (/api/assets)
    ↓
[ Golang API ]
    └─ 查詢資料庫 → 返回庫存數據
    ↓
[ Python 組織結果 ]
    └─ LLM 生成自然語言回應
    ↓
使用者收到: "高級飼料 A 有 150 公斤..."
```

## 4. 部署架構 (Docker Compose)

```yaml
services:
  mariadb:           # 資料庫 (Port 3306)
  golang-api:        # Golang API (Port 8080)
  python-ai:         # Python AI (Port 8000)
  nginx:             # 網頁伺服器 (Port 80)
```

**啟動流程:**
```
1. 啟動 mariadb - 等待資料庫就緒
2. 執行 init.sql - 初始化表結構與示例數據
3. 啟動 golang-api - 連接資料庫
4. 啟動 python-ai - 準備 AI 服務
5. 啟動 nginx - 配置反向代理
```

## 5. 擴展性考慮

### 5.1 水平擴展

```
可在多台伺服器上運行相同的容器:

[ Nginx 負載平衡器 ]
    ├─ Golang API 實例 1
    ├─ Golang API 實例 2
    ├─ Golang API 實例 3
    └─ Python AI 實例 1
        Python AI 實例 2
        ...
    ↓
[ 共享資料庫 ]
```

### 5.2 垂直扩展

- 增加容器資源 (CPU、內存)
- 優化 GORM 查詢
- 實施 Redis 快取層
- 資料庫索引優化

## 6. 監控與日誌

### 6.1 健康檢查端點

```
Golang API:   GET /health
Python API:   GET /health
Nginx:        GET /health
```

### 6.2 日誌管理

```
Golang:       stdout → Docker logs
Python:       stdout → Docker logs
Nginx:        /var/log/nginx/
MariaDB:      /var/log/mysql/
```

## 7. 安全性

### 7.1 當前實現 (MVP)

- ❌ 無認證 (全部開放)
- ✅ CORS 中間件
- ✅ 輸入驗證

### 7.2 建議實現 (生產環境)

- ✅ JWT Token 認證
- ✅ API Key 管理
- ✅ SQL 注入防護
- ✅ HTTPS/TLS
- ✅ WAF (Web Application Firewall)
- ✅ 基於角色的存取控制 (RBAC)

## 8. 性能指標

### 預期性能 (基於測試數據)

| 指標 | Golang | Python |
|------|--------|--------|
| 平均響應時間 | < 50ms | 200-500ms |
| QPS (吞吐量) | 5000+ | 100-200 |
| 內存占用 | 50-100MB | 200-500MB |
| 啟動時間 | < 2s | 5-10s |

### 優化建議

1. **資料庫層:**
   - 添加索引
   - 實施分區
   - 查詢優化

2. **應用層:**
   - 實施快取 (Redis)
   - 連接池優化
   - 非同步處理

3. **基礎設施:**
   - CDN 快取
   - 負載平衡
   - 自動擴展 (Kubernetes)

## 9. 故障恢復

### 9.1 服務隔離

每個微服務獨立運行，一個服務故障不會影響整個系統。

### 9.2 健康檢查

- Nginx 定時檢查後端服務
- Docker 自動重啟失敗容器
- 應用級監控告警

### 9.3 資料備份

```bash
# 定期備份 MariaDB
docker-compose exec mariadb mysqldump -u root -p \
  aquaculture_db > backup.sql
```

## 10. 技術決策理由

| 選擇 | 理由 |
|------|------|
| Golang API | 高性能、高併發、單一二進制部署 |
| Python AI | 豐富的 AI/ML 庫、快速開發 |
| MariaDB | 開源、社區支援、MySQL 兼容 |
| Docker | 標準化部署、環境一致性 |
| Nginx | 成熟、高性能、配置靈活 |

---

**澎湖數位老船長** 🐠 | 架構設計 v1.0 | 最後更新: 2026-04-07
