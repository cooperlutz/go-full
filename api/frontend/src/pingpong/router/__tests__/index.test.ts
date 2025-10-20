import { describe, it, expect, vi } from "vitest";
import pingpongRoutes from "../index";

// Mock the imported Vue components
vi.mock("~/pingpong/views/PingPongAppView.vue", () => ({
  default: "PingPongAppView",
}));
vi.mock("~/pingpong/views/PingPongAnalyticsView.vue", () => ({
  default: "PingPongAnalyticsView",
}));

describe("pingpongRoutes", () => {
  it("should export an array of routes", () => {
    expect(Array.isArray(pingpongRoutes)).toBe(true);
    expect(pingpongRoutes.length).toBe(3);
  });

  it("should have a redirect from /ping-pong to /ping-pong/app", () => {
    const route = pingpongRoutes.find((r) => r.path === "/ping-pong");
    expect(route).toBeDefined();
    expect(route?.redirect).toBe("/ping-pong/app");
  });

  it("should have /ping-pong/app route with PingPongAppView component", () => {
    const route = pingpongRoutes.find((r) => r.path === "/ping-pong/app");
    expect(route).toBeDefined();
    expect(route?.component).toBe("PingPongAppView");
  });

  it("should have /ping-pong/analytics route with PingPongAnalyticsView component", () => {
    const route = pingpongRoutes.find((r) => r.path === "/ping-pong/analytics");
    expect(route).toBeDefined();
    expect(route?.component).toBe("PingPongAnalyticsView");
  });
});
