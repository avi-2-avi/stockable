import AppNavbar from "@/components/navigation/AppNavbar.vue"
import AboutCard from "@/features/app/dashboard/components/AboutCard.vue"
import DonutChartCard from "@/features/app/dashboard/components/DonutChartCard.vue"
import NewsCard from "@/features/app/dashboard/components/NewsCard.vue"
import RecentRatingTable from "@/features/app/dashboard/components/RecentRatingTable.vue"
import DashboardPage from "@/features/app/dashboard/pages/DashboardPage.vue"
import { mount, VueWrapper } from "@vue/test-utils"
import axios from "axios"
import { beforeEach, describe, expect, it, vi } from "vitest"

vi.mock('axios')

let wrapper: VueWrapper<any>;

beforeEach(async () => {
  await axios.post("http://localhost:8085/api/login", { email: "test@mail.com", password: "password123" });
  wrapper = mount(DashboardPage);
});


describe("DashboardPage.vue", () => {
    it('renders the App Navbar', () => {
        expect(wrapper.findComponent(AppNavbar).exists()).toBe(true)
    })

    it('displays about card', () => {
        expect(wrapper.findComponent(AboutCard).exists()).toBe(true)
    })

    it('displays recent ratings', () => {
        expect(wrapper.findComponent(RecentRatingTable).exists()).toBe(true)
    })

    it('displays stocks CPI donut chart', () => {
        expect(wrapper.findAllComponents(DonutChartCard).length).toBe(2)
        expect(wrapper.findAllComponents(DonutChartCard)[0].text()).toContain('CPI Breakdown')
    })

    it('displays stocks ratings donut chart', () => {
        expect(wrapper.findAllComponents(DonutChartCard).length).toBe(2)
        expect(wrapper.findAllComponents(DonutChartCard)[1].text()).toContain('Ratings Breakdown')
    })

    it('displays news cards', () => {
        expect(wrapper.findAllComponents(NewsCard).length).toBe(3)
    })
})