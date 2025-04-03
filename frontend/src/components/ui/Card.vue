<template>
    <div :class="['w-full p-4 rounded-xl', cardClass, variantClasses]">
        <slot />
    </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'

const props = defineProps({
    cardClass: {
        type: String,
        default: ''
    },
    variant: {
        type: String,
        default: 'default',
        validator: (value: string) => ['default', 'filled', 'text'].includes(value)
    }
})

const variantClasses = computed(() => {
    switch (props.variant) {
        case 'filled':
            return 'bg-stock-200/20 border border-stock-500/50 shadow-sm'
        case 'text':
            return ''
        default:
            return 'border border-border bg-card shadow-sm'
    }
})
</script>