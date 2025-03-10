<template>
  <div class="h-screen flex flex-col items-center">
    <AppNavbar />
    <div class="flex flex-col mx-auto pt-22 flex-1 container w-full space-y-4 px-4">
      <Card card-class="flex items-center space-x-4">
        <p class="font-medium">Select data source:</p>
        <SelectDropdown id="source-select" v-model="selectedSource" :options="sources" position="top" />
      </Card>
      <slot />
    </div>
  </div>
</template>

<script setup lang="ts">
import AppNavbar from '@/features/app/components/AppNavbar.vue';
import Card from '@/components/Card.vue';
import SelectDropdown from '@/components/SelectDropdown.vue'
import { useSourceStore } from '@/store/sourceStore.ts'
import { storeToRefs } from 'pinia'
import { onMounted } from 'vue';

const sourceStore = useSourceStore();
const { sources, selectedSource } = storeToRefs(sourceStore);

defineProps({
  pageTitle: {
    type: String,
    default: 'Stockable',
  }
})

onMounted(() => {
  sourceStore.fetchSources();
})

</script>
