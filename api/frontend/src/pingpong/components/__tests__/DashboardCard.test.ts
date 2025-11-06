import { describe, it, expect } from "vitest";
import { mount } from "@vue/test-utils";
import DashboardCard from "../DashboardCard.vue";

describe("DashboardCard.vue", () => {
  it("renders the card title", () => {
    // Arrange & Act
    const wrapper = mount(DashboardCard);

    // Assert
    expect(wrapper.find(".card-title").text()).toBe("Ping Pong");
  });

  it("renders the logo image", () => {
    // Arrange & Act
    const wrapper = mount(DashboardCard);
    const img = wrapper.find("img");
    const imgData = img.element as HTMLImageElement;

    // Assert
    expect(imgData.src).toContain("data:image/svg+xml");
    expect(img.exists()).toBe(true);
    expect(img.attributes("width")).toBe("48");
    expect(img.attributes("height")).toBe("48");
  });

  it("contains the correct link", () => {
    // Arrange & Act
    const wrapper = mount(DashboardCard);

    // Assert
    const link = wrapper.find("a");
    expect(link.exists()).toBe(true);
    expect(link.attributes("href")).toBe("/ping-pong");
  });

  it("renders the description paragraph", () => {
    // Arrange & Act
    const wrapper = mount(DashboardCard);

    // Assert
    expect(wrapper.find("p").text()).toContain(
      "Application of the Ping Pong Module",
    );
  });
});
