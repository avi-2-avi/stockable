<template>
    <nav class="bg-blue-600 text-white p-4 flex justify-between items-center">
      <!-- Logo on the Left -->
      <div class="flex items-center">
        <img :src="logoSrc" alt="Logo" class="h-10 w-auto" />
      </div>

      <!-- Slot for Custom Components on the Right -->
      <div class="flex items-center space-x-4">
        <slot />
      </div>
    </nav>
</template>

<script setup lang="ts">
defineProps<{ logoSrc: string }>();
</script>
