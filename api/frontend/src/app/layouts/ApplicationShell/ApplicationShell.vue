<script setup lang="ts">
import { ref, h } from "vue";
import {
  Icon,
  LayoutDashboard,
  LibraryBig,
  type LucideIcon,
  PanelLeftClose,
  Settings,
  File,
  User,
} from "lucide-vue-next";
import { batBall } from "@lucide/lab";

import Footer from "~/app/layouts/ApplicationShell/Footer/FooterPrimary.vue";
import LogoName from "./NavBar/NavBarLogo.vue";
import CONFIG from "~/app/config";

// Create a functional component for the custom icon
const PingPongIcon = () => h(Icon, { name: "ping-pong", iconNode: batBall });

type SidebarItem = {
  name: string;
  url?: string;
  icon?: LucideIcon;
};

const topSidebarItems = ref<SidebarItem[]>([
  { name: "Dashboard", url: "/dashboard", icon: LayoutDashboard },
  { name: "Ping Pong", url: "/ping-pong", icon: PingPongIcon },
  { name: "Exam Library", url: "/exam-library", icon: LibraryBig },
]);
const bottomSidebarItems = ref<SidebarItem[]>([
  { name: "Docs", url: CONFIG.DOCS_URL, icon: File },
  { name: "Settings", url: "/settings", icon: Settings },
  { name: "Profile", url: "/profile", icon: User },
]);
</script>

<template>
  <div id="app-shell" class="antialiased">
    <div class="drawer lg:drawer-open">
      <input id="my-drawer-4" type="checkbox" class="drawer-toggle" />
      <div id="main-content" class="drawer-content">
        <!-- Navbar -->
        <nav class="navbar w-full bg-info fixed top-0 z-50">
          <label
            for="my-drawer-4"
            aria-label="open sidebar"
            class="btn btn-square btn-ghost"
          >
            <!-- Sidebar toggle icon -->
            <PanelLeftClose />
          </label>
          <div class="px-4">
            <LogoName />
          </div>
        </nav>
        <!-- Page content here -->
        <main>
          <div class="w-full max-w-7xl p-10 mt-10">
            <!-- Router View for Main Content -->
            <router-view v-slot="{ Component }">
              <component :is="Component" />
            </router-view>
          </div>
        </main>
        <Footer />
      </div>

      <div class="drawer-side is-drawer-close:overflow-visible">
        <label
          for="my-drawer-4"
          aria-label="close sidebar"
          class="drawer-overlay"
        ></label>
        <div
          class="flex min-h-full flex-col items-start bg-base-200 is-drawer-close:w-16 is-drawer-open:w-64"
        >
          <!-- Sidebar content here -->
          <ul id="sidebar-top" class="menu w-full grow">
            <li v-for="item in topSidebarItems" :key="item.name">
              <a :href="item.url" class="hover:link">
                <button>
                  <component
                    :is="item.icon"
                    :size="16"
                    class="h-6 w-6"
                  ></component>
                </button>
                <span class="is-drawer-close:hidden">
                  {{ item.name }}
                </span>
              </a>
            </li>
          </ul>
          <!-- Bottom Sidebar Items -->
          <ul id="sidebar-bottom" class="menu w-full">
            <li v-for="item in bottomSidebarItems" :key="item.name">
              <a :href="item.url" class="hover:link">
                <button>
                  <component
                    :is="item.icon"
                    :size="16"
                    class="h-6 w-6"
                  ></component>
                </button>
                <span class="is-drawer-close:hidden">
                  {{ item.name }}
                </span>
              </a>
            </li>
          </ul>
        </div>
      </div>
    </div>
  </div>
</template>
