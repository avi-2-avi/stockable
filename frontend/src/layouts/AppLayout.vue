<template>
  <div class="h-screen flex flex-col items-center">
    <Navbar />
    <div class="flex flex-col mx-auto pt-20 flex-1 container w-full space-y-2">
      <h3>{{ pageTitle }}</h3>
      <div>
        <p>Select the source:</p>
        <SelectDropdown id="source-select" label="Select a data source" v-model="selectedSource" :options="sources"
          position="top" />
      </div>
      <slot />
    </div>
  </div>
</template>

<script setup lang="ts">
import Navbar from '@/components/BaseNavbar.vue';
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
