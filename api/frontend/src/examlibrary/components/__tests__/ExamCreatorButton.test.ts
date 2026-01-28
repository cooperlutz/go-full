import { describe, it, expect, vi } from "vitest";
import { mount } from "@vue/test-utils";
import { createRouter, createWebHistory } from "vue-router";
import ExamCreatorButton from "../ExamCreatorButton.vue";

describe("ExamCreatorButton", () => {
  const mockPush = vi.fn();
  const router = createRouter({
    history: createWebHistory(),
    routes: [
      { path: "/", name: "Home", component: { template: "<div>Home</div>" } },
      {
        path: "/exam-creator",
        name: "ExamCreator",
        component: { template: "<div>Exam Creator</div>" },
      },
    ],
  });

  // Mock the router push method
  router.push = mockPush;

  it("renders the button with correct text", () => {
    const wrapper = mount(ExamCreatorButton, {
      global: {
        plugins: [router],
      },
    });

    const button = wrapper.find("button");
    expect(button.exists()).toBe(true);
    expect(button.text()).toBe("New Exam");
  });

  it("navigates to ExamCreator route when clicked", async () => {
    const wrapper = mount(ExamCreatorButton, {
      global: {
        plugins: [router],
      },
    });

    const button = wrapper.find("button");
    await button.trigger("click");

    expect(mockPush).toHaveBeenCalledWith({ name: "ExamCreator" });
  });
});
