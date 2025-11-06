/* STEP 5.3. Implement Frontend Component Tests
here we implement a the frontend UI component tests
we can use the test to describe what we want the component to ultimately look like and do before we start implementing
the actual component itself
*/
import PingPongLookup from "../PingPongLookup.vue";
import { describe, it, expect } from "vitest";
import { mount } from "@vue/test-utils";
import { nextTick } from "vue";

describe("PingPongLookup", () => {
  it("renders the lookup input and button", () => {
    // Arrange & Act
    const wrapper = mount(PingPongLookup);
    // Assert
    expect(wrapper.find("#pingpong-lookup").exists()).toBe(true);
    expect(wrapper.find("#lookup-button").exists()).toBe(true);
  });

  it("calls lookup with a provided value of a given pingpong id and displays the result", async () => {
    // Arrange
    const wrapper = mount(PingPongLookup);
    const input = wrapper.find("#pingpong-lookup");

    // Act
    await input.setValue("f660452b-4075-4eac-b87a-a5b1ce7bd428");
    await wrapper.find("#lookup-button").trigger("click");
    await new Promise((resolve) => setTimeout(resolve, 100));
    await nextTick();

    // Assert
    const output = wrapper.find("#pingpong-lookup-card");
    expect(output.text()).toContain("f660452b-4075-4eac-b87a-a5b1ce7bd428");
    expect(output.text()).toContain("pong");
  });

  it("displays an error message for a non-existent ID", async () => {
    // Arrange
    const wrapper = mount(PingPongLookup);
    const input = wrapper.find("#pingpong-lookup");

    // Act
    await input.setValue("non-existent-id");
    await wrapper.find("#lookup-button").trigger("click");
    await new Promise((resolve) => setTimeout(resolve, 100));
    await nextTick();

    // Assert
    const errorMessage = wrapper.find("#pingpong-lookup-error");
    const output = wrapper.find("#pingpong-lookup-card");
    expect(output.exists()).toBe(false);
    expect(errorMessage.text()).toContain("Error loading ping pong");
  });
});
