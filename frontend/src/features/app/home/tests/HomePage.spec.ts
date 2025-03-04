import { mount } from "@vue/test-utils";
import { describe, it, expect } from "vitest";
import HomePage from "../pages/HomePage.vue";
import Navbar from "../../../../components/Navbar.vue";
import SelectDropdown from "@/components/SelectDropdown.vue";
import DataTable from "../../components/DataTable.vue";

describe("HomePage.vue", () => {
  it("renders the Navbar", () => {
    const wrapper = mount(HomePage);
    expect(wrapper.findComponent(Navbar).exists()).toBe(true);
  });

  it("displays the page title", () => {
    const wrapper = mount(HomePage);
    expect(wrapper.find("h1").text()).toBe("🏠 Home Page");
  });

  it("renders the Select component", () => {
    const wrapper = mount(HomePage);
    expect(wrapper.findComponent(SelectDropdown).exists()).toBe(true);
  });

  it("renders the Data Table with filters, sort, and pagination", () => {
    const wrapper = mount(HomePage);
    expect(wrapper.findComponent(DataTable).exists()).toBe(true);
  });
});
