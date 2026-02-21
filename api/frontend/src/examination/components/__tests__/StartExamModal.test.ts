import StartExamModal from "../StartExamModal.vue";
import { describe, it, expect, vi } from "vitest";
import { mount } from "@vue/test-utils";
import { nextTick } from "vue";

describe("StartExamModal", () => {
  it("renders input and send button", () => {
    // Arrange & Act
    const wrapper = mount(StartExamModal, {
      props: {
        libraryExamId: "test-exam-id",
      },
    });

    // Assert
    expect(wrapper.find("#start-exam-modal-button").exists()).toBe(true);
  });

  it("opens modal and input box is accessible on button click", async () => {
    // Arrange
    const wrapper = mount(StartExamModal, {
      props: {
        libraryExamId: "test-exam-id",
      },
    });
    const button = wrapper.find("#start-exam-modal-button");

    // Act
    await button.trigger("click");
    await nextTick();

    // Assert
    const studentIdInput = wrapper.find("#confirm-start-exam-button");
    expect(studentIdInput.exists()).toBe(true);
  });

  it("sends correct data on start exam", async () => {
    const fetchSpy = vi.spyOn(window, "fetch");
    const wrapper = mount(StartExamModal, {
      props: {
        libraryExamId: "test-exam-id",
      },
    });
    const button = wrapper.find("#start-exam-modal-button");
    await button.trigger("click");
    await nextTick();
    const startButton = wrapper.find("#confirm-start-exam-button");
    await startButton.trigger("click");
    await nextTick();

    expect(fetchSpy).toHaveBeenCalledTimes(1);
    expect(fetchSpy).toHaveBeenCalledWith(
      expect.stringContaining("/api/iam/profile"),
      expect.objectContaining({
        method: "GET",
        headers: expect.objectContaining({
          Authorization: expect.stringContaining("Bearer "),
        }),
      }),
    );
  });
});
