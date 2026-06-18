<script setup>
import templates from "@/data/welcome-page/templates.js";
</script>

<template>
  <section id="templates" class="py-20 max-w-6xl mx-auto px-6">
    <div class="text-center mb-14">
      <div
        class="text-[12px] font-semibold text-primary-500 uppercase tracking-[0.12em] mb-3"
      >
        Challenge Templates
      </div>
      <h2 class="text-[2.2rem] font-extrabold text-primary-900 tracking-tight">
        Start with any prop firm's rules
      </h2>
      <p class="text-neutral-500 text-[15px] mt-3 max-w-md mx-auto">
        Pre-configured rule sets based on real prop firms. Or build your own
        from scratch.
      </p>
    </div>

    <div class="grid md:grid-cols-2 lg:grid-cols-4 gap-4">
      <div
        v-for="t in templates"
        :key="t.name"
        class="feature-card bg-white rounded-2xl border border-neutral-200 p-5"
      >
        <div class="flex items-center justify-between mb-4">
          <span class="text-[13px] font-bold text-primary-900">{{
            t.name
          }}</span>
          <span
            class="text-[10px] font-semibold text-primary-500 bg-primary-50 px-2 py-1 rounded-full"
            >{{ t.account }}</span
          >
        </div>
        <div class="space-y-2">
          <div class="flex justify-between text-[12px]">
            <span class="text-neutral-400">Max DD</span>
            <span class="font-mono font-semibold text-red-500">{{
              t.maxDD
            }}</span>
          </div>
          <div class="flex justify-between text-[12px]">
            <span class="text-neutral-400">Daily DD</span>
            <span class="font-mono font-semibold text-gold-600">{{
              t.dailyDD
            }}</span>
          </div>
          <div class="flex justify-between text-[12px]">
            <span class="text-neutral-400">Target</span>
            <span class="font-mono font-semibold text-success-600">{{
              t.target
            }}</span>
          </div>
          <div class="flex justify-between text-[12px]">
            <span class="text-neutral-400">Min Days</span>
            <span class="font-mono font-semibold text-primary-600">{{
              t.minDays
            }}</span>
          </div>
        </div>
        <button
          class="mt-4 w-full text-[12px] font-semibold text-primary-700 border border-primary-200 rounded-lg py-2 hover:bg-primary-50 transition-colors"
        >
          Use Template →
        </button>
      </div>
    </div>
  </section>
</template>

<style scope>
.feature-card {
  transition:
    box-shadow 0.2s,
    transform 0.2s;
}
.feature-card:hover {
  transform: translateY(-3px);
  box-shadow: 0 8px 32px 0 rgba(48, 45, 67, 0.1);
}
</style>
