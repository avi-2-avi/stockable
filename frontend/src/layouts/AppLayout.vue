<template>
  <div class="h-screen flex flex-col items-center">
    <AppNavbar />
    <div class="flex flex-col mx-auto pt-22 flex-1 container w-full space-y-4 px-4">
      <div class="flex flex-col lg:flex-row items-center gap-2.5 justify-between">
        <div class="text-center lg:text-left space-y-0.5">
          <h3>{{ pageTitle }}</h3>
          <p v-if="pageDescription" class="text-lg font-semibold opacity-70">{{ pageDescription }}</p>
        </div>
        <div class="flex space-x-5 items-center">
          <div class="flex w-[200px] items-center space-x-1">
            <SelectDropdown dropdown-class="text-stock-500 font-semibold" id="source-select"
              v-model="selectedSource" :options="sources" position="top" />
          </div>
          <CircleHelp @click="showModal = true"
            class="text-stock-500 hover:scale-110 transition-transform duration-200" />
        </div>
      </div>
      <slot />
    </div>
  </div>
  <SourceModal v-if="showModal" @close="showModal = false" />
</template>

<script setup lang="ts">
import AppNavbar from '@/components/navigation/AppNavbar.vue';
import SelectDropdown from '@/components/ui/SelectDropdown.vue'
import { useSourceStore } from '@/store/sourceStore.ts'
import { storeToRefs } from 'pinia'
import { onMounted, ref } from 'vue';
import { CircleHelp } from 'lucide-vue-next';
import SourceModal from '@/components/ui/SourceModal.vue';

const sourceStore = useSourceStore();
const { sources, selectedSource } = storeToRefs(sourceStore);

defineProps({
  pageTitle: {
    type: String,
    default: 'Stockable',
  },
  pageDescription: {
    type: String,
    optional: true,
  },
})

onMounted(() => {
  sourceStore.fetchSources();
})

const showModal = ref(false);

</script>
