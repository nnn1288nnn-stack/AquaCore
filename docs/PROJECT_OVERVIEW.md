# 專案總覽

此文檔整理 `project/` 目錄下的主要檔案與運行方式，讓專案結構更清楚。

## 目前整理結果

- `docker-compose.yml`：主 Docker 編排檔，啟動 MariaDB、Golang API、Python AI、Vue 前端、Nginx。
- `nginx/conf/nginx.conf`：目前實際使用的 Nginx 配置，負責靜態頁面服務與 `/api/` 代理。
- `web/index.html`：Vue 前端入口頁面。
- `web/src/`：Vue 源碼目錄，包含 `views/`、`router/`、`store/`。
- `web/src/services/api.js`：新增的前端 API 服務層，用於呼叫後端 `/api` 接口。
- `web/src/views/Environment.vue`：環境監控頁面，已新增資料輸入與資料庫寫入功能。
- `mysql-init/init.sql`：MariaDB 初始化資料表與樣本資料。
- `web/nginx.conf.legacy`：舊的重複配置，已移至備用，避免與主配置混淆。
- `web/legacy-nginx-index.html`：舊首頁備份。

## 前後端 API 連接

目前前端已經使用 `process.env.VUE_APP_API_BASE_URL` 指向 `/api`，因此：

- `GET /api/environmental-data`：讀取環境資料
- `POST /api/environmental-data`：寫入環境資料到 MariaDB

此配置在 Nginx 代理下會正確呼叫 Golang API。

## 資料輸入功能

已在 `web/src/views/Environment.vue` 中加入：

- 環境資料輸入表單
- API 呼叫 `createEnvironmentalData()`
- 送出後重新讀取歷史資料

這表示目前已支援「從前端要求資料庫輸入」功能。

## 天氣資料擴充建議

目前還未接入氣象局資料，但可按以下方式擴展：

1. 在 Golang API 新增 `/api/weather` 端點。
2. 於後端呼叫中央氣象局或氣象局開放資料 API。
3. 將天氣資訊與養殖環境資料結合顯示在前端。

## 啟動方式

```bash
cd /home/ouo/project
docker compose up -d
```

訪問：

- 前端：`http://localhost` 或 `http://192.168.50.75`
- Golang API：`http://localhost:8080`
- Python AI：`http://localhost:8000`


## 重要說明

- `nginx/conf/nginx.conf` 是主配置文件
- `web/nginx.conf.legacy` 和 `web/legacy-nginx-index.html` 為備用或歷史檔案，不會由 Docker Compose 直接使用
