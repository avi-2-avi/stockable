import AppNavbar from '@/components/navigation/AppNavbar.vue'
import CalculationsCard from '@/features/app/analyst-ratings/components/CalculationsCard.vue'
import CalculationsModal from '@/features/app/analyst-ratings/components/CalculationsModal.vue'
import HomeCard from '@/features/app/analyst-ratings/components/InfoCard.vue'
import RatingTable from '@/features/app/analyst-ratings/components/RatingTable.vue'
import AnalystRatingsPage from '@/features/app/analyst-ratings/pages/AnalystRatingsPage.vue'
import { mount, VueWrapper } from '@vue/test-utils'
import axios from 'axios'
import { describe, it, expect, vi, beforeEach } from 'vitest'

vi.mock('axios')

let wrapper: VueWrapper<any>;

beforeEach(async () => {
  await axios.post("http://localhost:8085/api/login", { email: "test@mail.com", password: "password123" });
  wrapper = mount(AnalystRatingsPage);
});

describe('AnalystRatingsPage.vue', () => {
  it('renders the App Navbar', () => {
    expect(wrapper.findComponent(AppNavbar).exists()).toBe(true)
  })
  it('displays the three home cards', () => {
    expect(wrapper.findAllComponents(HomeCard).length).toBe(3)
  })
  it('displays the calculations card', () => {
    expect(wrapper.findComponent(CalculationsCard).exists()).toBe(true)
  })
  it('opens the calculation modal when calculations card is clicked', async () => {
    await wrapper.findComponent(CalculationsCard).trigger('click')
    expect(wrapper.findComponent(CalculationsModal).exists()).toBe(true)
  })
  it('closes the calculation modal when the close button is clicked', async () => {
    await wrapper.findComponent(CalculationsCard).trigger('click')
    await wrapper.findComponent(CalculationsModal).find('button').trigger('click')
    expect(wrapper.findComponent(CalculationsModal).exists()).toBe(false)
  })
  it('renders the Data Table with filters, sort and pagination', () => {
    expect(wrapper.findComponent(RatingTable).exists()).toBe(true)
  })
})
