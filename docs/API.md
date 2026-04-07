# API 文檔 - 澎湖數位老船長

## 目錄

1. [基本信息](#基本信息)
2. [Golang API](#golang-api)
3. [Python AI API](#python-ai-api)
4. [認證與授權](#認證與授權)
5. [錯誤處理](#錯誤處理)
6. [使用範例](#使用範例)

---

## 基本信息

### API 基礎 URL

| 服務 | URL | 端口 |
|------|-----|------|
| **Golang API** | `http://localhost:8080` | 8080 |
| **Python AI** | `http://localhost:8000` | 8000 |
| **Nginx 代理** | `http://localhost` | 80 |

### 通用請求頭

```http
Content-Type: application/json
X-Request-ID: (optional) 唯一請求識別符
```

### 通用回應格式

**成功 (2xx)**
```json
{
  "status": "success",
  "code": 200,
  "data": { /* 實際數據 */ },
  "timestamp": "2026-04-07T14:30:00Z"
}
```

**失敗 (4xx/5xx)**
```json
{
  "status": "error",
  "code": 400,
  "message": "錯誤信息",
  "details": "詳細說明",
  "timestamp": "2026-04-07T14:30:00Z"
}
```

---

## Golang API

### 健康檢查

#### GET /health

檢查 API 服務狀態

**回應 (200 OK)**
```json
{
  "status": "healthy",
  "service": "golang-api",
  "version": "1.0.0"
}
```

---

### 儀表板

#### GET /api/dashboard

獲取儀表板概覽數據（水溫、鹽度、溶氧量、任務計數等）

**回應 (200 OK)**
```json
{
  "status": "success",
  "data": {
    "water_temperature": 24.5,
    "salinity": 30.2,
    "dissolved_oxygen": 7.5,
    "pending_tasks": 5,
    "low_stock_items": 2
  }
}
```

---

### 環境數據

#### GET /api/environmental-data

查詢環境數據列表

**查詢參數**
- `limit`: 返回記錄數 (預設 100)
- `offset`: 開始位置 (預設 0)
- `order`: 排序方式 (predefault: -recorded_at)

**回應 (200 OK)**
```json
{
  "status": "success",
  "data": [
    {
      "id": 1,
      "water_temperature": 24.5,
      "salinity": 30.2,
      "dissolved_oxygen": 7.5,
      "ph_level": 7.8,
      "ammonia": 0.02,
      "recorded_at": "2026-04-07T14:30:00Z",
      "notes": "正常運作"
    }
  ],
  "total": 150,
  "limit": 100,
  "offset": 0
}
```

#### POST /api/environmental-data

記錄新的環境數據

**請求體**
```json
{
  "water_temperature": 24.5,
  "salinity": 30.2,
  "dissolved_oxygen": 7.5,
  "ph_level": 7.8,
  "ammonia": 0.02,
  "notes": "正常運作"
}
```

**回應 (201 Created)**
```json
{
  "status": "success",
  "data": {
    "id": 1,
    "water_temperature": 24.5,
    "recorded_at": "2026-04-07T14:30:00Z"
  }
}
```

---

### 庫存管理 (Assets)

#### GET /api/assets

查詢所有資產/庫存

**查詢參數**
- `category`: 分類篩選 (飼料/網具/藥劑/工具)
- `low_stock`: 是否只顯示低庫存 (true/false)

**回應 (200 OK)**
```json
{
  "status": "success",
  "data": [
    {
      "id": 1,
      "name": "高級飼料 A",
      "category": "飼料",
      "quantity": 150,
      "unit": "公斤",
      "reorder_level": 50,
      "unit_cost": 45.00,
      "supplier": "飼料廠商 ABC",
      "status": "充足",
      "updated_at": "2026-04-07T14:30:00Z"
    }
  ]
}
```

#### POST /api/assets

建立新資產

**請求體**
```json
{
  "name": "高級飼料 A",
  "category": "飼料",
  "quantity": 150,
  "unit": "公斤",
  "reorder_level": 50,
  "unit_cost": 45.00,
  "supplier": "飼料廠商 ABC"
}
```

**回應 (201 Created)**
```json
{
  "status": "success",
  "data": {
    "id": 1,
    "name": "高級飼料 A"
  }
}
```

#### PUT /api/assets/:id

更新資產信息或庫存數量

**請求體 (至少需要一個欄位)**
```json
{
  "quantity": 100,
  "reorder_level": 40,
  "supplier": "新供應商"
}
```

**回應 (200 OK)**
```json
{
  "status": "success",
  "data": {
    "id": 1,
    "quantity": 100
  }
}
```

#### DELETE /api/assets/:id

刪除資產

**回應 (204 No Content)**
```
(空)
```

---

### 任務管理

#### GET /api/tasks

查詢所有任務

**查詢參數**
- `status`: 篩選狀態 (pending/in-progress/completed/cancelled)
- `assigned_to`: 篩選負責人 ID
- `priority`: 篩選優先級 (low/medium/high/urgent)

**回應 (200 OK)**
```json
{
  "status": "success",
  "data": [
    {
      "id": 1,
      "title": "檢查水質",
      "description": "每日水質參數檢測",
      "assigned_to": 2,
      "assigned_to_name": "李小姐",
      "status": "completed",
      "priority": "high",
      "due_date": "2026-04-07",
      "created_at": "2026-04-06T10:00:00Z",
      "completed_at": "2026-04-07T14:30:00Z"
    }
  ]
}
```

#### POST /api/tasks

建立新任務

**請求體**
```json
{
  "title": "檢查水質",
  "description": "每日水質參數檢測",
  "assigned_to": 2,
  "priority": "high",
  "due_date": "2026-04-08"
}
```

**回應 (201 Created)**
```json
{
  "status": "success",
  "data": {
    "id": 1,
    "title": "檢查水質"
  }
}
```

#### PUT /api/tasks/:id

更新任務

**請求體 (至少需要一個欄位)**
```json
{
  "status": "in-progress",
  "priority": "urgent",
  "assigned_to": 3
}
```

**回應 (200 OK)**
```json
{
  "status": "success",
  "data": {
    "id": 1,
    "status": "in-progress"
  }
}
```

#### DELETE /api/tasks/:id

刪除任務

**回應 (204 No Content)**
```
(空)
```

---

## Python AI API

### 健康檢查

#### GET /health

**回應 (200 OK)**
```json
{
  "status": "healthy",
  "service": "python-ai",
  "version": "1.0.0"
}
```

### 聊天 / AI 助手

#### POST /api/chat

發送消息給 AI Agent（澎湖數位老船長）

**請求體**
```json
{
  "message": "飼料庫存有多少?",
  "language": "zh-TW"
}
```

**回應 (200 OK)**
```json
{
  "status": "success",
  "reply": "根據系統記錄，高級飼料 A 目前有 150 公斤，已達到良好庫存水平。",
  "action": "check_inventory"
}
```

**常見 AI 指令範例**

| 指令 | 說明 |
|------|------|
| "飼料庫存有多少？" | 查詢庫存 |
| "創建一個新任務" | 建立任務 |
| "生成今日報表" | 生成日報 |
| "水溫是多少？" | 查詢環境數據 |

---

## 認證與授權

當前版本（MVP）**不需要認證**。在生產環境中，建議實現：

- JWT Token 認證
- API Key 管理
- 基於角色的訪問控制 (RBAC)

---

## 錯誤處理

### HTTP 狀態碼

| 狀態碼 | 說明 | 
|--------|------|
| 200 | OK - 請求成功 |
| 201 | Created - 資源已建立 |
| 204 | No Content - 刪除成功 |
| 400 | Bad Request - 請求參數錯誤 |
| 404 | Not Found - 資源不存在 |
| 500 | Internal Server Error - 伺服器錯誤 |

### 錯誤回應示例

```json
{
  "status": "error",
  "code": 400,
  "message": "無效的時間格式",
  "details": "due_date 應為 YYYY-MM-DD 格式",
  "request_id": "req_123abc",
  "timestamp": "2026-04-07T14:30:00Z"
}
```

---

## 使用範例

### 使用 cURL

**查詢環境數據**
```bash
curl -X GET http://localhost:8080/api/environmental-data \
  -H "Content-Type: application/json"
```

**記錄新的環境數據**
```bash
curl -X POST http://localhost:8080/api/environmental-data \
  -H "Content-Type: application/json" \
  -d '{
    "water_temperature": 24.5,
    "salinity": 30.2,
    "dissolved_oxygen": 7.5,
    "ph_level": 7.8
  }'
```

**與 AI 助手對話**
```bash
curl -X POST http://localhost:8000/api/chat \
  -H "Content-Type: application/json" \
  -d '{
    "message": "飼料庫存有多少?",
    "language": "zh-TW"
  }'
```

### 使用 Python requests

```python
import requests

# 查詢任務
response = requests.get('http://localhost:8080/api/tasks')
tasks = response.json()
print(tasks)

# 建立新任務
new_task = {
    'title': '檢查水質',
    'assigned_to': 2,
    'priority': 'high',
    'due_date': '2026-04-08'
}
response = requests.post('http://localhost:8080/api/tasks', json=new_task)
result = response.json()
print(result)
```

### 使用 JavaScript/Fetch

```javascript
// 查詢資產
fetch('http://localhost:8080/api/assets')
  .then(res => res.json())
  .then(data => console.log(data));

// 與 AI 對話
fetch('http://localhost:8000/api/chat', {
  method: 'POST',
  headers: { 'Content-Type': 'application/json' },
  body: JSON.stringify({
    message: '飼料庫存有多少?',
    language: 'zh-TW'
  })
})
  .then(res => res.json())
  .then(data => console.log(data.reply));
```

---

## 版本歷史

| 版本 | 日期 | 說明 |
|------|------|------|
| 1.0.0 | 2026-04-07 | MVP 初始版本 |

---

**澎湖數位老船長** 🐠 | 最後更新: 2026-04-07
