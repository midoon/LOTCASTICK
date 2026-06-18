<script setup>
import analyticsStats from "@/data/welcome-page/analyticsStats.js";

import cd from "@/data/welcome-page/calenderDays.js";

const calendarDays = cd();
</script>

<template>
  <section id="analytics" class="py-20 bg-neutral-100/60">
    <div class="max-w-6xl mx-auto px-6">
      <div class="text-center mb-14">
        <div
          class="text-[12px] font-semibold text-primary-500 uppercase tracking-[0.12em] mb-3"
        >
          Analytics
        </div>
        <h2
          class="text-[2.2rem] font-extrabold text-primary-900 tracking-tight"
        >
          8 dimensions of performance
        </h2>
      </div>

      <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
        <div
          v-for="stat in analyticsStats"
          :key="stat.label"
          class="bg-white rounded-2xl border border-neutral-200 p-5 feature-card"
        >
          <div
            class="text-[26px] font-extrabold font-mono mb-1"
            :class="stat.color"
          >
            {{ stat.value }}
          </div>
          <div class="text-[12px] font-semibold text-neutral-700 mb-1">
            {{ stat.label }}
          </div>
          <div class="text-[11.5px] text-neutral-400 leading-snug">
            {{ stat.sub }}
          </div>
        </div>
      </div>

      <!-- Calendar heatmap visual -->
      <div class="mt-8 bg-white rounded-2xl border border-neutral-200 p-6">
        <div
          class="text-[12px] font-semibold text-neutral-500 uppercase tracking-wider mb-4"
        >
          Daily P&amp;L Calendar — June 2026
        </div>
        <div class="grid grid-cols-7 gap-1.5">
          <div
            v-for="d in calendarDays"
            :key="d.day"
            class="aspect-square rounded-lg flex items-center justify-center text-[10px] font-mono font-medium transition-transform hover:scale-110 cursor-default"
            :class="d.cls"
            :title="d.label"
          >
            {{ d.day || "" }}
          </div>
        </div>
        <div class="flex items-center gap-3 mt-4">
          <span class="text-[11px] text-neutral-400">Loss</span>
          <div class="flex gap-1">
            <div class="w-4 h-2 rounded bg-red-400"></div>
            <div class="w-4 h-2 rounded bg-red-200"></div>
            <div class="w-4 h-2 rounded bg-neutral-200"></div>
            <div class="w-4 h-2 rounded bg-success-200"></div>
            <div class="w-4 h-2 rounded bg-success-400"></div>
          </div>
          <span class="text-[11px] text-neutral-400">Gain</span>
        </div>
      </div>
    </div>
  </section>
</template>

<style scoped>
/* Feature card hover */
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
