<script lang="ts" setup>
import { Apple } from "lucide-vue-next";
import DashboardCard from "~/app/components/DashboardCard/DashboardCard.vue";
</script>

<template>
  <DashboardCard
    title="Exam Grading"
    description="Grade completed exams"
    href="/grading"
    :icon="Apple"
  />
</template>
