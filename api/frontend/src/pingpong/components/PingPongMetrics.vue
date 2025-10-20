<script setup lang="ts">
import { onMounted, ref } from "vue";

import { type ChartData } from "chart.js";

import SectionDivider from "~/app/components/SectionDivider/SectionDivider.vue";
import LineChartComponent from "~/app/components/Charts/LineChart.vue";

import { usePingPongMetrics } from "~/pingpong/composables/usePingPongMetrics";

const graphData = ref<ChartData>({ labels: [], datasets: [] });

const {
  state,
  actions,
  totalNumberOfPingPongs,
  totalNumberOfPingPongsPerDay,
  totalNumberOfPings,
  totalNumberOfPongs,
} = usePingPongMetrics();

onMounted(async () => {
  await actions.fetchData();

  graphData.value = {
    labels: totalNumberOfPingPongsPerDay.value.labels || ["No Data"],
    datasets: [
      {
        label: "Number of PingPongs Created",
        backgroundColor: "#f87979",
        data: totalNumberOfPingPongsPerDay.value.values || [0],
      },
    ],
  };
});
</script>

<template>
  <div :v-if="!state.loading">
    <!-- section header -->
    <h2>Metrics</h2>

    <!-- Total Number of PingPongs -->
    <div class="stats shadow border border-secondary">
      <div class="stat place-items-center">
        <div class="stat-title">Total Number of PingPongs</div>
        <div class="stat-value">
          {{ totalNumberOfPingPongs }}
        </div>
      </div>

      <div class="stat place-items-center">
        <div class="stat-title">Total Number of Pings</div>
        <div class="stat-value">
          {{ totalNumberOfPings }}
        </div>
      </div>

      <div class="stat place-items-center">
        <div class="stat-title">Total Number of Pongs</div>
        <div class="stat-value">
          {{ totalNumberOfPongs }}
        </div>
      </div>
    </div>
  </div>

  <SectionDivider />

  <!-- Line Chart -->
  <div
    class="card w-full bg-base-100 shadow-md card-border border-secondary border-solid"
  >
    <div v-if="!graphData">Loading graph...</div>
    <LineChartComponent v-else :chartData="graphData"></LineChartComponent>
  </div>
</template>
