import { test, expect } from "vitest";
import { mount } from "@vue/test-utils";
import Error404View from "../Error404View.vue";

test("renders correctly", () => {
  const wrapper = mount(Error404View);
  expect(wrapper.exists()).toBe(true);
});
