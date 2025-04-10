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
  const sortBy = ref<string>('combined_prediction_index')
  const sortOrder = ref<'asc' | 'desc'>('desc')
  const filters = ref<Record<string, string>>({})

  const sourceStore = useSourceStore()
  const selectedSource = computed(() => sourceStore.selectedSource)

  const fetchRatings = async () => {
    if (!selectedSource.value) {
      console.warn('No source selected. Skipping fetchRatings.')
      return
    }

    loading.value = true
    error.value = null

    try {
      const token = localStorage.getItem('auth_token');
      if (!token) {
        throw new Error('No authentication token found');
      }

      const response = await axios.get(`${import.meta.env.VITE_API_URL}/api/ratings`, {
        params: {
          source_id: selectedSource.value,
          page: page.value,
          limit: limit.value,
          sort_by: sortBy.value,
          sort_order: sortOrder.value,
          ...filters.value,
        },
        headers: {
          'Authorization': `Bearer ${token}`, 
        },
        withCredentials: true,
      })
      ratings.value = response.data.body.ratings || []
      totalRatings.value = response.data.body.total
    } catch (err) {
      console.error('Error fetching ratings', err)
      error.value = 'Error fetching ratings'
    } finally {
      loading.value = false
    }
  }

  watch(
    selectedSource,
    (newSource) => {
      if (newSource) {
        fetchRatings()
      }
    },
    { immediate: true },
  )

  const setSorting = (column: string, order: 'asc' | 'desc') => {
    sortBy.value = column
    sortOrder.value = order
    fetchRatings()
  }

  const setFilter = (newFilters: Record<string, string>) => {
    const cleanedFilters = Object.fromEntries(
      Object.entries({ ...filters.value, ...newFilters }).filter(([_, v]) => v !== ''),
    )

    filters.value = cleanedFilters

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
