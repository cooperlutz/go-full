import PingPongAppView from "~/pingpong/views/PingPongAppView.vue";
import PingPongAnalyticsView from "~/pingpong/views/PingPongAnalyticsView.vue";

const pingpongRoutes = [
  {
    path: "/ping-pong",
    redirect: "/ping-pong/app",
  },
  {
    path: "/ping-pong/app",
    component: PingPongAppView,
  },
  {
    path: "/ping-pong/analytics",
    component: PingPongAnalyticsView,
  },
];

export default pingpongRoutes;
