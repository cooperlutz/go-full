<script setup lang="ts">
import PageHeader from "~/app/layouts/PageLayouts/PageHeader.vue";
import ExamLibraryTable from "../components/ExamLibraryTable.vue";
</script>

<template>
  <PageHeader title="Exam Library" :disable-menu="true" />
  <ExamLibraryTable />
</template>
