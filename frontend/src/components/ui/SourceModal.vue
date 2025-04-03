<template>
    <Modal @close="$emit('close')">
        <template #header>
            <h3 class="text-lg font-bold text-red-600">About Data Sources</h3>
        </template>

        <template #body>
            <p>Data sources are the data providers that provide the analyst ratings and stock price data used in the calculations.</p>
            <p>You can select the data source from the dropdown menu in the top right corner of the page.</p>
        </template>

        <template #footer>
            <Button @click="$emit('close')" label="Close" variant="solid" class="w-full" />
        </template>
    </Modal>
</template>

<script setup lang="ts">
import Modal from '@/components/ui/Modal.vue';
import Button from '@/components/ui/Button.vue';

defineEmits(["close"]);
</script>
