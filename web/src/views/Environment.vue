<template>
  <div>
    <h1>🌊 環境監測</h1>
    <p>透過後端 API 讀取並記錄養殖池環境數據。</p>

    <div class="row mb-4">
      <div class="col-md-6">
        <div class="card">
          <div class="card-header">
            <h5>新增環境數據</h5>
          </div>
          <div class="card-body">
            <form @submit.prevent="submitEnvironmentalData">
              <div class="mb-3">
                <label class="form-label">水溫 (℃)</label>
                <input type="number" step="0.1" class="form-control" v-model.number="form.water_temperature" required />
              </div>
              <div class="mb-3">
                <label class="form-label">pH 值</label>
                <input type="number" step="0.1" class="form-control" v-model.number="form.ph_level" required />
              </div>
              <div class="mb-3">
                <label class="form-label">溶氧量 (mg/L)</label>
                <input type="number" step="0.1" class="form-control" v-model.number="form.dissolved_oxygen" required />
              </div>
              <div class="mb-3">
                <label class="form-label">備註</label>
                <textarea class="form-control" rows="2" v-model="form.notes"></textarea>
              </div>
              <button type="submit" class="btn btn-primary" :disabled="loading">
                {{ loading ? '送出中...' : '送出資料' }}
              </button>
            </form>
            <div v-if="message" class="alert alert-info mt-3">{{ message }}</div>
            <div v-if="error" class="alert alert-danger mt-3">{{ error }}</div>
          </div>
        </div>
      </div>

      <div class="col-md-6">
        <div class="card">
          <div class="card-header">
            <h5>目前 API 端點</h5>
          </div>
          <div class="card-body">
            <p>目前已連線到：</p>
            <ul>
              <li><strong>GET</strong> /api/environmental-data</li>
              <li><strong>POST</strong> /api/environmental-data</li>
            </ul>
            <p>若要擴充天氣資料，可新增一個後端 `/api/weather` 端點，從氣象局開放資料抓取最新天氣資訊。</p>
          </div>
        </div>
      </div>
    </div>

    <div class="card">
      <div class="card-header">
        <h5>歷史環境資料</h5>
      </div>
      <div class="card-body p-0">
        <table class="table table-striped mb-0">
          <thead>
            <tr>
              <th>時間</th>
              <th>水溫</th>
              <th>pH</th>
              <th>溶氧</th>
              <th>備註</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="record in records" :key="record.id">
              <td>{{ formatDate(record.recorded_at) }}</td>
              <td>{{ record.water_temperature ?? '-' }}</td>
              <td>{{ record.ph_level ?? '-' }}</td>
              <td>{{ record.dissolved_oxygen ?? '-' }}</td>
              <td>{{ record.notes || '-' }}</td>
            </tr>
            <tr v-if="records.length === 0">
              <td colspan="5" class="text-center">目前尚無環境資料</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, onMounted } from 'vue'
import { getEnvironmentalData, createEnvironmentalData } from '../services/api'

export default {
  name: 'Environment',
  setup () {
    const records = ref([])
    const loading = ref(false)
    const error = ref('')
    const message = ref('')
    const form = ref({
      water_temperature: 0,
      ph_level: 7.0,
      dissolved_oxygen: 0,
      notes: ''
    })

    const fetchRecords = async () => {
      loading.value = true
      error.value = ''
      try {
        const response = await getEnvironmentalData()
        records.value = response.data.data || []
      } catch (err) {
        error.value = err.response?.data?.error || err.message || '讀取環境資料失敗'
      } finally {
        loading.value = false
      }
    }

    const submitEnvironmentalData = async () => {
      loading.value = true
      error.value = ''
      message.value = ''
      try {
        const payload = {
          water_temperature: form.value.water_temperature,
          ph_level: form.value.ph_level,
          dissolved_oxygen: form.value.dissolved_oxygen,
          notes: form.value.notes
        }
        await createEnvironmentalData(payload)
        message.value = '資料已送出，請稍後重新整理列表。'
        await fetchRecords()
      } catch (err) {
        error.value = err.response?.data?.error || err.message || '送出環境資料失敗'
      } finally {
        loading.value = false
      }
    }

    const formatDate = (value) => {
      return value ? new Date(value).toLocaleString() : '-'
    }

    onMounted(fetchRecords)

    return {
      records,
      loading,
      error,
      message,
      form,
      fetchRecords,
      submitEnvironmentalData,
      formatDate
    }
  }
}
</script>
