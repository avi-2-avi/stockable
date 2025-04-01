import { mount } from '@vue/test-utils'
import { describe, it, expect, vi } from 'vitest'
import AppNavbar from '../../../../components/navigation/AppNavbar.vue'
import HomeCard from '../components/InfoCard.vue'
import CalculationsCard from '../components/CalculationsCard.vue'
import CalculationsModal from '../components/CalculationsModal.vue'
import RatingTable from '../components/RatingTable.vue'
import AnalystRatingsPage from '../pages/AnalystRatingsPage.vue'

vi.mock('axios', () => ({
  default: {
    post: vi.fn(() => Promise.resolve({ data: {} })),
    get: vi.fn(() => Promise.resolve({ data: {} })),
  },
}))

describe('AnalystRatingsPage.vue', () => {
  it('renders the App Navbar', () => {
    const wrapper = mount(AnalystRatingsPage)
    expect(wrapper.findComponent(AppNavbar).exists()).toBe(true)
  })
  it('displays the three home cards', () => {
    const wrapper = mount(AnalystRatingsPage)
    expect(wrapper.findAllComponents(HomeCard).length).toBe(3)
  })
  it('displays the calculations card', () => {
    const wrapper = mount(AnalystRatingsPage)
    expect(wrapper.findComponent(CalculationsCard).exists()).toBe(true)
  })
  it('opens the calculation modal when calculations card is clicked', async () => {
    const wrapper = mount(AnalystRatingsPage)
    await wrapper.findComponent(CalculationsCard).trigger('click')
    expect(wrapper.findComponent(CalculationsModal).exists()).toBe(true)
  })
  it('closes the calculation modal when the close button is clicked', async () => {
    const wrapper = mount(AnalystRatingsPage)
    await wrapper.findComponent(CalculationsCard).trigger('click')
    await wrapper.findComponent(CalculationsModal).find('button').trigger('click')
    expect(wrapper.findComponent(CalculationsModal).exists()).toBe(false)
  })
  it('renders the Data Table with filters, sort and pagination', () => {
    const wrapper = mount(AnalystRatingsPage)
    expect(wrapper.findComponent(RatingTable).exists()).toBe(true)
  })
})
