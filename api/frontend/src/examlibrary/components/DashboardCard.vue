<script lang="ts" setup>
import DashboardCard from "~/app/components/DashboardCard/DashboardCard.vue";
import { LibraryBig } from "lucide-vue-next";
</script>

<template>
  <DashboardCard
    title="Exam Library"
    description="Explore the exam library"
    href="/exam-library"
    :icon="LibraryBig"
  />
</template>
