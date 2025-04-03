<template>
    <Modal @close="$emit('close')">
        <template #header>
            <h3 class="text-lg font-bold text-red-600">Authentication Error</h3>
        </template>

        <template #body>
            <p>{{ message }}. Please try again.</p>
        </template>

        <template #footer>
            <Button @click="$emit('close')" label="Close" variant="solid" class="w-full" />
        </template>
    </Modal>
</template>

<script setup lang="ts">
import Modal from '@/components/ui/Modal.vue';
import Button from '@/components/ui/Button.vue';

defineProps<{ message: string }>();
</script>
