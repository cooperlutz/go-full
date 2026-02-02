import { describe, it, expect } from "vitest";
import { mount } from "@vue/test-utils";
import FooterPrimary from "../ApplicationShell/Footer/FooterPrimary.vue";

describe("FooterPrimary.vue", () => {
  it("renders footer element", () => {
    const wrapper = mount(FooterPrimary);
    expect(wrapper.find("footer").exists()).toBe(true);
  });

  it("renders the correct description text", () => {
    const wrapper = mount(FooterPrimary);
    expect(wrapper.text()).toContain("Go Full");
    expect(wrapper.text()).toContain("Y2K Compliant");
  });
});
