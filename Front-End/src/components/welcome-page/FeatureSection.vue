<script setup>
import features from "@/data/welcome-page/features";
</script>

<template>
  <section id="features" class="py-20 max-w-6xl mx-auto px-6">
    <div class="text-center mb-14">
      <div
        class="text-[12px] font-semibold text-primary-500 uppercase tracking-[0.12em] mb-3"
      >
        Everything You Need
      </div>
      <h2
        class="text-[2.2rem] font-extrabold text-primary-900 tracking-tight leading-tight"
      >
        Built around how<br />prop challenges actually work
      </h2>
    </div>

    <div class="grid md:grid-cols-3 gap-5">
      <!-- Feature 1 -->
      <div
        class="feature-card bg-white rounded-2xl border border-neutral-200 p-6"
        v-for="f in features"
        :key="f.title"
      >
        <div
          class="w-10 h-10 rounded-xl flex items-center justify-center mb-4"
          :class="f.iconBg"
        >
          <span v-html="f.icon" class="text-lg"></span>
        </div>
        <h3 class="text-[15px] font-semibold text-primary-900 mb-2">
          {{ f.title }}
        </h3>
        <p class="text-[13.5px] text-neutral-500 leading-relaxed">
          {{ f.desc }}
        </p>
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
