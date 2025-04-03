<template>
    <div :class="[badgeColor, textColor, sizeClass]" class="px-2 py-1 min-w-[80px] text-sm font-bold text-center rounded-lg shadow">
        {{ categoryLabel }}
    </div>
</template>

<script setup lang="ts">
import { computed } from 'vue';

const props = defineProps<{ value: number; size?: 'normal' | 'small' }>();

const categories = [
    { min: 90, label: "+ Strong Buy", color: "bg-green-700", text: "text-white" },
    { min: 80, label: "Strong Buy", color: "bg-green-600", text: "text-white" },
    { min: 70, label: "Buy", color: "bg-green-300", text: "text-green-800" },
    { min: 60, label: "+ Hold", color: "bg-blue-500", text: "text-white" },
    { min: 50, label: "Hold", color: "bg-gray-500", text: "text-white" },
    { min: 40, label: "- Hold", color: "bg-yellow-400", text: "text-black" },
    { min: 30, label: "Sell", color: "bg-orange-500", text: "text-white" },
    { min: 20, label: "Strong Sell", color: "bg-red-500", text: "text-white" },
    { min: 0, label: "Avoid", color: "bg-red-700", text: "text-white" }
];

const safeValue = computed(() => isNaN(props.value) ? 0 : props.value);

const categoryData = computed(() => {
    return categories.find(c => safeValue.value >= c.min) || { label: "Unknown", color: "bg-gray-400", text: "text-white" };
});

const categoryLabel = computed(() => categoryData.value.label);
const badgeColor = computed(() => categoryData.value.color);
const textColor = computed(() => categoryData.value.text);

const sizeClass = computed(() => {
    return props.size === 'small' ? 'text-xs py-0.5 px-0.5 min-w-[60px]' : 'text-sm py-1 px-2 min-w-[80px]';
});
</script>
