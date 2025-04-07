import type { DashboardData } from '@/types/dashboard'
import axios from 'axios'
import { defineStore } from 'pinia'
import { reactive } from 'vue'

export const useDashboardStore = defineStore('useDashboard', () => {
  const dashboardRatings = reactive<DashboardData>({
    latest_ratings: [],
    donut_cpi_chart: [],
    donut_rating_chart: [],
  })

  const cachedDashboardRatings = reactive<DashboardData>({...dashboardRatings})

  const fetchDashboardRatings = async (sourceId: string) => {
    const token = localStorage.getItem('auth_token');
    if (!token) {
      throw new Error('No authentication token found');
    }

    try {
      const response = await axios.get(`${import.meta.env.VITE_API_URL}/api/ratings/dashboard`, {
        params: { source_id: sourceId },
        headers: {
          'Authorization': `Bearer ${token}`, 
        },
        withCredentials: true,
      })

      if (response.status !== 200) {
        throw new Error('Failed to fetch dashboard ratings')
      }

      const newDashboardRatings = response.data.body.data
      if(JSON.stringify(newDashboardRatings) !== JSON.stringify(dashboardRatings)) {
        Object.assign(dashboardRatings, response.data.body.data)
        Object.assign(cachedDashboardRatings, response.data.body.data)
      }
    } catch (error) {
      console.error('Error fetching dashboard ratings', error)
    }
  }

  return { dashboardRatings, cachedDashboardRatings, fetchDashboardRatings }
})
