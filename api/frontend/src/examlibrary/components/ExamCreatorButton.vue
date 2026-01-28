<script setup lang="ts">
import { useRouter } from "vue-router";

const router = useRouter();
function navigateToExamCreator() {
  router.push({ name: "ExamCreator" });
}
</script>

<template>
  <button class="btn btn-primary" @click="navigateToExamCreator">
    New Exam
  </button>
</template>
