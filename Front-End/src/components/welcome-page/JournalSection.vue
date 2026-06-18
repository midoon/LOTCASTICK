<script setup>
import trades from "@/data/welcome-page/trades.js";
import journalFeatures from "@/data/welcome-page/JournalFeatures.js";
</script>

<template>
  <section class="py-20 max-w-6xl mx-auto px-6">
    <div class="grid md:grid-cols-2 gap-16 items-center">
      <!-- Trade journal mock -->
      <div class="space-y-2.5 order-2 md:order-1">
        <!-- Table header -->
        <div
          class="grid grid-cols-5 px-4 py-2 text-[11px] font-semibold text-neutral-400 uppercase tracking-wider"
        >
          <span>Symbol</span><span>Dir</span><span>P&amp;L</span><span>R:R</span
          ><span>Session</span>
        </div>
        <div
          v-for="t in trades"
          :key="t.sym"
          class="bg-white rounded-xl border border-neutral-200 grid grid-cols-5 px-4 py-3.5 items-center feature-card"
        >
          <span class="text-[13px] font-semibold text-primary-900 font-mono">{{
            t.sym
          }}</span>
          <span
            class="text-[11px] font-bold px-2 py-0.5 rounded-md inline-block w-fit"
            :class="
              t.dir === 'L'
                ? 'bg-success-50 text-success-700'
                : 'bg-red-50 text-red-600'
            "
            >{{ t.dir === "L" ? "LONG" : "SHORT" }}</span
          >
          <span
            class="text-[13px] font-mono font-semibold"
            :class="t.pnl > 0 ? 'text-success-600' : 'text-red-500'"
            >{{ t.pnl > 0 ? "+" : "" }}${{ t.pnl }}</span
          >
          <span class="text-[13px] font-mono text-neutral-600">{{ t.rr }}</span>
          <span class="text-[11px] text-neutral-400 font-medium">{{
            t.session
          }}</span>
        </div>
      </div>

      <!-- Copy -->
      <div class="order-1 md:order-2">
        <div
          class="text-[12px] font-semibold text-primary-500 uppercase tracking-[0.12em] mb-3"
        >
          Trading Journal
        </div>
        <h2
          class="text-[2.2rem] font-extrabold text-primary-900 tracking-tight leading-tight mb-5"
        >
          Log trades.<br />
          Find your edge.
        </h2>
        <p class="text-neutral-500 text-[15px] leading-relaxed mb-6">
          Record every closed trade with entry/exit, lot size, session, and
          screenshots. Filter by symbol, strategy, or tags. Let the data tell
          you which setups are actually working.
        </p>
        <div class="space-y-3">
          <div
            v-for="j in journalFeatures"
            :key="j"
            class="flex items-center gap-3 text-[14px] text-neutral-600"
          >
            <div
              class="w-5 h-5 rounded-full bg-success-50 border border-success-200 flex items-center justify-center shrink-0"
            >
              <svg width="10" height="10" viewBox="0 0 10 10" fill="none">
                <polyline
                  points="2,5.5 4,7.5 8,3"
                  stroke="#17ab70"
                  stroke-width="1.6"
                  stroke-linecap="round"
                  stroke-linejoin="round"
                />
              </svg>
            </div>
            {{ j }}
          </div>
        </div>
      </div>
    </div>
  </section>
</template>
