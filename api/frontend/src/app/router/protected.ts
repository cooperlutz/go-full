// Route imports
import pingpongRoutes from "~/pingpong/router";
import examLibraryRoutes from "~/examlibrary/router";
// View imports
import HomeView from "~/app/views/HomeView.vue";

const appRoutes = [
  {
    path: "dashboard",
    component: HomeView,
  },
  {
    path: "settings",
    redirect: "coming-soon",
  },
];

const protectedRoutes = [...appRoutes, ...pingpongRoutes, ...examLibraryRoutes];

export default protectedRoutes;
