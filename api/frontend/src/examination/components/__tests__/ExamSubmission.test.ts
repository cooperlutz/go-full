import ExamSubmission from "../SubmitExamButton.vue";
import { describe, it, expect } from "vitest";
import { mount } from "@vue/test-utils";

const wrapper = mount(ExamSubmission, {
  props: {
    examId: "exam1",
  },
});

describe("ExamSubmission", () => {
  it("renders input and send button", () => {
    // Assert
    expect(wrapper.find("#exam-submission-button").exists()).toBe(true);
  });
});
