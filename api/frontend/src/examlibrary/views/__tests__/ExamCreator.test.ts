import { describe, it, expect } from "vitest";
import { mount } from "@vue/test-utils";
import ExamCreator from "../ExamCreator.vue";
import { nextTick } from "vue";

describe("ExamCreator", () => {
  it("renders an input for exam name and grade level", () => {
    // Arrange & Act
    const wrapper = mount(ExamCreator);

    // Assert
    expect(wrapper.find("#new-exam-name-input").exists()).toBe(true);
    expect(wrapper.find("#new-exam-grade-level-input").exists()).toBe(true);
  });

  it("creates and returns a new exam object on form fill and submission", async () => {
    // Arrange
    const wrapper = mount(ExamCreator);
    const nameInput = wrapper.find("#new-exam-name-input");
    const gradeLevelInput = wrapper.find("#new-exam-grade-level-input");

    // Act
    await nameInput.setValue("Sample Exam");
    await gradeLevelInput.setValue(5);

    await wrapper.find("#add-multiple-choice-button").trigger("click");
    await nextTick();

    const mcQuestionTextInput = wrapper.find("#mc-question-text-input");
    const mcOption1Input = wrapper.find("#mc-option-a-input");
    const mcOption2Input = wrapper.find("#mc-option-b-input");
    const mcOption3Input = wrapper.find("#mc-option-c-input");
    const mcOption4Input = wrapper.find("#mc-option-d-input");
    const mcCorrectOptionSelect = wrapper.find("#mc-correct-answer-input");
    const mcPossiblePointsInput = wrapper.find("#mc-possible-points-input");

    await mcQuestionTextInput.setValue("What is 2 + 2?");
    await mcOption1Input.setValue("3");
    await mcOption2Input.setValue("4");
    await mcOption3Input.setValue("5");
    await mcOption4Input.setValue("6");
    await mcCorrectOptionSelect.setValue("4");
    await mcPossiblePointsInput.setValue(10);

    await wrapper.find("#add-multiple-choice-confirm-button").trigger("click");

    // wait
    await nextTick();

    // Assert
    const newExamOutput = wrapper.find("#new-exam-inputs");
    expect(newExamOutput.exists()).toBe(true);

    // expect the output to contain the exam name and grade level
    expect(newExamOutput.text()).toContain("Sample Exam");
    expect(newExamOutput.text()).toContain("5");
    // expect the output to contain the added multiple choice question details
    expect(newExamOutput.text()).toContain("What is 2 + 2?");
    expect(newExamOutput.text()).toContain("3");
    expect(newExamOutput.text()).toContain("4");
    expect(newExamOutput.text()).toContain("5");
    expect(newExamOutput.text()).toContain("6");
    expect(newExamOutput.text()).toContain("4");
    expect(newExamOutput.text()).toContain("10");

    await wrapper.find("#add-short-answer-button").trigger("click");
    await nextTick();

    const saQuestionTextInput = wrapper.find("#sa-question-text-input");
    const saPossiblePointsInput = wrapper.find("#sa-possible-points-input");

    await saQuestionTextInput.setValue("Explain the theory of relativity.");
    await saPossiblePointsInput.setValue(15);

    await wrapper.find("#add-short-answer-confirm-button").trigger("click");

    // wait
    await nextTick();

    // Assert
    expect(newExamOutput.text()).toContain("Explain the theory of relativity.");
    expect(newExamOutput.text()).toContain("15");

    await wrapper.find("#add-essay-question-button").trigger("click");
    await nextTick();

    const essayQuestionTextInput = wrapper.find("#essay-question-text-input");
    const essayPossiblePointsInput = wrapper.find(
      "#essay-possible-points-input",
    );

    await essayQuestionTextInput.setValue(
      "Discuss the impacts of climate change.",
    );
    await essayPossiblePointsInput.setValue(20);

    await wrapper.find("#add-essay-question-confirm-button").trigger("click");

    // wait
    await nextTick();

    // Assert
    expect(newExamOutput.text()).toContain(
      "Discuss the impacts of climate change.",
    );
    expect(newExamOutput.text()).toContain("20");

    // Finally, submit the exam
    await wrapper.find("#save-exam-to-library-button").trigger("click");
  });
});
