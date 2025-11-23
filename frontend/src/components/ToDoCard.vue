<script setup>
import { ref, defineProps } from "vue";

const props = defineProps({
  todo: Object,
});

const onEdit = ref(false);
const onDelete = ref(false);
const showFull = ref(false);
const statusOptions = ref(["Pending", "In Progress", "Done"]);

function toggleOnEdit() {
  console.log("toggle edit")
  onEdit.value = !onEdit.value;
}

function toggleOnDelete() {
  onDelete.value = !onDelete.value;
}

function toggleShowFull() {
  showFull.value = !showFull.value;
}

// 這個顏色之後會被用在底下
const color = props.todo.color;
</script>

<template>
  <div class="container">
    <div v-if="onDelete">
      <h3>確定要刪除此代辦事項嗎？</h3>
      <button @click="toggleOnEdit(), $emit('deleteTodo', props.todo._id)">確定</button>
      <button @click="toggleOnDelete">取消</button>
    </div>

    <div v-if="showFull" class="full-view">
      <div class="full-view-header">
        <button @click="toggleShowFull" class="close-btn">✕</button>
        <button class="edit-btn" @click="toggleOnEdit">
          {{ onEdit ? "儲存" : "編輯" }}
        </button>
        <button class="delete-btn" @click="toggleOnDelete">刪除</button>
      </div>

      <div class="full-view-content">
        <div v-if="onEdit" class="edit-fields">
          <input class="field" v-model="props.todo.title" placeholder="標題" />
          <textarea
            class="field textarea-field"
            v-model="props.todo.content"
            placeholder="內容"
            rows="6"
          ></textarea>
          <select class="field" v-model="props.todo.status">
            <option v-for="(option, index) in statusOptions" :key="index">
              {{ option }}
            </option>
          </select>
        </div>

        <div v-else class="display-full">
          <div class="title-wrap-full">
            <div
              class="dot"
              :style="{ background: color || 'var(--accent)' }"
            ></div>
            <h2 class="full-title">{{ props.todo.title || "無標題" }}</h2>
          </div>
          <div class="content-full">
            <p class="full-content">{{ props.todo.content || "沒有內容" }}</p>
          </div>
          <div class="meta-info">
            <span class="status-full">{{ props.todo.status || "未設定" }}</span>
            <span class="date-info">{{
              new Date().toLocaleDateString("zh-TW")
            }}</span>
          </div>
        </div>
      </div>
    </div>
    <div v-else>
      <button
        @click="toggleShowFull"
        class="todo-card"
        :style="{ borderLeftColor: color || 'var(--accent)' }"
      >
        <header class="card-head">
          <div class="title-wrap">
            <div
              class="dot"
              :style="{ background: color || 'var(--accent)' }"
            ></div>
            <h3 class="card-title">{{ props.todo.title || "無標題" }}</h3>
          </div>
          <button class="edit-btn" @click="toggleOnEdit">
            {{ onEdit ? "儲存" : "編輯" }}
          </button>
        </header>

        <main class="card-body">
          <div v-if="onEdit" class="edit-fields">
            <input
              class="field"
              v-model="props.todo.title"
              placeholder="標題"
            />
            <input
              class="field"
              v-model="props.todo.content"
              placeholder="內容"
            />
            <input
              class="field"
              v-model="props.todo.status"
              placeholder="狀態"
            />
            <select class="field" v-model="props.todo.status">
              <option v-for="(option, index) in statusOptions" :key="index">
                {{ option }}
              </option>
            </select>
          </div>

          <div v-else class="display">
            <!-- 如果 content 最多只顯示十五個字 -->
            <p class="content">{{ props.todo.content || "沒有內容" }}</p>
            <span class="status">{{ props.todo.status || "未設定" }}</span>
          </div>
        </main>
      </button>
    </div>
  </div>
</template>

<style scoped>
/* 基本色階與按鈕變數 */
.todo-card {
  --card-bg: #ffffff;
  --text: #0f172a;
  --muted: #64748b;
  --accent: #3b82f6;
  --border: #e2e8f0;
  --border-hover: #cbd5e1;
  --shadow: rgba(0, 0, 0, 0.1);
  --hover-surface: #f8fafc;

  /* button / control variables */
  --btn-bg: transparent;
  --btn-border: var(--border);
  --btn-text: var(--accent);
  --btn-bg-hover: var(--accent);
  --btn-hover-text: #ffffff;
  --danger: #ef4444;
  --danger-contrast: #ffffff;
}

