<script setup>
import { ref } from "vue";
import { useRouter } from "vue-router";
import { useAuthStore } from "@/stores/auth";

import AuthLayout from "@/components/auth-page/AuthLayout.vue";

const authStore = useAuthStore();
const router = useRouter();

const login = ref({
  email: "",
  password: "",
  showPass: false,
});

const handleLogin = async () => {
  try {
    await authStore.login(login.value);

    router.push("/dashboard");
  } catch (error) {
    console.log(error);
    alert("Login gagal");
  }
};
</script>

<template>
  <AuthLayout>
    <div class="w-full max-w-100">
      <div class="mb-8">
        <h2 class="text-[1.75rem] font-extrabold">Welcome back</h2>

        <p class="text-neutral-400">Log in to your TradeSim account</p>
      </div>

      <div class="flex gap-3 mb-6">
        <button
          class="flex-1 flex items-center justify-center gap-2 border border-neutral-200 rounded-xl py-2.5 text-[13px] font-medium text-neutral-700 hover:bg-neutral-50 hover:border-neutral-300 transition-colors"
        >
          <!-- Google -->
          <svg width="16" height="16" viewBox="0 0 16 16">
            <path
              d="M15.68 8.18c0-.57-.05-1.11-.14-1.64H8v3.1h4.31a3.68 3.68 0 01-1.6 2.41v2h2.59c1.52-1.4 2.39-3.46 2.39-5.87z"
              fill="#4285F4"
            />
            <path
              d="M8 16c2.16 0 3.97-.72 5.3-1.94l-2.59-2a4.8 4.8 0 01-7.15-2.52H.96v2.07A8 8 0 008 16z"
              fill="#34A853"
            />
            <path
              d="M3.56 9.54A4.8 4.8 0 013.3 8c0-.54.09-1.06.26-1.54V4.39H.96A8 8 0 000 8c0 1.29.31 2.51.96 3.61l2.6-2.07z"
              fill="#FBBC05"
            />
            <path
              d="M8 3.18c1.22 0 2.31.42 3.17 1.24l2.38-2.38A8 8 0 00.96 4.39l2.6 2.07A4.77 4.77 0 018 3.18z"
              fill="#EA4335"
            />
          </svg>
          Google
        </button>
        <button
          class="flex-1 flex items-center justify-center gap-2 border border-neutral-200 rounded-xl py-2.5 text-[13px] font-medium text-neutral-700 hover:bg-neutral-50 hover:border-neutral-300 transition-colors"
        >
          <!-- Apple -->
          <svg width="14" height="16" viewBox="0 0 14 16" fill="currentColor">
            <path
              d="M13.23 11.56a7.31 7.31 0 01-.73 1.31c-.38.55-.7.93-.94 1.14-.38.35-.78.53-1.22.53-.31 0-.69-.09-1.12-.27-.43-.18-.83-.27-1.2-.27-.38 0-.79.09-1.23.27-.44.18-.8.27-1.08.28-.42.01-.83-.17-1.23-.54-.26-.22-.59-.62-.98-1.19A7.7 7.7 0 011 11.42 7.13 7.13 0 01.5 8.8c0-1 .22-1.86.65-2.59.34-.58.8-1.04 1.37-1.38a3.7 3.7 0 011.86-.52c.37 0 .85.11 1.46.34.61.23 1 .34 1.18.34.13 0 .57-.13 1.32-.4.71-.24 1.3-.35 1.8-.32 1.33.1 2.33.62 2.98 1.55-1.19.72-1.78 1.73-1.77 3.02 0 1.01.38 1.85 1.13 2.52zM10 .32c0 .79-.29 1.53-.86 2.2-.69.81-1.53 1.28-2.43 1.21a2.44 2.44 0 01-.02-.3c0-.76.33-1.57.91-2.23.29-.33.66-.61 1.12-.82.45-.21.88-.32 1.28-.33.01.09.01.18 0 .27z"
            />
          </svg>
          Apple
        </button>
      </div>

      <!-- Divider -->
      <div class="flex items-center gap-3 mb-6">
        <div class="flex-1 h-px bg-neutral-200"></div>
        <span class="text-[12px] text-neutral-400 font-medium">OR</span>
        <div class="flex-1 h-px bg-neutral-200"></div>
      </div>

      <div class="space-y-5">
        <!-- email -->
        <div>
          <label
            class="block text-[12.5px] font-semibold text-neutral-700 mb-1.5"
            >Email address</label
          >
          <input
            v-model="login.email"
            type="email"
            placeholder="Enter your email address"
            class="input-field w-full border border-neutral-200 rounded-xl px-4 py-3 text-[13.5px] text-neutral-900 placeholder-neutral-400 bg-white"
          />
        </div>

        <!-- password -->
        <div class="mb-5">
          <label
            class="block text-[12.5px] font-semibold text-neutral-700 mb-1.5"
            >Password</label
          >
          <div class="relative">
            <input
              v-model="login.password"
              :type="login.showPass ? 'text' : 'password'"
              placeholder="Enter your password"
              class="input-field w-full border border-neutral-200 rounded-xl px-4 py-3 text-[13.5px] text-neutral-900 placeholder-neutral-400 bg-white pr-11"
            />
            <button
              @click="login.showPass = !login.showPass"
              class="pass-toggle absolute right-3.5 top-1/2 -translate-y-1/2 text-neutral-400 hover:text-neutral-600 transition-colors"
            >
              <svg
                v-if="!login.showPass"
                width="16"
                height="16"
                viewBox="0 0 16 16"
                fill="none"
              >
                <path
                  d="M1 8s2.5-5 7-5 7 5 7 5-2.5 5-7 5-7-5-7-5z"
                  stroke="currentColor"
                  stroke-width="1.4"
                />
                <circle
                  cx="8"
                  cy="8"
                  r="2"
                  stroke="currentColor"
                  stroke-width="1.4"
                />
              </svg>
              <svg
                v-else
                width="16"
                height="16"
                viewBox="0 0 16 16"
                fill="none"
              >
                <path
                  d="M2 2l12 12M6.5 6.5A2 2 0 009.5 9.5M1 8s2.5-5 7-5c1.1 0 2.1.24 3 .66M15 8s-.7 1.4-2 2.5"
                  stroke="currentColor"
                  stroke-width="1.4"
                  stroke-linecap="round"
                />
              </svg>
            </button>
          </div>
        </div>

        <button
          @click="handleLogin"
          class="w-full bg-primary-800 text-white py-3 rounded-xl"
        >
          Log In
        </button>

        <p class="text-center text-neutral-400">
          Don't have an account?

          <RouterLink to="/register" class="font-semibold text-primary-700">
            Sign Up
          </RouterLink>
        </p>
      </div>
    </div>
  </AuthLayout>
</template>
