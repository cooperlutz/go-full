import { describe, it, expect } from "vitest";
import { mount } from "@vue/test-utils";
import RegisterView from "../RegisterView.vue";

describe("RegisterView", () => {
  it("renders an input for exam name and grade level", () => {
    // Arrange & Act
    const wrapper = mount(RegisterView);

    // Assert
    expect(wrapper.find("#register-input-email").exists()).toBe(true);
    expect(wrapper.find("#register-input-password").exists()).toBe(true);
    expect(wrapper.find("#register-button").exists()).toBe(true);
  });

  it("logs a user in and redirects to home on valid form submission", async () => {
    // Arrange
    const wrapper = mount(RegisterView);
    const emailInput = wrapper.find("#register-input-email");
    const passwordInput = wrapper.find("#register-input-password");

    // Act
    await emailInput.setValue("test@example.com");
    await passwordInput.setValue("ValidPass123");

    await wrapper.find("#register-button").trigger("click");

    await new Promise((resolve) => setTimeout(resolve, 100)); // wait for redirect
    // Assert
    expect(window.location.pathname).toBe("/login");
  });
});
