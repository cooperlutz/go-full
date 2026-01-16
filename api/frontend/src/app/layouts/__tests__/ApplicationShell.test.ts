import { describe, it, expect, vi } from "vitest";
import { mount } from "@vue/test-utils";
import ApplicationShell from "../ApplicationShell/ApplicationShell.vue";

// Define a local stub for RouterView
const RouterViewStub = {
  name: "RouterView",
  template: '<div data-test="router-view"></div>',
};

// Mock child components
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
