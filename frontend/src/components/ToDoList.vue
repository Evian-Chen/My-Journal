<script setup>
import { ref, onMounted, warn } from 'vue';
import ToDoCard from './ToDoCard.vue';
import axios from 'axios';

const onInputTodo = ref(false)
const cancelAlert = ref(false)
const warning = ref("")
const todo = ref({
    title: "",
    content: "",
    status: ""
})
const todos = ref([
    {
        todoId: 1, 
        title: "temp title",
        content: "temp content",
        status: "temp status",
        color: "#000000"
    }
])

function toggleAddTodo() {
    onInputTodo.value = !onInputTodo.value
    warning.value = ""
    cancelAlert.value = false
}

function cancelAddaddTodo() {
    // 先檢查todo有沒有東西，有東西會警告要儲存草稿還是直接取消
    if (todo.value.title || todo.value.content || todo.value.status) {
        // 有任何一個是有被填寫的
        cancelAlert.value = true
    } else {
        toggleAddTodo()
    }
}

function notSaveTodo() {
    todo.value = {
        title: "",
        content: "",
        status: ""
    }
    cancelAlert.value = false
    toggleAddTodo()
}

async function addTodo() {
    try {
        if (todo.value.title === "") {
            warning.value = "此欄位必填"
            return
        }

        const res = await axios.post("api/todo", {
            title: todo.value.title,
            content: todo.value.content,
            status: todo.value.status
        })
        toggleAddTodo()
        loadInAllTodos()
    } catch (err) {
        console.log("addTodo frontend error: ", err)
    }
}

async function loadInAllTodos() {
    // 載入所有todolist
    try {
        // 更新這邊的todos
        const res = await axios.get("api/todo")
        todos.value = res.data
        console.log("get all todolist: ", res)
    } catch (err) {
        console.log("addTodoList.vue onmounted frontend error: ", err)
    }
}

onMounted(async () => {
    loadInAllTodos()
})
</script>

<template>
  <div class="todolist-root">
    <div class="add-todo">
      <div class="header">
        <h1 class="page-title">Hello from the todolist page</h1>
        <div class="actions">
          <button class="primary add-btn" @click="toggleAddTodo()">
            新增一筆代辦事項
        </button>
        </div>
      </div>
    </div>

    <div class="todos-container">
      <div v-if="onInputTodo" class="input-panel">
        <p v-if="warning" class="warning">{{ warning }}</p>
        <input
          class="input-card"
          v-model="todo.title"
          placeholder="輸入標題..."
        />
        <input
          class="input-card"
          v-model="todo.content"
          placeholder="輸入內容..."
        />
        <input
          class="input-card"
          v-model="todo.status"
          placeholder="輸入狀態..."
        />

        <div class="row">
          <button class="muted" @click="cancelAddaddTodo">取消</button>
          <button class="primary" @click="addTodo">儲存</button>
        </div>

        <div v-if="cancelAlert" class="cancel-alert">
            <h3>你要儲存目前編輯嗎？</h3>
            <div class="row">
              <button class="primary" @click="toggleAddTodo()">儲存進度</button>
              <button class="danger" @click="notSaveTodo">刪除進度</button>
            </div>
        </div>
      </div>
      <div v-else>
        <div v-if="todos.length === 0" class="empty-list">
          <h1>沒有任何代辦事項</h1>
        </div>
        <div v-else class="list-grid">
          <div v-for="(todo, index) in todos" :key="todo.todoId ?? index">
            <ToDoCard :todo="todo" />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
/* 亮色系主題 */
.todolist-root {
  --bg: #f8fafc;
  --surface: #ffffff;
  --text: #0f172a;
  --muted: #64748b;
  --accent: #3b82f6;
  --danger: #ef4444;
  --border: #e2e8f0;
  --border-hover: #cbd5e1;
  --shadow: rgba(0, 0, 0, 0.1);
  --hover-surface: #f1f5f9;
}

