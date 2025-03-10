<template>
    <div class="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50">
        <div class="bg-white rounded-lg shadow-lg p-6 w-[400px]">
            <header class="flex justify-between items-center mb-4">
                <slot name="header"></slot>
                <button @click="$emit('close')" class="text-gray-500 hover:text-gray-700">✕</button>
            </header>
            <div>
                <slot name="body"></slot>
            </div>
            <footer class="mt-4 flex justify-end">
                <slot name="footer"></slot>
            </footer>
        </div>
    </div>
</template>

<script setup lang="ts">
defineEmits(["close"]);
</script>