/* 黑色主題覆寫變數，確保按鈕文字有良好對比 */
[data-theme="dark"] .todo-card {
  --card-bg: #0f172a; /* 更深色背景 */
  --text: #e6eef8;
  --muted: #94a3b8;
  --accent: #60a5fa;
  --border: #1f2a37;
  --border-hover: #334155;
  --shadow: rgba(0, 0, 0, 0.45);
  --hover-surface: #071327;

  --btn-bg: transparent;
  --btn-border: var(--border);
  --btn-text: var(--accent);
  --btn-bg-hover: var(--accent);
  --btn-hover-text: #001021; /* 深色主題 hover 後用深色文字（與淺色背景搭配） */
  --danger: #f87171; /* 微調 danger 以便在深色中顯示良好 */
  --danger-contrast: #08121a;
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

/* 統一按鈕樣式，使用變數確保在不同主題下可讀 */
button.edit-btn,
button.close-btn,
button.delete-btn {
  background: var(--btn-bg);
  border: 2px solid var(--btn-border);
  padding: 6px 10px;
  border-radius: 8px;
  cursor: pointer;
  color: var(--btn-text);
  transition: all 0.18s ease;
  font-weight: 600;
  font-size: 13px;
  display: inline-flex;
  align-items: center;
  gap: 6px;
}

button.edit-btn {
  padding: 6px 10px;
}

button.edit-btn:hover {
  background: var(--btn-bg-hover);
  color: var(--btn-hover-text);
  border-color: var(--btn-bg-hover);
  transform: translateY(-2px);
  box-shadow: 0 4px 12px var(--shadow);
}

button.close-btn {
  padding: 8px 12px;
  font-size: 16px;
  color: var(--muted);
  border-radius: 8px;
}

/* close 按鈕在 hover 時使用同一套變數，確保深色/亮色都有對比 */
button.close-btn:hover {
  background: var(--btn-bg-hover);
  color: var(--btn-hover-text);
  border-color: var(--btn-bg-hover);
  transform: translateY(-2px);
}

/* danger button 使用專屬變數 */
button.delete-btn {
  border: 2px solid var(--danger);
  color: var(--danger);
  padding: 6px 10px;
  font-size: 13px;
}

button.delete-btn:hover {
  background: var(--danger);
  color: var(--danger-contrast);
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(239, 68, 68, 0.25);
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
  box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.12);
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
  background: rgba(59, 130, 246, 0.08);
  padding: 5px 10px;
  border-radius: 999px;
  border: 2px solid var(--accent);
  font-weight: 600;
  transition: all 0.2s ease;
  cursor: default;
}

[data-theme="dark"] .status {
  background: rgba(96, 165, 250, 0.11);
}

.status:hover {
  background: var(--accent);
  color: var(--btn-hover-text);
  transform: scale(1.05);
  box-shadow: 0 2px 8px var(--shadow);
}

/* 完整視圖樣式 */
.full-view {
  background: var(--card-bg);
  border-radius: 16px;
  padding: 24px;
  border: 2px solid var(--border);
  color: var(--text);
  box-shadow: 0 12px 32px var(--shadow);
  max-width: 600px;
  margin: 0 auto;
  position: relative;
}

.full-view-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
  gap: 12px;
}

.full-view-content {
  min-height: 200px;
}

.title-wrap-full {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-bottom: 16px;
}

.full-title {
  margin: 0;
  font-size: 1.75rem;
  color: var(--text);
  font-weight: 700;
  line-height: 1.3;
}

.content-full {
  margin-bottom: 20px;
}

.full-content {
  color: var(--text);
  margin: 0;
  line-height: 1.6;
  font-size: 1.1rem;
  white-space: pre-wrap;
  word-wrap: break-word;
}

.meta-info {
  display: flex;
  justify-content: space-between;
  align-items: center;
  gap: 12px;
  margin-top: 20px;
  padding-top: 16px;
  border-top: 2px solid var(--border);
}

.status-full {
  display: inline-block;
  font-size: 0.9rem;
  color: var(--accent);
  background: rgba(59, 130, 246, 0.08);
  padding: 8px 16px;
  border-radius: 999px;
  border: 2px solid var(--accent);
  font-weight: 600;
  transition: all 0.2s ease;
}

[data-theme="dark"] .status-full {
  background: rgba(96, 165, 250, 0.11);
}

.date-info {
  color: var(--muted);
  font-size: 0.85rem;
  font-weight: 500;
}

.textarea-field {
  min-height: 120px;
  resize: vertical;
  font-family: inherit;
}

.display-full .field {
  margin-bottom: 16px;
}
</style>
