<script setup>
import ToDoCard from "./ToDoCard.vue";
import { ref } from "vue";

defineEmits(["deleteTodo"]);

const porps = defineProps({
  todos: Array,
});
const filters = [
  { key: "All", value: "All" },
  { key: "Pending", value: "Pending" },
  { key: "In Progress", value: "In Progress" },
  { key: "Done", value: "Done" },
];
const activeFilter = ref("All");

// TODO
const filterTodos = ref([]);

async function deleteTodo(id) {
  try {
    const res = await axios.delete(`/api/todo/${id}`);
    // reload page
    router.go();
  } catch (err) {
    console.log("err: ", err);
  }
}
</script>

<template>
  <div class="container">
    <button
      @click="activeFilter.value = filter.key"
      class="filter-tab"
      v-for="filter in filters"
      :key="filter.key"
    >
      {{ filter.value }}
    </button>

    <!-- TODO -->
    <div v-for="(todo, index) in filterTodos" :key="index">
      <ToDoCard :todo="todo" @deleteTodo="deleteTodo" />
    </div>
  </div>
</template>
