import QuestionGrading from "../QuestionGrading.vue";
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

describe("QuestionGrading", () => {
  beforeEach(() => {
    vi.clearAllMocks();
  });

  it("displays exam data when exam is loaded", async () => {
    const wrapper = mount(QuestionGrading);

    await flushPromises();
    await nextTick();

    const examDiv = wrapper.find("#question-grading");
    expect(examDiv.exists()).toBe(true);
  });
});
