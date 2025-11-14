import { createRouter, createWebHistory } from "vue-router";
import Home from "@/components/Home.vue";
import About from "@/components/About.vue";
import ToDoList from "@/components/ToDoList.vue"

const routes = [
  { path: "/", name: "Home", component: Home },
  { path: "/about", name: "About", component: About },
  { path: "/todolist", name: "ToDoList", component: ToDoList },
  
];

const router = createRouter({
  history: createWebHistory(),
  routes: routes,
});

export default router;
