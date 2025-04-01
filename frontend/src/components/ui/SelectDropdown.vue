<template>
  <div :class="['relative', customClass]">
    <label v-if="label" class="block text-left text-sm font-medium mb-1">
      {{ label }}
    </label>
    <div
      :class="['px-2 py-2 text-sm border border-border rounded-md bg-card dark:text-white flex justify-between items-center cursor-pointer', dropdownClass]"
      @click="toggleDropdown">
      {{ selectedLabel }}
      <ChevronDown class="w-4 h-4 text-foreground" />
    </div>

    <ul v-if="isOpen"
      :class="['absolute w-full bg-card border border-border rounded-md shadow-md z-50', dropdownPositionClass]">
      <li v-for="option in options" :key="option.value" @click="selectOption(option.value)"
        class="text-sm p-2 hover:bg-stock-200/50 cursor-pointer flex items-center">
        <span>{{ option.label }}</span>
      </li>
    </ul>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { ChevronDown } from "lucide-vue-next";
import type { Option } from '@/types/option.ts'

const props = defineProps({
  id: {
    type: String,
    default: 'select',
  },
  options: {
    type: Array as () => Option[],
    required: true,
  },
  modelValue: {
    type: String,
    required: true,
  },
  placeholder: {
    type: String,
    default: 'Select an option',
  },
  customClass: {
    type: String,
    default: 'w-full max-w-lg',
  },
  dropdownClass: {
    type: String,
    default: 'w-full max-w-lg',
  },
  position: {
    type: String as () => "top" | "bottom",
    default: "bottom",
  },
  label: {
    type: String,
    optional: true,
  },
})

const emit = defineEmits<{
  (e: 'update:modelValue', value: string): void
}>()

const isOpen = ref(false);
const selectedLabel = computed(() => props.options.find(o => o.value === props.modelValue)?.label || "Select an option");

const dropdownPositionClass = computed(() => props.position === "bottom" ? "bottom-full mb-1" : "top-full mt-1");

const toggleDropdown = () => {
  isOpen.value = !isOpen.value;
};

const selectOption = (value: string) => {
  emit('update:modelValue', value);
  isOpen.value = false;
};

</script>
