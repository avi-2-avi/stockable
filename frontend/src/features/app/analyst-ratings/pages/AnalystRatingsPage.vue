<template>
  <AppLayout pageTitle="Home">
    <div class="grid grid-cols-1 lg:grid-cols-3 gap-4">
      <HomeCard v-for="(indicator, index) in indicatorInformation" :key="index" :title="indicator.title"
        :value="indicator.value" :description="indicator.description" :icon="indicator.icon" />
    </div>
    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4">
      <CalculationsCard :card-class="'col-span-1 lg:col-span-2'" @open-modal="showModal = true" />
      <ExcelCard :table-data="ratings" :file-name="'Ratings'" :card-class="'col-span-1'" />
    </div>
    <RatingTable />
    <CalculationsModal v-if="showModal" @close="showModal = false" />
  </AppLayout>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue';
import CalculationsModal from '@/features/app/analyst-ratings/components/CalculationsModal.vue'
import HomeCard from '@/features/app/analyst-ratings/components/InfoCard.vue'
import CalculationsCard from '@/features/app/analyst-ratings/components/CalculationsCard.vue'
import ExcelCard from '../components/ExcelCard.vue';
import RatingTable from '@/features/app/analyst-ratings/components/RatingTable.vue'
import AppLayout from '@/layouts/AppLayout.vue'
import { ChartNoAxesCombined, CircleDollarSign, TrendingUp } from 'lucide-vue-next';
import { useIndicatorStore } from '@/store/indicatorStore';
import { formatNumberToPercentage, formatNumberToCurrency } from '@/utils/formater';
import { useRatingStore } from '@/store/ratingStore';
import { storeToRefs } from 'pinia';

const showModal = ref(false);

const ratingStore = useRatingStore();
const { ratings } = storeToRefs(ratingStore);

const {cachedIndicators} = useIndicatorStore()
const indicatorInformation = computed(() => [
  {
    title: "Buy Now Percentage",
    value: cachedIndicators
      ? `${formatNumberToPercentage(cachedIndicators.buy_now_percentage)}`
      : "Loading...",
    description: "Percentage of ratings indicating a strong buy signal.",
    icon: CircleDollarSign
  },
  {
    title: "Positive Target Adjustment",
    value: cachedIndicators
      ? `${formatNumberToPercentage(cachedIndicators.positive_target_adjustment_percentage)}`
      : "Loading...",
    description: "Percentage of ratings with an increased target price.",
    icon: ChartNoAxesCombined
  },
  {
    title: "Highest Rating Increment",
    value: cachedIndicators
      ? `${formatNumberToCurrency(cachedIndicators.highest_increment_in_target_price)}`
      : "Loading...",
    description: cachedIndicators
      ? `Highest $ increase stock: ${cachedIndicators.highest_increment_in_target_price_name} (${cachedIndicators.highest_increment_in_target_price_ticker})`
      : "Loading...",
    icon: TrendingUp
  }
]);
</script>
