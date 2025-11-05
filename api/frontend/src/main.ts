import "~/app/styles/app.css";

import { createApp } from "vue";

import router from "~/app/router";
import App from "~/app/App.vue";

async function prepareApp() {
  if (
    process.env.NODE_ENV === "development" ||
    process.env.NODE_ENV === "test"
  ) {
    const { worker } = await import("../test/mocks/browser");
    return worker.start();
  }

  return Promise.resolve();
}

export function createNewApp() {
  // Create a Vue App
  const app = createApp(App);

  // Register the router
  app.use(router);

  return app;
}

const app = createNewApp();

// mount the app
prepareApp().then(() => {
  app.mount("#app");
});
