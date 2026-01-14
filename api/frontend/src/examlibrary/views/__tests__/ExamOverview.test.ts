import { describe, it, expect, vi } from "vitest";
import { mount } from "@vue/test-utils";
import { nextTick } from "vue";
import ExamOverview from "../ExamOverview.vue";

describe("ExamOverview", () => {
  it("renders exam data in table rows", async () => {
    // Arrange
    vi.mock("vue-router", () => ({
      useRoute: () => ({
        params: { id: "f660452b-4075-4eac-b87a-a5b1ce7bd428" },
        path: "/examlibrary/exams/f660452b-4075-4eac-b87a-a5b1ce7bd428",
      }),
    }));
    const wrapper = mount(ExamOverview);
    await nextTick();

    // Act
    await new Promise((resolve) => setTimeout(resolve, 100));
    await nextTick();

    // Assert
    const card = wrapper.find(".card");
    expect(card.exists()).toBe(true);
    expect(card.find("h2").text()).toBe("Exam Details");
    const pTags = card.findAll("p");
    expect(pTags[0]?.text()).toBe("ID: f660452b-4075-4eac-b87a-a5b1ce7bd428");
    expect(pTags[1]?.text()).toBe("Name: Sample Exam");
    expect(pTags[2]?.text()).toBe("Grade Level: 6");
  });

  it("shows loading state initially", async () => {
    // Arrange & Act
    const wrapper = mount(ExamOverview);
    await nextTick();

    // Assert
    expect(wrapper.text()).toContain("Loading exam...");
  });

  it("shows error message on fetch failure", async () => {
    // Arrange
    global.fetch = vi.fn(() =>
      Promise.resolve({
        ok: false,
        statusText: "Internal Server Error",
      } as Response),
    ) as unknown as typeof fetch;

    const wrapper = mount(ExamOverview);

    // Act
    await new Promise((resolve) => setTimeout(resolve, 100));
    await nextTick();

    // Assert
    expect(wrapper.text()).toContain(
      "Error loading exam: ResponseError: Response returned an error code",
    );
  });
});
