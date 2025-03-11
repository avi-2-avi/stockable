<template>
    <div class="grid grid-cols-1 md:grid-cols-2 gap-8 md:gap-20 items-center">
        <div :class="reverse ? 'md:order-2' : ''" class="flex flex-col space-y-5 text-center md:text-left">
            <h2 class="text-2xl font-bold">
                <span class="text-stock-400">{{ highlight }}</span> {{ title }}
            </h2>
            <p class="text-lg text-foreground/80">{{ description }}</p>
        </div>
        <div class="flex justify-center items-center">
            <img :src="image" :alt="title" class="w-80 md:w-96" />
        </div>
    </div>
</template>

<script setup lang="ts">
defineProps({
    title: String,
    highlight: String,
    description: String,
    image: String,
    reverse: Boolean
});
</script>