<template>
  <BaseNavbar>
    <div class="flex items-center justify-between">
      <router-link :to="route.path" v-for="route in routes" :key="route.path"
        class="px-4 py-2 text-sm font-semibold text-foreground hover:text-stock-400/70 hidden sm:block"
        :class="{ 'text-stock-400': $route.path === route.path }">
        {{ route.name }}
      </router-link>
      <Button label="Log out" size="md" variant="solid" @click="logOut" />
    </div>
  </BaseNavbar>
</template>

<script setup lang="ts">
import Button from '@/components/Button.vue'
import BaseNavbar from '@/components/BaseNavbar.vue'

const routes = [
  {
    name: 'Home',
    path: '/app/home',
  },
  {
    name: 'Recommendations',
    path: '/app/recommendation',
  }
]

const logOut = () => {
  console.log('Logging out')
}
</script>
