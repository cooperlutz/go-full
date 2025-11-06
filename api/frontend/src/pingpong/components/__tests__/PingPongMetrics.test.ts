import PingPongMetrics from "../PingPongMetrics.vue";
import { describe, it, expect, vi } from "vitest";
import { mount } from "@vue/test-utils";
import { nextTick } from "vue";

// Mock the line-chart component used in PingPongMetrics.vue
vi.mock("~/app/components/Charts/LineChart.vue", () => ({
  default: {
    template: '<div class="line-chart-mock">Line Chart Mock</div>',
  },
}));

describe("PingPongMetrics", () => {
  it("renders correctly", () => {
    // Arrange & Act
    const wrapper = mount(PingPongMetrics);
    // Assert
    expect(wrapper.exists()).toBe(true);
  });

  it("fetches and displays metrics data", async () => {
    // Arrange
    const wrapper = mount(PingPongMetrics);

    // Act
    await new Promise((resolve) => setTimeout(resolve, 200));
    await nextTick();

    // Assert
    const metricsTotal = wrapper.find("#metrics-totalpingpongs");
    const metricsPings = wrapper.find("#metrics-totalpings");
    const metricsPongs = wrapper.find("#metrics-totalpongs");
    expect(metricsTotal.exists()).toBe(true);
    expect(metricsPings.exists()).toBe(true);
    expect(metricsPongs.exists()).toBe(true);
    expect(metricsTotal.text()).toContain("14");
    expect(metricsPings.text()).toContain("7");
    expect(metricsPongs.text()).toContain("7");
  });

  it("displays an error message when metrics fail to load", async () => {
    // Arrange
    // mock fetch to return an error
    global.fetch = vi.fn(() =>
      Promise.resolve({
        ok: false,
        statusText: "Internal Server Error",
      } as Response),
    ) as unknown as typeof fetch;

    const wrapper = mount(PingPongMetrics);

    // Act
    await new Promise((resolve) => setTimeout(resolve, 200));
    await nextTick();

    // Assert
    const errorMessage = wrapper.find("#metrics-error");
    expect(errorMessage.text()).toContain("Error loading metrics");
  });
});
