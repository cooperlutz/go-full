<script setup lang="ts">
import { useSubmitAnswer } from "../../composables/useExamination";

const props = defineProps<{
  examId: string;
  questionIndex: number;
  answer: string;
}>();
const emit = defineEmits(["question-answered"]);

const { submitAnswer, loading } = useSubmitAnswer();
</script>

<template>
  <div class="flex justify-end">
    <button
      id="record-answer-button"
      class="btn btn-outline"
      @click="submitAnswer(props.examId, props.questionIndex, props.answer).then(() => emit('question-answered'))"
      :disabled="loading"
    >
      Record Answer
    </button>
  </div>
</template>
