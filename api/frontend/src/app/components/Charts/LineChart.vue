<template>
  <div class="container">
    <LineChart :data="chartData" :options="chartOptions" />
  </div>
</template>

<script lang="ts">
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

export default {
  name: "LineChartComponent",
  components: {
    LineChart,
  },
  props: {
    chartData: {
      type: Object,
      required: true,
    },
    chartOptions: {
      type: Object,
      default: () => ({
        scales: {
          y: {
            beginAtZero: true,
          },
        },
      }),
    },
  },
};
</script>
