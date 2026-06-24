<script setup>
import { ref, computed } from "vue";
import AuthLayout from "@/components/auth-page/AuthLayout.vue";

const register = ref({
  name: "",
  email: "",
  password: "",
  currency: "USD",
  timezone: "Asia/Jakarta",
  agreed: false,
});

function handleRegister() {
  console.log(register.value);
}

const passwordStrength = computed(() => {
  let score = 0;

  const p = register.value.password;

  if (p.length >= 8) score++;
  if (/[A-Z]/.test(p)) score++;
  if (/[0-9]/.test(p)) score++;
  if (/[^A-Za-z0-9]/.test(p)) score++;

  return score;
});

const strengthColor = computed(() => {
  const m = {
    1: "bg-red-400",
    2: "bg-gold-400",
    3: "bg-gold-500",
    4: "bg-success-500",
  };
  return m[passwordStrength.value] || "bg-neutral-200";
});

const strengthTextColor = computed(() => {
  const m = {
    1: "text-red-500",
    2: "text-gold-600",
    3: "text-gold-600",
    4: "text-success-600",
  };
  return m[passwordStrength.value] || "";
});

const strengthLabel = computed(() => {
  return ["", "Weak", "Fair", "Good", "Strong"][passwordStrength.value] || "";
});
</script>

