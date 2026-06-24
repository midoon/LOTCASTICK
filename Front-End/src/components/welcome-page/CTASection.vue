<script setup></script>

<template>
  <section class="py-16 max-w-6xl mx-auto px-6">
    <div
      class="bg-primary-800 rounded-3xl p-10 md:p-14 relative overflow-hidden"
    >
      <div
        class="absolute -top-20 -right-20 w-72 h-72 rounded-full bg-primary-600/30 blur-3xl pointer-events-none"
      ></div>
      <div class="relative z-10 text-center">
        <h2
          class="text-[2rem] md:text-[2.6rem] font-extrabold text-neutral-50 tracking-tight mb-4"
        >
          Ready to pass your<br />next challenge?
        </h2>
        <p class="text-primary-300 text-[15px] mb-8 max-w-md mx-auto">
          Simulate before you trade. Build discipline. Understand your rules
          deeply. Then go live with confidence.
        </p>
        <RouterLink
          to="/register"
          class="bg-neutral-50 text-primary-900 text-[14px] font-bold px-8 py-3.5 rounded-xl hover:bg-white transition-colors shadow-lg"
        >
          Create Your First Simulation →
        </RouterLink>
      </div>
    </div>
  </section>
</template>
