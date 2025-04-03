<template>
    <div class="p-2 border-b border-border last:border-b-0">
        <div class="flex space-x-2.5 pb-2">
            <ClockIcon :size="24" class="text-stock-500 mt-0.5 p-1 bg-stock-500/20 rounded-full" />
            <div>
                <p class="text-lg font-semibold">{{ rating.company }}</p>
                <div class="flex justify-between items-center space-x-4 mt-1">
                    <CPIBadge size="small" class="rounded-full" :value="rating.combined_prediction_index" />
                    <small class="text-xs font-medium text-foreground/70">Rated on {{ formatStringToDate(rating.rated_at) }} at {{ formatStringToTime(rating.rated_at) }}</small>
                </div>
            </div>
        </div>
    </div>
</template>

<script lang="ts" setup>
import { ClockIcon } from 'lucide-vue-next';
import CPIBadge from '@/components/ui/CPIBadge.vue';
import type { Rating } from '@/types/rating';
import { formatStringToDate, formatStringToTime } from '@/utils/formater';

const { rating } = defineProps<{ rating: Rating }>()
</script>
