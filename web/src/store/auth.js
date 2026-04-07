import { defineStore } from 'pinia'
import axios from 'axios'

export const useAuthStore = defineStore('auth', {
  state: () => ({
    user: null,
    token: localStorage.getItem('auth_token') || null,
    isAuthenticated: !!localStorage.getItem('auth_token'),
    loading: false,
    error: null
  }),

  getters: {
    fullName: (state) => {
      if (!state.user) return ''
      return `${state.user.firstName} ${state.user.lastName}`
    },

    userRole: (state) => state.user?.role || 'user',

    isAdmin: (state) => state.user?.role === 'admin'
  },

  actions: {
    async login(credentials) {
      this.loading = true
      this.error = null

      try {
        const response = await axios.post('/api/auth/login', credentials)

        if (response.data.success) {
          const { token, user } = response.data.data

          this.token = token
          this.user = user
          this.isAuthenticated = true

          // 存储到localStorage
          localStorage.setItem('auth_token', token)

          // 设置axios默认header
          axios.defaults.headers.common['Authorization'] = `Bearer ${token}`

          // 记录登录日志
          await this.logAction('login', `用戶 ${user.email} 登入系統`)

          return { success: true }
        } else {
          throw new Error(response.data.message || '登入失敗')
        }
      } catch (error) {
        this.error = error.response?.data?.message || error.message || '登入失敗'
        return { success: false, error: this.error }
      } finally {
        this.loading = false
      }
    },

    async register(userData) {
      this.loading = true
      this.error = null

      try {
        const response = await axios.post('/api/auth/register', userData)

        if (response.data.success) {
          // 记录注册日志
          await this.logAction('registration', `新用戶註冊: ${userData.email}`)

          return { success: true, message: response.data.message }
        } else {
          throw new Error(response.data.message || '註冊失敗')
        }
      } catch (error) {
        this.error = error.response?.data?.message || error.message || '註冊失敗'
        return { success: false, error: this.error }
      } finally {
        this.loading = false
      }
    },

    async logout() {
      try {
        // 记录登出日志
        await this.logAction('logout', `用戶 ${this.user?.email} 登出系統`)

        // 调用登出API（如果有）
        await axios.post('/api/auth/logout')
      } catch (error) {
        console.error('登出API調用失敗:', error)
      }

      // 清除本地状态
      this.user = null
      this.token = null
      this.isAuthenticated = false

      // 清除localStorage
      localStorage.removeItem('auth_token')

      // 清除axios默认header
      delete axios.defaults.headers.common['Authorization']
    },

    async checkAuth() {
      if (!this.token) {
        this.isAuthenticated = false
        return
      }

      try {
        const response = await axios.get('/api/auth/me')

        if (response.data.success) {
          this.user = response.data.data.user
          this.isAuthenticated = true
        } else {
          // Token无效，清除状态
          this.logout()
        }
      } catch (error) {
        // Token无效，清除状态
        this.logout()
      }
    },

    async refreshToken() {
      try {
        const response = await axios.post('/api/auth/refresh')

        if (response.data.success) {
          const { token } = response.data.data

          this.token = token
          localStorage.setItem('auth_token', token)
          axios.defaults.headers.common['Authorization'] = `Bearer ${token}`

          return { success: true }
        }
      } catch (error) {
        // 刷新失败，登出
        this.logout()
        return { success: false, error: error.message }
      }
    },

    async logAction(action, message) {
      try {
        await axios.post('/api/logs', {
          type: action,
          message: message,
          level: 'info',
          user: this.user?.email || 'unknown',
          timestamp: new Date().toISOString(),
          service: 'frontend'
        })
      } catch (error) {
        console.error('記錄操作日誌失敗:', error)
      }
    }
  }
})