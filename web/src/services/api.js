import axios from 'axios'

const baseURL = process.env.VUE_APP_API_BASE_URL || '/api'

const apiClient = axios.create({
  baseURL,
  timeout: 15000,
  headers: {
    'Content-Type': 'application/json'
  }
})

// ==================== 認證相關 ====================

export function registerUser (payload) {
  // 將前端表單數據轉換為後端期望的格式
  const userData = {
    name: `${payload.firstName} ${payload.lastName}`,
    email: payload.email,
    phone: payload.phone,
    preferred_language: 'zh-TW',
    role: 'operator',
    status: 'active'
  }
  return apiClient.post('/users', userData)
}

export function loginUser (payload) {
  return apiClient.post('/auth/login', payload)
}

// ==================== 環境數據相關 ====================

export function getEnvironmentalData () {
  return apiClient.get('/environmental-data')
}

export function createEnvironmentalData (payload) {
  return apiClient.post('/environmental-data', payload)
}

// ==================== 資產/庫存相關 ====================

export function getAssets () {
  return apiClient.get('/assets')
}

export function createAsset (payload) {
  return apiClient.post('/assets', payload)
}

// ==================== 任務相關 ====================

export function getTasks () {
  return apiClient.get('/tasks')
}

export function createTask (payload) {
  return apiClient.post('/tasks', payload)
}

// ==================== 用戶相關 ====================

export function getUsers () {
  return apiClient.get('/users')
}

export default apiClient
