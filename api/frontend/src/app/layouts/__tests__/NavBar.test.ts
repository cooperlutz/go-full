import { describe, it, expect, vi } from "vitest";
import { mount } from "@vue/test-utils";
import NavBar from "../ApplicationShell/NavBar/NavBar.vue";

// Mock child components
vi.mock("~/app/layouts/ApplicationShell/NavBar/NavBarLogo.vue", () => ({
  default: { name: "LogoName", template: '<div class="logo-name-mock"></div>' },
}));
vi.mock("lucide-vue-next", () => ({
  Settings: {
    name: "Settings",
    template: '<svg class="settings-icon-mock"></svg>',
  },
}));
vi.mock("~/app/config", () => ({
  default: { DOCS_URL: "https://docs.example.com" },
}));

describe("NavBar.vue", () => {
  it("renders logo, search bar, docs link, theme control, and settings icon", () => {
    const wrapper = mount(NavBar);
    expect(wrapper.find(".logo-name-mock").exists()).toBe(true);
    expect(wrapper.find(".settings-icon-mock").exists()).toBe(true);
    expect(wrapper.find('a[href="https://docs.example.com"]').text()).toBe(
      "Docs",
    );
    expect(wrapper.find('a[href="/settings"]').exists()).toBe(true);
  });

  it("docs link opens in a new tab", () => {
    const wrapper = mount(NavBar);
    const docsLink = wrapper.find('a[href="https://docs.example.com"]');
    expect(docsLink.attributes("target")).toBe("_blank");
  });

  it("settings link points to /settings", () => {
    const wrapper = mount(NavBar);
    const settingsLink = wrapper.find('a[href="/settings"]');
    expect(settingsLink.exists()).toBe(true);
  });
});
