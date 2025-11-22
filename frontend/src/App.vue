<script setup>
import HelloWorld from './components/ToDoList.vue'
import { ref, onMounted } from 'vue'

// TODO: 搞懂 local storage
// 下一步，先把兩邊串起來，把後端的東西先寫好
// TODO: 下一步，設計login頁面，使用google

const theme = ref(localStorage.getItem('theme') || 'light')

function applyTheme() {
  document.documentElement.setAttribute('data-theme', theme.value)
}

function toggleTheme() {
  theme.value = theme.value === 'light' ? 'dark' : 'light'
  applyTheme()
  localStorage.setItem('theme', theme.value)
}

onMounted(() => {
  applyTheme()
})
</script>

<template>
  <div class="app-root">
    <div class="nav">
      <div class="nav-inner">
        <router-link to="/">Home - My Journal</router-link>
        <span class="divider">|</span>
        <router-link to="/journal">Journal</router-link>
        <span class="divider">|</span>
        <router-link to="/todolist">To Do List</router-link>
        <button
          class="switch-appearance"
          @click="toggleTheme"
          :aria-label="theme === 'light' ? '啟用深色模式' : '啟用淺色模式'"
          title="切換主題"
        >
          <span v-if="theme === 'light'">🌙</span>
          <span v-else>☀️</span>
        </button>
      </div>
    </div>

    <main class="view-root">
      <router-view></router-view>
    </main>

    <footer class="site-footer">
      <div class="footer-inner">
        © 2025 My Journal
      </div>
    </footer>
  </div>

  <!-- <HelloWorld msg="Vite + Vue" /> -->
</template>

<style scoped>
.app-root {
  display: flex;
  flex-direction: column;
  align-items: stretch;
  min-height: 100vh;
  gap: 16px;
  padding: 0; /* remove page padding so nav/footer can be full-bleed */
  box-sizing: border-box;
}

.view-root {
  flex: 1 1 auto;
  max-width: 980px;
  margin: 0 auto;
  width: 100%;
  padding: 18px 12px; /* keep content comfortable from edges */
}

/* Full-bleed navbar */
.nav {
  width: 100%;
  padding: 12px 0; /* vertical padding for full-width band */
  background: linear-gradient(90deg, #e6f7ff 0%, #e6fff2 100%); /* 淡藍到淡綠 */
  box-shadow: 0 4px 14px rgba(16,24,40,0.04);
}

.nav-inner {
  max-width: 980px;
  margin: 0 auto;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 0.5rem;
  padding: 0 14px;
  font-weight: 600;
  font-size: 14px;
}

.nav a {
  color: #0b63b7;
  text-decoration: none;
  padding: 6px 10px;
  border-radius: 8px;
  transition: background-color 160ms ease, color 160ms ease, transform 120ms ease;
}

.nav a:hover {
  background: rgba(11,99,183,0.08);
  color: #044b86;
  transform: translateY(-2px);
}

.nav a.router-link-active {
  background: rgba(34,197,94,0.12); /* 淡綠的 active */
  color: #14532d;
  box-shadow: inset 0 -3px 0 rgba(34,197,94,0.14);
}

.divider {
  color: rgba(3,37,65,0.35);
  padding: 0 6px;
}

.switch-appearance {
  margin-left: 8px;
  background: transparent;
  border: 0;
  padding: 6px 8px;
  border-radius: 8px;
  cursor: pointer;
  font-size: 14px;
  line-height: 1;
  transition: background-color 140ms ease, transform 120ms ease;
}

.switch-appearance:hover {
  background: rgba(11,99,183,0.06);
  transform: translateY(-2px);
}

/* Footer full-bleed */
.site-footer {
  width: 100%;
  padding: 12px 0;
  background: linear-gradient(90deg, #f0fcf8 0%, #f2f9ff 100%);
  border-top: 1px solid rgba(3,37,65,0.04);
}

.footer-inner {
  max-width: 980px;
  margin: 0 auto;
  padding: 0 14px;
  text-align: center;
  color: rgba(3,37,65,0.7);
  font-size: 13px;
}

/* 小螢幕調整 */
@media (max-width: 520px) {
  .nav-inner {
    flex-wrap: wrap;
    gap: 6px 10px;
    padding: 0 10px;
  }

  .divider { display: none; }
  .view-root { padding: 12px 10px; }
  .footer-inner { font-size: 12px; }
}
</style>

/* Global (non-scoped) styles for full-bleed reliability and dark theme overrides */
<style>
/* ensure no default margin prevents full-bleed */
html, body, #app {
  height: 100%;
  margin: 0;
}

/* Dark theme overrides */
[data-theme="dark"] {
  background-color: #071923;
  color: #dbeee6;
}

[data-theme="dark"] .nav {
  background: linear-gradient(90deg,#042035 0%, #02292a 100%);
  box-shadow: 0 4px 18px rgba(0,0,0,0.6);
}

[data-theme="dark"] .nav a {
  color: #9ad4ff;
}

[data-theme="dark"] .nav a:hover {
  background: rgba(154,212,255,0.06);
  color: #cceeff;
}

[data-theme="dark"] .nav a.router-link-active {
  background: rgba(16,185,129,0.12);
  color: #a7f3d0;
  box-shadow: inset 0 -3px 0 rgba(16,185,129,0.18);
}

[data-theme="dark"] .site-footer {
  background: linear-gradient(90deg,#021617 0%, #042035 100%);
  border-top-color: rgba(255,255,255,0.03);
}

[data-theme="dark"] .footer-inner { color: rgba(255,255,255,0.75); }
</style>
