<template>
  <div class="register-page">
    <div class="container">
      <div class="row justify-content-center">
        <div class="col-md-8 col-lg-6">
          <div class="register-card">
            <div class="text-center mb-4">
              <router-link to="/" class="register-logo">
                <i class="bi bi-water"></i>
                <span>澎湖數位老船長</span>
              </router-link>
              <h2 class="register-title">創建帳號</h2>
              <p class="text-muted">加入智慧養殖管理系統</p>
            </div>

            <form @submit.prevent="handleRegister">
              <div class="row">
                <div class="col-md-6 mb-3">
                  <label for="firstName" class="form-label">姓氏</label>
                  <input
                    type="text"
                    class="form-control"
                    id="firstName"
                    v-model="form.firstName"
                    required
                  >
                </div>
                <div class="col-md-6 mb-3">
                  <label for="lastName" class="form-label">名字</label>
                  <input
                    type="text"
                    class="form-control"
                    id="lastName"
                    v-model="form.lastName"
                    required
                  >
                </div>
              </div>

              <div class="mb-3">
                <label for="email" class="form-label">電子郵件</label>
                <input
                  type="email"
                  class="form-control"
                  id="email"
                  v-model="form.email"
                  required
                >
                <div class="form-text">我們不會與他人分享您的電子郵件。</div>
              </div>

              <div class="mb-3">
                <label for="phone" class="form-label">手機號碼</label>
                <input
                  type="tel"
                  class="form-control"
                  id="phone"
                  v-model="form.phone"
                  placeholder="0912345678"
                >
              </div>

              <div class="mb-3">
                <label for="farmName" class="form-label">養殖場名稱</label>
                <input
                  type="text"
                  class="form-control"
                  id="farmName"
                  v-model="form.farmName"
                  placeholder="例如：澎湖海鮮養殖場"
                >
              </div>

              <div class="mb-3">
                <label for="farmType" class="form-label">養殖類型</label>
                <select class="form-select" id="farmType" v-model="form.farmType">
                  <option value="">請選擇</option>
                  <option value="fish">魚類養殖</option>
                  <option value="shellfish">貝類養殖</option>
                  <option value="seaweed">海藻養殖</option>
                  <option value="mixed">綜合養殖</option>
                </select>
              </div>

              <div class="mb-3">
                <label for="password" class="form-label">密碼</label>
                <input
                  type="password"
                  class="form-control"
                  id="password"
                  v-model="form.password"
                  required
                  minlength="8"
                >
                <div class="form-text">密碼至少需要8個字符。</div>
              </div>

              <div class="mb-3">
                <label for="confirmPassword" class="form-label">確認密碼</label>
                <input
                  type="password"
                  class="form-control"
                  id="confirmPassword"
                  v-model="form.confirmPassword"
                  required
                >
              </div>

              <div class="mb-3 form-check">
                <input
                  type="checkbox"
                  class="form-check-input"
                  id="agreeTerms"
                  v-model="form.agreeTerms"
                  required
                >
                <label class="form-check-label" for="agreeTerms">
                  我同意<a href="#" class="text-primary">服務條款</a>和<a href="#" class="text-primary">隱私政策</a>
                </label>
              </div>

              <div class="d-grid mb-3">
                <button
                  type="submit"
                  class="btn btn-primary btn-lg"
                  :disabled="loading || !isFormValid"
                >
                  <span v-if="loading" class="spinner-border spinner-border-sm me-2"></span>
                  {{ loading ? '註冊中...' : '創建帳號' }}
                </button>
              </div>
            </form>

            <div class="text-center">
              <p class="mb-0">
                已經有帳號了？
                <router-link to="/login" class="text-primary">立即登入</router-link>
              </p>
            </div>

            <div v-if="error" class="alert alert-danger mt-3">
              {{ error }}
            </div>

            <div v-if="success" class="alert alert-success mt-3">
              <i class="bi bi-check-circle me-2"></i>
              註冊成功！請檢查您的電子郵件以驗證帳號。
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'Register',
  data() {
    return {
      form: {
        firstName: '',
        lastName: '',
        email: '',
        phone: '',
        farmName: '',
        farmType: '',
        password: '',
        confirmPassword: '',
        agreeTerms: false
      },
      loading: false,
      error: null,
      success: false
    }
  },
  computed: {
    isFormValid() {
      return (
        this.form.firstName &&
        this.form.lastName &&
        this.form.email &&
        this.form.password.length >= 8 &&
        this.form.password === this.form.confirmPassword &&
        this.form.agreeTerms
      )
    }
  },
  methods: {
    async handleRegister() {
      if (!this.isFormValid) return

      this.loading = true
      this.error = null
      this.success = false

      try {
        // 准备注册数据
        const registerData = {
          firstName: this.form.firstName,
          lastName: this.form.lastName,
          email: this.form.email,
          phone: this.form.phone,
          farmName: this.form.farmName,
          farmType: this.form.farmType,
          password: this.form.password,
          timestamp: new Date().toISOString()
        }

        // 调用注册 API
        const response = await axios.post('/api/auth/register', registerData)

        if (response.data.success) {
          this.success = true

          // 记录注册日志
          await this.logRegistration(registerData)

          // 3秒后跳转到登录页面
          setTimeout(() => {
            this.$router.push('/login')
          }, 3000)
        } else {
          throw new Error(response.data.message || '註冊失敗')
        }
      } catch (error) {
        this.error = error.response?.data?.message || error.message || '註冊失敗，請稍後再試'
      } finally {
        this.loading = false
      }
    },

    async logRegistration(userData) {
      try {
        await axios.post('/api/logs', {
          type: 'registration',
          message: `新用戶註冊: ${userData.email}`,
          level: 'info',
          user: userData.email,
          service: 'frontend'
        })
      } catch (error) {
        console.error('記錄註冊日誌失敗:', error)
      }
    }
  }
}
</script>

