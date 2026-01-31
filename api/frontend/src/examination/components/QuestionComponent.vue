<script setup lang="ts">
import { onMounted } from "vue";

import MultipleChoice from "./Questions/MultipleChoice.vue";
import ShortAnswer from "./Questions/ShortAnswer.vue";
import EssayQuestion from "./Questions/EssayQuestion.vue";
import { useGetQuestion } from "../composables/useExamination";

const props = defineProps<{
  examId: string;
  questionIndex: number;
}>();

const { getQuestion, examQuestion, loading, error } = useGetQuestion();

onMounted(async () => {
  await getQuestion(props.examId, props.questionIndex);
});
</script>

<template>
  <div
    class="card w-full bg-base-100 shadow-lg card-border border-secondary border-solid"
  >
    <div class="card-body">
      <MultipleChoice
        v-if="examQuestion?.questionType === 'multiple-choice'"
        :question="examQuestion"
      />
      <ShortAnswer
        v-else-if="examQuestion?.questionType === 'short-answer'"
        :question="examQuestion"
      />
      <EssayQuestion
        v-else-if="examQuestion?.questionType === 'essay'"
        :question="examQuestion"
      />
      <div v-if="loading">Loading question...</div>
      <div v-if="error">Error loading question: {{ error.message }}</div>
    </div>
  </div>
</template>
