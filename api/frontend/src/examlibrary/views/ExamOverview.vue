<script setup lang="ts">
/*
This view displays an overview of a specific exam, fetching its details based on the exam ID from the route parameters.
It uses the useFindExamByID composable to retrieve the exam data and handles loading and error states.
*/
import { useRoute } from "vue-router";
import { onMounted, ref } from "vue";
import PageHeader from "~/app/layouts/PageLayouts/PageHeader.vue";

import StartExamModal from "~/examination/components/StartExamModal.vue";
import { useFindExamByID } from "~/examlibrary/composables/useGetFindOne";
import type { Exam } from "../services";

const exam = ref<Exam>();
const route = useRoute();
const examId = route.params.id as string;
const { error, loading, findExam } = useFindExamByID();

onMounted(async () => {
  exam.value = await findExam(examId);
});
</script>

<template>
  <PageHeader title="Exam Overview" :disable-menu="true" />
  <div v-if="loading">Loading exam...</div>
  <div v-else-if="error" id="exam-overview-error">
    Error loading exam: {{ error }}
  </div>
  <div v-else>
    <div
      class="card w-full bg-base-100 shadow-lg card-border border-secondary border-solid"
    >
      <div class="card-body">
        <h2 class="card-title">Exam Details</h2>
        <p><b>ID:</b> {{ exam?.id }}</p>
        <p><b>Name:</b> {{ exam?.name }}</p>
        <p><b>Grade Level:</b> {{ exam?.gradeLevel }}</p>
        <div class="mt-6">
          <StartExamModal :libraryExamId="exam?.id ?? ''"/>
        </div>
      </div>
    </div>
  </div>
</template>
