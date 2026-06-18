<script setup>
import Navbar from "@/components/welcome-page/Navbar.vue";
import HeroSection from "@/components/welcome-page/HeroSection.vue";
import FeatureSection from "@/components/welcome-page/FeatureSection.vue";
import RulesSection from "@/components/welcome-page/RulesSection.vue";
import JournalSection from "@/components/welcome-page/JournalSection.vue";
import AnalyticsSection from "@/components/welcome-page/AnalyticsSection.vue";
import TemplateSection from "@/components/welcome-page/TemplateSection.vue";
import CTASection from "@/components/welcome-page/CTASection.vue";
import FooterSection from "@/components/welcome-page/FooterSection.vue";
</script>

<template>
  <div class="body-welcome">
    <Navbar />

    <HeroSection />

    <FeatureSection />

    <RulesSection />

    <JournalSection />

    <AnalyticsSection />

    <TemplateSection />

    <CTASection />

    <FooterSection />
  </div>
</template>

<style scoped>
* {
  box-sizing: border-box;
}

.body-welcome {
  background: #f6f6f6;
  color: #0f0f10;
  font-family: "Inter", sans-serif;
  -webkit-font-smoothing: antialiased;
}
</style>
