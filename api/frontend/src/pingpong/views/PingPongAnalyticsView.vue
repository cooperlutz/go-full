<script setup lang="ts">
import PageHeader from "~/app/layouts/PageLayouts/PageHeader.vue";
import SectionDivider from "~/app/components/SectionDivider/SectionDivider.vue";

import PingPongTable from "../components/PingPongTable.vue";
import PingPongMetrics from "../components/PingPongMetrics.vue";
</script>

<template>
  <PageHeader
    title="Ping Pong Analytics"
    description="Analytics and insights for the Ping Pong domain"
  />
  <PingPongMetrics />
  <SectionDivider />
  <PingPongTable />
</template>
