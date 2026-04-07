import { defineStore } from 'pinia'

export const useAppStore = defineStore('app', {
  state: () => ({
    // 应用配置
    config: {
      apiBaseUrl: process.env.VUE_APP_API_BASE_URL || '/api',
      theme: 'light',
      language: 'zh-TW'
    },

    // UI 状态
    ui: {
      sidebarCollapsed: false,
      loading: false,
      notifications: []
    },

    // 缓存数据
    cache: {
      dashboard: null,
      environment: null,
      assets: null,
      tasks: null
    },

    // 错误状态
    error: null,

    // 最后更新时间
    lastUpdated: {
      dashboard: null,
      environment: null,
      assets: null,
      tasks: null
    }
  }),

  getters: {
    isLoading: (state) => state.ui.loading,

    hasError: (state) => !!state.error,

    unreadNotifications: (state) => state.ui.notifications.filter(n => !n.read).length,

    cacheValid: (state) => (key) => {
      const lastUpdated = state.lastUpdated[key]
      if (!lastUpdated) return false

      const now = new Date()
      const cacheTime = new Date(lastUpdated)
      const diffMinutes = (now - cacheTime) / (1000 * 60)

      // 缓存5分钟有效
      return diffMinutes < 5
    }
  },

  actions: {
    // UI 状态管理
    setLoading(loading) {
      this.ui.loading = loading
    },

    setSidebarCollapsed(collapsed) {
      this.ui.sidebarCollapsed = collapsed
    },

    // 错误处理
    setError(error) {
      this.error = error
      console.error('应用错误:', error)
    },

    clearError() {
      this.error = null
    },

    // 通知管理
    addNotification(notification) {
      const newNotification = {
        id: Date.now(),
        read: false,
        timestamp: new Date().toISOString(),
        ...notification
      }

      this.ui.notifications.unshift(newNotification)

      // 最多保留50条通知
      if (this.ui.notifications.length > 50) {
        this.ui.notifications = this.ui.notifications.slice(0, 50)
      }
    },

    markNotificationAsRead(id) {
      const notification = this.ui.notifications.find(n => n.id === id)
      if (notification) {
        notification.read = true
      }
    },

    clearNotifications() {
      this.ui.notifications = []
    },

    // 缓存管理
    setCache(key, data) {
      this.cache[key] = data
      this.lastUpdated[key] = new Date().toISOString()
    },

    getCache(key) {
      if (this.cacheValid(key)) {
        return this.cache[key]
      }
      return null
    },

    clearCache(key = null) {
      if (key) {
        this.cache[key] = null
        this.lastUpdated[key] = null
      } else {
        // 清除所有缓存
        Object.keys(this.cache).forEach(k => {
          this.cache[k] = null
          this.lastUpdated[k] = null
        })
      }
    },

    // 配置管理
    updateConfig(newConfig) {
      this.config = { ...this.config, ...newConfig }
      // 可以在这里持久化配置到localStorage
      localStorage.setItem('app_config', JSON.stringify(this.config))
    },

    loadConfig() {
      const savedConfig = localStorage.getItem('app_config')
      if (savedConfig) {
        try {
          const parsedConfig = JSON.parse(savedConfig)
          this.config = { ...this.config, ...parsedConfig }
        } catch (error) {
          console.error('加载配置失败:', error)
        }
      }
    },

    // 主题切换
    setTheme(theme) {
      this.config.theme = theme
      document.documentElement.setAttribute('data-theme', theme)

      // 保存到localStorage
      localStorage.setItem('theme', theme)
    },

    loadTheme() {
      const savedTheme = localStorage.getItem('theme') || 'light'
      this.setTheme(savedTheme)
    },

    // 语言设置
    setLanguage(language) {
      this.config.language = language
      // 这里可以集成i18n库
      localStorage.setItem('language', language)
    },

    // 初始化应用
    async initialize() {
      this.loadConfig()
      this.loadTheme()

      // 设置全局错误处理
      window.addEventListener('unhandledrejection', (event) => {
        this.setError({
          type: 'promise_rejection',
          message: event.reason?.message || '未处理的Promise錯誤',
          details: event.reason
        })
      })

      window.addEventListener('error', (event) => {
        this.setError({
          type: 'javascript_error',
          message: event.message,
          details: {
            filename: event.filename,
            lineno: event.lineno,
            colno: event.colno
          }
        })
      })
    }
  }
})