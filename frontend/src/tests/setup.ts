import { createPinia } from 'pinia'
import { vi } from 'vitest'
import { config } from '@vue/test-utils'
import { createRouter, createWebHistory } from 'vue-router'
import { routes } from '@/router'

const router = createRouter({
  history: createWebHistory(),
  routes,
})

const pinia = createPinia()

config.global.plugins = [pinia, router]

config.global.mocks = {
  $route: { path: '/' },
  $router: { push: vi.fn(), replace: vi.fn() },
}

Object.defineProperty(window, 'matchMedia', {
  writable: true,
  value: vi.fn().mockImplementation((query) => ({
    matches: query.includes('dark'),
    media: query,
    onchange: null,
    addListener: vi.fn(),
    removeListener: vi.fn(),
    addEventListener: vi.fn(),
    removeEventListener: vi.fn(),
    dispatchEvent: vi.fn(),
  })),
})
