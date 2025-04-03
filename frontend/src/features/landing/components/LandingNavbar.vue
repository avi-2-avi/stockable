<template>
  <BaseNavbar>
    <Button label="Sign in" size="md" @click="redirectToLogin" />
  </BaseNavbar>
</template>

<script setup lang="ts">
import BaseNavbar from '@/components/navigation/BaseNavbar.vue'
import Button from '@/components/ui/Button.vue'
import { useRouter } from 'vue-router'

const router = useRouter()

const redirectToLogin = () => {
  router.push('/auth')
}
</script>
