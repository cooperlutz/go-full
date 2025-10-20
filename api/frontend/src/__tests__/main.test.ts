import { describe, it, expect } from "vitest";
import { createNewApp } from "../main";

// import router from "~/app/router";

describe("createNewApp", () => {
  it("should create and return a Vue app instance", () => {
    const app = createNewApp();
    expect(app).toBeDefined();
    // You can add more specific assertions here, e.g., checking if plugins are installed
  });
});
