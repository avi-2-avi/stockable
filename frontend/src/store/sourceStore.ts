import { ref } from 'vue'
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
      const response = await axios.get(import.meta.env.VITE_API_URL + '/sources')

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

      if (selectedSource.value) {
        await indicatorStore.fetchIndicators(selectedSource.value)
      }
    } catch (error) {
      console.error('Error fetching sources', error)
    }
  }

  const setSelectedSource = (source_id: string) => {
    selectedSource.value = source_id
    indicatorStore.fetchIndicators(source_id)
  }

  return { sources, selectedSource, fetchSources, setSelectedSource }
})
