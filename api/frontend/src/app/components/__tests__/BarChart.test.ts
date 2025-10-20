import { describe, it, expect } from "vitest";
import { defineComponent } from "vue";
import Bar from "../Charts/BarChart.vue";
import { mount } from "@vue/test-utils";

describe("Bar.vue", () => {
  it("renders Bar component with correct props", () => {
    // Define a test component to mount Bar with props
    const TestComponent = defineComponent({
      components: { Bar },
      template: `<Bar ref="barChart" :data="data" :options="options" />`,
      setup() {
        const data = {
          labels: ["January", "February", "March"],
          datasets: [{ data: [40, 20, 12] }],
        };
        const options = { responsive: true };
        return { data, options };
      },
    });

    const wrapper = mount(TestComponent);
    expect(wrapper.findComponent(Bar).exists()).toBe(true);
  });
});
