<script lang="ts" setup>
import { AppWindow, ChartColumnStacked } from "lucide-vue-next";
import { useRoute } from "vue-router";

const route = useRoute();
const currentPathRoot = route.path.split("/")[1];
</script>

<template>
  <ul class="menu menu-sm lg:menu-horizontal rounded-box">
    <li>
      <a class="border" :href="`/${currentPathRoot}/app`">
        <AppWindow :size="16" class="h-5 w-5 mr-3" />
        Application
      </a>
    </li>
    <li>
      <a class="border" :href="`/${currentPathRoot}/analytics`">
        <ChartColumnStacked :size="16" class="h-5 w-5 mr-3" />
        Analytics
      </a>
    </li>
  </ul>
</template>
