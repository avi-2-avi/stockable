<template>
    <button @click="handleClick" :class="buttonClasses" :disabled="disabled"
        class="cursor-pointer flex items-center font-semibold justify-center rounded-md transition-all focus:outline-none">
        <slot name="icon-start"></slot>
        <span class="mx-1">{{ label }}</span>
        <slot name="icon-end"></slot>
    </button>
</template>

<script setup lang="ts">
import { computed, type PropType } from 'vue'

const sizes = ['xs', 'sm', 'md', 'lg', 'xl'] as const;
const variants = ['solid', 'subtle', 'surface', 'outline', 'ghost', 'plain', 'warn'] as const;

const props = defineProps({
    label: String,
    size: {
        type: String as PropType<typeof sizes[number]>,
        default: 'md'
    },
    variant: {
        type: String as PropType<typeof variants[number]>,
        default: 'solid'
    },
    disabled: Boolean,
})

const emit = defineEmits(["click"]);
const handleClick = (event: Event) => {
    if (!props.disabled) {
        emit("click", event);
    }
};

const sizeClasses: Record<typeof sizes[number], string> = {
    xs: 'px-2 py-1 text-xs',
    sm: 'px-3 py-1 text-sm',
    md: 'px-4 py-2 text-sm',
    lg: 'px-5 py-3 text-lg',
    xl: 'px-6 py-4 text-xl',
}

const variantClasses: Record<typeof variants[number], string> = {
    solid: 'bg-stock-500 text-white hover:bg-stock-500',
    subtle: 'bg-stock-200 hover:bg-stock-300 text-black',
    surface: 'bg-stock-200 border border-border text-black hover:bg-stock-300',
    outline: 'border border-border text-foreground hover:bg-stock-200',
    ghost: 'text-foreground hover:bg-stock-500/30',
    plain: 'text-foreground',
    warn: 'bg-stock-600 text-white hover:bg-stock-600',
}

const buttonClasses = computed(() => {
    return [
        'flex items-center justify-center rounded-md transition-all',
        sizeClasses[props.size],
        variantClasses[props.variant],
        props.disabled ? 'opacity-50 cursor-not-allowed' : '',
    ]
})

</script>