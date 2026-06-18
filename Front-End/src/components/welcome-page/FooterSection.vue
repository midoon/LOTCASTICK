<script setup></script>

<template>
  <footer class="border-t border-neutral-200 py-8">
    <div
      class="max-w-6xl mx-auto px-6 flex flex-col md:flex-row items-center justify-between gap-4"
    >
      <div class="flex items-center gap-2">
        <div
          class="w-6 h-6 rounded-md bg-primary-800 flex items-center justify-center"
        >
          <svg width="11" height="11" viewBox="0 0 16 16" fill="none">
            <polyline
              points="1,12 5,7 8,9 12,4 15,6"
              stroke="#d1cede"
              stroke-width="2"
              stroke-linecap="round"
              stroke-linejoin="round"
              fill="none"
            />
            <circle cx="15" cy="6" r="2" fill="#1ccf88" />
          </svg>
        </div>
        <span class="text-[13px] font-semibold text-primary-900">TradeSim</span>
      </div>
      <div class="text-[12px] text-neutral-400">
        © 2026 LotCastisck · All financial calculations are for simulation
        purposes only
      </div>
      <div class="flex items-center gap-5 text-[12px] text-neutral-500">
        <a href="#" class="hover:text-primary-700 transition-colors">Privacy</a>
        <a href="#" class="hover:text-primary-700 transition-colors">Terms</a>
        <a href="#" class="hover:text-primary-700 transition-colors">Contact</a>
      </div>
    </div>
  </footer>
</template>
