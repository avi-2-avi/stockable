<template>
    <Card :cardClass="'cursor-pointer hover:bg-black/5 transition ' + cardClass" @click="$emit('open-modal')">
        <div class="flex items-center space-x-4">
            <div class="p-2.5 bg-stock-500 rounded-full">
                <CalculatorIcon class="w-6 h-6 text-base" />
            </div>
            <div>
                <p class="text-lg font-medium font-outfit">Curious about CPI?</p>
                <p class="text-sm text-foreground/80">Click here to learn how we calculate the CPI.</p>
            </div>
        </div>
    </Card>
</template>

<script setup lang="ts">
import { Calculator as CalculatorIcon } from 'lucide-vue-next';
import Card from '@/components/Card.vue';

defineProps({
    cardClass: {
        type: String,
        default: ''
    }
});

defineEmits(["open-modal"]);
</script>