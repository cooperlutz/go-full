// View imports
import HomeView from "../views/HomeView.vue";
// Route imports
import pingpongRoutes from "~/pingpong/router";
import examLibraryRoutes from "~/examlibrary/router";
import examinationRoutes from "~/examination/router";
import gradingRoutes from "~/grading/router";
import { protectedAuthRoutes } from "~/iam/router";
import reportingRoutes from "~/reporting/router";

const appRoutes = [
  {
    path: "dashboard",
    component: HomeView,
  },
];

const protectedRoutes = [
  ...appRoutes,
  ...pingpongRoutes,
  ...examLibraryRoutes,
  ...examinationRoutes,
  ...protectedAuthRoutes,
  ...gradingRoutes,
  ...reportingRoutes,
];

export default protectedRoutes;
