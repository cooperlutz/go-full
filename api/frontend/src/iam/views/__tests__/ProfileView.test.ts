import { describe, it, expect } from "vitest";
import { mount } from "@vue/test-utils";
import ProfileView from "../ProfileView.vue";

describe("ProfileView", () => {
  it("renders properly", () => {
    const wrapper = mount(ProfileView);
    expect(wrapper.exists()).toBe(true);
  });

  it("displays the user ID", async () => {
    const wrapper = mount(ProfileView);
    // wait for onMounted async operations
    await new Promise((resolve) => setTimeout(resolve, 0));
    const userId = wrapper.find("#user-id");
    expect(userId.exists()).toBe(true);
    expect(userId.text()).toBe("1f23abc456def7890ghi");
  });

  it("displays the user email", async () => {
    const wrapper = mount(ProfileView);
    // wait for onMounted async operations
    await new Promise((resolve) => setTimeout(resolve, 0));
    const userEmail = wrapper.find("#user-email");
    expect(userEmail.exists()).toBe(true);
    expect(userEmail.text()).toBe("email@example.com");
  });
});
