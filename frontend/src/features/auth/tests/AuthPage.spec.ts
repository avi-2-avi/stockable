import { mount } from '@vue/test-utils'
import { describe, it, expect, vi } from 'vitest'
import AuthPage from '../pages/LoginPage.vue'
import BaseNavbar from '@/components/navigation/BaseNavbar.vue'

vi.mock('axios', () => ({
  default: {
    post: vi.fn(() => Promise.resolve({ data: {} })),
    get: vi.fn(() => Promise.resolve({ data: {} })),
  },
}))

describe('AuthPage.vue', () => {
  it('renders the Base Navbar', () => {
    const wrapper = mount(AuthPage)
    expect(wrapper.findComponent(BaseNavbar).exists()).toBe(true)
  })
  it('displays the login form', () => {
    const wrapper = mount(AuthPage)
    expect(wrapper.find('form').exists()).toBe(true)
  })
  it('displays the email input', () => {
    const wrapper = mount(AuthPage)
    expect(wrapper.find('input[type="email"]').exists()).toBe(true)
  })
  it('displays the password input', () => {
    const wrapper = mount(AuthPage)
    expect(wrapper.find('input[type="password"]').exists()).toBe(true)
  })
  it('displays the login button', () => {
    const wrapper = mount(AuthPage)
    expect(wrapper.find('button').exists()).toBe(true)
  })
  it('displays the register link', () => {
    const wrapper = mount(AuthPage)
    expect(wrapper.find('a').exists()).toBe(true)
  })
  it('displays register title when register link is clicked', async () => {
    const wrapper = mount(AuthPage)
    await wrapper.find('a').trigger('click')
    expect(wrapper.find('h2').text()).toBe('Create an Account')
  })
})
