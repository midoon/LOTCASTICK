<script setup>
import drawdownRules from "@/data/welcome-page/drawdownRules.js";
</script>

<template>
  <section id="rules" class="py-20 bg-primary-950 relative overflow-hidden">
    <div
      class="absolute inset-0 hero-noise pointer-events-none opacity-50"
    ></div>
    <div
      class="absolute top-0 right-0 w-96 h-96 rounded-full bg-primary-700/20 blur-3xl pointer-events-none"
    ></div>

    <div class="max-w-6xl mx-auto px-6 relative z-10">
      <div class="grid md:grid-cols-2 gap-16 items-center">
        <!-- Left: copy -->
        <div>
          <div
            class="text-[12px] font-semibold text-primary-400 uppercase tracking-[0.12em] mb-3"
          >
            Risk Engine
          </div>
          <h2
            class="text-[2.2rem] font-extrabold text-neutral-50 tracking-tight leading-tight mb-5"
          >
            Real-Time Drawdown.<br />
            <span class="text-primary-300">Zero Surprises.</span>
          </h2>
          <p class="text-neutral-400 text-[15px] leading-relaxed mb-8">
            Every trade mutation triggers instant recalculation of your equity
            curve, high-water mark, and all drawdown rules. Violations are
            detected the moment they happen — not at end of day.
          </p>

          <div class="space-y-4">
            <div
              v-for="rule in drawdownRules"
              :key="rule.type"
              class="flex items-start gap-4"
            >
              <div
                class="w-2 h-2 rounded-full mt-2 shrink-0"
                :class="rule.dot"
              ></div>
              <div>
                <div class="text-[14px] font-semibold text-neutral-200">
                  {{ rule.type }}
                </div>
                <div
                  class="text-[13px] text-neutral-500 leading-relaxed mt-0.5"
                >
                  {{ rule.desc }}
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Right: code card -->
        <div
          class="bg-primary-900/60 border border-primary-700/50 rounded-2xl overflow-hidden"
        >
          <div
            class="flex items-center gap-2 px-4 py-3 border-b border-primary-700/40"
          >
            <div class="w-3 h-3 rounded-full bg-red-500/70"></div>
            <div class="w-3 h-3 rounded-full bg-gold-400/70"></div>
            <div class="w-3 h-3 rounded-full bg-success-500/70"></div>
            <span class="ml-3 text-[11px] font-mono text-primary-400"
              >account_status.json</span
            >
          </div>
          <pre
            class="p-5 text-[12px] font-mono leading-relaxed text-primary-200 overflow-x-auto no-scrollbar"
          ><code><span class="text-neutral-500">{</span>
  <span class="text-gold-400">"status"</span>: <span class="text-success-400">"ACTIVE"</span>,
  <span class="text-gold-400">"equity"</span>: <span class="text-primary-300">104280.00</span>,
  <span class="text-gold-400">"drawdown_used_pct"</span>: <span class="text-primary-300">2.5</span>,
  <span class="text-gold-400">"drawdown_remaining_pct"</span>: <span class="text-success-400">7.5</span>,
  <span class="text-gold-400">"daily_drawdown_used_pct"</span>: <span class="text-primary-300">0.8</span>,
  <span class="text-gold-400">"profit_achieved_pct"</span>: <span class="text-success-400">4.28</span>,
  <span class="text-gold-400">"trading_days_completed"</span>: <span class="text-primary-300">6</span>,
  <span class="text-gold-400">"trading_days_required"</span>: <span class="text-primary-300">10</span>,
  <span class="text-gold-400">"consistency_score_pct"</span>: <span class="text-primary-300">22.3</span>,
  <span class="text-gold-400">"violations"</span>: <span class="text-success-400">[]</span>
<span class="text-neutral-500">}</span></code></pre>
        </div>
      </div>
    </div>
  </section>
</template>

<style scoped></style>
