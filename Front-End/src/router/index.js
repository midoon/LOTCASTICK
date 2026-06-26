import { createRouter, createWebHistory } from "vue-router";
import { useAuthStore } from "@/stores/auth";

import WelcomePage from "@/pages/WelcomePage.vue";
import DashboardPage from "@/pages/dashboard/DashboardPage.vue";
import LoginPage from "@/pages/auth/LoginPage.vue";
import RegisterPage from "@/pages/auth/RegisterPage.vue";
import AuthLayout from "../layout/AuthLayout.vue";
import DashboardLayout from "../layout/DashboardLayout.vue";

const routes = [
  {
    path: "/",
    name: "Welcome",
    component: WelcomePage,
  },

  {
    path: "/",
    component: AuthLayout,
    children: [
      {
        path: "login",
        name: "Login",
        component: LoginPage,
      },
      {
        path: "register",
        name: "Register",
        component: RegisterPage,
      },
    ],
  },

  {
    path: "/",
    component: DashboardLayout,
    meta: {
      requiresAuth: true,
    },
    children: [
      {
        path: "dashboard",
        name: "Dashboard",
        component: DashboardPage,
      },
    ],
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

router.beforeEach((to, from) => {
  const authSotre = useAuthStore();

  if (to.meta.requiresAuth && !authSotre.isAuthenticated) {
    return "/";
  }
});

export default router;
