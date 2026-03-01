<script setup lang="ts">
import { onMounted, ref } from "vue";
import { useRoute, useRouter } from "vue-router";

import PageHeader from "~/app/layouts/PageLayouts/PageHeader.vue";

import {
  useGetExamQuestion,
  useGradeExamQuestion,
} from "../composables/useGrading";

const emit = defineEmits(["question-graded"]);

const pointsToGive = ref(0);
const graderComments = ref("");

const { question, loading, error, getExamQuestion } = useGetExamQuestion();
const { gradeExamQuestion } = useGradeExamQuestion();
const route = useRoute();
const router = useRouter();
const examId = route.params.examId as string;
const questionIndex = Number(route.params.questionIndex);

function gradeQuestion() {
  gradeExamQuestion(
    examId,
    questionIndex,
    pointsToGive.value,
    graderComments.value,
  );
  emit("question-graded");
  // navigate back to exam grading page after grading question
  router.push(`/grading/exam/${examId}`);
}

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
          <form class="fieldset">
            <fieldset class="fieldset">
              <textarea
                class="textarea"
                id="grader-comments"
                placeholder="Type feedback here"
                v-model="graderComments"
              ></textarea>
            </fieldset>
            <fieldset class="fieldset">
              <input
                type="number"
                id="points-to-give"
                placeholder="Type points here"
                class="input validator"
                min="0"
                required
                :max="question.pointsPossible"
                v-model="pointsToGive"
              />
            </fieldset>
            <button
              class="btn btn-primary"
              @click="gradeQuestion"
              id="save-feedback-and-points"
            >
              Save
            </button>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>
