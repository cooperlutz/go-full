import GradingView from "../GradingView.vue";
import { describe, it, expect, vi, beforeEach } from "vitest";
import { mount, flushPromises } from "@vue/test-utils";
import { nextTick } from "vue";

vi.mock("vue-router", () => ({
  useRoute: () => ({
    path: "/exam/5d9abb80-0706-42ad-8131-33627d3e6b17/question/1",
    params: {
      examId: "5d9abb80-0706-42ad-8131-33627d3e6b17",
      questionIndex: "1",
    },
  }),
  useRouter: () => ({
    path: "/exam",
  }),
}));

describe("GradingView", () => {
  beforeEach(() => {
    vi.clearAllMocks();
  });

  it("renders the table of ungraded exams", async () => {
    const wrapper = mount(GradingView);

    await flushPromises();
    await nextTick();

    const headers = wrapper.findAll("th");
    expect(headers).toHaveLength(4);
    expect(headers[0]?.text()).toBe("Exam ID");
    expect(headers[1]?.text()).toBe("Total Points Earned");
    expect(headers[2]?.text()).toBe("Total Points Possible");
    expect(headers[3]?.text()).toBe("Grading State");

    const rows = wrapper.findAll("tbody tr");
    expect(rows.length).toBeGreaterThan(0);
    const firstRowCells = rows[0]?.findAll("td");
    expect(firstRowCells?.length).toBe(4);
    expect(firstRowCells?.[0]?.text()).toBe(
      "5d9abb80-0706-42ad-8131-33627d3e6b17",
    );
    expect(firstRowCells?.[1]?.text()).toBe("85");
    expect(firstRowCells?.[2]?.text()).toBe("100");
    expect(firstRowCells?.[3]?.text()).toBe("not-started");
  });
});
