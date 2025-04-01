<template>
  <BaseNavbar>
    <div class="flex items-center justify-between space-x-2.5">
      <router-link :to="route.path" v-for="route in routes" :key="route.path"
        class="px-4 py-2 text-sm text-foreground hover:text-stock-500 hidden sm:block"
        :class="{ 'font-bold': $route.path === route.path }">
        {{ route.name }}
      </router-link>
      <Button label="Log out" size="md" variant="solid" @click="logOut" />
    </div>
  </BaseNavbar>
</template>

<script setup lang="ts">
import Button from '@/components/ui/Button.vue'
import BaseNavbar from '@/components/navigation/BaseNavbar.vue'
import { useAuthStore } from '@/store/authStore'

const authStore = useAuthStore()

const routes = [
  {
    name: 'Dashboard',
    path: '/app/dashboard',
  },
  {
    name: 'Analyst Ratings',
    path: '/app/analyst-ratings',
  }
]

const logOut = async () => {
  await authStore.logout()
}
</script>
