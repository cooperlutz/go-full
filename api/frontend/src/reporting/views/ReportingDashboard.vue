<script setup lang="ts">
import PageHeader from "~/app/layouts/PageLayouts/PageHeader.vue";
import MetricsStats from "../components/MetricsStats.vue";
</script>

<template>
  <div id="reporting-dashboard">
    <PageHeader title="Reporting Dashboard" :disable-menu="true" />

    <div>
      <MetricsStats />
    </div>
  </div>
</template>