/* 黑色系主題 */
[data-theme="dark"] .todolist-root {
  --bg: #0f172a;
  --surface: #1e293b;
  --text: #f1f5f9;
  --muted: #94a3b8;
  --accent: #60a5fa;
  --danger: #f87171;
  --border: #475569;
  --border-hover: #64748b;
  --shadow: rgba(0, 0, 0, 0.3);
  --hover-surface: #334155;
}

.todolist-root {
  background: var(--bg);
  padding: 20px;
  min-height: 100vh;
  color: var(--text);
  transition: background-color 0.3s ease, color 0.3s ease;
}

.header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 12px;
  flex-wrap: wrap;
}

.page-title {
  margin: 0;
  font-size: 1.15rem;
  color: var(--text);
  transition: color 0.3s ease;
}

.actions {
  display: flex;
  gap: 8px;
}

.primary {
  background: var(--accent);
  color: white;
  border: 2px solid var(--accent);
  padding: 8px 12px;
  border-radius: 8px;
  cursor: pointer;
  transition: all 0.2s ease;
  font-weight: 500;
  box-shadow: 0 2px 4px var(--shadow);
}

.primary:hover {
  background: var(--accent);
  opacity: 0.85;
  transform: translateY(-2px);
  box-shadow: 0 6px 16px var(--shadow);
}

.primary:active {
  transform: translateY(0);
  box-shadow: 0 2px 6px var(--shadow);
}

.add-btn {
  background: var(--accent);
  color: white;
  border: 2px solid var(--accent);
}

.muted {
  background: transparent;
  border: 2px solid var(--border);
  padding: 8px 12px;
  border-radius: 8px;
  cursor: pointer;
  color: var(--text);
  transition: all 0.2s ease;
  font-weight: 500;
}

.muted:hover {
  background: var(--hover-surface);
  border-color: var(--border-hover);
  transform: translateY(-2px);
  box-shadow: 0 4px 12px var(--shadow);
}

.muted:active {
  transform: translateY(0);
  box-shadow: 0 1px 4px var(--shadow);
}

.danger {
  background: var(--danger);
  color: white;
  border: 2px solid var(--danger);
  padding: 8px 12px;
  border-radius: 8px;
  cursor: pointer;
  transition: all 0.2s ease;
  font-weight: 500;
  box-shadow: 0 2px 4px var(--shadow);
}

.danger:hover {
  background: var(--danger);
  opacity: 0.85;
  transform: translateY(-2px);
  box-shadow: 0 6px 16px var(--shadow);
}

.danger:active {
  transform: translateY(0);
  box-shadow: 0 2px 6px var(--shadow);
}

.warning {
    color:#ef4444;
}

.input-panel {
  display: flex;
  flex-direction: column;
  gap: 10px;
  background: var(--surface);
  padding: 16px;
  border-radius: 12px;
  border: 2px solid var(--border);
  box-shadow: 0 4px 12px var(--shadow);
  transition: all 0.3s ease;
}

.input-card {
  padding: 10px 12px;
  border-radius: 8px;
  border: 2px solid var(--border);
  background: var(--surface);
  color: var(--text);
  transition: all 0.2s ease;
  font-size: 14px;
}

.input-card:hover {
  border-color: var(--border-hover);
}

.input-card:focus {
  outline: none;
  border-color: var(--accent);
  box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.15);
}

.input-card::placeholder {
  color: var(--muted);
}

.row {
  display: flex;
  gap: 8px;
}

.cancel-alert {
  margin-top: 10px;
  padding: 12px;
  border-radius: 10px;
  background: var(--hover-surface);
  border: 2px solid var(--border);
  box-shadow: 0 2px 8px var(--shadow);
  transition: all 0.3s ease;
}

.cancel-alert h3 {
  color: var(--text);
  margin: 0 0 8px 0;
}

.todos-container {
  margin-top: 14px;
}

.list-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(260px, 1fr));
  gap: 12px;
}

.empty-list {
  text-align: center;
  padding: 24px;
  color: var(--muted);
}

.empty-list h1 {
  color: var(--muted);
  font-weight: 400;
}

@media (max-width: 640px) {
  .todolist-root {
    padding: 14px;
  }
  
  .list-grid {
    grid-template-columns: 1fr;
  }
}
</style>
