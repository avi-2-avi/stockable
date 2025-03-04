<template>
  <DefaultLayout>
    <p class="headline">Home</p>
    <h1 class="text-stock-500">Home</h1>
    <button @click="toggleDarkMode" class="px-4 py-2 mt-4 bg-stock-500 text-white rounded">
      Toggle Dark Mode
    </button>
    <div>
      <p class="text-2xl font-open-sans">HOLA QUE HACE</p>
      <div class="bg-base h-20 w-20">
      </div>
    </div>
    <SelectDropdown />
    <DataTable />
  </DefaultLayout>
</template>

<script setup lang="ts">
import { ref, onMounted } from "vue";

import SelectDropdown from '@/components/SelectDropdown.vue'
import DataTable from '@/features/app/components/DataTable.vue'
import DefaultLayout from '@/layouts/DefaultLayout.vue'

const isDark = ref(false);

const toggleDarkMode = () => {
  isDark.value = !isDark.value;
  if (isDark.value) {
    document.documentElement.classList.add("dark");
    localStorage.setItem("theme", "dark");
  } else {
    document.documentElement.classList.remove("dark");
    localStorage.setItem("theme", "light");
  }
};

onMounted(() => {
  const savedTheme = localStorage.getItem("theme");
  if (savedTheme === "dark" || (!savedTheme && window.matchMedia("(prefers-color-scheme: dark)").matches)) {
    isDark.value = true;
    document.documentElement.classList.add("dark");
  }
});
</script>
