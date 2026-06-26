<script setup></script>

<template>
  <div class="min-h-screen flex">
    <!-- LEFT PANEL -->
    <div class="hidden lg:flex lg:w-[46%] panel-bg">
      <div
        class="hidden lg:flex w-full panel-bg flex-col justify-between p-12 relative overflow-hidden"
      >
        <!-- Decorative blobs -->
        <div
          class="absolute top-0 right-0 w-80 h-80 rounded-full bg-primary-500/10 blur-3xl pointer-events-none"
        ></div>
        <div
          class="absolute bottom-20 -left-16 w-64 h-64 rounded-full bg-success-500/8 blur-3xl pointer-events-none"
        ></div>

        <!-- Top: brand -->
        <div>
          <RouterLink
            class="flex items-center gap-2.5 mb-16 hover:cursor-pointer"
            to="/"
          >
            <div
              class="w-9 h-9 rounded-xl bg-white/10 flex items-center justify-center"
            >
              <svg width="16" height="16" viewBox="0 0 16 16" fill="none">
                <polyline
                  points="1,12 5,7 8,9 12,4 15,6"
                  stroke="#d1cede"
                  stroke-width="1.8"
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  fill="none"
                />
                <circle cx="15" cy="6" r="1.5" fill="#1ccf88" />
              </svg>
            </div>
            <span class="font-bold text-white text-[16px] tracking-tight"
              >LotCastick</span
            >
          </RouterLink>

          <!-- Headline -->
          <h1
            class="text-[2.4rem] font-extrabold text-white leading-[1.1] tracking-tight mb-4"
          >
            Simulate.<br />
            Analyze.<br />
            <span class="text-primary-300">Pass.</span>
          </h1>
          <p class="text-primary-300 text-[14.5px] leading-relaxed max-w-xs">
            The prop firm challenge simulator built for serious traders. No live
            capital, full discipline.
          </p>
        </div>

        <!-- Middle: mock dashboard card -->
        <div class="my-10">
          <div
            class="bg-white/8 border border-white/10 rounded-2xl p-5 backdrop-blur-sm"
          >
            <!-- Card header -->
            <div class="flex items-center justify-between mb-4">
              <div>
                <div
                  class="text-[10px] text-primary-300 font-medium uppercase tracking-widest mb-1"
                >
                  FTMO $100K · Phase 1
                </div>
                <div class="flex items-center gap-2">
                  <span class="text-[20px] font-bold text-white font-mono"
                    >$104,280</span
                  >
                  <span
                    class="text-[10px] font-semibold text-success-400 bg-success-500/15 px-2 py-0.5 rounded-full"
                    >+4.28%</span
                  >
                </div>
              </div>
              <div
                class="w-7 h-7 rounded-lg bg-success-500/20 flex items-center justify-center"
              >
                <svg width="12" height="12" viewBox="0 0 12 12" fill="none">
                  <polyline
                    points="2,8 4.5,5 7,6.5 10,2"
                    stroke="#1ccf88"
                    stroke-width="1.6"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                  />
                </svg>
              </div>
            </div>

            <!-- Mini equity SVG -->
            <svg
              class="w-full chart-glow mb-4"
              height="56"
              viewBox="0 0 280 56"
              preserveAspectRatio="none"
            >
              <defs>
                <linearGradient id="gfill" x1="0" y1="0" x2="0" y2="1">
                  <stop offset="0%" stop-color="#615a87" stop-opacity="0.3" />
                  <stop offset="100%" stop-color="#615a87" stop-opacity="0" />
                </linearGradient>
              </defs>
              <line
                x1="0"
                y1="44"
                x2="280"
                y2="44"
                stroke="#aea9c6"
                stroke-width="0.8"
                stroke-dasharray="4,3"
                opacity="0.4"
              />
              <line
                x1="0"
                y1="18"
                x2="280"
                y2="18"
                stroke="#1ccf88"
                stroke-width="0.8"
                stroke-dasharray="4,3"
                opacity="0.5"
              />
              <path
                d="M0,44 C20,42 35,45 55,40 C75,35 90,37 110,31 C130,25 145,33 165,27 C185,21 200,18 220,13 C240,8 260,11 280,6 L280,56 L0,56 Z"
                fill="url(#gfill)"
              />
              <path
                d="M0,44 C20,42 35,45 55,40 C75,35 90,37 110,31 C130,25 145,33 165,27 C185,21 200,18 220,13 C240,8 260,11 280,6"
                stroke="#857fa9"
                stroke-width="1.8"
                fill="none"
                stroke-linecap="round"
              />
            </svg>

            <!-- Progress bars -->
            <div class="space-y-2.5">
              <div>
                <div class="flex justify-between text-[10px] font-medium mb-1">
                  <span class="text-primary-300">Profit Target</span>
                  <span class="text-white/70 font-mono">4.28% / 10%</span>
                </div>
                <div class="h-1 bg-white/10 rounded-full">
                  <div
                    class="h-full bg-primary-400 rounded-full"
                    style="width: 43%"
                  ></div>
                </div>
              </div>
              <div>
                <div class="flex justify-between text-[10px] font-medium mb-1">
                  <span class="text-primary-300">Trading Days</span>
                  <span class="text-white/70 font-mono">6 / 10</span>
                </div>
                <div class="h-1 bg-white/10 rounded-full">
                  <div
                    class="h-full bg-gold-400 rounded-full"
                    style="width: 60%"
                  ></div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Bottom: testimonial / trust -->
        <!-- <div class="space-y-3">
          <div class="flex items-center gap-2.5">
            <div
              class="w-1.5 h-1.5 rounded-full bg-success-400 animate-pulse-dot"
            ></div>
            <span class="text-[12px] text-primary-300"
              >Real-time risk engine · &lt;200ms recalc</span
            >
          </div>
          <div class="flex items-center gap-2.5">
            <div class="w-1.5 h-1.5 rounded-full bg-primary-400"></div>
            <span class="text-[12px] text-primary-300"
              >FTMO, MyFundedFx, TopStep templates</span
            >
          </div>
          <div class="flex items-center gap-2.5">
            <div class="w-1.5 h-1.5 rounded-full bg-gold-400"></div>
            <span class="text-[12px] text-primary-300"
              >Deep analytics — 8 performance dimensions</span
            >
          </div>
        </div> -->
      </div>
    </div>

    <!-- RIGHT PANEL -->
    <div class="flex-1 flex justify-center items-center bg-neutral-50 px-6">
      <RouterView />
    </div>
  </div>
