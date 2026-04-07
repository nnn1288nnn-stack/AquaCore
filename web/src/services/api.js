import axios from 'axios'

const baseURL = process.env.VUE_APP_API_BASE_URL || '/api'

const apiClient = axios.create({
  baseURL,
  timeout: 15000,
  headers: {
    'Content-Type': 'application/json'
  }
})

export function getEnvironmentalData () {
  return apiClient.get('/environmental-data')
}

export function createEnvironmentalData (payload) {
  return apiClient.post('/environmental-data', payload)
}

export function getAssets () {
  return apiClient.get('/assets')
}

export function getTasks () {
  return apiClient.get('/tasks')
}

export default apiClient
