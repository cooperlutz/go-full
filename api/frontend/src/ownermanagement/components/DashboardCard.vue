<script lang="ts" setup>
import { Apple } from "lucide-vue-next";
import DashboardCard from "~/app/components/DashboardCard/DashboardCard.vue";
</script>

<template>
  <!-- TODO: update description and icon -->
  <DashboardCard
    title="OwnerManagement"
    description="Placeholder Value"
    href="/owner-management"
    :icon="Apple"
  />
</template>
