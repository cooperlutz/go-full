<script lang="ts" setup>
import PageHeaderBreadcrumbs from "../../components/PageHeader/PageHeaderBreadcrumbs.vue";
import PageHeaderMenu from "../../components/PageHeader/PageHeaderMenu.vue";
import SectionDivider from "../../components/SectionDivider/SectionDivider.vue";

defineProps({
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
    <PageHeaderMenu v-if="!disableMenu" />
  </div>
  <!-- page header -->
  <div class="columns-2 w-full">
    <!-- page title -->
    <h1 class="text-4xl font-bold tracking-tight sm:text-3xl py-2">
      {{ title }}
    </h1>
  </div>
  <!-- spacer -->
  <div class="mt-2" />
  <!-- intro text -->
  <div class="flex flex-col">
    <p class="text-lg">
      {{ subtitle }}
    </p>
  </div>

  <div class="mt-2" />
  <div class="flex flex-col">
    <p class="text-md text-base-content/70">
      {{ description }}
    </p>
  </div>
  <!-- spacer with divider -->
  <SectionDivider />
</template>