<template>
  <AuthLayout>
    <div class="w-full max-w-100">
      <h2 class="text-[1.75rem] font-extrabold">Create an account</h2>

      <div class="space-y-4 mt-8">
        <!-- Social buttons -->
        <div class="flex gap-3 mb-6">
          <button
            class="flex-1 flex items-center justify-center gap-2 border border-neutral-200 rounded-xl py-2.5 text-[13px] font-medium text-neutral-700 hover:bg-neutral-50 hover:border-neutral-300 transition-colors"
          >
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
            <svg width="14" height="16" viewBox="0 0 14 16" fill="currentColor">
              <path
                d="M13.23 11.56a7.31 7.31 0 01-.73 1.31c-.38.55-.7.93-.94 1.14-.38.35-.78.53-1.22.53-.31 0-.69-.09-1.12-.27-.43-.18-.83-.27-1.2-.27-.38 0-.79.09-1.23.27-.44.18-.8.27-1.08.28-.42.01-.83-.17-1.23-.54-.26-.22-.59-.62-.98-1.19A7.7 7.7 0 011 11.42 7.13 7.13 0 01.5 8.8c0-1 .22-1.86.65-2.59.34-.58.8-1.04 1.37-1.38a3.7 3.7 0 011.86-.52c.37 0 .85.11 1.46.34.61.23 1 .34 1.18.34.13 0 .57-.13 1.32-.4.71-.24 1.3-.35 1.8-.32 1.33.1 2.33.62 2.98 1.55-1.19.72-1.78 1.73-1.77 3.02 0 1.01.38 1.85 1.13 2.52zM10 .32c0 .79-.29 1.53-.86 2.2-.69.81-1.53 1.28-2.43 1.21a2.44 2.44 0 01-.02-.3c0-.76.33-1.57.91-2.23.29-.33.66-.61 1.12-.82.45-.21.88-.32 1.28-.33.01.09.01.18 0 .27z"
              />
            </svg>
            Apple
          </button>
        </div>

        <!-- Divider -->
        <div class="flex items-center gap-3 mb-5">
          <div class="flex-1 h-px bg-neutral-200"></div>
          <span class="text-[12px] text-neutral-400 font-medium">OR</span>
          <div class="flex-1 h-px bg-neutral-200"></div>
        </div>

        <!-- Display Name -->
        <div class="mb-3.5">
          <label
            class="block text-[12.5px] font-semibold text-neutral-700 mb-1.5"
            >Display name</label
          >
          <input
            v-model="register.name"
            type="text"
            placeholder="e.g. John Trader"
            class="input-field w-full border border-neutral-200 rounded-xl px-4 py-3 text-[13.5px] text-neutral-900 placeholder-neutral-400 bg-white"
          />
        </div>

        <!-- Email -->
        <div class="mb-3.5">
          <label
            class="block text-[12.5px] font-semibold text-neutral-700 mb-1.5"
            >Email address</label
          >
          <input
            v-model="register.email"
            type="email"
            placeholder="Enter your email address"
            class="input-field w-full border border-neutral-200 rounded-xl px-4 py-3 text-[13.5px] text-neutral-900 placeholder-neutral-400 bg-white"
          />
        </div>

        <!-- Password -->
        <div class="mb-3.5">
          <label
            class="block text-[12.5px] font-semibold text-neutral-700 mb-1.5"
            >Password</label
          >
          <div class="relative">
            <input
              v-model="register.password"
              :type="register.showPass ? 'text' : 'password'"
              placeholder="Minimum 8 characters"
              class="input-field w-full border border-neutral-200 rounded-xl px-4 py-3 text-[13.5px] text-neutral-900 placeholder-neutral-400 bg-white pr-11"
            />
            <button
              @click="register.showPass = !register.showPass"
              class="pass-toggle absolute right-3.5 top-1/2 -translate-y-1/2 text-neutral-400 hover:text-neutral-600 transition-colors"
            >
              <svg
                v-if="!register.showPass"
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
          <!-- Password strength -->
          <div v-if="register.password.length > 0" class="mt-2 flex gap-1">
            <div
              v-for="i in 4"
              :key="i"
              class="flex-1 h-1 rounded-full transition-colors duration-300"
              :class="i <= passwordStrength ? strengthColor : 'bg-neutral-200'"
            ></div>
          </div>
          <p
            v-if="register.password.length > 0"
            class="text-[11px] mt-1"
            :class="strengthTextColor"
          >
            {{ strengthLabel }}
          </p>
        </div>

        <!-- Currency + Timezone side by side -->
        <div class="grid grid-cols-2 gap-3 mb-5">
          <!-- Currency -->
          <div>
            <label
              class="block text-[12.5px] font-semibold text-neutral-700 mb-1.5"
              >Currency</label
            >
            <div class="relative">
              <select
                v-model="register.currency"
                class="input-field w-full border border-neutral-200 rounded-xl px-4 py-3 text-[13px] text-neutral-900 bg-white cursor-pointer pr-8"
              >
                <option value="USD">🇺🇸 USD</option>
                <option value="EUR">🇪🇺 EUR</option>
                <option value="GBP">🇬🇧 GBP</option>
                <option value="SGD">🇸🇬 SGD</option>
                <option value="AUD">🇦🇺 AUD</option>
                <option value="CAD">🇨🇦 CAD</option>
                <option value="JPY">🇯🇵 JPY</option>
              </select>
            </div>
          </div>
          <!-- Timezone -->
          <div>
            <label
              class="block text-[12.5px] font-semibold text-neutral-700 mb-1.5"
              >Timezone</label
            >
            <div class="relative">
              <select
                v-model="register.timezone"
                class="input-field w-full border border-neutral-200 rounded-xl px-4 py-3 text-[13px] text-neutral-900 bg-white cursor-pointer pr-8"
              >
                <optgroup label="Asia">
                  <option value="Asia/Jakarta">Asia/Jakarta</option>
                  <option value="Asia/Singapore">Asia/Singapore</option>
                  <option value="Asia/Tokyo">Asia/Tokyo</option>
                  <option value="Asia/Seoul">Asia/Seoul</option>
                  <option value="Asia/Shanghai">Asia/Shanghai</option>
                  <option value="Asia/Dubai">Asia/Dubai</option>
                  <option value="Asia/Karachi">Asia/Karachi</option>
                  <option value="Asia/Kolkata">Asia/Kolkata</option>
                  <option value="Asia/Bangkok">Asia/Bangkok</option>
                  <option value="Asia/Makassar">Asia/Makassar</option>
                  <option value="Asia/Jayapura">Asia/Jayapura</option>
                </optgroup>
                <optgroup label="Europe">
                  <option value="Europe/London">Europe/London</option>
                  <option value="Europe/Berlin">Europe/Berlin</option>
                  <option value="Europe/Paris">Europe/Paris</option>
                  <option value="Europe/Amsterdam">Europe/Amsterdam</option>
                </optgroup>
                <optgroup label="Americas">
                  <option value="America/New_York">America/New_York</option>
                  <option value="America/Chicago">America/Chicago</option>
                  <option value="America/Denver">America/Denver</option>
                  <option value="America/Los_Angeles">
                    America/Los_Angeles
                  </option>
                  <option value="America/Toronto">America/Toronto</option>
                </optgroup>
                <optgroup label="Pacific">
                  <option value="Australia/Sydney">Australia/Sydney</option>
                  <option value="Pacific/Auckland">Pacific/Auckland</option>
                </optgroup>
                <optgroup label="UTC">
                  <option value="UTC">UTC</option>
                </optgroup>
              </select>
            </div>
          </div>
        </div>

        <button
          @click="handleRegister"
          class="w-full bg-primary-800 text-white py-3 rounded-xl"
        >
          Create Account
        </button>

        <p class="text-center text-neutral-400">
          Already have an account?

          <RouterLink to="/login" class="font-semibold text-primary-700">
            Log In
          </RouterLink>
        </p>
      </div>
    </div>
  </AuthLayout>
</template>
