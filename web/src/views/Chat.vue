<template>
  <div>
    <h1>🤖 AI 助手</h1>
    <p>與 AI 助手對話，獲取養殖建議和幫助。</p>

    <div class="chat-container">
      <div class="chat-messages" ref="messagesContainer">
        <div
          v-for="message in messages"
          :key="message.id"
          class="message"
          :class="{ 'user-message': message.sender === 'user', 'ai-message': message.sender === 'ai' }"
        >
          {{ message.text }}
        </div>
      </div>

      <div class="chat-input">
        <div class="input-group">
          <input
            type="text"
            class="form-control"
            v-model="newMessage"
            @keyup.enter="sendMessage"
            placeholder="輸入您的問題..."
          >
          <button class="btn btn-primary" @click="sendMessage" :disabled="!newMessage.trim()">
            發送
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'Chat',
  data() {
    return {
      messages: [
        {
          id: 1,
          sender: 'ai',
          text: '您好！我是澎湖數位老船長的 AI 助手，有什麼可以幫助您的嗎？'
        }
      ],
      newMessage: ''
    }
  },
  methods: {
    async sendMessage() {
      if (!this.newMessage.trim()) return

      // 添加用户消息
      const userMessage = {
        id: Date.now(),
        sender: 'user',
        text: this.newMessage
      }
      this.messages.push(userMessage)

      const messageToSend = this.newMessage
      this.newMessage = ''

      try {
        // 调用 AI API
        const response = await axios.post('/api/chat', {
          message: messageToSend,
          language: 'zh-TW'
        })

        // 添加 AI 回复
        const aiMessage = {
          id: Date.now() + 1,
          sender: 'ai',
          text: response.data.reply
        }
        this.messages.push(aiMessage)

        this.$nextTick(() => {
          this.scrollToBottom()
        })
      } catch (error) {
        console.error('发送消息失败:', error)
        const errorMessage = {
          id: Date.now() + 1,
          sender: 'ai',
          text: '抱歉，處理您的請求時出現錯誤。請稍後再試。'
        }
        this.messages.push(errorMessage)
      }
    },

    scrollToBottom() {
      const container = this.$refs.messagesContainer
      container.scrollTop = container.scrollHeight
    }
  }
}
</script>

<style scoped>
.chat-container {
  height: 600px;
  display: flex;
  flex-direction: column;
  border: 1px solid #ddd;
  border-radius: 10px;
  overflow: hidden;
}

.chat-messages {
  flex: 1;
  padding: 1rem;
  overflow-y: auto;
  background: #f8f9fa;
}

.message {
  margin-bottom: 1rem;
  padding: 0.75rem 1rem;
  border-radius: 10px;
  max-width: 70%;
}

.user-message {
  background: #007bff;
  color: white;
  margin-left: auto;
  text-align: right;
}

.ai-message {
  background: white;
  border: 1px solid #ddd;
}

.chat-input {
  padding: 1rem;
  background: white;
  border-top: 1px solid #ddd;
}
</style>