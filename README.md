# 澎湖數位老船長 - 養殖漁業管理系統 (MVP)

> 自主養殖經營中樞，讓您輕鬆記錄數據、觀察資訊、提升效率

## 🚀 快速開始

**立即啟動系統**: 請參考 [部署指南](DEPLOYMENT.md) 進行完整安裝和配置。

**快速測試**: 請參考 [快速啟動指南](QUICKSTART.md) 進行基本功能驗證。

## 🎯 專案概述

澎湖數位老船長是一個**微服務架構**的養殖漁業管理系統，專為澎湖地區養殖漁民打造。系統整合了：
- 🐠 **環境數據監測**（水溫、鹽度、溶氧量）
- 📦 **庫存管理**（飼料、網具、藥劑）  
- ✅ **任務分配**（日常工作、巡查紀錄）
- 🤖 **AI 助手**（自動生成報表、決策建議）
- 💬 **LINE Bot 整合**（即時查詢、語音命令）

## 🏗️ 技術架構

```
┌─────────────────────────────────────────────────────┐
│              LINE LIFF (前端介面)                   │
│          HTML/JavaScript/CSS 極簡 UI                │
└─────────────────┬───────────────────────────────────┘
                  │ HTTP
          ┌───────┴────────┐
          │                │
    ┌─────▼──────┐    ┌────▼──────────┐
    │ Nginx      │    │ Golang API    │ (Port 8080)
    │(Port 80)   │    │ - CRUD 業務   │
    └──────────┐ │    │ - 資料驗證    │
               │ │    │ - 高併發      │
               │ │    └────┬──────────┘
               │ │         │
               │ │    ┌────▼──────────┐
               │ │    │ Python FastAPI│ (Port 8000)
               │ │    │ - AI Agent    │
               │ │    │ - LLM 串接    │
               │ │    │ - Tool Calling│
               │ │    └───────┬───────┘
               │ │            │
               │ └────────────┼────────┐
               │              │        │
               └──────────────┼────────┤
                              │        │
                         ┌────▼──────┴──┐
                         │   MariaDB     │
                         │   (Port 3306) │
                         └───────────────┘
```

## 📋 技術棧 (Tech Stack)

| 層級 | 技術 | 用途 |
|------|------|------|
| **前端** | HTML5/JavaScript/CSS | LINE LIFF 介面 |
| **網頁** | Nginx | 靜態文件、反向代理 |
| **後端 A** | Golang + Gin + GORM | 高併發 API、CRUD、業務邏輯 |
| **後端 B** | Python 3.10+ FastAPI + LangChain | AI Agent、LLM 串接、Tool Calling |
| **資料庫** | MariaDB 10.11 | 數據持久化、ORM 支援 |
| **部署** | Docker + Docker Compose | 容器編排、一鍵啟動 |
| **AI 模型** | OpenAI GPT-4o / Gemini | 自然語言理解、決策建議 |

## 🚀 快速開始

### 前置要求
- Docker & Docker Compose
- Git
- Linux/macOS (或 Windows WSL2)

### 一鍵啟動

```bash
cd ~/project

# 1. 複製環境配置
cp .env.example .env

# 2. 啟動所有服務
docker-compose up -d

# 3. 檢查服務狀態
docker-compose ps

# 4. 查看日誌
docker-compose logs -f
```

### 存取服務

| 服務 | URL | 說明 |
|------|-----|------|
| **網頁首頁** | http://localhost | Nginx 靜態頁面 |
| **Golang API** | http://localhost:8080/api | 業務 API 入口 |
| **Python AI** | http://localhost:8000/api | AI 服務入口 |
| **MariaDB** | localhost:3306 | 資料庫連線 |

## 📁 項目結構

```
project/
├── .env                              # 環境配置
├── docker-compose.yml                # Docker 編排
├── README.md                         # 本文件
├── 
├── golang-api/                       # 🚀 Golang 微服務 (Port 8080)
│   ├── Dockerfile                    # Go 容器定義
│   ├── go.mod                        # Go 依賴管理
│   ├── main.go                       # 入口程式
│   ├── config/                       # 配置管理
│   ├── models/                       # 數據模型 (ORM)
│   ├── handlers/                     # HTTP 處理器
│   ├── routes/                       # 路由定義
│   ├── middleware/                   # 中間件 (CORS、驗證)
│   └── utils/                        # 工具函數
│
├── python-ai/                        # 🤖 Python AI 微服務 (Port 8000)
│   ├── Dockerfile                    # Python 容器定義
│   ├── requirements.txt               # 依賴列表
│   ├── main.py                       # FastAPI 應用 
│   ├── app/                          # FastAPI 應用模組
│   ├── agents/                       # AI Agent 邏輯
│   ├── tools/                        # Agent Tools (Function Calling)
│   ├── models/                       # 數據模型
│   └── utils/                        # 工具函數
│
├── nginx/                            # 🌐 Nginx 網頁服務
│   ├── Dockerfile                    # Nginx 容器定義
│   ├── conf/                         # Nginx 配置
│   │   └── nginx.conf               # 主配置文件
│   └── html/                         # 靜態文件
│       └── index.html               # 首頁
│
├── mysql-init/                       # 🗄️ 數據庫初始化
│   └── init.sql                      # SQL 初始化腳本
│
├── web/                              # 💬 LINE LIFF 前端
│   ├── index.html                    # LINE LIFF 首頁
│   ├── static/
│   │   ├── css/
│   │   ├── js/
│   │   └── images/
│   └── .liffrc.json                  # LIFF 配置
│
├── scripts/                          # 🔧 工具腳本
│   ├── start.sh                      # 啟動脚本
│   └── stop.sh                       # 停止腳本
│
└── docs/                             # 📖 文檔
    ├── API.md                        # API 文檔
    ├── ARCHITECTURE.md               # 架構設計文檔
    └── DEPLOYMENT.md                 # 部署指南
```

