<script setup lang="ts">
import { onMounted } from "vue";
import { CircleX } from "lucide-vue-next";

import { type ExamMetadata } from "~/examlibrary/services";
import { useGetFindAllExams } from "~/examlibrary/composables/useGetFindAllExams";

const { error, loading, allExams, fetchData } = useGetFindAllExams();

const examTableHeaders: Record<keyof ExamMetadata, string> = {
  id: "Exam ID",
  name: "Name",
  gradeLevel: "Grade Level",
};

onMounted(async () => {
  await fetchData();
});
</script>

<template>
  <div
    class="card w-full bg-base-100 shadow-lg card-border border-secondary border-solid"
  >
    <div v-if="loading" id="exam-table-loading">
      <div class="skeleton h-32 w-full"></div>
    </div>

    <table
      class="table table-xs"
      v-else-if="!loading && !error && allExams != null && allExams.length > 0"
      id="exam-library-table"
    >
      <thead>
        <tr>
          <th v-for="header in examTableHeaders" :key="header">
            {{ header }}
          </th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="entity in allExams" :key="entity.id">
          <td>
            <a :href="`/exam-library/${entity.id}`" class="link">{{
              entity.id
            }}</a>
          </td>
          <td>{{ entity.name }}</td>
          <td>{{ entity.gradeLevel }}</td>
        </tr>
      </tbody>
    </table>

    <div v-else-if="error" id="exam-table-error">
      <div role="alert" class="alert alert-error">
        <CircleX />
        Error loading exams, library may be empty. Error message: {{ error }}
      </div>
    </div>

    <div v-else id="exam-table-unknown-state">How did you get here?</div>
  </div>
</template>
