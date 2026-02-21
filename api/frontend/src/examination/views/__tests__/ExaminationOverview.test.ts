import { describe, it, expect, vi } from "vitest";
import { mount } from "@vue/test-utils";
import ReportingDashboard from "../ExaminationOverview.vue";

vi.mock("vue-router", () => ({
  useRoute: () => ({
    path: "/exam/5d9abb80-0706-42ad-8131-33627d3e6b17",
    params: { id: "5d9abb80-0706-42ad-8131-33627d3e6b17" },
  }),
}));

describe("ExaminationOverview.vue", () => {
  it("renders title, subtitle, and description", () => {
    const wrapper = mount(ReportingDashboard);

    expect(wrapper.text()).toContain("Exam Overview");
  });
});
