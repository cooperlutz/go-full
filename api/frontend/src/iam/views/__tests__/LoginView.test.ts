import { describe, it, expect } from "vitest";
import { mount } from "@vue/test-utils";
import LoginView from "../LoginView.vue";

describe("LoginView", () => {
  it("renders an input for exam name and grade level", () => {
    // Arrange & Act
    const wrapper = mount(LoginView);

    // Assert
    expect(wrapper.find("#login-input-email").exists()).toBe(true);
    expect(wrapper.find("#login-input-password").exists()).toBe(true);
    expect(wrapper.find("#login-button").exists()).toBe(true);
    expect(wrapper.find("#login-register-link").exists()).toBe(true);
  });

  it("logs a user in and redirects to home on valid form submission", async () => {
    // Arrange
    const wrapper = mount(LoginView);
    const emailInput = wrapper.find("#login-input-email");
    const passwordInput = wrapper.find("#login-input-password");

    // Act
    await emailInput.setValue("test@example.com");
    await passwordInput.setValue("ValidPass123");

    await wrapper.find("#login-button").trigger("click");
    await new Promise((resolve) => setTimeout(resolve, 100)); // wait for redirect
    // Assert
    expect(window.location.pathname).toBe("/dashboard");
  });
});
