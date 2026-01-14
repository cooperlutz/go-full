import { describe, it, expect } from "vitest";
import { mount } from "@vue/test-utils";
import SideBar from "../ApplicationShell/SideBar/SideBar.vue";

describe("SideBar.vue", () => {
  it("renders two sidebar items, each with an icon and name", () => {
    const wrapper = mount(SideBar);
    const sidebarItems = wrapper.findAll(".sidebar-item");
    expect(sidebarItems.length).toBe(3);
    sidebarItems.forEach((item) => {
      expect(item.find("svg").exists()).toBe(true);
      expect(item.text()).toBeTruthy();
    });
  });
});
