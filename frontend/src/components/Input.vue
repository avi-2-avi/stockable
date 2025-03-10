<template>
    <div class="w-full">
        <label v-if="label" class="block text-left text-sm font-medium mb-1">
            {{ label }}
        </label>
        <div class="relative">
            <input :type="type" :placeholder="placeholder" v-model="model" class="w-full px-2 py-2 border rounded-md focus:outline-none transition-all
                border-border focus:border-stock-300 focus:ring-1 focus:ring-stock-300"
                :class="{ 'border-red-400': errorMessage }" />
            <p v-if="errorMessage" class="text-red-400 text-sm text-left mt-1">{{ errorMessage }}</p>
        </div>
    </div>
</template>

<script setup lang="ts">
import { defineProps, defineModel } from 'vue';

defineProps({
    label: {
        type: String,
        optional: true
    },
    type: {
        type: String,
        default: 'text'
    },
    placeholder: String,
    errorMessage: String
});

const model = defineModel<string>();
</script>
