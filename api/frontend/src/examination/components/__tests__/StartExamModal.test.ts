import StartExamModal from "../StartExamModal.vue";
import { describe, it, expect, vi } from "vitest";
import { mount } from "@vue/test-utils";
import { nextTick } from "vue";

describe("StartExamModal", () => {
  it("renders input and send button", () => {
    // Arrange & Act
    const wrapper = mount(StartExamModal, {
      props: {
        examId: "test-exam-id",
      },
    });

    // Assert
    expect(wrapper.find("#start_exam_button").exists()).toBe(true);
  });

  it("opens modal and input box is accessible on button click", async () => {
    // Arrange
    const wrapper = mount(StartExamModal, {
      props: {
        examId: "test-exam-id",
      },
    });
    const button = wrapper.find("#start_exam_button");

    // Act
    await button.trigger("click");
    await nextTick();

    // Assert
    const studentIdInput = wrapper.find("#student-id-input");
    expect(studentIdInput.exists()).toBe(true);
  });

  it("sends correct data on start exam", async () => {
    const fetchSpy = vi.spyOn(window, "fetch");
    const wrapper = mount(StartExamModal, {
      props: {
        examId: "test-exam-id",
      },
    });
    const button = wrapper.find("#start_exam_button");
    await button.trigger("click");
    await nextTick();
    const studentIdInput = wrapper.find("#student-id-input");
    await studentIdInput.setValue("123e4567-e89b-12d3-a456-426614174000");
    const startButton = wrapper.find("#start-button");
    await startButton.trigger("click");
    await nextTick();

    expect(fetchSpy).toHaveBeenCalledTimes(1);
    expect(fetchSpy).toHaveBeenCalledWith(
      expect.stringContaining("/examination/api/v1/exams"),
      expect.objectContaining({
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          studentId: "123e4567-e89b-12d3-a456-426614174000",
          libraryExamId: "test-exam-id",
        }),
      }),
    );
  });
});
