# My-Journal

包含兩大核心功能：

1. **日記（Journal System）**
2. **待辦事項（Todo List System）**

後端使用：

* Go
* MongoDB
* Redis（快取，可選）
* JWT

前端使用：

* Vue（兩個 Tab UI）

---

# 🌟 一、產品使用情境（User Scenario）

使用者每天打開系統可以：

* 左邊 Tab：寫今天的日記（自由書寫）
* 右邊 Tab：管理 Todo（今天要做什麼）

兩者互相獨立
但共用同一個使用者帳號

未來也可以擴充：

* 從 Todo 自動生成日記摘要
* 從日記分析情緒推估 Todo 生產力

---

# 🌟 二、主要功能整理（FRD）

## ✔ 1. 使用者系統（User）

* 註冊
* 登入
* JWT 驗證
* 取得自己的 profile

---

## ✔ 2. 日記系統（Journal）

每篇日記包含：

* 標題
* 內容（純文字）
* 標籤（可選）
* 心情（emoji：😀 😐 😞）
* 日期（預設今天）
* 建立/更新時間

日記功能：

* 新增
* 讀取（單篇）
* 編輯
* 刪除
* 搜尋（關鍵字/標籤/心情/日期）

---

## ✔ 3. Todo 系統（Todo List）

每個 Todo 包含：

* item 名稱
* 是否完成
* 所屬日期（預設今天）
* 排序（optional）
* 建立時間

Todo 功能：

* 新增 todo item
* 列表 todo items
* 標記完成/未完成
* 編輯 todo item
* 刪除 item
* 依日期查詢（例如今天 / 本週）

Todo 不需搜尋系統（日記才有）

---

# 🌟 三、前端（Vue）視覺架構

```
----------------------------------------
|  Personal Journal & Todo Manager     |
----------------------------------------
|   [ Journal ]   [ Todo ]             |
----------------------------------------
```

## 【Tab 1：Journal（日記）】

UI：

* 日期選擇器（預設今天）
* 標題欄位
* 內容 textarea
* 標籤選擇器（chips）
* 心情 emoji
* 「儲存」 / 「刪除」

## 【Tab 2：Todo】

UI：

* “新增代辦” 輸入框
* 待辦清單 (checkbox + item)
* 排序（選做）
* 完成顯示灰色

---

# 🌟 四、MongoDB 資料模型（Model）

---

## 1. User

```
User {
  _id: ObjectId
  email: string
  passwordHash: string
  createdAt: time
}
```

---

## 2. Diary

```
Diary {
  _id: ObjectId
  userId: ObjectId
  title: string
  content: string
  tags: [string]
  mood: string (happy | neutral | sad)
  date: string (YYYY-MM-DD)
  createdAt: time
  updatedAt: time
}
```

---

## 3. Todo

```
Todo {
  _id: ObjectId
  userId: ObjectId
  content: string
  done: bool
  date: string (YYYY-MM-DD)
  createdAt: time
}
```

---

# 🌟 五、API 設計（更新版本）

## 🔐 Auth

### POST /api/auth/signup

### POST /api/auth/login

### GET /api/auth/me

（不變）

---

# 📘 Journal（日記 API）

## 1. POST /api/diaries

Body：

```json
{
  "title": "今天心情不錯",
  "content": "寫了很多程式。",
  "tags": ["life"],
  "mood": "happy",
  "date": "2025-02-21"
}
```

---

## 2. GET /api/diaries?keyword=&tags=&mood=&start=&end=

支援搜尋

---

## 3. GET /api/diaries/:id

## 4. PUT /api/diaries/:id

## 5. DELETE /api/diaries/:id

（與舊版相同）

---

# 📗 Todo（代辦 API）⭐新加的

## 1. POST /api/todos

Body：

```json
{
  "content": "寫日記系統後端",
  "date": "2025-02-21"
}
```

回應：

```json
{
  "id": "abc123",
  "content": "寫日記系統後端",
  "done": false,
  "date": "2025-02-21"
}
```

---

## 2. GET /api/todos?date=2025-02-21

回傳：

```json
{
  "items": [
    { "id": "1", "content": "買牛奶", "done": false },
    { "id": "2", "content": "寫Go練習", "done": true }
  ]
}
```

---

## 3. PATCH /api/todos/:id

Body：

```json
{
  "done": true
}
```

或

```json
{
  "content": "重寫代辦"
}
```

---

## 4. DELETE /api/todos/:id

---

# 🌟 六、後端架構（更新）

```
backend/
│── cmd/
│   └── server/main.go
│
│── internal/
│   ├── controllers/
│   │   ├── diary_controller.go
│   │   └── todo_controller.go   # ⭐新加
│   │
│   ├── services/
│   │   ├── diary_service.go
│   │   └── todo_service.go      # ⭐新加
│   │
│   ├── repositories/
│   │   ├── diary_repository.go
│   │   ├── todo_repository.go   # ⭐新加
│   │   └── mongo/
│   │       ├── diary_mongo.go
│   │       └── todo_mongo.go    # ⭐新加
│   │
│   ├── models/
│   │   ├── diary.go
│   │   └── todo.go               # ⭐新加
│   │
│   ├── middleware/
│   └── config/
│
└── go.mod
```

Todo 會遵循同樣 architecture。

---

# 🌟 七、前端（Vue）架構（更新）

```
frontend/
│── src/
│   ├── pages/
│   │   ├── JournalTab.vue
│   │   └── TodoTab.vue     # ⭐新加 
│   │
│   ├── components/
│   │   ├── TodoItem.vue    # ⭐新加
│   │   └── DiaryEditor.vue
│   │
│   ├── api/
│   │   ├── diary.js
│   │   └── todo.js         # ⭐新加
│   │
│   ├── router/
│   └── App.vue  ← 裡面兩個 Tab
```

---

# 🌟 八、功能整合（前後端）

## Journal Tab 功能

* 顯示日期
* 載入該日期的日記（如果有）
* 新增 / 修改 / 刪除日記

## Todo Tab 功能

* 顯示該日期的 Todo 列表
* 新增 Todo
* 刪除 Todo
* 修改 Todo（done / content）

---

# 🌟 九、未來擴充（選做）

* Todo 自動生成「今日摘要」推薦放入日記
* 心情分析（AI）
* 日記打卡統計
* Todo 生產力統計
* Tag cloud（日記標籤雲）
