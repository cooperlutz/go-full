<script setup lang="ts">
import PageHeader from "~/app/layouts/PageLayouts/PageHeader.vue";
import PingPongButtons from "../components/PingPongButtons/PingPongButtons.vue";
import PingPongInput from "../components/PingPongInput.vue";
</script>

<template>
  <!-- Header -->
  <PageHeader
    title="Ping Pong Application"
    description="The application of the Ping Pong domain. Will you ping or pong?"
  />

  <!-- Layout -->
  <div class="flex flex-col w-full gap-4">
    <p>Welcome to the Ping Pong application!</p>

    <PingPongButtons />

    <PingPongInput />
  </div>
</template>
