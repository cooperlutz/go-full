import { describe, it, expect, vi } from "vitest";
import examLibraryRoutes from "../index";

// Mock the imported Vue components
vi.mock("~/examlibrary/views/ExamLibrary.vue", () => ({
  default: "ExamLibrary",
}));

describe("examLibraryRoutes", () => {
  it("should export an array of routes", () => {
    expect(Array.isArray(examLibraryRoutes)).toBe(true);
    expect(examLibraryRoutes.length).toBeGreaterThan(0);
  });

  it("should render", () => {
    const route = examLibraryRoutes.find((r) => r.path === "/exam-library");
    expect(route).toBeDefined();
    expect(route?.component).toBe("ExamLibrary");
  });
});
