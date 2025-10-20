<script setup lang="ts">
import { useShowCreatePingPongResponse } from "~/pingpong/composables/usePingPongToast";
import { useSendPingPong } from "~/pingpong/composables/usePingPong";

const { sendPingPong } = useSendPingPong();

// clickPong sends a "pong" message to the API and shows a notification with the response
const clickPong = async () => {
  const response = await sendPingPong("pong");
  useShowCreatePingPongResponse(response?.message);
};
</script>

<template>
  <div
    id="pong-button"
    class="btn btn-info btn-xl text-xs"
    @click="clickPong()"
  >
    pong
  </div>
</template>
