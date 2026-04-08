import { createRouter, createWebHistory } from "vue-router";

import { useAuthState } from "../../iam/stores/useAuthState";
// Layout imports
import ShellComponent from "../layouts/ApplicationShell/ApplicationShell.vue";
// View imports
import ComingSoon from "../views/ComingSoon.vue";
import Error404View from "../views/Error404View.vue";
// Route imports
import INSECURE_ROUTES from "./insecure";
import protectedRoutes from "./protected";

/**
 * @constant {Array} routes - The complete set of routes for the application.
 * @description This array combines secure, admin, and insecure routes into a single routing configuration.
 */
const routes = [
  {
    path: "/",
    component: ShellComponent,
    redirect: "/dashboard",
    meta: { requiresAuth: true },
    children: [...protectedRoutes],
  },
  {
    path: "/coming-soon",
    component: ComingSoon,
  },
  {
    path: "/:pathMatch(.*)",
    name: "not-found",
    component: Error404View,
    meta: { requiresAuth: true },
  },
  ...INSECURE_ROUTES,
];

/**
 * @module router
 * @description This module provides a collection of utility functions for common tasks.
 * It includes functions for string manipulation, array operations, and date formatting.
 */
const router = createRouter({
  history: createWebHistory(),
  routes,
});

// Navigation guard to check authentication
// This will redirect users to the login page if they try to access a route that requires authentication
router.beforeEach((to, _from, next) => {
  const { isAuthenticated } = useAuthState();
  if (isAuthenticated()) {
    next();
  } else if (to.meta.requiresAuth) {
    next("/login");
  } else {
    next();
  }
});

export default router;
