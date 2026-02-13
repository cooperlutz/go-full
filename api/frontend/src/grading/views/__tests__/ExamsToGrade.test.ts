import ExamsToGrade from "../ExamsToGrade.vue";
import { describe, it, expect, vi, beforeEach } from "vitest";
import { mount, flushPromises } from "@vue/test-utils";
import { nextTick } from "vue";

vi.mock("vue-router", () => ({
  useRoute: () => ({
    path: "/exam/5d9abb80-0706-42ad-8131-33627d3e6b17/question/1",
    params: {
      examId: "5d9abb80-0706-42ad-8131-33627d3e6b17",
      questionIndex: "1",
    },
  }),
}));

describe("ExamsToGrade", () => {
  beforeEach(() => {
    vi.clearAllMocks();
  });

  it("displays exam data when exam is loaded", async () => {
    const wrapper = mount(ExamsToGrade);

    await flushPromises();
    await nextTick();

    const examDiv = wrapper.find("#exams-to-grade");
    expect(examDiv.exists()).toBe(true);
  });

  it("renders the list of ungraded exams", async () => {
    const wrapper = mount(ExamsToGrade);

    await flushPromises();
    await nextTick();

    const examDiv = wrapper.find("#grading-ungraded-exams-list");
    expect(examDiv.text()).toBe(
      "Grade Exam: 5d9abb80-0706-42ad-8131-33627d3e6b17",
    );
  });
});
