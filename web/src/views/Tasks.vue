<template>
  <div>
    <h1>📋 任務管理</h1>
    <p>管理養殖作業任務和維護計劃。</p>

    <div class="mb-3">
      <button class="btn btn-primary">新增任務</button>
    </div>

    <div class="card">
      <div class="card-header">
        <h5>任務清單</h5>
      </div>
      <div class="card-body">
        <div class="task-item" v-for="task in tasks" :key="task.id">
          <div class="d-flex justify-content-between align-items-center">
            <div>
              <h6>{{ task.title }}</h6>
              <p class="mb-0 text-muted">{{ task.description }}</p>
            </div>
            <div class="text-end">
              <small class="text-muted">{{ task.dueDate }}</small>
              <br>
              <span class="badge" :class="task.status === 'pending' ? 'bg-warning' : 'bg-success'">
                {{ task.status === 'pending' ? '待處理' : '已完成' }}
              </span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Tasks',
  data() {
    return {
      tasks: [
        {
          id: 1,
          title: '檢查水泵',
          description: '檢查養殖池水泵運作狀態',
          dueDate: '2026-04-08',
          status: 'pending'
        },
        {
          id: 2,
          title: '添加飼料',
          description: '為 A 區添加 50kg 飼料',
          dueDate: '2026-04-07',
          status: 'completed'
        }
      ]
    }
  }
}
</script>

<style scoped>
.task-item {
  padding: 1rem;
  border-bottom: 1px solid #eee;
}

.task-item:last-child {
  border-bottom: none;
}
</style>