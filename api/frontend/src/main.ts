import "~/app/styles/app.css";

import { createApp } from "vue";
// import { createPinia } from "pinia";
// import piniaPluginPersistedstate from "pinia-plugin-persistedstate";

import router from "~/app/router";
import App from "~/app/App.vue";

export function createNewApp() {
  // Create a Vue App
  const app = createApp(App);

  // setup plugins
  // const pinia = createPinia();
  // pinia.use(piniaPluginPersistedstate);

  // register plugins
  // app.use(pinia);

  // Register the router
  app.use(router);

  return app;
}

const app = createNewApp();
// mount the app
app.mount("#app");
