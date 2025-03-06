<template>
    <div class="border border-border bg-card p-4 rounded-xl shadow-sm" :class="cardClass">
        <slot />
    </div>
</template>

<script setup lang="ts">
defineProps({
    cardClass: {
        type: String,
        default: 'w-full'
    }
})
</script>
