import ExaminationView from "../ExaminationView.vue";
import { describe, it, expect, vi, beforeEach } from "vitest";
import { mount, flushPromises } from "@vue/test-utils";
import { nextTick } from "vue";

vi.mock("vue-router", () => ({
  useRoute: () => ({
    path: "/exam/5d9abb80-0706-42ad-8131-33627d3e6b17/question/1",
    params: { id: "5d9abb80-0706-42ad-8131-33627d3e6b17", index: "1" },
  }),
  useRouter: () => ({
    push: vi.fn(),
  }),
}));

describe("ExaminationView", () => {
  beforeEach(() => {
    vi.clearAllMocks();
  });

  it("displays exam data when exam is loaded", async () => {
    const wrapper = mount(ExaminationView);

    await flushPromises();
    await nextTick();

    const examDiv = wrapper.find("#examination-view-component");
    expect(examDiv.exists()).toBe(true);
  });
});
