import HealthView from "~/app/views/HealthView.vue";

const INSECURE_ROUTES = [
  {
    path: "/health",
    component: HealthView,
  },
];

export default INSECURE_ROUTES;
