<template>
  <div id="app">
    <!-- 加载状态 -->
    <div v-if="appStore.isLoading" class="loading-overlay">
      <div class="spinner-border text-primary" role="status">
        <span class="visually-hidden">載入中...</span>
      </div>
    </div>

    <!-- 错误提示 -->
    <div v-if="appStore.hasError" class="alert alert-danger alert-dismissible fade show position-fixed" style="z-index: 9999; top: 20px; right: 20px; min-width: 300px;">
      <strong>錯誤：</strong> {{ appStore.error.message || appStore.error }}
      <button type="button" class="btn-close" @click="appStore.clearError()"></button>
    </div>

    <!-- 导航栏 -->
    <nav class="navbar navbar-expand-lg navbar-dark bg-primary" v-if="!isHomePage">
      <div class="container">
        <router-link class="navbar-brand" to="/dashboard">
          🐠 澎湖數位老船長
        </router-link>
        <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNav">
          <span class="navbar-toggler-icon"></span>
        </button>
        <div class="collapse navbar-collapse" id="navbarNav">
          <ul class="navbar-nav me-auto">
            <li class="nav-item">
              <router-link class="nav-link" to="/dashboard">儀表板</router-link>
            </li>
            <li class="nav-item">
              <router-link class="nav-link" to="/environment">環境監測</router-link>
            </li>
            <li class="nav-item">
              <router-link class="nav-link" to="/assets">庫存管理</router-link>
            </li>
            <li class="nav-item">
              <router-link class="nav-link" to="/tasks">任務管理</router-link>
            </li>
            <li class="nav-item">
              <router-link class="nav-link" to="/chat">AI 助手</router-link>
            </li>
          </ul>
          <ul class="navbar-nav">
            <li class="nav-item dropdown">
              <a class="nav-link dropdown-toggle" href="#" role="button" data-bs-toggle="dropdown">
                <i class="bi bi-bell me-1"></i>
                <span v-if="appStore.unreadNotifications > 0" class="badge bg-danger">{{ appStore.unreadNotifications }}</span>
              </a>
              <ul class="dropdown-menu">
                <li v-if="appStore.ui.notifications.length === 0">
                  <a class="dropdown-item" href="#">沒有新通知</a>
                </li>
                <li v-for="notification in appStore.ui.notifications.slice(0, 5)" :key="notification.id">
                  <a class="dropdown-item" href="#" @click="markAsRead(notification.id)">
                    <small class="text-muted">{{ formatTime(notification.timestamp) }}</small><br>
                    {{ notification.message }}
                  </a>
                </li>
                <li v-if="appStore.ui.notifications.length > 5">
                  <hr class="dropdown-divider">
                  <li><a class="dropdown-item text-center" href="#" @click="appStore.clearNotifications()">清除所有</a></li>
                </li>
              </ul>
            </li>
            <li class="nav-item dropdown">
              <a class="nav-link dropdown-toggle" href="#" role="button" data-bs-toggle="dropdown">
                <i class="bi bi-person-circle me-1"></i>
                {{ authStore.fullName || '用戶' }}
              </a>
              <ul class="dropdown-menu">
                <li><a class="dropdown-item" href="#"><i class="bi bi-person me-2"></i>個人資料</a></li>
                <li><a class="dropdown-item" href="#"><i class="bi bi-gear me-2"></i>設定</a></li>
                <li><hr class="dropdown-divider"></li>
                <li><a class="dropdown-item" href="#" @click="handleLogout"><i class="bi bi-box-arrow-right me-2"></i>登出</a></li>
              </ul>
            </li>
          </ul>
        </div>
      </div>
    </nav>

    <!-- 主内容区域 -->
    <main :class="mainClass">
      <router-view />
    </main>
  </div>
</template>

<script>
import { useAuthStore } from './store/auth'
import { useAppStore } from './store/app'

export default {
  name: 'App',
  setup() {
    const authStore = useAuthStore()
    const appStore = useAppStore()

    return {
      authStore,
      appStore
    }
  },
  computed: {
    isHomePage() {
      return this.$route.name === 'Home'
    },
    mainClass() {
      return this.isHomePage ? '' : 'container mt-4'
    }
  },
  methods: {
    async handleLogout() {
      await this.authStore.logout()
      this.$router.push('/')
    },
    markAsRead(id) {
      this.appStore.markNotificationAsRead(id)
    },
    formatTime(timestamp) {
      const date = new Date(timestamp)
      return date.toLocaleString('zh-TW', {
        month: 'short',
        day: 'numeric',
        hour: '2-digit',
        minute: '2-digit'
      })
    }
  }
}
</script>

<style>
/* 全局样式 */
body {
  font-family: 'Microsoft JhengHei', 'PingFang TC', sans-serif;
  background-color: #f8f9fa;
}

/* 加载覆盖层 */
.loading-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(255, 255, 255, 0.8);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 9999;
}

/* 导航栏样式 */
.navbar-brand {
  font-weight: bold;
  font-size: 1.25rem;
}

.nav-link {
  font-weight: 500;
  transition: color 0.3s ease;
}

.nav-link:hover {
  color: rgba(255, 255, 255, 0.8) !important;
}

.dropdown-menu {
  border: none;
  box-shadow: 0 0.5rem 1rem rgba(0, 0, 0, 0.15);
}

/* 响应式设计 */
@media (max-width: 768px) {
  .navbar-brand {
    font-size: 1.1rem;
  }

  .nav-link {
    padding: 0.5rem 1rem;
  }
}
</style>