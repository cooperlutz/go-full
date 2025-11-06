<script setup lang="ts">
import { onMounted } from "vue";

import { type PingPongRaw } from "~/pingpong/services";
import { useFindAllPingPongs } from "~/pingpong/composables/usePingPong";

const { error, loading, allPingPongs, fetchData } = useFindAllPingPongs();

// Define table headers explicitly
const pingPongTableHeaders: Record<keyof PingPongRaw, string> = {
  id: "Ping Pong ID",
  message: "Message",
  createdAt: "Created At",
  updatedAt: "Updated At",
  deletedAt: "Deleted At",
  deleted: "Deleted",
};

onMounted(async () => {
  await fetchData();
});
</script>

<template>
  <div
    class="card w-full bg-base-100 shadow-lg card-border border-secondary border-solid"
  >
    <div class="card-body">
      <h2 class="card-title">Ping Pongs</h2>
      <p>All Ping Pongs</p>
    </div>
    <table class="table table-xs" v-if="!loading && !error">
      <thead>
        <tr>
          <th v-for="header in pingPongTableHeaders" :key="header">
            {{ header }}
          </th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="entity in allPingPongs?.pingpongs" :key="entity.id">
          <td>{{ entity.id }}</td>
          <td>{{ entity.message }}</td>
          <td>{{ entity.createdAt }}</td>
          <td>{{ entity.updatedAt }}</td>
          <td>{{ entity.deletedAt }}</td>
          <td>{{ entity.deleted }}</td>
        </tr>
      </tbody>
    </table>
    <div v-else-if="loading">Loading ping pongs...</div>
    <div v-else-if="error" id="pingpong-table-error">
      Error loading ping pongs: {{ error }}
    </div>
  </div>
</template>
