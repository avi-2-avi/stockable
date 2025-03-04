<template>
    <div class="flex flex-col h-screen">
        <LandingNavbar />
        <main class="flex-1">
            <slot />
        </main>
        <LandingFooter />
    </div>
</template>

<script setup lang="ts">
import LandingFooter from '@/features/landing/components/LandingFooter.vue';
import LandingNavbar from '@/features/landing/components/LandingNavbar.vue';

</script>