<script setup lang="ts">
import { useShowCreatePingPongResponse } from "~/pingpong/composables/usePingPongToast";
import { useSendPingPong } from "~/pingpong/composables/usePingPong";

const { sendPingPong } = useSendPingPong();

// clickPing sends a "ping" message to the API and shows a notification with the response
const clickPing = async () => {
  const response = await sendPingPong("ping");
  useShowCreatePingPongResponse(response?.message);
};
</script>

<template>
  <div
    id="ping-button"
    class="btn btn-error btn-xl text-xs"
    @click="clickPing()"
  >
    ping
  </div>
</template>
