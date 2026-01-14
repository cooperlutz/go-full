import { describe, it, expect } from "vitest";
import { mount } from "@vue/test-utils";
import DashboardCard from "../DashboardCard.vue";

describe("DashboardCard.vue", () => {
  it("renders the card title", () => {
    // Arrange & Act
    const wrapper = mount(DashboardCard);

    // Assert
    expect(wrapper.find(".card-title").text()).toBe("Exam Library");
  });

  it("contains the correct link", () => {
    // Arrange & Act
    const wrapper = mount(DashboardCard);

    // Assert
    const link = wrapper.find("a");
    expect(link.exists()).toBe(true);
    expect(link.attributes("href")).toBe("/exam-library");
  });

  it("renders the description paragraph", () => {
    // Arrange & Act
    const wrapper = mount(DashboardCard);

    // Assert
    expect(wrapper.find("p").text()).toContain("Explore the exam library");
  });
});
