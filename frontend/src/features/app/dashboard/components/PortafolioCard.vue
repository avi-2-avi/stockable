<script setup lang="ts">
import Button from '@/components/ui/Button.vue';
import Card from '@/components/ui/Card.vue';
import { ChevronRight } from 'lucide-vue-next';

</script>

<template>
    <Card class="grid grid-cols-2 gap-5 p-5">
        <div class="justify-between">
            <p>Portafolio Value</p>
            <h3><span class="text-xl">$</span> 14,200.00</h3>
        </div>
        <div class="flex flex-col items-end justify-between">
            <Button label="Details" size="md" variant="solid">
                <template #icon-end>
                    <ChevronRight :size="20" />
                </template>
            </Button>
            <div>
                <small>Last updated: Today</small>
            </div>
        </div>
    </Card>
</template>