<script setup lang="ts">
import { onMounted } from "vue";

import SectionDivider from "~/app/components/SectionDivider/SectionDivider.vue";

import { useGetMetrics } from "../composables/useReporting";

const {
  error,
  loading,
  getMetrics,
  numberOfExamsInProgress,
  numberOfExamsCompleted,
  numberOfExamsBeingGraded,
  numberOfExamsGradingCompleted,
} = useGetMetrics();

onMounted(async () => {
  await getMetrics();
});
</script>

<template>
  <div v-if="!loading && !error">
    <!-- section header -->
    <h2>Metrics</h2>

    <div class="stats shadow border border-secondary">
      <div id="metrics-num-exams-in-progress" class="stat place-items-center">
        <div class="stat-title">Number of Exams In Progress</div>
        <div class="stat-value">
          {{ numberOfExamsInProgress?.metricValue || 0 }}
        </div>
      </div>
    </div>
    <div class="stats shadow border border-secondary">
      <div id="metrics-num-exams-completed" class="stat place-items-center">
        <div class="stat-title">Number of Exams Completed</div>
        <div class="stat-value">
          {{ numberOfExamsCompleted?.metricValue || 0 }}
        </div>
      </div>
    </div>
    <div class="stats shadow border border-secondary">
      <div id="metrics-num-exams-being-graded" class="stat place-items-center">
        <div class="stat-title">Number of Exams Being Graded</div>
        <div class="stat-value">
          {{ numberOfExamsBeingGraded?.metricValue || 0 }}
        </div>
      </div>
    </div>
    <div class="stats shadow border border-secondary">
      <div
        id="metrics-num-exams-grading-completed"
        class="stat place-items-center"
      >
        <div class="stat-title">Number of Exams Grading Completed</div>
        <div class="stat-value">
          {{ numberOfExamsGradingCompleted?.metricValue || 0 }}
        </div>
      </div>
    </div>
  </div>

  <div v-else-if="loading">Loading metrics...</div>

  <div v-else-if="error" id="metrics-error">
    Error loading metrics: {{ error }}
  </div>

  <SectionDivider />
</template>
