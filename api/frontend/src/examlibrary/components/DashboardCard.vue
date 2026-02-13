<script lang="ts" setup>
import { LibraryBig } from "lucide-vue-next";
</script>

<template>
  <a href="/exam-library">
    <div
      class="card w-full bg-base-100 card-xl shadow-sm border border-neutral min-h-48"
    >
      <div class="card-body">
        <div class="flex gap-4">
          <LibraryBig />
          <h2 class="card-title">Exam Library</h2>
        </div>
        <hr />
        <p>Explore the exam library</p>
      </div>
    </div>
  </a>
</template>
