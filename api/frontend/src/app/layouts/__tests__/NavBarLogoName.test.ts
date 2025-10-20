import { describe, it, expect } from "vitest";
import { mount } from "@vue/test-utils";
import NavBarLogoName from "../ApplicationShell/NavBar/NavBarLogo.vue";

describe("NavBarLogoName", () => {
  it("renders the logo name text", () => {
    const wrapper = mount(NavBarLogoName);
    expect(wrapper.text()).toContain("Go Full");
  });

  it("renders an image with correct src and dimensions", () => {
    const wrapper = mount(NavBarLogoName);
    const img = wrapper.find("img");
    expect(img.exists()).toBe(true);
    expect(img.attributes("src")).toBe("/image.png");
    expect(img.attributes("height")).toBe("32");
    expect(img.attributes("width")).toBe("32");
  });

  it("renders a link to the homepage", () => {
    const wrapper = mount(NavBarLogoName);
    const a = wrapper.find("a");
    expect(a.exists()).toBe(true);
    expect(a.attributes("href")).toBe("/");
  });
});
