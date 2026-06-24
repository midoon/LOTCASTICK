<script setup></script>

<template>
  <section
    class="relative min-h-screen flex items-center overflow-hidden pt-16 hero-noise"
  >
    <!-- Background circles -->
    <div
      class="absolute -top-32 -right-48 w-150 h-150 rounded-full bg-primary-100/60 blur-3xl pointer-events-none"
    ></div>
    <div
      class="absolute bottom-0 -left-32 w-100 h-100 rounded-full bg-success-50/70 blur-3xl pointer-events-none"
    ></div>

    <div
      class="max-w-6xl mx-auto px-6 w-full grid md:grid-cols-2 gap-12 items-center py-20"
    >
      <!-- Left: copy -->
      <div class="animate-fade-up">
        <!-- Badge -->
        <div
          class="inline-flex items-center gap-2 px-3 py-1.5 rounded-full border border-primary-200 bg-primary-50 text-primary-600 text-[12px] font-medium mb-7"
        >
          <span
            class="w-1.5 h-1.5 rounded-full bg-success-500 pulse-dot inline-block"
          ></span>
          Prop Firm Challenge Simulator
        </div>

        <h1
          class="text-[2.75rem] md:text-[3.5rem] font-extrabold leading-[1.08] tracking-tight text-primary-900 mb-6"
        >
          Master the<br />
          <span class="gradient-text">Prop Firm Rules.</span><br />
          Pass with Confidence.
        </h1>

        <p class="text-neutral-500 text-[16px] leading-relaxed mb-8 max-w-md">
          Simulate any prop firm challenge — FTMO, MyFundedFx, TopStep — with a
          real-time risk engine, equity tracking, and deep trade analytics. No
          live capital at risk.
        </p>

        <div class="flex items-center gap-3 flex-wrap">
          <RouterLink
            to="/register"
            class="bg-primary-800 text-neutral-50 text-[14px] font-light px-6 py-3 rounded-xl hover:bg-primary-700 transition-colors shadow-md shadow-primary-900/10"
          >
            Start Simulating Free
          </RouterLink>
          <button
            class="border border-neutral-300 bg-white text-neutral-700 text-[14px] font-medium px-6 py-3 rounded-xl hover:border-primary-300 hover:text-primary-700 transition-colors"
          >
            See How It Works
          </button>
        </div>

        <!-- Stats row -->
        <div class="mt-10 flex items-center gap-6 flex-wrap">
          <div>
            <div class="text-[22px] font-bold text-primary-900 font-mono">
              20
            </div>
            <div class="text-[12px] text-neutral-400 font-medium">
              Active Sims/User
            </div>
          </div>
          <div class="w-px h-8 bg-neutral-200"></div>
          <div>
            <div class="text-[22px] font-bold text-primary-900 font-mono">
              4+
            </div>
            <div class="text-[12px] text-neutral-400 font-medium">
              Drawdown Types
            </div>
          </div>
          <div class="w-px h-8 bg-neutral-200"></div>
          <div>
            <div class="text-[22px] font-bold text-success-600 font-mono">
              &lt;200ms
            </div>
            <div class="text-[12px] text-neutral-400 font-medium">
              Risk Engine Speed
            </div>
          </div>
        </div>
      </div>

      <!-- Right: mock dashboard card -->
      <div class="animate-fade-up-delay2 relative">
        <!-- Main card -->
        <div
          class="glass-card rounded-2xl border border-neutral-200 shadow-xl shadow-primary-900/8 overflow-hidden"
        >
          <!-- Card header -->
          <div
            class="px-5 pt-5 pb-4 border-b border-neutral-100 flex items-center justify-between"
          >
            <div>
              <div
                class="text-[11px] text-neutral-400 font-medium uppercase tracking-widest mb-1"
              >
                FTMO $100K — Phase 1
              </div>
              <div class="flex items-center gap-2">
                <span
                  class="text-[22px] font-bold text-primary-900 font-mono count-anim"
                  >$104,280</span
                >
                <span
                  class="text-[11px] font-semibold text-success-600 bg-success-50 px-2 py-0.5 rounded-full"
                  >+4.28%</span
                >
              </div>
            </div>
            <div
              class="w-8 h-8 rounded-lg bg-success-50 border border-success-200 flex items-center justify-center"
            >
              <span class="text-success-600 text-xs font-bold">✓</span>
            </div>
          </div>

          <!-- Equity chart -->
          <div class="px-5 py-4">
            <svg
              class="w-full chart-glow"
              height="90"
              viewBox="0 0 340 90"
              preserveAspectRatio="none"
            >
              <defs>
                <linearGradient id="eq-fill" x1="0" y1="0" x2="0" y2="1">
                  <stop offset="0%" stop-color="#615a87" stop-opacity="0.18" />
                  <stop offset="100%" stop-color="#615a87" stop-opacity="0" />
                </linearGradient>
              </defs>
              <!-- Floor line -->
              <line
                x1="0"
                y1="72"
                x2="340"
                y2="72"
                stroke="#d1cede"
                stroke-width="1"
                stroke-dasharray="5,4"
              />
              <!-- Equity path fill -->
              <path
                d="M0,70 C30,68 50,72 70,65 C90,58 110,60 130,52 C150,44 160,55 180,48 C200,41 220,38 240,30 C260,22 280,28 300,20 C315,14 330,16 340,10 L340,90 L0,90 Z"
                fill="url(#eq-fill)"
              />
              <!-- Equity path stroke -->
              <path
                d="M0,70 C30,68 50,72 70,65 C90,58 110,60 130,52 C150,44 160,55 180,48 C200,41 220,38 240,30 C260,22 280,28 300,20 C315,14 330,16 340,10"
                stroke="#615a87"
                stroke-width="2.2"
                fill="none"
                stroke-linecap="round"
                stroke-linejoin="round"
              />
              <!-- Profit target dotted line -->
              <line
                x1="0"
                y1="30"
                x2="340"
                y2="30"
                stroke="#1ccf88"
                stroke-width="1"
                stroke-dasharray="5,4"
                opacity="0.7"
              />
              <!-- Label -->
              <text
                x="4"
                y="26"
                font-size="8"
                fill="#1ccf88"
                font-family="JetBrains Mono"
                opacity="0.9"
              >
                target
              </text>
              <text
                x="4"
                y="68"
                font-size="8"
                fill="#aea9c6"
                font-family="JetBrains Mono"
              >
                floor
              </text>
            </svg>
          </div>

          <!-- Progress bars -->
          <div class="px-5 pb-5 space-y-3">
            <div>
              <div class="flex justify-between text-[11px] font-medium mb-1.5">
                <span class="text-neutral-500">Profit Target</span>
                <span class="text-primary-700 font-mono">4.28% / 10%</span>
              </div>
              <div class="h-1.5 bg-neutral-100 rounded-full overflow-hidden">
                <div
                  class="progress-fill h-full bg-primary-500 rounded-full"
                  style="--fill: 43%"
                ></div>
              </div>
            </div>
            <div>
              <div class="flex justify-between text-[11px] font-medium mb-1.5">
                <span class="text-neutral-500">Trading Days</span>
                <span class="text-primary-700 font-mono">6 / 10 days</span>
              </div>
              <div class="h-1.5 bg-neutral-100 rounded-full overflow-hidden">
                <div
                  class="progress-fill h-full bg-gold-500 rounded-full"
                  style="--fill: 60%"
                ></div>
              </div>
            </div>
            <div>
              <div class="flex justify-between text-[11px] font-medium mb-1.5">
                <span class="text-neutral-500">Daily Drawdown Used</span>
                <span class="text-success-600 font-mono">0.8% / 5%</span>
              </div>
              <div class="h-1.5 bg-neutral-100 rounded-full overflow-hidden">
                <div
                  class="progress-fill h-full bg-success-400 rounded-full"
                  style="--fill: 16%"
                ></div>
              </div>
            </div>
          </div>
        </div>

        <!-- Floating mini card: last trade -->
        <div
          class="absolute bg-primary-50 -bottom-6 -left-6 glass-card rounded-xl border border-neutral-200 shadow-lg px-4 py-3 items-center gap-3 hidden md:flex"
        >
          <div
            class="w-7 h-7 rounded-lg bg-success-50 flex items-center justify-center text-success-600"
          >
            <svg width="14" height="14" viewBox="0 0 14 14" fill="none">
              <polyline
                points="2,10 5,6 8,8 12,3"
                stroke="currentColor"
                stroke-width="1.8"
                stroke-linecap="round"
                stroke-linejoin="round"
              />
              <polyline
                points="10,3 12,3 12,5"
                stroke="currentColor"
                stroke-width="1.8"
                stroke-linecap="round"
                stroke-linejoin="round"
              />
            </svg>
          </div>
          <div>
            <div class="text-[11px] font-semibold text-primary-900">
              XAUUSD LONG
            </div>
            <div class="text-[10px] text-success-600 font-mono font-medium">
              +$320 · R:R 2.4
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<style scoped>
.hero-noise {
  background-image: url("data:image/svg+xml,%3Csvg viewBox='0 0 200 200' xmlns='http://www.w3.org/2000/svg'%3E%3Cfilter id='n'%3E%3CfeTurbulence type='fractalNoise' baseFrequency='0.75' numOctaves='4' stitchTiles='stitch'/%3E%3C/filter%3E%3Crect width='100%25' height='100%25' filter='url(%23n)' opacity='0.04'/%3E%3C/svg%3E");
}

/* Fade up */
.animate-fade-up {
  animation: fadeUp 0.8s ease forwards;
}

.animate-fade-up-delay2 {
  opacity: 0;
  animation: fadeUp 0.8s ease forwards;
  animation-delay: 0.3s;
}

/* Dot pulse */
.pulse-dot {
  animation: pulseDot 1.5s infinite;
}

/* Progress bar */
.progress-fill {
  width: var(--fill);
  animation: fillBar 1s ease;
}

@keyframes fadeUp {
  from {
    opacity: 0;
    transform: translateY(25px);
  }

  to {
    opacity: 1;
    transform: translateY(0);
  }
}

@keyframes pulseDot {
  0%,
  100% {
    transform: scale(1);
    opacity: 1;
  }

  50% {
    transform: scale(1.4);
    opacity: 0.6;
  }
}

@keyframes fillBar {
  from {
    width: 0;
  }

  to {
    width: var(--fill);
  }
}
</style>
