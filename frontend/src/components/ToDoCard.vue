<script setup>
import { ref, defineProps } from 'vue'
const props = defineProps({
    todo: Object
})

const onEdit = ref(false)

function toggleOnEdit() {
    onEdit.value = !onEdit.value
}

// 這個顏色之後會被用在底下
const color = props.todo.color
</script>


<template>
  <div class="todo-card" :style="{ borderLeftColor: color || 'var(--accent)' }">
    <header class="card-head">
      <div class="title-wrap">
        <div class="dot" :style="{ background: color || 'var(--accent)' }"></div>
        <h3 class="card-title">{{ props.todo.title || '無標題' }}</h3>
      </div>
      <button class="edit-btn" @click="toggleOnEdit">{{ onEdit ? '儲存' : '編輯' }}</button>
    </header>

    <main class="card-body">
      <div v-if="onEdit" class="edit-fields">
        <input class="field" v-model="props.todo.title" placeholder="標題" />
        <input class="field" v-model="props.todo.content" placeholder="內容" />
        <input class="field" v-model="props.todo.status" placeholder="狀態" />
      </div>

      <div v-else class="display">
        <p class="content">{{ props.todo.content || '沒有內容' }}</p>
        <span class="status">{{ props.todo.status || '未設定' }}</span>
      </div>
    </main>
  </div>
</template>

<style scoped>
/* 亮色系主題 */
.todo-card {
  --card-bg: #ffffff;
  --text: #0f172a;
  --muted: #64748b;
  --accent: #3b82f6;
  --border: #e2e8f0;
  --border-hover: #cbd5e1;
  --shadow: rgba(0, 0, 0, 0.1);
  --hover-surface: #f8fafc;
}

/* 黑色系主題 */
[data-theme="dark"] .todo-card {
  --card-bg: #1e293b;
  --text: #f1f5f9;
  --muted: #94a3b8;
  --accent: #60a5fa;
  --border: #475569;
  --border-hover: #64748b;
  --shadow: rgba(0, 0, 0, 0.3);
  --hover-surface: #334155;
}

.todo-card {
  background: var(--card-bg);
  border-radius: 12px;
  padding: 14px;
  border: 2px solid var(--border);
  border-left: 4px solid var(--accent);
  color: var(--text);
  box-shadow: 0 2px 8px var(--shadow);
  transition: all 0.25s ease;
  cursor: pointer;
}

.todo-card:hover {
  transform: translateY(-4px);
  box-shadow: 0 8px 20px var(--shadow);
  border-color: var(--border-hover);
  border-left-width: 5px;
  background: var(--hover-surface);
}

.todo-card:active {
  transform: translateY(-1px);
  box-shadow: 0 4px 12px var(--shadow);
}

.card-head {
  display: flex;
  justify-content: space-between;
  align-items: center;
  gap: 8px;
}

.title-wrap {
  display: flex;
  align-items: center;
  gap: 8px;
}

.dot {
  width: 10px;
  height: 10px;
  border-radius: 50%;
  box-shadow: 0 0 0 2px var(--border);
  transition: all 0.2s ease;
}

.todo-card:hover .dot {
  transform: scale(1.2);
  box-shadow: 0 0 0 2px var(--accent);
}

.card-title {
  margin: 0;
  font-size: 1rem;
  color: var(--text);
  font-weight: 600;
  transition: color 0.3s ease;
}

.edit-btn {
  background: transparent;
  border: 2px solid var(--border);
  padding: 6px 10px;
  border-radius: 8px;
  cursor: pointer;
  color: var(--accent);
  transition: all 0.2s ease;
  font-weight: 600;
  font-size: 13px;
}

.edit-btn:hover {
  background: var(--accent);
  color: white;
  border-color: var(--accent);
  transform: translateY(-2px);
  box-shadow: 0 4px 12px var(--shadow);
}

.edit-btn:active {
  transform: translateY(0);
  box-shadow: 0 2px 6px var(--shadow);
}

.card-body {
  margin-top: 8px;
}

.field {
  width: 100%;
  padding: 10px;
  border-radius: 8px;
  border: 2px solid var(--border);
  background: var(--card-bg);
  color: var(--text);
  margin-bottom: 8px;
  transition: all 0.2s ease;
  font-size: 14px;
}

.field:hover {
  border-color: var(--border-hover);
}

.field:focus {
  outline: none;
  border-color: var(--accent);
  box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.15);
}

.field::placeholder {
  color: var(--muted);
}

.display .content {
  color: var(--muted);
  margin: 0 0 8px 0;
  line-height: 1.5;
}

.status {
  display: inline-block;
  font-size: 0.82rem;
  color: var(--accent);
  background: rgba(59, 130, 246, 0.1);
  padding: 5px 10px;
  border-radius: 999px;
  border: 2px solid var(--accent);
  font-weight: 600;
  transition: all 0.2s ease;
  cursor: default;
}

[data-theme="dark"] .status {
  background: rgba(96, 165, 250, 0.15);
}

.status:hover {
  background: var(--accent);
  color: white;
  transform: scale(1.05);
  box-shadow: 0 2px 8px var(--shadow);
}
</style>