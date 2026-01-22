import { describe, it, expect, vi } from "vitest";
import examinationRoutes from "../index";

// Mock the imported Vue components
vi.mock("~/examination/views/ExaminationView.vue", () => ({
  default: "ExaminationView",
}));

describe("examinationRoutes", () => {
  it("should export an array of routes", () => {
    expect(Array.isArray(examinationRoutes)).toBe(true);
    expect(examinationRoutes.length).toBe(1);
  });

  it("should render", () => {
    const route = examinationRoutes.find((r) => r.path === "/exam/:id");
    expect(route).toBeDefined();
    expect(route?.component).toBe("ExaminationView");
  });
});
