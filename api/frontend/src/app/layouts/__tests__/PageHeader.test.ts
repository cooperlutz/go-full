import { describe, it, expect } from "vitest";
import { mount } from "@vue/test-utils";
import PageHeader from "../PageLayouts/PageHeader.vue";

// Mock child components
const stubs = {
  PageHeaderBreadcrumbs: { template: '<div data-testid="breadcrumbs"/>' },
  PageHeaderMenu: { template: '<div data-testid="menu"/>' },
  SectionDivider: { template: '<div data-testid="divider"/>' },
};

describe("PageHeader.vue", () => {
  it("renders title, subtitle, and description", () => {
    const wrapper = mount(PageHeader, {
      props: {
        title: "Test Title",
        subtitle: "Test Subtitle",
        description: "Test Description",
      },
      global: { stubs },
    });
    expect(wrapper.text()).toContain("Test Title");
    expect(wrapper.text()).toContain("Test Subtitle");
    expect(wrapper.text()).toContain("Test Description");
  });

  it("shows PageHeaderMenu by default", () => {
    const wrapper = mount(PageHeader, {
      props: { title: "", subtitle: "", description: "" },
      global: { stubs },
    });
    expect(wrapper.find('[data-testid="menu"]').exists()).toBe(true);
  });

  it("hides PageHeaderMenu when disableMenu is true", () => {
    const wrapper = mount(PageHeader, {
      props: { disableMenu: true },
      global: { stubs },
    });
    expect(wrapper.find('[data-testid="menu"]').exists()).toBe(false);
  });

  it("always renders breadcrumbs and divider", () => {
    const wrapper = mount(PageHeader, {
      props: {},
      global: { stubs },
    });
    expect(wrapper.find('[data-testid="breadcrumbs"]').exists()).toBe(true);
    expect(wrapper.find('[data-testid="divider"]').exists()).toBe(true);
  });
});
