<script setup lang="ts">
import { onMounted } from "vue";
import { useGetUngradedExams } from "../composables/useGrading";
import PageHeader from "~/app/layouts/PageLayouts/PageHeader.vue";

const { ungradedExams, loading, error, getUngradedExams } =
  useGetUngradedExams();

onMounted(async () => {
  await getUngradedExams();
});
</script>
<template>
  <div id="exams-to-grade">
    <PageHeader title="Exams to Grade" :disable-menu="true" />

    <div v-if="loading">
      <div class="skeleton h-32 w-full"></div>
    </div>

    <div
      v-else-if="
        !loading && !error && ungradedExams != null && ungradedExams.length > 0
      "
    >
      <div
        v-for="exam in ungradedExams"
        :key="exam.examId"
        class="card shadow-sm"
        id="grading-ungraded-exams-list"
      >
        <a :href="`/grading/exam/${exam.examId}`"
          >Grade Exam: {{ exam.examId }}</a
        >
      </div>
    </div>
    <div v-else-if="error">Error: {{ error }}</div>
  </div>
</template>
