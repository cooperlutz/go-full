<script lang="ts" setup>
import { defineProps } from "vue";

import PageHeaderBreadcrumbs from "~/app/components/PageHeader/PageHeaderBreadcrumbs.vue";
import PageHeaderMenu from "~/app/components/PageHeader/PageHeaderMenu.vue";
import SectionDivider from "~/app/components/SectionDivider/SectionDivider.vue";

const props = defineProps({
  title: String,
  subtitle: String,
  description: String,
  disableMenu: {
    type: Boolean,
    default: false,
  },
});
</script>

<template>
  <div class="flex flex-row justify-between w-full mb-4">
    <PageHeaderBreadcrumbs />
    <PageHeaderMenu v-if="!props.disableMenu" />
  </div>
  <!-- page header -->
  <div class="columns-2 w-full">
    <!-- page title -->
    <h1 class="text-4xl font-bold tracking-tight sm:text-3xl py-2">
      {{ props.title }}
    </h1>
  </div>
  <!-- spacer -->
  <div class="mt-2" />
  <!-- intro text -->
  <div class="flex flex-col">
    <p class="text-lg">
      {{ props.subtitle }}
    </p>
  </div>

  <div class="mt-2" />
  <div class="flex flex-col">
    <p class="text-md text-base-content/70">
      {{ props.description }}
    </p>
  </div>
  <!-- spacer with divider -->
  <SectionDivider />
</template>
