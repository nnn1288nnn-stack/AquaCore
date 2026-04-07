<template>
  <div>
    <div class="d-flex justify-content-between align-items-center mb-4">
      <h1>📊 儀表板</h1>
      <small class="text-muted">最後更新: {{ lastUpdate }}</small>
    </div>

    <!-- 指标卡片 -->
    <div class="row mb-4">
      <div class="col-md-3 mb-3" v-for="metric in metrics" :key="metric.id">
        <div class="card h-100">
          <div class="card-body text-center">
            <div class="metric-icon mb-2">{{ metric.icon }}</div>
            <h5 class="card-title">{{ metric.title }}</h5>
            <p class="metric-value">{{ metric.value }}</p>
            <small class="text-muted">{{ metric.unit }}</small>
          </div>
        </div>
      </div>
    </div>

    <!-- 图表区域 -->
    <div class="row">
      <div class="col-md-8">
        <div class="card">
          <div class="card-header">
            <h5>環境趨勢圖</h5>
          </div>
          <div class="card-body">
            <canvas id="environmentChart" height="200"></canvas>
          </div>
        </div>
      </div>

      <div class="col-md-4">
        <div class="card">
          <div class="card-header">
            <h5>系統狀態</h5>
          </div>
          <div class="card-body">
            <div class="status-item" v-for="status in systemStatus" :key="status.name">
              <span class="status-dot" :class="status.status"></span>
              {{ status.name }}
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
  name: 'Dashboard',
  data() {
    return {
      lastUpdate: new Date().toLocaleString('zh-TW'),
      metrics: [
        { id: 1, title: '水溫', value: '24.5', unit: '°C', icon: '🌡️' },
        { id: 2, title: '鹽度', value: '30.2', unit: 'ppt', icon: '🧂' },
        { id: 3, title: '溶氧量', value: '7.5', unit: 'mg/L', icon: '💧' },
        { id: 4, title: '待辦任務', value: '5', unit: '項', icon: '📋' }
      ],
      systemStatus: [
        { name: 'Nginx 服務', status: 'online' },
        { name: 'MariaDB 資料庫', status: 'online' },
        { name: 'Go API 服務', status: 'offline' },
        { name: 'Python AI 服務', status: 'offline' }
      ]
    }
  },
  mounted() {
    this.loadDashboardData()
    this.initChart()
  },
  methods: {
    async loadDashboardData() {
      try {
        // 这里可以调用实际的 API 获取数据
        // const response = await axios.get('/api/dashboard')
        // this.metrics = response.data.metrics
        // this.systemStatus = response.data.status

        // 暂时使用模拟数据
        this.lastUpdate = new Date().toLocaleString('zh-TW')
      } catch (error) {
        console.error('加载仪表板数据失败:', error)
      }
    },

    initChart() {
      // 这里可以初始化图表库如 Chart.js
      // 暂时留空
    }
  }
}
</script>

<style scoped>
.metric-icon {
  font-size: 2rem;
}

.metric-value {
  font-size: 2rem;
  font-weight: bold;
  color: #007bff;
}

.status-item {
  display: flex;
  align-items: center;
  margin-bottom: 0.5rem;
}

.status-dot {
  width: 10px;
  height: 10px;
  border-radius: 50%;
  margin-right: 0.5rem;
}

.status-dot.online {
  background-color: #28a745;
}

.status-dot.offline {
  background-color: #dc3545;
}

.card {
  border-radius: 10px;
  box-shadow: 0 2px 10px rgba(0,0,0,0.1);
}
</style>