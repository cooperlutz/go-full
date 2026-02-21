import { describe, it, expect, vi } from "vitest";
import { mount } from "@vue/test-utils";
import ReportingDashboard from "../ReportingDashboard.vue";

vi.mock("vue-router", () => ({
  useRoute: () => ({
    path: "/reporting",
  }),
}));

describe("ReportingDashboard.vue", () => {
  it("renders title, subtitle, and description", () => {
    const wrapper = mount(ReportingDashboard);

    expect(wrapper.text()).toContain("Reporting Dashboard");
  });
});
