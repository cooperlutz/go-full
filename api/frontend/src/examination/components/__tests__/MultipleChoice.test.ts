import MultipleChoice from "../Questions/MultipleChoice.vue";
import { describe, it, expect } from "vitest";
import { mount } from "@vue/test-utils";
import { type Question } from "~/examination/services";

const sampleQuestion: Question = {
  questionId: "q1",
  examId: "exam1",
  questionType: "multiple-choice",
  questionText: "choose the correct option.",
  questionIndex: 1,
  responseOptions: ["Option A", "Option B", "Option C", "Option D"],
  answered: false,
};

const wrapper = mount(MultipleChoice, {
  props: {
    question: sampleQuestion,
  },
});

describe("MultipleChoice", () => {
  it("renders input and submit button", () => {
    // Assert
    expect(wrapper.find("#multiple-choice-radio-option-0").exists()).toBe(true);
    expect(wrapper.find("#record-answer-button").exists()).toBe(true);
  });

  it("allows user to select an option from the radio input", () => {
    const inputOption = wrapper.find("#multiple-choice-radio-option-0");
    inputOption.setValue("Option A");
    inputOption.trigger("click");
    const submitButton = wrapper.find("#record-answer-button");
    submitButton.trigger("click");
    // Assert
    expect((inputOption.element as HTMLInputElement).value).toBe("Option A");
  });
});
