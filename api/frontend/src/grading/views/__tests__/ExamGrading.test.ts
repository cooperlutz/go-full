import ExamGrading from "../ExamGrading.vue";
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

describe("ExamGrading", () => {
  beforeEach(() => {
    vi.clearAllMocks();
  });

  it("displays exam data when exam is loaded", async () => {
    const wrapper = mount(ExamGrading);

    await flushPromises();
    await nextTick();

    const examDiv = wrapper.find("#exam-grading-component");
    expect(examDiv.exists()).toBe(true);
  });

  it("shows loading state initially", async () => {
    // Arrange & Act
    const wrapper = mount(ExamGrading);
    await nextTick();

    // Assert
    expect(wrapper.text()).toContain("Loading...");
  });

  it("renders the list of questions that need to be graded", async () => {
    const wrapper = mount(ExamGrading);

    await flushPromises();
    await nextTick();

    const examDiv = wrapper.find("#grading-ungraded-questions-list");
    expect(examDiv.text()).toContain("1 - Graded: false");
  });
});
