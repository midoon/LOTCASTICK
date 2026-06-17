import { createRouter, createWebHistory } from "vue-router";
import WelcomePage from "@/pages/WelcomePage.vue";
import DashboardPage from "@/pages/dashboard/DashboardPage.vue";
import LoginPage from "@/pages/auth/LoginPage.vue";
import RegisterPage from "@/pages/auth/RegisterPage.vue";

const routes = [
  {
    path: "/",
    name: "Welcome",
    component: WelcomePage,
  },

  {
    path: "/login",
    name: "Login",
    component: LoginPage,
  },
  {
    path: "/register",
    name: "Register",
    component: RegisterPage,
  },
  {
    path: "/dashboard",
    name: "Dashboard",
    component: DashboardPage,
  },
];

export const router = createRouter({
  history: createWebHistory(),
  routes,
});
