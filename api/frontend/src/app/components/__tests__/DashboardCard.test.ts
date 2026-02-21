import { describe, it, expect } from "vitest";
import { mount } from "@vue/test-utils";
import DashboardCard from "../DashboardCard/DashboardCard.vue";
import { LibraryBig } from "lucide-vue-next";

const wrapper = mount(DashboardCard, {
  props: {
    title: "Exam Library",
    description: "Explore the exam library",
    href: "/exam-library",
    icon: LibraryBig,
  },
});

describe("DashboardCard.vue", () => {
  it("renders the card title", () => {
    // Assert
    expect(wrapper.find(".card-title").text()).toBe("Exam Library");
  });

  it("contains the correct link", () => {
    // Assert
    const link = wrapper.find("a");
    expect(link.exists()).toBe(true);
    expect(link.attributes("href")).toBe("/exam-library");
  });

  it("renders the description paragraph", () => {
    // Assert
    expect(wrapper.find("p").text()).toContain("Explore the exam library");
  });
});
