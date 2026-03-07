<script setup lang="ts">
import { onMounted } from "vue";

import PageHeader from "~/app/layouts/PageLayouts/PageHeader.vue";

import { useGetFindIncompleteExams } from "../composables/useGrading";
import ExamsToGradeTable from "../components/ExamsToGradeTable.vue";

const { incompleteExams, loading, error, getFindIncompleteExams } =
  useGetFindIncompleteExams();

onMounted(async () => {
  await getFindIncompleteExams();
});
</script>
<template>
  <div id="exams-to-grade">
    <PageHeader title="Grading" :disable-menu="true" />

    <div v-if="loading">
      <div class="skeleton h-32 w-full"></div>
    </div>

    <div
      v-else-if="
        !loading &&
        !error &&
        incompleteExams != null &&
        incompleteExams.length > 0
      "
    >
      <ExamsToGradeTable :exams="incompleteExams" />
    </div>
  </div>
</template>
