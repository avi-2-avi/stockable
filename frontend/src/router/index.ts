import { createRouter, createWebHistory } from 'vue-router'
import LandingPage from '../features/landing/pages/LandingPage.vue'
import LoginPage from '../features/auth/pages/LoginPage.vue'
import { useAuthStore } from '@/store/authStore'
import DashboardPage from '@/features/app/dashboard/pages/DashboardPage.vue'
import AnalystRatingsPage from '@/features/app/analyst-ratings/pages/AnalystRatingsPage.vue'

export const routes = [
  {
    path: '/',
    name: 'Landing',
    component: LandingPage,
    meta: { requiresAuth: false },
  },
  {
    path: '/auth',
    name: 'Auth',
    component: LoginPage,
    meta: { requiresAuth: false },
  },
  {
    path: '/app/dashboard',
    name: 'Dashboard',
    component: DashboardPage,
    meta: { requiresAuth: true },
  },
  {
    path: '/app/analyst-ratings',
    name: 'AnalystRatings',
    component: AnalystRatingsPage,
    meta: { requiresAuth: true },
  },
  {
    path: '/:pathMatch(.*)*',
    redirect: '/',
  },
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
})

router.beforeEach((to, _, next) => {
  const authStore = useAuthStore()

  if (to.meta.requiresAuth && !authStore.isAuthenticated) {
    next('/auth')
  } else if (to.path === '/auth' && authStore.isAuthenticated) {
    next('/app/dashboard')
  } else {
    next()
  }
})

export default router
