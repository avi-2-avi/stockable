<template>
  <AppLayout pageTitle="Home">
    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4 mt-2.5">
      <HomeCard v-for="(card, index) in cardInformation" :key="index" :title="card.title" :value="card.value"
        :description="card.description" :icon="card.icon" />
    </div>
    <RatingTable />
  </AppLayout>
</template>

<script setup lang="ts">
import HomeCard from '@/features/app/components/HomeCard.vue'
import RatingTable from '@/features/app/components/RatingTable.vue'
import AppLayout from '@/layouts/AppLayout.vue'
import { ChartNoAxesCombined, CircleDollarSign, TrendingUp } from 'lucide-vue-next';

const cardInformation = [
  {
    title: "Average Increment",
    value: "7%",
    description: "This is the average increment in the last 24 hours",
    icon: TrendingUp
  },
  {
    title: "Highest Increment",
    value: "10",
    description: "This is the highest increment in the last 24 hours",
    icon: ChartNoAxesCombined
  },
  {
    title: "Lowest Increment",
    value: "5",
    description: "This is the lowest increment in the last 24 hours",
    icon: CircleDollarSign
  },
]
</script>
