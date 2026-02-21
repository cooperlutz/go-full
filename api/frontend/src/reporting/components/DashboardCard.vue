<script setup lang="ts">
import DashboardCard from "~/app/components/DashboardCard/DashboardCard.vue";
import { ChartBarStacked } from "lucide-vue-next";
</script>
<template>
  <DashboardCard
    title="Reporting"
    description="View reports and analytics on exam performance"
    href="/reporting"
    :icon="ChartBarStacked"
  />
</template>
