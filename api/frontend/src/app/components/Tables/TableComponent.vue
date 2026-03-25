<script setup lang="ts">
// accepts a generic array of objects with unknown key value pairs
const props = defineProps<{
  title: string;
  items: Array<Record<string, unknown>>; // array of objects with unknown key value pairs
}>();
</script>

<template>
  <div
    class="card w-full bg-base-100 shadow-lg card-border border-secondary border-solid"
  >
    <div class="card-body">
      <h2 class="card-title">{{ props.title }}</h2>
    </div>
    <table class="table table-xs px-2 py-2">
      <thead>
        <tr>
          <th v-for="header in Object.keys(props.items[0] || {})" :key="header">
            {{ header }}
          </th>
        </tr>
      </thead>
      <tbody>
        <tr
          v-for="(item, index) in props.items"
          :key="index"
          class="hover:bg-base-300"
        >
          <td v-for="(value, key) in item" :key="key">{{ value }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>
