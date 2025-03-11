<template>
  <AppLayout pageTitle="Home">
    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4">
      <HomeCard v-for="(indicator, index) in indicatorInformation" :key="index" :title="indicator.title"
        :value="indicator.value" :description="indicator.description" :icon="indicator.icon" />
    </div>
    <CalculationsCard @open-modal="showModal = true" />
    <RatingTable />
    <CalculationsModal v-if="showModal" @close="showModal = false" />
  </AppLayout>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue';
import CalculationsCard from '@/features/app/components/CalculationsCard.vue'
import CalculationsModal from '@/features/app/components/CalculationsModal.vue'
import HomeCard from '@/features/app/components/HomeCard.vue'
import RatingTable from '@/features/app/components/RatingTable.vue'
import AppLayout from '@/layouts/AppLayout.vue'
import { ChartNoAxesCombined, CircleDollarSign, TrendingUp } from 'lucide-vue-next';
import { useIndicatorStore } from '@/store/indicatorStore';
import { formatNumberToPercentage, formatNumberToCurrency } from '@/utils/formater';

const showModal = ref(false);

const indicatorStore = useIndicatorStore()
const indicatorInformation = computed(() => [
  {
    title: "Buy Now Percentage",
    value: indicatorStore.indicators
      ? `${formatNumberToPercentage(indicatorStore.indicators.buy_now_percentage)}`
      : "Loading...",
    description: "Percentage of ratings indicating a strong buy signal.",
    icon: CircleDollarSign
  },
  {
    title: "Positive Target Adjustment",
    value: indicatorStore.indicators
      ? `${formatNumberToPercentage(indicatorStore.indicators.positive_target_adjustment_percentage)}`
      : "Loading...",
    description: "Percentage of ratings with an increased target price.",
    icon: ChartNoAxesCombined
  },
  {
    title: "Highest Rating Increment",
    value: indicatorStore.indicators
      ? `${formatNumberToCurrency(indicatorStore.indicators.highest_increment_in_target_price)}`
      : "Loading...",
    description: indicatorStore.indicators
      ? `Highest $ increase stock: ${indicatorStore.indicators.highest_increment_in_target_price_name} (${indicatorStore.indicators.highest_increment_in_target_price_ticker})`
      : "Loading...",
    icon: TrendingUp
  }
]);
</script>
