import { describe, it, expect, vi } from "vitest";

import { mount } from "@vue/test-utils";
import PingPongAppView from "../PingPongAppView.vue";

// Mock the currentroute.path
vi.mock("vue-router", () => ({
  useRoute: () => ({
    path: "/ping-pong/app",
  }),
}));

describe("PingPongAppView.vue", () => {
  it("should render correctly", () => {
    const wrapper = mount(PingPongAppView);
    expect(wrapper.exists()).toBe(true);
  });
});
