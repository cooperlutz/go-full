import { describe, it, expect } from "vitest";
import { createNewApp } from "../main";

describe("createNewApp", () => {
  it("should create and return a Vue app instance", () => {
    const app = createNewApp();
    expect(app).toBeDefined();
  });
});
