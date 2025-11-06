import PingButton from "../PingPongButtons/PingButton.vue";
import { describe, it, expect } from "vitest";
import { mount } from "@vue/test-utils";
import { nextTick } from "vue";

describe("PingButton", () => {
  it('renders a button with text "ping"', () => {
    // Arrange & Act
    const wrapper = mount(PingButton);
    // Assert
    expect(wrapper.text()).toContain("ping");
  });

  it("calls sendPingPong with 'ping' and shows the response message on click", async () => {
    // Arrange
    const wrapper = mount(PingButton);
    await wrapper.find("#ping-button").trigger("click");

    // Act
    await new Promise((resolve) => setTimeout(resolve, 100));
    await nextTick();

    // Assert
    const swalPopup = document.querySelector(".swal2-container");
    expect(swalPopup).not.toBeNull();
    expect(swalPopup?.textContent).toContain("Pong!");
    expect(
      swalPopup?.querySelector(".swal2-popup")?.getAttribute("style"),
    ).toContain("background: #0ea5e9;");
  });
});