## 📡 核心 API 端點

### Golang API (Port 8080)

```
## 儀表板
GET  /api/dashboard              - 取得儀表板概覽

## 環境數據
GET  /api/environmental-data     - 查詢環境數據
POST /api/environmental-data     - 記錄環境數據

## 庫存管理  
GET  /api/assets                 - 查詢庫存清單
POST /api/assets                 - 新增物品
PUT  /api/assets/:id             - 更新庫存數量
DELETE /api/assets/:id           - 刪除物品

## 任務管理
GET  /api/tasks                  - 查詢任務列表
POST /api/tasks                  - 建立新任務
PUT  /api/tasks/:id              - 更新任務狀態
DELETE /api/tasks/:id            - 刪除任務
```

### Python AI (Port 8000)

```
## 聊天與對話
POST /api/chat                   - 發送訊息給 AI Agent

## Agent Tools (內部使用)
POST /api/agent/check-inventory  - 查詢庫存
POST /api/agent/create-task      - 建立任務  
POST /api/agent/generate-report  - 生成報表
```

## 🤖 AI Agent 工作流程

```
使用者 (LINE Bot)
    ↓
    發送自然語言請求 (如："飼料庫存有多少?")
    ↓
Python FastAPI (/api/chat)
    ↓
LangChain Agent
    ├─ 理解用戶意圖 (GPT-4o / Gemini)
    ├─ 選擇適當的 Tool
    └─ 執行 Tool Calling
    ↓
Golang API (HTTP 請求)
    └─ 查詢/更新資料庫
    ↓
Python AI 組織返回結果
    ↓
生成自然語言回應
    ↓
回傳至 LINE Bot
```

## 🗄️ 資料表結構

### Users (用戶)
```sql
CREATE TABLE users (
  id INT PRIMARY KEY AUTO_INCREMENT,
  name VARCHAR(100) NOT NULL,
  phone VARCHAR(20),
  language VARCHAR(10) DEFAULT 'zh-TW',
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### Environmental_Data (環境數據)
```sql
CREATE TABLE environmental_data (
  id INT PRIMARY KEY AUTO_INCREMENT,
  water_temperature DECIMAL(5,2),
  salinity DECIMAL(5,2),
  dissolved_oxygen DECIMAL(5,2),
  recorded_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### Assets (庫存)
```sql
CREATE TABLE assets (
  id INT PRIMARY KEY AUTO_INCREMENT,
  name VARCHAR(100) NOT NULL,
  quantity INT DEFAULT 0,
  unit VARCHAR(50),
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
```

### Tasks (任務)
```sql
CREATE TABLE tasks (
  id INT PRIMARY KEY AUTO_INCREMENT,
  description TEXT,
  assigned_to INT,
  status ENUM('pending', 'completed'),
  due_date DATE,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (assigned_to) REFERENCES users(id)
);
```

## ⚙️ 環境配置 (.env)

查看 [.env 範本](.env) 瞭解可用選項：

```bash
# 數據庫
DB_ROOT_PASSWORD=penghu2024aqua!
DB_NAME=aquaculture_db
DB_USER=appuser

# Golang
GO_PORT=8080

# Python  
PYTHON_PORT=8000

# AI 模型鑰匙
OPENAI_API_KEY=sk-xxx
GEMINI_API_KEY=xxx
```

## 🐛 故障排除

### 服務無法連線
```bash
# 檢查狀態
docker-compose ps

# 查看日誌
docker-compose logs golang-api
docker-compose logs python-ai
docker-compose logs mariadb
```

### 資料庫連線失敗
```bash
# 重啟資料庫
docker-compose restart mariadb

# 檢查數據庫初始化
docker-compose exec mariadb mysql -u root -p aquaculture_db < /init/init.sql
```

## 📚 更多文檔

- [API 文檔](docs/API.md) - 詳細 API 端點說明
- [架構設計](docs/ARCHITECTURE.md) - 系統架構深度分析
- [部署指南](docs/DEPLOYMENT.md) - 生產環境部署

## 👥 貢獻指南

歡迎提交 Issue 或 Pull Request！

## 📄 授權

MIT License - 詳見 [LICENSE](LICENSE)

---

**澎湖數位老船長** 🐠 | Made for Penghu Aquaculture with ❤️
# AquaCore
# AquaCore
