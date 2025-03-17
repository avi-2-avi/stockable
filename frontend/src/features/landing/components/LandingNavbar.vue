<template>
  <BaseNavbar>
    <Button label="Sign in" size="md" @click="redirectToLogin" />
  </BaseNavbar>
</template>

<script setup lang="ts">
import BaseNavbar from '@/components/BaseNavbar.vue'
import Button from '@/components/Button.vue'
import { useRouter } from 'vue-router'

const router = useRouter()

const redirectToLogin = () => {
  router.push('/auth')
}
</script>
