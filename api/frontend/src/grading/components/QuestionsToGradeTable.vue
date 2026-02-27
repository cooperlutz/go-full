<script setup lang="ts">
import { useRouter } from "vue-router";

import { type Question } from "../services";

const props = defineProps<{
  questions: Array<Question>;
}>();

const router = useRouter();

const questionTableHeaders: Record<
  keyof Omit<
    Question,
    "feedback" | "examId" | "providedAnswer" | "questionIndex"
  >,
  string
> = {
  questionId: "Question ID",
  questionType: "Question Type",
  graded: "Graded",
  pointsEarned: "Points Earned",
  pointsPossible: "Points Possible",
};
</script>

<template>
  <div
    class="card w-full"
  >
    <div class="card-body">
      <h2 class="card-title">Questions to Grade</h2>
    </div>
    <table class="table table-xs px-2 py-2" id="grading-grade-questions-table">
      <thead>
        <tr>
          <th v-for="header in questionTableHeaders" :key="header">
            {{ header }}
          </th>
        </tr>
      </thead>
      <tbody>
        <tr
          v-for="(question, index) in props.questions"
          :key="question.questionId"
          class="hover:bg-base-300"
        >
          <td>{{ question.questionId }}</td>
          <td>{{ question.questionType }}</td>
          <td>{{ question.graded }}</td>
          <td>{{ question.pointsEarned }}</td>
          <td>{{ question.pointsPossible }}</td>
          <td>
            <button
              class="btn btn-sm btn-primary"
              @click="
                () =>
                  router.push(
                    `/grading/exam/${question.examId}/question/${question.questionIndex}`,
                  )
              "
              :id="`grading-grade-question-button-${index}`"
            >
              Grade
            </button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>
