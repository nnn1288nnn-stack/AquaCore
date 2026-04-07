<template>
  <div class="home-page">
    <!-- 導航欄 -->
    <header class="hero-header">
      <nav class="navbar navbar-expand-lg navbar-dark">
        <div class="container">
          <router-link class="navbar-brand d-flex align-items-center" to="/">
            <i class="bi bi-water me-2"></i>
            <span class="brand-text">澎湖數位老船長</span>
          </router-link>

          <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNav">
            <span class="navbar-toggler-icon"></span>
          </button>

          <div class="collapse navbar-collapse" id="navbarNav">
            <ul class="navbar-nav ms-auto">
              <li class="nav-item">
                <a class="nav-link" href="#features">功能特色</a>
              </li>
              <li class="nav-item">
                <a class="nav-link" href="#about">關於我們</a>
              </li>
              <li class="nav-item">
                <a class="nav-link" href="#contact">聯絡我們</a>
              </li>
              <li class="nav-item ms-2">
                <router-link class="btn btn-outline-light" to="/login">登入</router-link>
              </li>
              <li class="nav-item ms-2">
                <router-link class="btn btn-primary" to="/register">免費註冊</router-link>
              </li>
            </ul>
          </div>
        </div>
      </nav>
    </header>

    <!-- 英雄區塊 -->
    <section class="hero-section">
      <div class="container">
        <div class="row align-items-center">
          <div class="col-lg-6">
            <h1 class="hero-title">
              智慧養殖，<br>
              <span class="text-primary">數位轉型</span>
            </h1>
            <p class="hero-subtitle">
              整合 AI 技術與物聯網，為澎湖養殖業者提供全方位的智慧管理解決方案，
              提升生產效率，降低營運成本，實現永續發展。
            </p>
            <div class="hero-actions">
              <router-link class="btn btn-primary btn-lg me-3" to="/register">
                <i class="bi bi-rocket-takeoff me-2"></i>
                立即開始
              </router-link>
              <a class="btn btn-outline-primary btn-lg" href="#features">
                <i class="bi bi-info-circle me-2"></i>
                了解更多
              </a>
            </div>
          </div>
          <div class="col-lg-6">
            <div class="hero-image">
              <img src="/assets/images/hero-dashboard.png" alt="智慧養殖儀表板" class="img-fluid">
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- 功能特色 -->
    <section id="features" class="features-section">
      <div class="container">
        <div class="section-header text-center mb-5">
          <h2>核心功能</h2>
          <p class="text-muted">整合多項先進技術，提供完整的養殖管理解決方案</p>
        </div>

        <div class="row g-4">
          <div class="col-md-6 col-lg-3">
            <div class="feature-card">
              <div class="feature-icon">
                <i class="bi bi-graph-up"></i>
              </div>
              <h5>即時監測</h5>
              <p>24小時監控水質參數、設備狀態，及時發現異常情況</p>
            </div>
          </div>

          <div class="col-md-6 col-lg-3">
            <div class="feature-card">
              <div class="feature-icon">
                <i class="bi bi-robot"></i>
              </div>
              <h5>AI 助手</h5>
              <p>智慧分析數據，提供養殖建議和預警通知</p>
            </div>
          </div>

          <div class="col-md-6 col-lg-3">
            <div class="feature-card">
              <div class="feature-icon">
                <i class="bi bi-box-seam"></i>
              </div>
              <h5>庫存管理</h5>
              <p>自動追蹤飼料、藥品庫存，智能補貨提醒</p>
            </div>
          </div>

          <div class="col-md-6 col-lg-3">
            <div class="feature-card">
              <div class="feature-icon">
                <i class="bi bi-calendar-check"></i>
              </div>
              <h5>任務管理</h5>
              <p>工作任務排程、進度追蹤、團隊協作</p>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- 數據統計 -->
    <section class="stats-section bg-light">
      <div class="container">
        <div class="row text-center">
          <div class="col-md-3">
            <div class="stat-item">
              <div class="stat-number">500+</div>
              <div class="stat-label">養殖戶</div>
            </div>
          </div>
          <div class="col-md-3">
            <div class="stat-item">
              <div class="stat-number">10萬+</div>
              <div class="stat-label">數據點</div>
            </div>
          </div>
          <div class="col-md-3">
            <div class="stat-item">
              <div class="stat-number">99.9%</div>
              <div class="stat-label">正常運行</div>
            </div>
          </div>
          <div class="col-md-3">
            <div class="stat-item">
              <div class="stat-number">24/7</div>
              <div class="stat-label">技術支持</div>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- 關於我們 -->
    <section id="about" class="about-section">
      <div class="container">
        <div class="row align-items-center">
          <div class="col-lg-6">
            <h2>關於澎湖數位老船長</h2>
            <p>
              我們是專注於養殖漁業數位化的科技公司，結合 AI、物聯網和大數據技術，
              為傳統養殖業帶來革命性的改變。
            </p>
            <p>
              我們的團隊由資深養殖專家、AI 工程師和物聯網專才組成，
              致力於為澎湖及全台灣的養殖業者提供最先進的智慧管理解決方案。
            </p>
            <ul class="about-list">
              <li><i class="bi bi-check-circle text-success"></i> 10年養殖業經驗</li>
              <li><i class="bi bi-check-circle text-success"></i> AI 驅動的決策支持</li>
              <li><i class="bi bi-check-circle text-success"></i> 物聯網即時監測</li>
              <li><i class="bi bi-check-circle text-success"></i> 本土化服務支持</li>
            </ul>
          </div>
          <div class="col-lg-6">
            <div class="about-image">
              <img src="/assets/images/about-team.png" alt="我們的團隊" class="img-fluid rounded">
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- 聯絡我們 -->
    <section id="contact" class="contact-section bg-light">
      <div class="container">
        <div class="row">
          <div class="col-lg-8 mx-auto text-center">
            <h2>開始您的數位養殖之旅</h2>
            <p class="mb-4">
              聯絡我們，了解更多關於智慧養殖解決方案的資訊，
              或安排免費的系統演示和諮詢服務。
            </p>
            <div class="contact-info">
              <div class="row">
                <div class="col-md-4">
                  <div class="contact-item">
                    <i class="bi bi-telephone"></i>
                    <h6>電話</h6>
                    <p>(06) 123-4567</p>
                  </div>
                </div>
                <div class="col-md-4">
                  <div class="contact-item">
                    <i class="bi bi-envelope"></i>
                    <h6>郵件</h6>
                    <p>info@penghu-aquaculture.com</p>
                  </div>
                </div>
                <div class="col-md-4">
                  <div class="contact-item">
                    <i class="bi bi-geo-alt"></i>
                    <h6>地址</h6>
                    <p>澎湖縣馬公市xxx路xxx號</p>
                  </div>
                </div>
              </div>
            </div>
            <router-link class="btn btn-primary btn-lg mt-4" to="/register">
              <i class="bi bi-person-plus me-2"></i>
              立即註冊體驗
            </router-link>
          </div>
        </div>
      </div>
    </section>

    <!-- 頁腳 -->
    <footer class="footer">
      <div class="container">
        <div class="row">
          <div class="col-md-6">
            <div class="footer-brand">
              <i class="bi bi-water me-2"></i>
              <span>澎湖數位老船長</span>
            </div>
            <p class="footer-desc">
              智慧養殖，永續漁業。整合 AI 與物聯網技術，
              為養殖業者提供全方位的數位化解決方案。
            </p>
          </div>
          <div class="col-md-3">
            <h6>快速連結</h6>
            <ul class="footer-links">
              <li><router-link to="/login">登入系統</router-link></li>
              <li><router-link to="/register">註冊帳號</router-link></li>
              <li><a href="#features">功能介紹</a></li>
              <li><a href="#contact">聯絡我們</a></li>
            </ul>
          </div>
          <div class="col-md-3">
            <h6>技術支持</h6>
            <ul class="footer-links">
              <li><a href="#">使用說明</a></li>
              <li><a href="#">常見問題</a></li>
              <li><a href="#">技術支援</a></li>
              <li><a href="#">隱私政策</a></li>
            </ul>
          </div>
        </div>
        <hr>
        <div class="footer-bottom">
          <p>&copy; 2026 澎湖數位老船長. All rights reserved.</p>
        </div>
      </div>
    </footer>
  </div>
