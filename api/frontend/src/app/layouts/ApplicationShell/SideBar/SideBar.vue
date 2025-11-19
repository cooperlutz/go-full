<script lang="ts" setup>
import { ref, h } from "vue";
import { Icon, LayoutDashboard, type LucideIcon } from "lucide-vue-next";
import { batBall } from "@lucide/lab";

// Create a functional component for the custom icon
const PingPongIcon = () => h(Icon, { name: "ping-pong", iconNode: batBall });

type SidebarItem = {
  name: string;
  url?: string;
  icon?: LucideIcon;
  children?: SidebarItem[];
};

const sidebarItems = ref<SidebarItem[]>([
  { name: "Dashboard", url: "/dashboard", icon: LayoutDashboard },
  { name: "Ping Pong", url: "/ping-pong", icon: PingPongIcon },
]);
</script>

<template>
  <!-- Sidebar -->
  <aside
    id="sidebar"
    class="fixed top-0 h-screen overflow-y-auto bg-base-200 w-64 border-r shadow-sm z-49"
  >
    <!-- Sidebar Header -->
    <div class="py-8" />

    <!-- Sidebar Content -->
    <div class="py-4">
      <!-- Navigation Section -->
      <div class="px-4 mb-4">
        <nav class="mt-2 space-y-1">
          <div
            v-for="item in sidebarItems"
            :key="item.name"
            class="sidebar-item mb-2"
          >
            <!-- if the sidebar item has no children -->
            <a
              v-if="!item.children"
              :href="item.url"
              class="flex items-center px-4 py-2 text-sm font-medium hover:link hover:text-info rounded-lg"
            >
              <component
                :is="item.icon"
                :size="16"
                class="h-5 w-5 mr-3"
              ></component>
              <span class="ml-1">
                {{ item.name }}
              </span>
            </a>
          </div>
        </nav>
      </div>
    </div>
  </aside>

  <slot name="content"></slot>
</template>
