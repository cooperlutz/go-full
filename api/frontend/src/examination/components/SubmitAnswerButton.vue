<script setup lang="ts">
import { useSubmitAnswer } from "../composables/useExamination";

const props = defineProps<{
  examId: string;
  questionIndex: number;
  answer: string;
}>();

const { submitAnswer, loading } = useSubmitAnswer();
</script>

<template>
  <div class="flex justify-end">
    <button
      id="submit-question-button"
      class="btn btn-outline"
      @click="submitAnswer(props.examId, props.questionIndex, props.answer)"
      :disabled="loading"
    >
      Submit
    </button>
  </div>
</template>
