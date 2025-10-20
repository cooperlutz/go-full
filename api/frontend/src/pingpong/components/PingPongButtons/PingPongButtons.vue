<script setup lang="ts">
import PingButton from "./PingButton.vue";
import PongButton from "./PongButton.vue";
</script>

<template>
  <div
    class="card w-full bg-base-300 shadow-md card-border border-secondary border-solid"
  >
    <div class="card-body">
      <div class="card-title">Ping Pong Buttons</div>
      <div class="card-actions">
        <PingButton />
        <PongButton />
      </div>
    </div>
  </div>
</template>
