<script setup lang="ts">
import { onMounted, ref } from "vue";
import { useRoute } from "vue-router";
import {
  useGetExamQuestion,
  useGradeExamQuestion,
} from "../composables/useGrading";
import PageHeader from "~/app/layouts/PageLayouts/PageHeader.vue";

const pointsToGive = ref(0);
const graderComments = ref("");

const { question, loading, error, getExamQuestion } = useGetExamQuestion();
const { gradeExamQuestion } = useGradeExamQuestion();
const route = useRoute();
const examId = route.params.examId as string;
const questionIndex = Number(route.params.questionIndex);

onMounted(() => {
  getExamQuestion(examId, questionIndex);
});
</script>

<template>
  <PageHeader title="Grade Question" :disable-menu="true" />
  <div id="question-grading">
    <div v-if="loading">Loading...</div>
    <div v-else-if="error">Error: {{ error }}</div>
    <div v-else-if="question">
      <div class="card border border-neutral">
        <div class="card-body shadow-sm">
          <h2>Question {{ question.questionIndex + 1 }}</h2>
          <p>Exam ID: {{ question.examId }}</p>
          <p>Index: {{ question.questionIndex }}</p>
          <p>Type: {{ question.questionType }}</p>
          <p>Points Possible: {{ question.pointsPossible }}</p>
        </div>
      </div>
      <div class="grid grid-cols-2">
        <div class="col-span-1">
          <h2>Student Answer:</h2>
          <p>{{ question.providedAnswer }}</p>
        </div>
        <div class="col-span-1">
          <textarea
            class="textarea"
            placeholder="Type feedback here"
            v-model="graderComments"
          ></textarea>
          <input
            type="number"
            placeholder="Type points here"
            class="input"
            v-model="pointsToGive"
          />
          <button
            class="btn btn-primary"
            @click="
              gradeExamQuestion(
                examId,
                question.questionIndex,
                pointsToGive,
                graderComments,
              )
            "
          >
            Send
          </button>
        </div>
      </div>
    </div>
  </div>
</template>
