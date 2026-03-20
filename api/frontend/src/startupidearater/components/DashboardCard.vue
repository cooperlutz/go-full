<script lang="ts" setup>
import { Apple } from "lucide-vue-next";
import DashboardCard from "~/app/components/DashboardCard/DashboardCard.vue";
</script>

<template>
  <!-- TODO: update description and icon -->
  <DashboardCard
    title="StartupIdeaRater"
    description="Placeholder Value"
    href="/startup-idea-rater"
    :icon="Apple"
  />
</template>