<style scoped>
.register-page {
  min-height: 100vh;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  padding: 60px 0;
}

.register-card {
  background: white;
  border-radius: 15px;
  box-shadow: 0 20px 40px rgba(0,0,0,0.1);
  padding: 3rem;
}

.register-logo {
  display: inline-flex;
  align-items: center;
  font-size: 1.5rem;
  font-weight: bold;
  color: #007bff;
  text-decoration: none;
  margin-bottom: 1rem;
}

.register-logo i {
  font-size: 2rem;
  margin-right: 0.5rem;
}

.register-title {
  color: #333;
  font-weight: 700;
  margin-bottom: 0.5rem;
}

.form-label {
  font-weight: 600;
  color: #555;
}

.form-control, .form-select {
  border-radius: 8px;
  border: 2px solid #e9ecef;
  padding: 0.75rem;
  font-size: 1rem;
  transition: border-color 0.3s ease;
}

.form-control:focus, .form-select:focus {
  border-color: #007bff;
  box-shadow: 0 0 0 0.2rem rgba(0, 123, 255, 0.25);
}

.btn-primary {
  background: linear-gradient(45deg, #007bff, #0056b3);
  border: none;
  border-radius: 50px;
  padding: 0.75rem 2rem;
  font-weight: 600;
  transition: all 0.3s ease;
}

.btn-primary:hover:not(:disabled) {
  background: linear-gradient(45deg, #0056b3, #004085);
  transform: translateY(-1px);
  box-shadow: 0 10px 25px rgba(0,0,0,0.2);
}

.btn-primary:disabled {
  opacity: 0.7;
  cursor: not-allowed;
}

.alert {
  border-radius: 8px;
  border: none;
}

.alert-success {
  background: #d4edda;
  color: #155724;
}

.alert-danger {
  background: #f8d7da;
  color: #721c24;
}

.form-text {
  font-size: 0.875rem;
  color: #6c757d;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .register-card {
    padding: 2rem;
    margin: 1rem;
  }

  .register-title {
    font-size: 1.75rem;
  }
}
</style>