import PongButton from "../PingPongButtons/PongButton.vue";
import { describe, it, expect } from "vitest";
import { mount } from "@vue/test-utils";
import { nextTick } from "vue";

describe("PongButton", () => {
  it('renders a button with text "pong"', () => {
    // Arrange & Act
    const wrapper = mount(PongButton);

    // Assert
    expect(wrapper.text()).toContain("pong");
  });

  it("calls sendPingPong with 'pong' and shows the response message on click", async () => {
    // Arrange
    const wrapper = mount(PongButton);

    // Act
    await wrapper.find("#pong-button").trigger("click");
    await new Promise((resolve) => setTimeout(resolve, 100));
    await nextTick();

    // Assert
    const swalPopup = document.querySelector(".swal2-container");
    expect(swalPopup).not.toBeNull();
    expect(swalPopup?.textContent).toContain("Ping!");
    expect(
      swalPopup?.querySelector(".swal2-popup")?.getAttribute("style"),
    ).toContain("background: #e11d48;");
  });
});
