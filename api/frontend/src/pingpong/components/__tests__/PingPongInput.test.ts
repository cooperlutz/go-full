import PingPongInput from "../PingPongInput.vue";
import { describe, it, expect } from "vitest";
import { mount } from "@vue/test-utils";
import { nextTick } from "vue";

describe("PingPongInput", () => {
  it("renders input and send button", () => {
    // Arrange & Act
    const wrapper = mount(PingPongInput);

    // Assert
    expect(wrapper.find("#pingpong-input").exists()).toBe(true);
    expect(wrapper.find("#send-button").exists()).toBe(true);
  });

  it("calls sendPingPong with the input value and shows the response", async () => {
    // Arrange
    const wrapper = mount(PingPongInput);
    const input = wrapper.find("#pingpong-input");
    await input.setValue("ping");

    // Act
    await wrapper.find("#send-button").trigger("click");
    await new Promise((resolve) => setTimeout(resolve, 100));
    await nextTick();
    const swalPopup = document.querySelector(".swal2-container");

    // Assert
    expect(swalPopup).not.toBeNull();
    expect(swalPopup?.textContent).toContain("Pong!");
  });

  it("does something when you send nonsense", async () => {
    // Arrange
    const wrapper = mount(PingPongInput);
    const input = wrapper.find("#pingpong-input");
    await input.setValue("blahasdfkasdfkhj");

    // Act
    await wrapper.find("#send-button").trigger("click");
    await new Promise((resolve) => setTimeout(resolve, 100));
    await nextTick();
    const swalPopup = document.querySelector(".swal2-container");

    // Assert
    expect(swalPopup).not.toBeNull();
    expect(swalPopup?.textContent).toContain("undefined");
  });
});
