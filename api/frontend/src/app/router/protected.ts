// Route imports
import pingpongRoutes from "~/pingpong/router";
import examLibraryRoutes from "~/examlibrary/router";
import examinationRoutes from "~/examination/router";
import gradingRoutes from "~/grading/router";
import { protectedAuthRoutes } from "~/iam/router";
// View imports
import HomeView from "~/app/views/HomeView.vue";

const appRoutes = [
  {
    path: "dashboard",
    component: HomeView,
  },
  {
    path: "reporting",
    redirect: "coming-soon",
  },
];

const protectedRoutes = [
  ...appRoutes,
  ...pingpongRoutes,
  ...examLibraryRoutes,
  ...examinationRoutes,
  ...protectedAuthRoutes,
  ...gradingRoutes,
];

export default protectedRoutes;
