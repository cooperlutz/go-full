<script setup lang="ts">
import { h } from "vue";
import { Icon } from "lucide-vue-next";
import { batBall } from "@lucide/lab";
const PingPongIcon = () => h(Icon, { name: "ping-pong", iconNode: batBall });
</script>

<template>
  <a href="/ping-pong">
    <div
      class="card w-full bg-base-100 card-xl shadow-sm border border-neutral min-h-48"
    >
      <div class="card-body">
        <div class="flex gap-4">
          <PingPongIcon />
          <h2 class="card-title">Ping Pong</h2>
        </div>
        <hr />
        <p>Application of the Ping Pong Module</p>
      </div>
    </div>
  </a>
</template>
