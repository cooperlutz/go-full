<script lang="ts" setup>
import { Settings } from "lucide-vue-next";

import CONFIG from "~/app/config";

import LogoName from "./NavBarLogo.vue";
</script>

<template>
  <div id="navbar" class="fixed top-0 left-0 w-full z-50 bg-info">
    <div class="mx-auto max-w-7xl px-6 sm:px-6 lg:px-8">
      <div class="relative flex h-16 justify-between">
        <div class="flex flex-1 items-stretch justify-start">
          <a class="flex flex-shrink-0 items-center" href="#">
            <LogoName class="mr-8" />
          </a>
        </div>
        <div
          class="flex-shrink-0 flex px-2 py-3 items-center space-x-8 text-base-100"
        >
          <a
            :href="CONFIG.DOCS_URL"
            target="_blank"
            class="hover:text-accent-content text-sm font-medium"
          >
            Docs
          </a>
          <a
            href="/settings"
            class="hover:text-accent-content text-sm font-medium"
          >
            <Settings :size="20" />
          </a>
        </div>
      </div>
    </div>
  </div>
</template>
