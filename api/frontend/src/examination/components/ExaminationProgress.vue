<script setup lang="ts">
import { onMounted } from "vue";
import { useRoute } from "vue-router";

import { useExaminationProgress } from "../composables/useExamination";

const { progressPercentage, progress, loading, error, getExamProgress } =
  useExaminationProgress();

const examIdFromRoute = useRoute().params.id as string;

onMounted(async () => {
  const examId = examIdFromRoute;
  await getExamProgress(examId);
});
</script>

<template>
  <div class="flex flex-col items-center space-y-2">
    Answered: {{ progress?.answeredQuestions || 0 }} /
    {{ progress?.totalQuestions || 0 }} ({{ progressPercentage.toFixed(2) }}%)
    <progress
      v-if="progress !== null"
      class="progress w-56"
      :value="progress.answeredQuestions"
      :max="progress.totalQuestions"
    ></progress>

    <div v-else-if="loading">
      <progress class="progress w-56"></progress>
    </div>
    <div v-else-if="error" id="examination-progress-error">
      <div role="alert" class="alert alert-error">
        <span>Error! Task failed successfully.</span>
      </div>
    </div>
    <div v-else>No progress available.</div>
  </div>
</template>
