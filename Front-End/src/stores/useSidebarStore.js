// stores/useSidebarStore.js
import { defineStore } from "pinia";
import { ref, computed } from "vue";

export const useSidebarStore = defineStore("sidebar", () => {
  // ─── State ───────────────────────────────────────────────────────────────
  const isOpen = ref(true);
  const activeRoute = ref("dashboard");
  const windowWidth = ref(
    typeof window !== "undefined" ? window.innerWidth : 1280,
  );

  // ─── Getters ──────────────────────────────────────────────────────────────
  const isMobile = computed(() => windowWidth.value < 768);

  // Nilai ini dipakai di App.vue / layout untuk offset main content
  // Contoh: :class="sidebarContentOffset"
  const contentOffset = computed(() => {
    if (isMobile.value) return "ml-0";
    return isOpen.value ? "ml-64" : "ml-[68px]";
  });

  // ─── Actions ──────────────────────────────────────────────────────────────
  function toggleSidebar() {
    isOpen.value = !isOpen.value;
  }

  function openSidebar() {
    isOpen.value = true;
  }

  function closeSidebar() {
    isOpen.value = false;
  }

  function setRoute(route) {
    activeRoute.value = route;
    // Auto-close on mobile setelah navigasi
    if (isMobile.value) closeSidebar();
  }

  function syncWindowWidth(width) {
    windowWidth.value = width;
    // Auto-collapse saat masuk mobile, auto-expand saat kembali desktop
    if (width < 768) {
      isOpen.value = false;
    } else {
      isOpen.value = true;
    }
  }

  return {
    // state
    isOpen,
    activeRoute,
    windowWidth,
    // getters
    isMobile,
    contentOffset,
    // actions
    toggleSidebar,
    openSidebar,
    closeSidebar,
    setRoute,
    syncWindowWidth,
  };
});