</template>

<script>
export default {
  name: 'Home',
  mounted() {
    // 平滑滾動到錨點
    const hash = window.location.hash
    if (hash) {
      setTimeout(() => {
        const element = document.querySelector(hash)
        if (element) {
          element.scrollIntoView({ behavior: 'smooth' })
        }
      }, 100)
    }
  }
}
</script>

<style scoped>
/* 英雄區塊樣式 */
.hero-header {
  position: fixed;
  top: 0;
  width: 100%;
  z-index: 1030;
  background: rgba(0, 123, 255, 0.95);
  backdrop-filter: blur(10px);
  box-shadow: 0 2px 20px rgba(0,0,0,0.1);
}

.navbar-brand .brand-text {
  font-weight: bold;
  font-size: 1.25rem;
}

.hero-section {
  padding: 120px 0 80px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  min-height: 100vh;
  display: flex;
  align-items: center;
}

.hero-title {
  font-size: 3.5rem;
  font-weight: 700;
  line-height: 1.2;
  margin-bottom: 1.5rem;
}

.hero-subtitle {
  font-size: 1.25rem;
  line-height: 1.6;
  margin-bottom: 2rem;
  opacity: 0.9;
}

.hero-actions .btn {
  padding: 0.75rem 2rem;
  font-weight: 600;
  border-radius: 50px;
  transition: all 0.3s ease;
}

