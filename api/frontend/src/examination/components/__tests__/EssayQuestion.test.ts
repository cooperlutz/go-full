import EssayQuestion from "../Questions/EssayQuestion.vue";
import { describe, it, expect } from "vitest";
import { mount } from "@vue/test-utils";
import { type Question } from "~/examination/services";

const sampleQuestion: Question = {
  questionId: "q1",
  examId: "exam1",
  questionType: "essay-question",
  questionText: "Explain the theory of relativity.",
  questionIndex: 1,
  answered: false,
};

const wrapper = mount(EssayQuestion, {
  props: {
    question: sampleQuestion,
  },
});

describe("EssayQuestion", () => {
  it("renders input and send button", () => {
    // Assert
    expect(wrapper.find("#record-answer-button").exists()).toBe(true);
  });

  it("allows user to input an essay answer", () => {
    const inputOption = wrapper.find("#essay-question-input");
    inputOption.setValue("this is my essay answer");
    const submitButton = wrapper.find("#record-answer-button");
    submitButton.trigger("click");

    // Assert
    expect((inputOption.element as HTMLInputElement).value).toBe(
      "this is my essay answer",
    );
  });
});
