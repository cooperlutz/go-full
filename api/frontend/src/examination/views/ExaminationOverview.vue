<script setup lang="ts">
import { useRoute } from "vue-router";

import PageHeader from "~/app/layouts/PageLayouts/PageHeader.vue";

const route = useRoute();
const examId = route.params.id as string;
</script>

<template>
  <PageHeader title="Exam Overview" :disable-menu="true" />
  <h2 class="card-title">Instructions</h2>
  <p>take the exam. do a good job</p>
  <button
    id="go-to-first-question-button"
    class="btn btn-primary mt-4"
    @click="$router.push(`/exam/${examId}/question/1`)"
  >
    Go to Question 1
  </button>
</template>
