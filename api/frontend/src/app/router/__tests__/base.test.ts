import router from "../index.ts"; // adjust the path as necessary

import { expect, describe, it } from "vitest";
// import { shallowMount } from "@vue/test-utils";
// // import { createRouter, createWebHistory } from 'vue-router'
// import ShellComponent from "~/app/layouts/ApplicationShell/ApplicationShell.vue";

// import HealthView from "~/app/views/HealthView.vue";

describe("HealthView", () => {
  it("navigates to /health when button is clicked", async () => {
    // const routes = [
    //   // { path: '/', component: ShellComponent},
    //   { path: '/health', component: HealthView }
    // ]
    // const router = createRouter({
    //   history: createWebHistory(),
    //   routes
    // })

    // const wrapper = shallowMount(HealthView, {
    //   global: {
    //     plugins: [router]
    //   }
    // })

    await router.push("/health"); // Initialize router to a known state
    // await wrapper.push('/health') // Simulate button click that triggers navigation

    expect(router.currentRoute.value.path).toBe("/health");
  });
});
