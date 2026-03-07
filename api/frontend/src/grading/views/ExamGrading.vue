<script setup lang="ts">
import { onMounted } from "vue";
import { useRoute } from "vue-router";

import PageHeader from "~/app/layouts/PageLayouts/PageHeader.vue";

import QuestionsToGradeTable from "../components/QuestionsToGradeTable.vue";
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
      <div class="card border border-neutral">
        <div class="card-body shadow-sm">
          <p><b>Exam ID:</b> {{ exam.examId }}</p>
          <p><b>Grading State:</b> {{ exam.state }}</p>
          <p><b>Number of Questions:</b> {{ exam.questions.length }}</p>
          <p><b>Total Points Earned:</b> {{ exam.totalPointsEarned }}</p>
          <p><b>Total Points Possible:</b> {{ exam.totalPointsPossible }}</p>
        </div>
        <div class="p-6">
          <QuestionsToGradeTable :questions="exam.questions" />
        </div>
      </div>
    </div>
  </div>
</template>
