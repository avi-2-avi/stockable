import { defineStore } from 'pinia'
import { computed, ref, watch } from 'vue'
import { useSourceStore } from './sourceStore'
import axios from 'axios'
import type { Rating } from '@/types/rating'

export const useRatingStore = defineStore('rating', () => {
  const ratings = ref<Rating[]>([])
  const totalRatings = ref<number>(0)
  const loading = ref(false)
  const error = ref<string | null>(null)

  const page = ref(1)
  const limit = ref(10)
  const sortBy = ref<string>('rated_at')
  const sortOrder = ref<'asc' | 'desc'>('desc')
  const filters = ref<Record<string, string>>({})

  const sourceStore = useSourceStore()
  const selectedSource = computed(() => sourceStore.selectedSource)

  watch(
    selectedSource,
    (newSource) => {
      if (newSource) {
        fetchRatings()
      }
    },
    { immediate: true },
  )

  const fetchRatings = async () => {
    if (!selectedSource.value) {
      console.warn('No source selected. Skipping fetchRatings.')
      return
    }

    loading.value = true
    error.value = null

    try {
      const response = await axios.get(import.meta.env.VITE_API_URL + '/ratings', {
        params: {
          source_id: selectedSource.value,
          page: page.value,
          limit: limit.value,
          sort_by: sortBy.value,
          sort_order: sortOrder.value,
          ...filters.value,
        },
      })
      ratings.value = response.data.body.ratings
      totalRatings.value = response.data.body.total
    } catch (err) {
      console.error('Error fetching ratings', err)
      error.value = 'Error fetching ratings'
    } finally {
      loading.value = false
    }
  }

  const setSorting = (column: string, order: 'asc' | 'desc') => {
    sortBy.value = column
    sortOrder.value = order
    fetchRatings()
  }

  const setFilter = (key: string, value: string) => {
    if (value) {
      filters.value[key] = value
    } else {
      delete filters.value[key]
    }
    fetchRatings()
  }

  const setPage = (newPage: number) => {
    page.value = newPage
    fetchRatings()
  }

  const setLimit = (newLimit: number) => {
    limit.value = newLimit
    fetchRatings()
  }

  return {
    ratings,
    totalRatings,
    loading,
    error,
    page,
    limit,
    sortBy,
    sortOrder,
    filters,
    fetchRatings,
    setSorting,
    setFilter,
    setPage,
    setLimit,
  }
})
