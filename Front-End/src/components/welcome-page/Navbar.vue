<script setup></script>

<template>
  <div>
    <nav
      class="fixed top-0 left-0 right-0 z-50 nav-blur border-b border-neutral-200/60"
    >
      <div
        class="max-w-6xl mx-auto px-6 h-16 flex items-center justify-between"
      >
        <!-- Logo -->
        <div class="flex items-center gap-2.5">
          <div
            class="w-8 h-8 rounded-lg bg-primary-800 flex items-center justify-center"
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
          <span
            class="font-semibold text-primary-900 tracking-tight text-[15px]"
            >Lot<span class="text-primary-500">Castick</span></span
          >
        </div>

        <!-- Nav links (desktop) -->
        <div
          class="hidden md:flex items-center gap-8 text-[13.5px] font-medium text-neutral-500"
        >
          <a href="#features" class="hover:text-primary-800 transition-colors"
            >Features</a
          >
          <a href="#rules" class="hover:text-primary-800 transition-colors"
            >Challenge Rules</a
          >
          <a href="#analytics" class="hover:text-primary-800 transition-colors"
            >Analytics</a
          >
          <a href="#templates" class="hover:text-primary-800 transition-colors"
            >Templates</a
          >
        </div>

        <!-- CTA -->
        <div class="flex items-center gap-2.5">
          <button
            class="hidden md:block text-[13.5px] font-medium text-neutral-700 hover:text-primary-800 transition-colors px-4 py-2"
          >
            Log In
          </button>
          <button
            class="text-[13.5px] font-semibold bg-primary-800 text-neutral-50 px-4 py-2 rounded-lg hover:bg-primary-700 transition-colors"
          >
            Get Started Free
          </button>
        </div>
      </div>
    </nav>
  </div>
</template>

<style scoped>
.nav-blur {
  backdrop-filter: blur(16px);
  -webkit-backdrop-filter: blur(16px);
  background: rgba(246, 246, 246, 0.88);
}
</style>
