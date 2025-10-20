<script setup lang="ts">
import { ref } from "vue";

import { useShowCreatePingPongResponse } from "~/pingpong/composables/usePingPongToast";
import { useSendPingPong } from "~/pingpong/composables/usePingPong";

const inputValue = ref("");
const { sendPingPong } = useSendPingPong();

// clickSend sends a custom message to the API and shows a notification with the response
const clickSend = async (msg: string) => {
  const response = await sendPingPong(msg);
  useShowCreatePingPongResponse(response?.message);
};
</script>

<template>
  <div
    class="card w-full bg-base-300 shadow-md card-border border-secondary border-solid"
  >
    <div class="card-title">Ping Pong Input</div>
    <div class="card-body flex flex-row">
      <label class="input">
        <input
          id="pingpong-input"
          type="text"
          class="grow"
          placeholder="ping or pong"
          v-model="inputValue"
        />
      </label>
      <div
        id="send-button"
        class="btn btn-m text-xs"
        @click="clickSend(inputValue)"
      >
        Send
      </div>
    </div>
  </div>
</template>
