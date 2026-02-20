<script setup lang="ts">
import { useRouter } from "vue-router";

import { type Exam } from "../services";

const props = defineProps<{
  exams: Array<Exam>;
}>();

const router = useRouter();

const examTableHeaders: Record<keyof Omit<Exam, "questions">, string> = {
  examId: "Exam ID",
  totalPointsEarned: "Total Points Earned",
  totalPointsPossible: "Total Points Possible",
  gradingCompleted: "Grading Completed",
};
</script>

<template>
  <div
    class="card w-full bg-base-100 shadow-lg card-border border-secondary border-solid"
  >
    <div class="card-body">
      <h2 class="card-title">Exams to Grade</h2>
    </div>
    <table class="table table-xs px-2 py-2" id="grading-ungraded-exams-table">
      <thead>
        <tr>
          <th v-for="header in examTableHeaders" :key="header">
            {{ header }}
          </th>
        </tr>
      </thead>
      <tbody>
        <tr
          v-for="exam in props.exams"
          :key="exam.examId"
          @click="() => router.push(`/grading/exam/${exam.examId}`)"
          class="hover:bg-base-300"
        >
          <td>{{ exam.examId }}</td>
          <td>{{ exam.totalPointsEarned }}</td>
          <td>{{ exam.totalPointsPossible }}</td>
          <td>{{ exam.gradingCompleted }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>
