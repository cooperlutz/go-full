import QuestionNav from "../QuestionNav.vue";
import { describe, it, expect } from "vitest";
import { mount } from "@vue/test-utils";
import { type Question } from "~/examination/services";

const sampleQuestionShort: Question = {
  questionId: "q1",
  examId: "exam1",
  questionType: "short-answer",
  questionText: "Provide a brief explanation.",
  questionIndex: 1,
  answered: false,
};
const sampleQuestionMC: Question = {
  questionId: "q1",
  examId: "exam1",
  questionType: "multiple-choice",
  questionText: "choose the correct option.",
  questionIndex: 1,
  responseOptions: ["Option A", "Option B", "Option C", "Option D"],
  answered: false,
};
const sampleQuestions: Question[] = [sampleQuestionShort, sampleQuestionMC];

const wrapper = mount(QuestionNav, {
  props: {
    examId: "exam1",
    questions: sampleQuestions,
  },
});

describe("ExamSubmission", () => {
  it("renders input and send button", () => {
    // Assert
    expect(wrapper.find("#question-nav-item-1").exists()).toBe(true);
  });

  it("allows user to navigate to question via nav item", () => {
    const navItem = wrapper.find("#question-nav-item-1");
    navItem.trigger("click");

    // Assert
    expect(navItem.attributes("href")).toBe("/exam/exam1/question/1");
    expect(navItem.text()).toBe("Question 1");
  });
});
