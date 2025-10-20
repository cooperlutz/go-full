import { describe, it, expect, vi } from "vitest";
import { mount } from "@vue/test-utils";
import ApplicationShell from "../ApplicationShell/ApplicationShell.vue";

// Define a local stub for RouterView
const RouterViewStub = {
  name: "RouterView",
  template: '<div data-test="router-view"></div>',
};

// Mock child components
vi.mock("~/app/layouts/ApplicationShell/NavBar/NavBar.vue", () => ({
  default: { name: "NavBar", template: '<nav data-test="navbar"></nav>' },
}));
vi.mock("~/app/layouts/ApplicationShell/SideBar/SideBar.vue", () => ({
  default: {
    name: "SideBar",
    template: '<aside data-test="sidebar"><slot name="content"></slot></aside>',
  },
}));
vi.mock("~/app/layouts/ApplicationShell/Footer/FooterPrimary.vue", () => ({
  default: {
    name: "FooterPrimary",
    template: '<footer data-test="footer"></footer>',
  },
}));

describe("ApplicationShell.vue", () => {
  it("renders NavBar, SideBar, and Footer", () => {
    const wrapper = mount(ApplicationShell, {
      global: {
        stubs: {
          RouterView: RouterViewStub,
        },
      },
    });
    expect(wrapper.find('[data-test="navbar"]').exists()).toBe(true);
    expect(wrapper.find('[data-test="sidebar"]').exists()).toBe(true);
    expect(wrapper.find('[data-test="footer"]').exists()).toBe(true);
  });

  it("renders main content area with correct classes and id", () => {
    const wrapper = mount(ApplicationShell, {
      global: {
        stubs: {
          RouterView: RouterViewStub,
        },
      },
    });
    const mainContent = wrapper.find("#main-content");
    expect(mainContent.exists()).toBe(true);
    expect(mainContent.classes()).toContain("min-h-screen");
    expect(mainContent.classes()).toContain("overflow-y-auto");
    expect(mainContent.classes()).toContain("bg-base-100");
  });

  it("renders router-view component", () => {
    const wrapper = mount(ApplicationShell, {
      global: {
        stubs: {
          RouterView: RouterViewStub,
        },
      },
    });
    expect(wrapper.findComponent(RouterViewStub).exists()).toBe(true);
  });
});
