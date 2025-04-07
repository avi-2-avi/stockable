import { ref, computed } from 'vue'
import { defineStore } from 'pinia'
import axios from 'axios'
import { useRouter } from 'vue-router'
import type { User } from '@/types/user'
import { jwtDecode } from 'jwt-decode'

interface LoginResponse {
  data: {
    body: {
      token: string
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

    const decodeToken = (token: string) => {
      try {
        const decoded = jwtDecode<{ User: { id: string, full_name: string, email: string, role_id: number }, exp: number, iat: number }>(token)
        return decoded.User
      } catch (error) {
        console.error('Failed to decode JWT', error)
        return null
      }
    }

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

        const token = response.data.body.token
        if (token) {
          const decodedUser = decodeToken(token)
          if (decodedUser) {
            const userData: User = {
              id: decodedUser.id,
              full_name: decodedUser.full_name,
              email: decodedUser.email,
              role_id: decodedUser.role_id || 0,
            }
            user.value = userData 
            localStorage.setItem('auth_token', token)
          }
        }

        router.push('/app/dashboard')
      } catch (error) {
        console.error('Login failed')
        throw new Error('Invalid credentials')
      }
    }

    const register = async (fullName: string, email: string, password: string, role_name: string) => {
      try {
        await axios.post(`${import.meta.env.VITE_API_URL}/api/auth/register`, {
          full_name: fullName,
          email,
          password,
          role_name,
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

        localStorage.removeItem('auth_token')
        document.cookie = 'auth_token=; Max-Age=0; path=/;'

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
