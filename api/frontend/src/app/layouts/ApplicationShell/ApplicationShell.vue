<script setup lang="ts">
import NavBar from "~/app/layouts/ApplicationShell/NavBar/NavBar.vue";
import Footer from "~/app/layouts/ApplicationShell/Footer/FooterPrimary.vue";

import SideBar from "~/app/layouts/ApplicationShell/SideBar/SideBar.vue";
</script>

<template>
  <div id="app-shell" class="antialiased">
    <NavBar />
    <SideBar>
      <template v-slot:content>
        <main>
          <div class="flex pt-8 overflow-hidden">
            <div
              id="main-content"
              class="relative w-full h-full min-h-screen overflow-y-auto ml-64 flex flex-col items-start bg-base-100"
            >
              <div class="w-full max-w-7xl p-10">
                <!-- Router View for Main Content -->
                <router-view v-slot="{ Component }">
                  <component :is="Component" />
                </router-view>
              </div>
              <Footer />
            </div>
          </div>
        </main>
      </template>
    </SideBar>
  </div>
</template>
