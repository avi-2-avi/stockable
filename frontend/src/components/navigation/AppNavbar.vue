<template>
  <BaseNavbar>
    <div class="flex items-center justify-between space-x-5">
      <div class="flex space-x-1.5">
        <router-link :to="route.path" v-for="route in filteredRoutes" :key="route.path"
          class="p-2 text-sm text-foreground hover:text-stock-500 hidden sm:block" :class="{
            'font-bold': $route.path === route.path,
            'font-medium': $route.path !== route.path
          }">
          {{ route.name }}
        </router-link>
      </div>

      <div class="relative">
        <div @click="toggleDropdown"
          class="h-10 w-10 bg-stock-500 hover:bg-stock-500/80 text-white rounded-full flex items-center justify-center font-semibold cursor-pointer">
          {{ userInitials }}
        </div>
        <div v-if="isDropdownOpen"
          class="absolute right-0 mt-2 w-48 bg-base border border-border rounded-lg shadow-lg overflow-hidden">
          <div>
            <button @click="logOut" class="flex items-center w-full px-4 py-2 hover:bg-stock-500/20">
              <LogOutIcon class="w-5 h-5 mr-4 text-stock-500" />
              Logout
            </button>
          </div>
        </div>
      </div>
    </div>
  </BaseNavbar>
</template>

<script setup lang="ts">
import BaseNavbar from '@/components/navigation/BaseNavbar.vue'
import { useAuthStore } from '@/store/authStore'
import { computed, ref } from 'vue'
import { LogOutIcon } from 'lucide-vue-next'

const authStore = useAuthStore()
const isAdmin = computed(() => authStore.user?.role_id === 1) 


const userInitials = computed(() => {
  return authStore.user?.full_name?.slice(0, 2).toUpperCase() || 'US'
})

const routes = [
  {
    name: 'Dashboard',
    path: '/app/dashboard',
    meta: { requiresAdmin: false },
  },
  {
    name: 'Analyst Ratings',
    path: '/app/analyst-ratings',
    meta: { requiresAdmin: false },
  },
  {
    name: 'User Management',
    path: '/app/admin/users',
    meta: { requiresAdmin: true },
  }
]

const filteredRoutes = computed(() => {
  return routes.filter(route => !route.meta || !route.meta.requiresAdmin || isAdmin.value)
})

const isDropdownOpen = ref(false)

const toggleDropdown = () => {
  isDropdownOpen.value = !isDropdownOpen.value
}

const logOut = async () => {
  await authStore.logout()
}

</script>
