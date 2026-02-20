<script setup lang="ts">
import { onMounted } from "vue";
import { useRoute } from "vue-router";

import PageHeader from "~/app/layouts/PageLayouts/PageHeader.vue";

import { useGetExam } from "../composables/useGrading";

const { exam, loading, error, getExam } = useGetExam();
const route = useRoute();
const examId = route.params.examId as string;

onMounted(() => {
  getExam(examId);
});
</script>
<template>
  <div>
    <PageHeader title="Exam Grading" :disable-menu="true" />
    <div v-if="loading">Loading...</div>
    <div v-else-if="error">Error: {{ error }}</div>
    <div v-else-if="exam" id="exam-grading-component">
      <h2>Exam ID: {{ exam.examId }}</h2>
      <p>Grading Completed: {{ exam.gradingCompleted }}</p>
      <p>Number of Questions: {{ exam.questions.length }}</p>
      <ul
        class="list bg-base-100 rounded-box shadow-md border border-neutral"
        id="grading-ungraded-questions-list"
      >
        <li class="p-4 pb-2 text-xs opacity-60 tracking-wide">
          Exam Questions
        </li>
        <li
          class="list-row"
          v-for="question in exam.questions"
          :key="question.questionId"
        >
          <div>
            <a
              :href="`/grading/exam/${examId}/question/${question.questionIndex}`"
            >
              {{ question.questionId }}
            </a>
            - Graded: {{ question.graded }}
          </div>
        </li>
      </ul>
    </div>
  </div>
</template>
