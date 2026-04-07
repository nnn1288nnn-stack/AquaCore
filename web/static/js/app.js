// 澎湖數位老船長 - 前端應用邏輯

// API 基礎 URL
const API_BASE_URL = window.location.origin;
const GOLANG_API = `${API_BASE_URL}/api`;
const PYTHON_AI_API = `${API_BASE_URL}/api/ai`;

// 通知函數
function showNotification(message, type = 'info') {
    console.log(`[${type.toUpperCase()}] ${message}`);
    alert(message);
}

// API 呼叫封裝
async function apiCall(method, endpoint, data = null) {
    try {
        const options = {
            method,
            headers: {
                'Content-Type': 'application/json',
            }
        };

        if (data) {
            options.body = JSON.stringify(data);
        }

        const response = await fetch(endpoint, options);

        if (!response.ok) {
            throw new Error(`HTTP ${response.status}: ${response.statusText}`);
        }

        return await response.json();
    } catch (error) {
        console.error('API Error:', error);
        showNotification(`請求失敗: ${error.message}`, 'error');
        throw error;
    }
}

// 初始化應用
document.addEventListener('DOMContentLoaded', () => {
    console.log('澎湖數位老船長 - 應用啟動');
    initializeEventListeners();
    loadDashboardData();
});

// 初始化事件監聽器
function initializeEventListeners() {
    const sendBtn = document.getElementById('send-btn');
    const userInput = document.getElementById('user-input');

    if (sendBtn) {
        sendBtn.addEventListener('click', handleChatMessage);
    }

    if (userInput) {
        userInput.addEventListener('keypress', (e) => {
            if (e.key === 'Enter') {
                handleChatMessage();
            }
        });
    }

    // 錄音按鈕 (如果使用語音)
    const recordBtn = document.getElementById('record-btn');
    if (recordBtn) {
        recordBtn.addEventListener('click', startVoiceInput);
    }
}

// 加載儀表板數據
async function loadDashboardData() {
    try {
        // 模擬數據加載 (實際應該從 API 獲取)
        console.log('加載儀表板數據...');

        // 在實際應用中，應該呼叫:
        // const data = await apiCall('GET', `${GOLANG_API}/dashboard`);
    } catch (error) {
        console.error('加載失敗:', error);
    }
}

// 聊天消息處理
async function handleChatMessage() {
    const userInput = document.getElementById('user-input');
    const message = userInput.value.trim();

    if (!message) return;

    // 顯示用戶消息
    displayMessage(message, 'user');
    userInput.value = '';

    try {
        // 發送到 Python AI API
        const response = await apiCall('POST', `${PYTHON_AI_API}/chat`, {
            message: message
        });

        // 顯示 AI 回應
        displayMessage(response.reply, 'bot');
    } catch (error) {
        displayMessage('抱歉，我無法處理您的請求。請稍後再試。', 'bot');
    }
}

// 顯示聊天消息
function displayMessage(text, sender) {
    const messagesContainer = document.getElementById('chat-messages');
    const messageDiv = document.createElement('div');
    messageDiv.className = `message message-${sender}`;
    messageDiv.innerHTML = `<p>${escapeHtml(text)}</p>`;
    messagesContainer.appendChild(messageDiv);
    messagesContainer.scrollTop = messagesContainer.scrollHeight;
}

// HTML 逸出 (防止 XSS)
function escapeHtml(text) {
    const map = {
        '&': '&amp;',
        '<': '&lt;',
        '>': '&gt;',
        '"': '&quot;',
        "'": '&#039;'
    };
    return text.replace(/[&<>"']/g, m => map[m]);
}

// 語音輸入 (如果支援)
async function startVoiceInput() {
    if (!('webkitSpeechRecognition' in window)) {
        showNotification('您的瀏覽器不支援語音輸入', 'error');
        return;
    }

    const recognition = new webkitSpeechRecognition();
    recognition.lang = 'zh-TW';

    recognition.onstart = () => {
        console.log('語音輸入開始...');
    };

    recognition.onresult = (event) => {
        let transcript = '';
        for (let i = event.resultIndex; i < event.results.length; i++) {
            transcript += event.results[i][0].transcript;
        }
        document.getElementById('user-input').value = transcript;
        handleChatMessage();
    };

    recognition.onerror = (event) => {
        console.error('語音識別錯誤:', event.error);
        showNotification(`語音識別失敗: ${event.error}`, 'error');
    };

    recognition.start();
}

// 記錄環境數據
async function recordEnvironmentalData() {
    const waterTemp = document.querySelector('input[placeholder="24.5"]').value;
    const salinity = document.querySelector('input[placeholder="30.2"]').value;
    const dissolvedOxygen = document.querySelector('input[placeholder="7.5"]').value;

    if (!waterTemp || !salinity || !dissolvedOxygen) {
        showNotification('請填入所有環境數據', 'error');
        return;
    }

    try {
        const data = {
            water_temperature: parseFloat(waterTemp),
            salinity: parseFloat(salinity),
            dissolved_oxygen: parseFloat(dissolvedOxygen),
            recorded_at: new Date().toISOString()
        };

        await apiCall('POST', `${GOLANG_API}/environmental-data`, data);
        showNotification('環境數據已記錄', 'success');

        // 清空表單
        document.querySelector('input[placeholder="24.5"]').value = '';
        document.querySelector('input[placeholder="30.2"]').value = '';
        document.querySelector('input[placeholder="7.5"]').value = '';
    } catch (error) {
        console.error('記錄失敗:', error);
    }
}

// 工具函數: 格式化日期
function formatDate(date) {
    const options = { year: 'numeric', month: '2-digit', day: '2-digit' };
    return new Date(date).toLocaleDateString('zh-TW', options);
}

// 工具函數: 格式化時間
function formatTime(date) {
    const options = { hour: '2-digit', minute: '2-digit', second: '2-digit' };
    return new Date(date).toLocaleTimeString('zh-TW', options);
}

// 導出給 HTML 使用
window.aquaculture = {
    recordEnvironmentalData,
    handleChatMessage
};

console.log('✅ 應用邏輯已載入');
