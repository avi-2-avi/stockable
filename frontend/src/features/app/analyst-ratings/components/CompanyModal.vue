<template>
    <Modal @close="$emit('close')">
        <template #header>
            <h3 class="text-lg font-bold text-red-600">{{ company }} ({{ ticker }})</h3>
        </template>

        <template #body class="mt-2">
            <p v-if="loading">Loading company information...</p>
            <p v-else>{{ description }}</p>
        </template>

        <template #footer>
            <Button @click="$emit('close')" label="Close" variant="solid" class="w-full" />
        </template>
    </Modal>
</template>

<script setup lang="ts">
import Modal from '@/components/ui/Modal.vue';
import Button from '@/components/ui/Button.vue';
import { useCompanyInfoStore } from '@/store/companyInfoStore';
import { onMounted, ref } from 'vue';

defineEmits(["close"]);
const props = defineProps<{
    ticker: string;
    company: string;
}>();

const companyInfoStore = useCompanyInfoStore();
const description = ref('');
const loading = ref(true);

onMounted(async () => {
    description.value = await companyInfoStore.fetchCompanyDescription(props.ticker, props.company);
    loading.value = false;
});
</script>
