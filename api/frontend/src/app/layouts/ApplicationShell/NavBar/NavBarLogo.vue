<script lang="ts" setup>
import CONFIG from "../../../config";
</script>

<template>
  <a class="text-2xl font-bold flex text-base-100" href="/">
    <div class="mr-2">
      <h1>{{ CONFIG.SITE_NAME }}</h1>
    </div>
    <img src="/logo.png" height="32" width="32" />
  </a>
</template>
