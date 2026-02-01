import ShortAnswer from "../Questions/ShortAnswer.vue";
import { describe, it, expect } from "vitest";
import { mount } from "@vue/test-utils";
import { type Question } from "~/examination/services";

const sampleQuestion: Question = {
  questionId: "q1",
  examId: "exam1",
  questionType: "short-answer",
  questionText: "Provide a brief explanation.",
  questionIndex: 1,
  answered: false,
};

const wrapper = mount(ShortAnswer, {
  props: {
    question: sampleQuestion,
  },
});

describe("ShortAnswer", () => {
  it("renders input and send button", () => {
    // Assert
    expect(wrapper.find("#record-answer-button").exists()).toBe(true);
    expect(wrapper.find("#short-answer-input").exists()).toBe(true);
  });

  it("allows user to input an answer", () => {
    const inputOption = wrapper.find("#short-answer-input");
    inputOption.setValue("this is my short answer");
    const submitButton = wrapper.find("#record-answer-button");
    submitButton.trigger("click");

    // Assert
    expect((inputOption.element as HTMLInputElement).value).toBe(
      "this is my short answer",
    );
  });
});
