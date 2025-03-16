import { ref, watch } from 'vue'
import type { Option } from '@/types/option.ts'
import { defineStore } from 'pinia'
import axios from 'axios'
import { useIndicatorStore } from './indicatorStore'

interface SourceResponse {
  id: string
  name: string
}

export const useSourceStore = defineStore('source', () => {
  const sources = ref<Option[]>([])
  const selectedSource = ref<string>('')
  const indicatorStore = useIndicatorStore()

  const fetchSources = async () => {
    try {
      const response = await axios.get(`${import.meta.env.VITE_API_URL}/api/sources`)

      if (response.status !== 200) {
        throw new Error('Failed to fetch sources')
      }

      sources.value = response.data.body.map((source: SourceResponse) => ({
        value: source.id.toString(),
        label: source.name,
      }))

      if (!selectedSource.value && sources.value.length) {
        selectedSource.value =
          sources.value.find((source) => source.label === 'TruAdapter')?.value ||
          sources.value[0].value
      }
    } catch (error) {
      console.error('Error fetching sources', error)
    }
  }

  const setSelectedSource = (source_id: string) => {
    selectedSource.value = source_id
    indicatorStore.fetchIndicators(source_id)
  }

  watch(selectedSource, async (newSource) => {
    if (newSource) {
      await indicatorStore.fetchIndicators(newSource)
    }
  })

  return { sources, selectedSource, fetchSources, setSelectedSource }
})
