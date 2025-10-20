<script lang="ts" setup>
import { useRoute } from "vue-router";

function parsePathToBreadcrumbs(path: string): string[] {
  const segments = path.split("/").filter((segment) => segment.length > 0);
  const breadcrumbs: string[] = [];

  segments.forEach((segment) => {
    // remove dashes and capitalize words
    segment = segment.replace(/-/g, " ");
    // capitalize first letter of each word
    segment = segment.replace(/\b\w/g, (char) => char.toUpperCase());
    breadcrumbs.push(segment);
  });

  return breadcrumbs;
}

const currentRoute = useRoute();
</script>

<template>
  <div class="breadcrumbs text-sm">
    <ul>
      <li
        v-for="(breadcrumb, index) in parsePathToBreadcrumbs(currentRoute.path)"
        :key="index"
      >
        <a>{{ breadcrumb }}</a>
      </li>
    </ul>
  </div>
</template>
