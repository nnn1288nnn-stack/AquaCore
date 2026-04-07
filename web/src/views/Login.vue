<template>
  <div class="row justify-content-center">
    <div class="col-md-6 col-lg-4">
      <div class="card shadow">
        <div class="card-body p-5">
          <div class="text-center mb-4">
            <h2 class="card-title">🐠 登入系統</h2>
            <p class="text-muted">澎湖數位老船長</p>
          </div>

          <form @submit.prevent="handleLogin">
            <div class="mb-3">
              <label for="username" class="form-label">用戶名稱</label>
              <input
                type="text"
                class="form-control"
                id="username"
                v-model="credentials.username"
                required
              >
            </div>

            <div class="mb-3">
              <label for="password" class="form-label">密碼</label>
              <input
                type="password"
                class="form-control"
                id="password"
                v-model="credentials.password"
                required
              >
            </div>

            <div class="d-grid">
              <button
                type="submit"
                class="btn btn-primary btn-lg"
                :disabled="loading"
              >
                <span v-if="loading" class="spinner-border spinner-border-sm me-2"></span>
                {{ loading ? '登入中...' : '登入' }}
              </button>
            </div>
          </form>

          <div v-if="error" class="alert alert-danger mt-3">
            {{ error }}
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'Login',
  data() {
    return {
      credentials: {
        username: '',
        password: ''
      },
      loading: false,
      error: null
    }
  },
  methods: {
    async handleLogin() {
      this.loading = true
      this.error = null

      try {
        // 这里应该调用实际的登录 API
        // 暂时模拟登录
        if (this.credentials.username && this.credentials.password) {
          // 模拟 API 调用延迟
          await new Promise(resolve => setTimeout(resolve, 1000))

          // 存储认证令牌
          localStorage.setItem('auth_token', 'fake_token_' + Date.now())

          // 记录登录日志到后端
          await this.logLogin()

          // 跳转到仪表板
          this.$router.push('/dashboard')
        } else {
          throw new Error('请输入用户名和密码')
        }
      } catch (error) {
        this.error = error.message || '登录失败'
      } finally {
        this.loading = false
      }
    },

    async logLogin() {
      try {
        await axios.post('/api/logs', {
          type: 'login',
          username: this.credentials.username,
          timestamp: new Date().toISOString(),
          ip: window.location.hostname
        })
      } catch (error) {
        console.error('记录登录日志失败:', error)
      }
    }
  }
}
</script>

<style scoped>
.card {
  border: none;
  border-radius: 15px;
}

.card-title {
  color: #007bff;
  font-weight: bold;
}

.btn-primary {
  background: linear-gradient(45deg, #007bff, #0056b3);
  border: none;
  border-radius: 10px;
}

.btn-primary:hover {
  background: linear-gradient(45deg, #0056b3, #004085);
}
</style>