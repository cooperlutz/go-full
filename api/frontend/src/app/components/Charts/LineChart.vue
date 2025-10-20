<script lang="ts">
import { defineComponent, type PropType } from "vue";
import {
  LineElement,
  PointElement,
  CategoryScale,
  Chart as ChartJS,
  Legend,
  LinearScale,
  Title,
  Tooltip,
} from "chart.js";
import { Line as LineChart } from "vue-chartjs";

ChartJS.register(
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Legend,
);

ChartJS.defaults.font.family = "JetBrains Mono";

export default defineComponent({
  name: "LineChartComponent",
  components: {
    // cast to a neutral type to avoid leaking vue-chartjs types into the SFC export
    LineChart: LineChart as unknown as Record<string, unknown>,
  },
  props: {
    chartData: {
      type: Object as PropType<Record<string, unknown>>,
      required: true,
    },
    chartOptions: {
      type: Object as PropType<Record<string, unknown>>,
      default: () => ({
        scales: {
          y: {
            beginAtZero: true,
          },
        },
      }),
    },
  },
});
</script>

<template>
  <div class="container">
    <LineChart :data="chartData" :options="chartOptions" />
  </div>
</template>
