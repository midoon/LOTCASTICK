<template>
  <div class="flex min-h-screen bg-neutral-50">
    <AppSidebar />
    <main
      :class="[
        'flex-1 min-h-screen transition-all duration-300',
        sidebar.contentOffset,
      ]"
    >
      <RouterView />
    </main>
  </div>
</template>

<script setup>
import AppSidebar from "@/components/dashboard-page/sidebar/AppSidebar.vue";
import { useSidebarStore } from "@/stores/useSidebarStore";

const sidebar = useSidebarStore();
</script>