</template>

<style scoped>
* {
  box-sizing: border-box;
}
body {
  font-family: "Inter", sans-serif;
  -webkit-font-smoothing: antialiased;
}

/* Custom select arrow */
select {
  appearance: none;
  -webkit-appearance: none;
  background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='12' height='12' viewBox='0 0 12 12'%3E%3Cpath d='M2 4l4 4 4-4' stroke='%2397969c' stroke-width='1.5' fill='none' stroke-linecap='round' stroke-linejoin='round'/%3E%3C/svg%3E");
  background-repeat: no-repeat;
  background-position: right 12px center;
}

/* Input focus ring */
.input-field:focus {
  outline: none;
  border-color: #615a87;
  box-shadow: 0 0 0 3px rgba(97, 90, 135, 0.12);
}
.input-field {
  transition:
    border-color 0.15s,
    box-shadow 0.15s;
}

/* Left panel chart decorative */
.chart-glow {
  filter: drop-shadow(0 0 10px rgba(97, 90, 135, 0.4));
}

/* Password toggle */
.pass-toggle {
  cursor: pointer;
  user-select: none;
}

/* Tab active underline */
.tab-active {
  border-bottom: 2px solid #302d43;
  color: #302d43;
}
.tab-inactive {
  border-bottom: 2px solid transparent;
  color: #97969c;
}
.tab-inactive:hover {
  color: #5a595f;
}

/* Scrollbar */
::-webkit-scrollbar {
  width: 4px;
}
::-webkit-scrollbar-track {
  background: transparent;
}
::-webkit-scrollbar-thumb {
  background: #d8d8da;
  border-radius: 4px;
}

/* Left panel gradient */
.panel-bg {
  background: linear-gradient(145deg, #201e2c 0%, #302d43 50%, #464162 100%);
}
</style>
