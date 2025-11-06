<script setup lang="ts">
/* STEP 5.4. Implement Frontend Component
here we implement the actual logic of the component
*/
import { ref } from "vue";

import { useFindPingPongByID } from "~/pingpong/composables/usePingPong";
import type { PingPongRaw } from "../services";

const lookupInput = ref("");
const lookupOutput = ref<PingPongRaw>();
const { lookup, loading, error } = useFindPingPongByID();

const clickLookup = async (id: string) => {
  const response = await lookup(id);
  lookupOutput.value = response;
};
</script>

<template>
  <div
    class="card w-full bg-base-300 shadow-md card-border border-secondary border-solid"
  >
    <div class="card-body">
      <div class="card-title">Ping Pong Lookup</div>
      <label class="input">
        <input
          id="pingpong-lookup"
          type="text"
          class="grow"
          placeholder="pingpong id"
          v-model="lookupInput"
        />
      </label>
      <div class="card-actions">
        <div
          id="lookup-button"
          class="btn btn-m text-xs"
          @click="clickLookup(lookupInput)"
        >
          Lookup
        </div>
      </div>
      <div v-if="!loading && !error && lookupOutput">
        <div class="card w-full card-border border-secondary border-solid">
          <div id="pingpong-lookup-card" class="card-body">
            <div
              class="overflow-x-auto rounded-box border border-base-content/5 bg-base-100"
            >
              <p><b>ID:</b> {{ lookupOutput?.id }}</p>
              <p><b>Message:</b> {{ lookupOutput?.message }}</p>
              <p><b>CreatedAt:</b> {{ lookupOutput?.createdAt }}</p>
              <p><b>UpdatedAt</b>: {{ lookupOutput?.updatedAt }}</p>
              <p><b>Deleted:</b> {{ lookupOutput?.deleted }}</p>
              <p><b>DeletedAt:</b> {{ lookupOutput?.deletedAt }}</p>
            </div>
          </div>
        </div>
      </div>
      <div v-else-if="loading">
        <span class="loading loading-bars loading-md"></span>
      </div>
      <div v-else-if="error">
        <div
          id="pingpong-lookup-error"
          class="card w-full card-border border-secondary border-solid"
        >
          <div class="card-body">
            <div
              class="overflow-x-auto rounded-box border border-base-content/5 bg-base-100"
            >
              <div role="alert" class="alert alert-error">
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  class="h-6 w-6 shrink-0 stroke-current"
                  fill="none"
                  viewBox="0 0 24 24"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z"
                  />
                </svg>
                <span>Error loading ping pong</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
