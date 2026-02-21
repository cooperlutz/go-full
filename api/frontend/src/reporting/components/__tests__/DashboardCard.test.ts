import { describe, it, expect } from "vitest";
import { mount } from "@vue/test-utils";
import DashboardCard from "../DashboardCard.vue";

const wrapper = mount(DashboardCard);

describe("DashboardCard.vue", () => {
  it("renders the card title", () => {
    // Assert
    expect(wrapper.find(".card-title").text()).toBe("Reporting");
  });

  it("contains the correct link", () => {
    // Assert
    const link = wrapper.find("a");
    expect(link.exists()).toBe(true);
    expect(link.attributes("href")).toBe("/reporting");
  });

  it("renders the description paragraph", () => {
    // Assert
    expect(wrapper.find("p").text()).toContain(
      "View reports and analytics on exam performance",
    );
  });
});
