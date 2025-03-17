import { reactive } from 'vue'
import { defineStore } from 'pinia'
import axios from 'axios'
import type { IndicatorData } from '@/types/indicator'

export const useIndicatorStore = defineStore('indicator', () => {
  const indicators = reactive<IndicatorData>({
    buy_now_percentage: 0,
    positive_target_adjustment_percentage: 0,
    highest_increment_in_target_price: 0,
    highest_increment_in_target_price_name: '',
    highest_increment_in_target_price_ticker: ''
  })

  const cachedIndicators = reactive<IndicatorData>({ ...indicators })

  const fetchIndicators = async (sourceId: string) => {
    try {
      Object.assign(indicators, {
        buy_now_percentage: 0,
        positive_target_adjustment_percentage: 0,
        highest_increment_in_target_price: 0,
        highest_increment_in_target_price_name: '',
        highest_increment_in_target_price_ticker: ''
      })

      const response = await axios.get(
        `${import.meta.env.VITE_API_URL}/api/ratings/indicators?source_id=${sourceId}`,
      )

      if (response.status !== 200) {
        throw new Error('Failed to fetch indicators')
      }

      const newIndicators = response.data.body.indicators

      if(JSON.stringify(newIndicators) !== JSON.stringify(indicators)) {
        Object.assign(indicators, response.data.body.indicators)
        Object.assign(cachedIndicators, response.data.body.indicators)
      }
    } catch (error) {
      console.error('Error fetching indicators', error)
    }
  }

  return { indicators, cachedIndicators, fetchIndicators }
})
