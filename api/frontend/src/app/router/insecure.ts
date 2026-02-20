import HealthView from "../views/HealthView.vue";

import { publicAuthRoutes } from "~/iam/router";

const INSECURE_ROUTES = [
  {
    path: "/health",
    component: HealthView,
  },
  ...publicAuthRoutes,
];

export default INSECURE_ROUTES;
