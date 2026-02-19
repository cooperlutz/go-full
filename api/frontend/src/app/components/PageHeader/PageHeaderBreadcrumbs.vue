<script lang="ts" setup>
import { useRoute } from "vue-router";

type Breadcrumb = {
  name: string;
  link: string;
};

function parsePathToBreadcrumbs(path: string): Breadcrumb[] {
  const segments = path.split("/").filter((segment) => segment.length > 0);
  const breadcrumbs: Breadcrumb[] = [];

  segments.forEach((segment) => {
    // capitalize first letter of each word
    const name = segment.replace(/\b\w/g, (char) => char.toUpperCase());
    // create link by joining segments up to the current segment
    const link = "/" + segments.slice(0, segments.indexOf(segment) + 1).join("/");
    breadcrumbs.push({ name, link });
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
        <a :href="breadcrumb.link">{{ breadcrumb.name }}</a>
      </li>
    </ul>
  </div>
</template>
