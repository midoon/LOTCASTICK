import { defineStore } from "pinia";
import api from "@/services/api";

export const useAuthStore = defineStore("auth", {
  state: () => ({
    userId: JSON.parse(localStorage.getItem("userId")) || null,
    accessToken: localStorage.getItem("accessToken") || null,
    refreshToken: localStorage.getItem("refreshToken") || null,
  }),

  getters: {
    isAuthenticated: (state) => !!state.accessToken,
  },

  actions: {
    async register(payload) {
      const response = await api.post("/register", payload);

      return response.data;
    },

    async login(payload) {
      const response = await api.post("/login", payload);

      const { access_token, refresh_token, user_id } = response.data;

      this.accessToken = access_token;
      this.refreshToken = refresh_token;
      this.userId = user_id;

      localStorage.setItem("accessToken", access_token);
      localStorage.setItem("refreshToken", refresh_token);
      localStorage.setItem("userId", JSON.stringify(user_id));
    },

    logout() {
      this.userId = null;
      this.accessToken = null;
      this.refreshToken = null;

      localStorage.removeItem("accessToken");
      localStorage.removeItem("refreshToken");
      localStorage.removeItem("userId");
    },
  },
});
