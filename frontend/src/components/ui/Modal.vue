<template>
    <div class="fixed inset-0 flex items-center justify-center bg-black/80 z-[100]">
        <div class="bg-base rounded-lg shadow-lg p-6 w-[400px]">
            <header class="flex justify-between items-center mb-4">
                <slot name="header"></slot>
                <button @click="$emit('close')" class="text-black/70 hover:text-black/80">✕</button>
            </header>
            <div>
                <slot name="body"></slot>
            </div>
            <footer class="mt-4 flex justify-end space-x-2">
                <slot name="footer"></slot>
            </footer>
        </div>
    </div>
</template>

<script setup lang="ts">
defineEmits(["close"]);
</script>