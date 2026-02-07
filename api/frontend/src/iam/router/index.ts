import LoginView from "../views/LoginView.vue";
import RegisterView from "../views/RegisterView.vue";
import ProfileView from "../views/ProfileView.vue";

const publicAuthRoutes = [
  {
    path: "/login",
    component: LoginView,
  },
  {
    path: "/register",
    component: RegisterView,
  },
];

const protectedAuthRoutes = [
  {
    path: "/profile",
    component: ProfileView,
  },
];

export { publicAuthRoutes, protectedAuthRoutes };
