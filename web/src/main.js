import { createApp } from 'vue'
import { createPinia } from 'pinia'
import App from './App.vue'
import router from './router'
import { useAppStore } from './store/app'
import 'bootstrap/dist/css/bootstrap.min.css'
import 'bootstrap/dist/js/bootstrap.bundle.min.js'

// 创建Vue应用
const app = createApp(App)

// 创建并使用Pinia
const pinia = createPinia()
app.use(pinia)
app.use(router)

// 初始化应用状态
const appStore = useAppStore()
appStore.initialize()

// 挂载应用
app.mount('#app')