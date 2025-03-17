import { ref, computed } from 'vue'
import { defineStore } from 'pinia'
import axios from 'axios'
import { useRouter } from 'vue-router'
import type { User } from '@/types/user'

interface LoginResponse {
  data: {
    body: {
      user: User
    }
    error?: string
    status: string
    message?: string
  }
}

export const useAuthStore = defineStore(
  'auth',
  () => {
    const user = ref<User | null>(null)
    const router = useRouter()

    const isAuthenticated = computed(() => !!user.value)

    const login = async (email: string, password: string) => {
      try {
        const response: LoginResponse = await axios.post(
          `${import.meta.env.VITE_API_URL}/api/auth/login`,
          {
            email,
            password,
          },
          { withCredentials: true },
        )

        user.value = response.data.body.user
        router.push('/app/home')
      } catch (error) {
        console.error('Login failed')
        throw new Error('Invalid credentials')
      }
    }

    const register = async (fullName: string, email: string, password: string) => {
      try {
        await axios.post(`${import.meta.env.VITE_API_URL}/api/auth/register`, {
          full_name: fullName,
          email,
          password,
        })
        await login(email, password)
      } catch (error) {
        console.error('Registration failed')
        throw new Error('Registration failed')
      }
    }

    const logout = async () => {
      try {
        await axios.post(
          `${import.meta.env.VITE_API_URL}/api/auth/logout`,
          {},
          { withCredentials: true },
        )
        user.value = null
        router.push('/auth')
      } catch (error) {
        console.error('Logout failed')
      }
    }

    return { user, isAuthenticated, login, register, logout }
  },
  {
    persist: {
      storage: localStorage,
    },
  },
)
