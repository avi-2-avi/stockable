<template>
    <div class="flex flex-col h-screen">
        <BaseNavbar />
        <main class="flex-1 container mx-auto">
            <slot />
        </main>
        <LandingFooter />
    </div>
</template>

<script setup lang="ts">
import BaseNavbar from '@/components/navigation/BaseNavbar.vue';
import LandingFooter from '@/features/landing/components/LandingFooter.vue';
</script>