<script setup lang="ts">
const props = defineProps<{
  totalQuestions: number;
  answeredQuestions: number;
}>();

function calculateProgressPercentage(answered: number, total: number): number {
  if (total === 0) return 0;
  return (answered / total) * 100;
}
</script>

<template>
  <div class="flex flex-col space-y-2">
    Answered: {{ props.answeredQuestions || 0 }} /
    {{ props.totalQuestions || 0 }} ({{ calculateProgressPercentage(props.answeredQuestions, props.totalQuestions).toFixed(2) }}%)
    <progress
      v-if="props.totalQuestions !== 0 && props.answeredQuestions !== undefined"
      class="progress w-56 mt-1"
      :value="props.answeredQuestions"
      :max="props.totalQuestions"
    ></progress>
  </div>
</template>
