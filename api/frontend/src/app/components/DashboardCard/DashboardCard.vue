<script lang="ts" setup>
import { type LucideIcon } from "lucide-vue-next";

const props = defineProps<{
  title: string;
  description: string;
  href: string;
  icon: LucideIcon;
}>();
</script>

<template>
  <a :href="props.href">
    <div
      class="card w-full bg-base-100 card-xl shadow-sm border border-neutral min-h-48"
    >
      <div class="card-body">
        <div class="flex gap-4">
          <component :is="props.icon" />
          <h2 class="card-title">{{ props.title }}</h2>
        </div>
        <hr />
        <p>{{ props.description }}</p>
      </div>
    </div>
  </a>
</template>
