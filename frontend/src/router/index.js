import { createRouter, createWebHistory } from "vue-router";
import Home from "@/components/Home.vue";
import Journal from "@/components/Journal.vue";
import ToDoList from "@/components/ToDoList.vue"

const routes = [
  { path: "/", name: "Home", component: Home },
  { path: "/journal", name: "Journal", component: Journal },
  { path: "/todolist", name: "ToDoList", component: ToDoList },
  
];

const router = createRouter({
  history: createWebHistory(),
  routes: routes,
});

export default router;
