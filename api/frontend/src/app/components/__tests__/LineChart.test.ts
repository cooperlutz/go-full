import { describe, it, expect } from "vitest";
import { defineComponent } from "vue";
import LineChart from "../Charts/LineChart.vue";
import { mount } from "@vue/test-utils";

describe("LineChart.vue", () => {
  it("renders LineChart component with correct props", () => {
    // Define a test component to mount LineChart with props
    const TestComponent = defineComponent({
      components: { LineChart },
      template: `<LineChart ref="lineChart" :chartData="data" :options="options" />`,
      setup() {
        const data = {
          labels: ["January", "February", "March"],
          datasets: [{ label: "My Dataset", data: [40, 20, 12] }],
        };
        const options = { responsive: true };
        return { data, options };
      },
    });

    const wrapper = mount(TestComponent);
    expect(wrapper.findComponent(LineChart).exists()).toBe(true);
  });
});
