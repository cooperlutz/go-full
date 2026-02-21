<script setup lang="ts">
import PageHeader from "../layouts/PageLayouts/PageHeader.vue";

import DashboardCard from "~/pingpong/components/DashboardCard.vue";
import ExamLibraryDashboardCard from "~/examlibrary/components/DashboardCard.vue";
import GradingDashboardCard from "~/grading/components/DashboardCard.vue";
import ReportingDashboardCard from "~/reporting/components/DashboardCard.vue";
</script>

<template>
  <PageHeader title="Dashboard" :disable-menu="true" />
  <div
    id="apps-grid"
    class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4 mb-4"
  >
    <DashboardCard />
    <ExamLibraryDashboardCard />
    <GradingDashboardCard />
    <ReportingDashboardCard />
  </div>
</template>
