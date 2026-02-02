<script setup lang="ts">
import { onMounted } from "vue";
import { useRoute } from "vue-router";

import QuestionComponent from "../components/QuestionComponent.vue";
import ExaminationProgress from "../components/ExaminationProgress.vue";
import QuestionNav from "../components/QuestionNav.vue";
import SubmitExamButton from "../components/SubmitExamButton.vue";
import { useGetExam } from "../composables/useExamination";

const route = useRoute();
const questionIndex = Number(route.params.index);
const examId = route.params.id as string;

const { exam, loading, error, getExam } = useGetExam();

onMounted(async () => {
  await getExam(examId);
});
</script>

<template>
  <div v-if="loading">Loading exam details...</div>
  <div v-if="error">Error loading exam details: {{ error.message }}</div>
  <div v-if="!exam">No exam data available.</div>
  <card
    v-if="exam"
    class="w-96 bg-base-100 shadow-xl"
    id="examination-view-component"
  >
    <div class="card-body">
      <div class="grid grid-cols-5 gap-4 mb-4">
        <h1 class="card-title col-span-3">Examination Component</h1>
        <ExaminationProgress class="col-span-2 items-end" />
        <div class="col-span-4">
          <QuestionComponent
            :examId="exam.examId"
            :questionIndex="questionIndex"
          />
        </div>
        <div class="col-span-1">
          <QuestionNav
            v-if="exam.questions"
            :examId="exam.examId"
            :questions="exam.questions"
          />
          <div class="flex justify-center mt-4">
            <SubmitExamButton :examId="exam.examId" />
          </div>
        </div>
      </div>
    </div>
  </card>
</template>
