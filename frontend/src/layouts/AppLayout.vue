<template>
    <div class="h-screen flex flex-col">
        <Navbar />
        <div class="flex-1 container mx-auto p-4">
            <slot />
        </div>
    </div>
</template>

<script setup lang="ts">
import Navbar from '@/components/BaseNavbar.vue';

</script>
