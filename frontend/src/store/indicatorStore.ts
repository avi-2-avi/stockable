import { ref } from 'vue'
import { defineStore } from 'pinia'
import axios from 'axios'
import type { IndicatorData } from '@/types/indicator'

export const useIndicatorStore = defineStore('indicator', () => {
  const indicators = ref<IndicatorData | null>(null)

  const fetchIndicators = async (sourceId: string) => {
    console.log(`Fetching indicators for source: ${sourceId}`) // Debugging
    try {
      const response = await axios.get(
        `${import.meta.env.VITE_API_URL}/ratings/indicators?source_id=${sourceId}`,
      )

      if (response.status !== 200) {
        throw new Error('Failed to fetch indicators')
      }

      console.log('Indicators fetched successfully:', response.data.body.indicators) // Debugging
      indicators.value = response.data.body.indicators
    } catch (error) {
      console.error('Error fetching indicators', error)
    }
  }
  return { indicators, fetchIndicators }
})
