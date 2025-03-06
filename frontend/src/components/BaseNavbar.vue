<template>
  <nav class="w-full fixed top-0 left-0 z-50 bg-stock-200/50 backdrop-blur-lg p-4 flex justify-between items-center">
    <div class="flex items-center">
      <div @click="redirectToDefault" class="flex items-center cursor-pointer">
        <img src="@/assets/logo.svg" alt="Logo" class="h-10 w-auto" />
        <p class="logo pl-2.5">Stockable</p>
      </div>
    </div>

    <div class="flex items-center space-x-4">
      <slot />

      <button @click="toggleTheme" class="p-2 rounded transition text-foreground hidden sm:block">
        <Sun v-if="isDark" class="w-6 h-6" />
        <Moon v-else class="w-6 h-6" />
      </button>
    </div>
  </nav>
</template>

<script setup lang="ts">
import { ref, onMounted } from "vue";
import { Sun, Moon } from "lucide-vue-next";
import { useRouter } from "vue-router";

const router = useRouter();

const isDark = ref(false);

const toggleTheme = () => {
  isDark.value = !isDark.value;
  document.documentElement.classList.toggle("dark", isDark.value);
  localStorage.setItem("theme", isDark.value ? "dark" : "light");
}

onMounted(() => {
  const savedTheme = localStorage.getItem("theme");
  if (savedTheme === "dark" || (!savedTheme && window.matchMedia("(prefers-color-scheme: dark)").matches)) {
    isDark.value = true;
    document.documentElement.classList.add("dark");
  }
})

const redirectToDefault = () => {
  router.push('/');
}

</script>
