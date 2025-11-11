<script setup lang="ts">
import PageHeader from "~/app/layouts/PageLayouts/PageHeader.vue";
import DashboardCard from "~/pingpong/components/DashboardCard.vue";
</script>

<template>
  <PageHeader title="Dashboard" :disable-menu="true" />
  <div
    id="apps-grid"
    class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4 mb-4"
  >
    <DashboardCard />
  </div>
</template>
