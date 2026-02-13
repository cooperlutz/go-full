<script lang="ts" setup>
import { Apple } from "lucide-vue-next";
</script>

<template>
  <a href="/grading">
    <div
      class="card w-full bg-base-100 card-xl shadow-sm border border-neutral min-h-48"
    >
      <div class="card-body">
        <div class="flex gap-4">
          <Apple />
          <h2 class="card-title">Exam Grading</h2>
        </div>
        <hr />
        <p>Grade completed exams</p>
      </div>
    </div>
  </a>
</template>
