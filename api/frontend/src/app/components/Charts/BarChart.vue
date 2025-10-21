<script lang="ts">
import { defineComponent } from "vue";
import {
  BarElement,
  CategoryScale,
  Chart as ChartJS,
  Legend,
  LinearScale,
  Title,
  Tooltip,
} from "chart.js";
import { Bar } from "vue-chartjs";

ChartJS.register(
  CategoryScale,
  LinearScale,
  BarElement,
  Title,
  Tooltip,
  Legend,
);

export default defineComponent({
  name: "App",
  components: {
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    Bar: Bar as any,
  },
  data(): {
    data: { labels: string[]; datasets: { data: number[] }[] };
    options: { responsive: boolean };
  } {
    return {
      data: {
        labels: ["January", "February", "March"],
        datasets: [{ data: [40, 20, 12] }],
      },
      options: {
        responsive: true,
      },
    };
  },
});
</script>

<template>
  <Bar :data="data" :options="options" />
</template>
