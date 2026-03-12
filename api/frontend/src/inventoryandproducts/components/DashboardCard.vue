<script lang="ts" setup>
import { Apple } from "lucide-vue-next";
import DashboardCard from "~/app/components/DashboardCard/DashboardCard.vue";
</script>

<template>
  <!-- TODO: update description and icon -->
  <DashboardCard
    title="InventoryAndProducts"
    description="Placeholder Value"
    href="/inventory-and-products"
    :icon="Apple"
  />
</template>
