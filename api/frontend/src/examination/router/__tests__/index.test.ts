import { describe, it, expect, vi } from "vitest";
import examinationRoutes from "../index";

// Mock the imported Vue components
vi.mock("~/examination/views/ExaminationOverview.vue", () => ({
  default: "ExaminationOverview",
}));

describe("examinationRoutes", () => {
  it("should export an array of routes", () => {
    expect(Array.isArray(examinationRoutes)).toBe(true);
    expect(examinationRoutes.length).toBe(3);
  });

  it("should render", () => {
    const route = examinationRoutes.find((r) => r.path === "/exam/:id");
    expect(route).toBeDefined();
    expect(route?.component).toBe("ExaminationOverview");
  });
});
