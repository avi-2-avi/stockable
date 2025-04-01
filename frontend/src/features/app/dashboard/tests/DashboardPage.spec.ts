import { mount } from "@vue/test-utils"
import { describe, expect, it } from "vitest"
import DashboardPage from "../pages/DashboardPage.vue"
import AppNavbar from "../../../../components/navigation/AppNavbar.vue"
import PortafolioCard from "../components/PortafolioCard.vue"

describe("DashboardPage.vue", () => {
    it('renders the App Navbar', () => {
        const wrapper = mount(DashboardPage)
        expect(wrapper.findComponent(AppNavbar).exists()).toBe(true)
    })

    it('displays stocks CPI donut chart', () => {
        const wrapper = mount(DashboardPage)
        expect(wrapper.findComponent(PortafolioCard).exists()).toBe(true)
    })

})