.hero-actions .btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 10px 25px rgba(0,0,0,0.2);
}

.hero-image {
  text-align: center;
}

/* 功能特色樣式 */
.features-section {
  padding: 80px 0;
}

.section-header h2 {
  font-size: 2.5rem;
  font-weight: 700;
  margin-bottom: 1rem;
}

.feature-card {
  background: white;
  padding: 2rem;
  border-radius: 15px;
  text-align: center;
  box-shadow: 0 5px 15px rgba(0,0,0,0.08);
  transition: transform 0.3s ease, box-shadow 0.3s ease;
  height: 100%;
}

.feature-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 15px 35px rgba(0,0,0,0.15);
}

.feature-icon {
  width: 80px;
  height: 80px;
  background: linear-gradient(45deg, #007bff, #0056b3);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  margin: 0 auto 1.5rem;
  color: white;
  font-size: 2rem;
}

.feature-card h5 {
  font-weight: 600;
  margin-bottom: 1rem;
}

/* 統計數據樣式 */
.stats-section {
  padding: 60px 0;
}

.stat-item {
  padding: 2rem 0;
}

.stat-number {
  font-size: 3rem;
  font-weight: 700;
  color: #007bff;
  margin-bottom: 0.5rem;
}

.stat-label {
  font-size: 1.1rem;
  color: #6c757d;
  font-weight: 500;
}

/* 關於我們樣式 */
.about-section {
  padding: 80px 0;
}

.about-section h2 {
  font-size: 2.5rem;
  font-weight: 700;
  margin-bottom: 2rem;
}

.about-list {
  list-style: none;
  padding: 0;
}

.about-list li {
  padding: 0.5rem 0;
  font-size: 1.1rem;
}

.about-image {
  text-align: center;
}

/* 聯絡我們樣式 */
.contact-section {
  padding: 80px 0;
}

.contact-info {
  margin: 2rem 0;
}

.contact-item {
  padding: 1.5rem;
  background: white;
  border-radius: 10px;
  box-shadow: 0 2px 10px rgba(0,0,0,0.05);
}

.contact-item i {
  font-size: 2rem;
  color: #007bff;
  margin-bottom: 1rem;
}

.contact-item h6 {
  font-weight: 600;
  margin-bottom: 0.5rem;
}

/* 頁腳樣式 */
.footer {
  background: #212529;
  color: white;
  padding: 60px 0 30px;
}

.footer-brand {
  font-size: 1.5rem;
  font-weight: bold;
  margin-bottom: 1rem;
}

.footer-desc {
  color: #adb5bd;
  line-height: 1.6;
  margin-bottom: 2rem;
}

.footer-links {
  list-style: none;
  padding: 0;
}

.footer-links li {
  margin-bottom: 0.5rem;
}

.footer-links a {
  color: #adb5bd;
  text-decoration: none;
  transition: color 0.3s ease;
}

.footer-links a:hover {
  color: #007bff;
}

.footer-bottom {
  text-align: center;
  padding-top: 30px;
  border-top: 1px solid #495057;
  color: #6c757d;
}

/* 響應式設計 */
@media (max-width: 768px) {
  .hero-title {
    font-size: 2.5rem;
  }

  .hero-subtitle {
    font-size: 1rem;
  }

  .section-header h2 {
    font-size: 2rem;
  }

  .feature-card {
    margin-bottom: 2rem;
  }

  .stat-number {
    font-size: 2.5rem;
  }
}
</style>