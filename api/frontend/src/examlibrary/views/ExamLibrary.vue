<script setup lang="ts">
import PageHeader from "~/app/layouts/PageLayouts/PageHeader.vue";

import ExamLibraryTable from "../components/ExamLibraryTable.vue";
import ExamCreatorButton from "../components/ExamCreatorButton.vue";
</script>

<template>
  <PageHeader title="Exam Library" :disable-menu="true" />
  <ExamCreatorButton />
  <ExamLibraryTable />
</template>
