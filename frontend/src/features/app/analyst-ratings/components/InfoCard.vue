<template>
    <Card card-class="bg-stock-200/20 border border-stock-500/50">
        <div class="flex items-center space-x-5">
            <div class="min-w-12 min-h-12 bg-stock-500 rounded-full flex items-center justify-center">
                <component :is="icon" class="w-6 h-6 text-white" />
            </div>
            <div>
                <p class="font-outfit text-lg font-medium">{{ title }}</p>
                <p class="font-outfit text-3xl font-medium">{{ value }}</p>
                <small>{{ description }}</small>
            </div>

        </div>
    </Card>
</template>

<script setup lang="ts">
import Card from '@/components/ui/Card.vue'

defineProps({
    title: String,
    value: String,
    description: String,
    icon: Function
})
</script>