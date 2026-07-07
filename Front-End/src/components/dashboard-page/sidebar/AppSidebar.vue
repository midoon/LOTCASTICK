<script setup>
import { onMounted, onUnmounted, defineComponent, h } from "vue";
import { useSidebarStore } from "@/stores/useSidebarStore";

const sidebar = useSidebarStore();

// ─── Sync window width ke store ───────────────────────────────────────────
function onResize() {
  sidebar.syncWindowWidth(window.innerWidth);
}

onMounted(() => {
  window.addEventListener("resize", onResize);
  sidebar.syncWindowWidth(window.innerWidth);
});
onUnmounted(() => {
  window.removeEventListener("resize", onResize);
});

// ─── NavSection ───────────────────────────────────────────────────────────
const NavSection = defineComponent({
  props: { label: String, collapsed: Boolean },
  setup(props, { slots }) {
    return () =>
      h("div", { class: "mb-1" }, [
        !props.collapsed
          ? h(
              "p",
              {
                class:
                  "px-4 pb-1 pt-2 text-[9px] font-bold uppercase tracking-widest text-primary-500 select-none",
              },
              props.label,
            )
          : h("div", { class: "mx-4 my-2 h-px bg-primary-800/60" }),
        slots.default?.(),
      ]);
  },
});

// ─── NavItem ──────────────────────────────────────────────────────────────
const NavItem = defineComponent({
  props: {
    icon: String,
    label: String,
    active: Boolean,
    collapsed: Boolean,
    badge: Number,
  },
  emits: ["click"],
  setup(props, { emit }) {
    const icons = {
      dashboard: `<path d="M2 2h5v5H2V2zm7 0h5v5H9V2zM2 9h5v5H2V9zm7 3h1v-2h2v-1H9v3zm2 2v-1h2v1h-2zm0-4h2v1h-2V10z" fill="currentColor"/>`,
      simulations: `<path d="M3 5a2 2 0 012-2h6a2 2 0 012 2v8a2 2 0 01-2 2H5a2 2 0 01-2-2V5zm2 0v8h6V5H5zm2 2h2v4H7V7z" fill="currentColor"/>`,
      trades: `<path d="M2 12l4-4 3 3 5-6" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round" fill="none"/><circle cx="13" cy="5" r="1.5" fill="currentColor"/>`,
      calendar: `<rect x="2" y="3" width="12" height="11" rx="1.5" stroke="currentColor" stroke-width="1.4" fill="none"/><path d="M2 7h12M6 2v2M10 2v2" stroke="currentColor" stroke-width="1.4" stroke-linecap="round"/>`,
      analytics: `<path d="M2 13V9m3 4V7m3 6V5m3 8V3" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" fill="none"/>`,
      reports: `<rect x="3" y="2" width="10" height="12" rx="1.5" stroke="currentColor" stroke-width="1.4" fill="none"/><path d="M6 6h4M6 9h4M6 12h2" stroke="currentColor" stroke-width="1.3" stroke-linecap="round"/>`,
      strategies: `<path d="M8 2L2 6v6l6 3 6-3V6L8 2z" stroke="currentColor" stroke-width="1.4" fill="none" stroke-linejoin="round"/><path d="M8 5v6M5 7l3-2 3 2" stroke="currentColor" stroke-width="1.2" stroke-linecap="round"/>`,
      tags: `<path d="M9 2H4a1 1 0 00-1 1v5l6 6 5-5-6-7z" stroke="currentColor" stroke-width="1.4" fill="none" stroke-linejoin="round"/><circle cx="6" cy="6" r="1" fill="currentColor"/>`,
      import: `<path d="M8 2v8M5 7l3 3 3-3" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round" fill="none"/><path d="M3 12v1a1 1 0 001 1h8a1 1 0 001-1v-1" stroke="currentColor" stroke-width="1.4" stroke-linecap="round" fill="none"/>`,
      templates: `<path d="M2 3h12v3H2V3zm0 5h7v6H2V8zm9 0h3v2h-3V8zm0 4h3v2h-3v-2z" fill="currentColor"/>`,
      settings: `<path d="M8 5a3 3 0 100 6 3 3 0 000-6z" stroke="currentColor" stroke-width="1.3" fill="none"/><path d="M8 1v2M8 13v2M1 8h2M13 8h2M3.05 3.05l1.41 1.41M11.54 11.54l1.41 1.41M3.05 12.95l1.41-1.41M11.54 4.46l1.41-1.41" stroke="currentColor" stroke-width="1.3" stroke-linecap="round" fill="none"/>`,
      profile: `<circle cx="8" cy="6" r="3" stroke="currentColor" stroke-width="1.4" fill="none"/><path d="M2 14c0-3.314 2.686-5 6-5s6 1.686 6 5" stroke="currentColor" stroke-width="1.4" stroke-linecap="round" fill="none"/>`,
    };

    return () =>
      h(
        "button",
        {
          class: [
            "group w-full flex items-center gap-3 px-3 py-2 mx-1 rounded-lg",
            "transition-all duration-150 text-left relative",
            props.collapsed ? "justify-center" : "",
            props.active
              ? "bg-primary-700/80 text-neutral-50 shadow-sm"
              : "text-primary-300 hover:bg-primary-800/60 hover:text-neutral-200",
          ].join(" "),
          style: "width: calc(100% - 8px)",
          title: props.collapsed ? props.label : undefined,
          onClick: () => emit("click"),
        },
        [
          h(
            "span",
            {
              class: [
                "shrink-0 w-8 h-8 flex items-center justify-center rounded-md transition-colors",
                props.active
                  ? "bg-primary-600/60 text-neutral-50"
                  : "text-primary-400 group-hover:text-neutral-200",
              ].join(" "),
            },
            [
              h("svg", {
                width: 16,
                height: 16,
                viewBox: "0 0 16 16",
                innerHTML: icons[props.icon] ?? "",
              }),
            ],
          ),

          !props.collapsed
            ? h(
                "span",
                { class: "flex-1 flex items-center justify-between min-w-0" },
                [
                  h(
                    "span",
                    { class: "text-sm font-medium truncate leading-none" },
                    props.label,
                  ),
                  props.badge
                    ? h(
                        "span",
                        {
                          class:
                            "ml-2 shrink-0 text-[10px] font-bold px-1.5 py-0.5 rounded-full bg-primary-500/50 text-primary-200",
                        },
                        String(props.badge),
                      )
                    : null,
                ],
              )
            : null,

          props.collapsed && props.active
            ? h("span", {
                class:
                  "absolute right-0.5 top-1/2 -translate-y-1/2 w-0.5 h-4 rounded-full bg-primary-400",
              })
            : null,
        ],
      );
  },
});
</script>

<template>
  <!-- Mobile Overlay -->
  <Transition name="fade">
    <div
      v-if="sidebar.isOpen && sidebar.isMobile"
      class="fixed inset-0 bg-primary-950/60 backdrop-blur-sm z-40"
      @click="sidebar.closeSidebar"
    />
  </Transition>

  <!-- Sidebar -->
  <aside
    :class="[
      'fixed top-0 left-0 h-screen z-50 flex flex-col',
      'bg-primary-900 border-r border-primary-800/60',
      'transition-all duration-300 ease-in-out',
      sidebar.isOpen ? 'w-64' : 'w-17',
      sidebar.isMobile && !sidebar.isOpen
        ? '-translate-x-full w-64'
        : 'translate-x-0',
    ]"
  >
    <!-- Header: Logo + Toggle -->
    <div
      class="flex items-center h-16 px-4 border-b border-primary-800/60 shrink-0"
    >
      <div class="flex items-center gap-3 min-w-0 flex-1">
        <!-- Logo Icon -->
        <div
          class="shrink-0 w-8 h-8 rounded-lg bg-primary-500 flex items-center justify-center shadow-lg"
        >
          <svg
            width="18"
            height="18"
            viewBox="0 0 18 18"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              d="M2 13L6.5 8L9.5 11L13 6L16 9"
              stroke="white"
              stroke-width="1.8"
              stroke-linecap="round"
              stroke-linejoin="round"
            />
            <circle cx="16" cy="9" r="1.5" fill="#6fecba" />
          </svg>
        </div>

        <!-- Logo Text -->
        <Transition name="label">
          <div
            v-if="sidebar.isOpen"
            class="flex flex-col leading-none overflow-hidden"
          >
            <span
              class="text-neutral-50 font-semibold text-sm tracking-tight whitespace-nowrap"
              >Lotcastick</span
            >
            <span
              class="text-primary-400 text-[10px] tracking-widest uppercase whitespace-nowrap"
              >Prop Simulator</span
            >
          </div>
        </Transition>
      </div>

      <!-- Toggle Button (desktop only) -->
      <button
        v-if="!sidebar.isMobile"
        @click="sidebar.toggleSidebar"
        class="shrink-0 w-7 h-7 rounded-md flex items-center justify-center text-primary-400 hover:text-neutral-50 hover:bg-primary-800 transition-colors"
        :title="sidebar.isOpen ? 'Collapse sidebar' : 'Expand sidebar'"
      >
        <svg
          width="14"
          height="14"
          viewBox="0 0 14 14"
          fill="none"
          :class="[
            'transition-transform duration-300',
            sidebar.isOpen ? '' : 'rotate-180',
          ]"
        >
          <path
            d="M9 11L5 7L9 3"
            stroke="currentColor"
            stroke-width="1.6"
            stroke-linecap="round"
            stroke-linejoin="round"
          />
        </svg>
      </button>
    </div>

    <!-- Scrollable Nav Area -->
    <nav
      class="flex-1 overflow-y-auto overflow-x-hidden py-4 space-y-1 scrollbar-hide"
    >
      <!-- MAIN -->
      <NavSection label="Main" :collapsed="!sidebar.isOpen">
        <NavItem
          :collapsed="!sidebar.isOpen"
          icon="dashboard"
          label="Dashboard"
          :active="sidebar.activeRoute === 'dashboard'"
          @click="sidebar.setRoute('dashboard')"
        />
        <NavItem
          :collapsed="!sidebar.isOpen"
          icon="simulations"
          label="Simulations"
          :active="sidebar.activeRoute === 'simulations'"
          :badge="3"
          @click="sidebar.setRoute('simulations')"
        />
        <NavItem
          :collapsed="!sidebar.isOpen"
          icon="trades"
          label="Trades"
          :active="sidebar.activeRoute === 'trades'"
          @click="sidebar.setRoute('trades')"
        />
        <NavItem
          :collapsed="!sidebar.isOpen"
          icon="calendar"
          label="Calendar"
          :active="sidebar.activeRoute === 'calendar'"
          @click="sidebar.setRoute('calendar')"
        />
        <NavItem
          :collapsed="!sidebar.isOpen"
          icon="analytics"
          label="Analytics"
          :active="sidebar.activeRoute === 'analytics'"
          @click="sidebar.setRoute('analytics')"
        />
        <NavItem
          :collapsed="!sidebar.isOpen"
          icon="reports"
          label="Reports"
          :active="sidebar.activeRoute === 'reports'"
          @click="sidebar.setRoute('reports')"
        />
      </NavSection>

      <!-- TOOLS -->
      <NavSection label="Tools" :collapsed="!sidebar.isOpen">
        <NavItem
          :collapsed="!sidebar.isOpen"
          icon="strategies"
          label="Strategies"
          :active="sidebar.activeRoute === 'strategies'"
          @click="sidebar.setRoute('strategies')"
        />
        <NavItem
          :collapsed="!sidebar.isOpen"
          icon="tags"
          label="Tags"
          :active="sidebar.activeRoute === 'tags'"
          @click="sidebar.setRoute('tags')"
        />
        <NavItem
          :collapsed="!sidebar.isOpen"
          icon="import"
          label="Import Trades"
          :active="sidebar.activeRoute === 'import'"
          @click="sidebar.setRoute('import')"
        />
        <NavItem
          :collapsed="!sidebar.isOpen"
          icon="templates"
          label="Templates"
          :active="sidebar.activeRoute === 'templates'"
          @click="sidebar.setRoute('templates')"
        />
      </NavSection>

      <!-- SETTINGS -->
      <NavSection label="Settings" :collapsed="!sidebar.isOpen">
        <NavItem
          :collapsed="!sidebar.isOpen"
          icon="settings"
          label="Settings"
          :active="sidebar.activeRoute === 'settings'"
          @click="sidebar.setRoute('settings')"
        />
        <NavItem
          :collapsed="!sidebar.isOpen"
          icon="profile"
          label="Profile"
          :active="sidebar.activeRoute === 'profile'"
          @click="sidebar.setRoute('profile')"
        />
      </NavSection>
    </nav>

    <!-- Active Simulation Card -->
    <Transition name="sim-card">
      <div
        v-if="sidebar.isOpen"
        class="mx-3 mb-3 rounded-xl bg-primary-800/70 border border-primary-700/50 p-3 shrink-0"
      >
        <div class="flex items-center gap-2 mb-2">
          <div
            class="w-6 h-6 rounded-md bg-gold-500/20 flex items-center justify-center"
          >
            <svg width="12" height="12" viewBox="0 0 12 12" fill="none">
              <path
                d="M6 1L7.5 4.5H11L8 6.5L9 10L6 8L3 10L4 6.5L1 4.5H4.5L6 1Z"
                fill="#d9ab1c"
              />
            </svg>
          </div>
          <div class="flex-1 min-w-0">
            <p
              class="text-neutral-50 text-[11px] font-semibold truncate leading-tight"
            >
              FTMO Challenge
            </p>
            <p class="text-primary-400 text-[10px] leading-tight">Phase 1</p>
          </div>
          <span
            class="shrink-0 text-[9px] font-bold px-1.5 py-0.5 rounded-full bg-success-500/15 text-success-400 uppercase tracking-wide"
            >Active</span
          >
        </div>

        <div class="space-y-1.5">
          <div>
            <div class="flex justify-between items-center mb-0.5">
              <span class="text-primary-400 text-[9px] uppercase tracking-wide"
                >Profit</span
              >
              <span class="text-neutral-300 text-[9px] font-medium">44.2%</span>
            </div>
            <div class="h-1 bg-primary-700/60 rounded-full overflow-hidden">
              <div
                class="h-full bg-success-500 rounded-full"
                style="width: 44.2%"
              ></div>
            </div>
          </div>
          <div>
            <div class="flex justify-between items-center mb-0.5">
              <span class="text-primary-400 text-[9px] uppercase tracking-wide"
                >Drawdown</span
              >
              <span class="text-neutral-300 text-[9px] font-medium">44.2%</span>
            </div>
            <div class="h-1 bg-primary-700/60 rounded-full overflow-hidden">
              <div
                class="h-full bg-gold-500 rounded-full"
                style="width: 44.2%"
              ></div>
            </div>
          </div>
        </div>

        <div
          class="mt-2 pt-2 border-t border-primary-700/40 flex items-center justify-between"
        >
          <span class="text-neutral-400 text-[10px]">$44,200 / $100,000</span>
          <div class="pulse-dot w-1.5 h-1.5 rounded-full bg-success-400"></div>
        </div>
      </div>
    </Transition>

    <!-- Collapsed: mini avatar -->
    <div v-if="!sidebar.isOpen" class="mb-3 mx-auto shrink-0">
      <div
        class="w-8 h-8 rounded-full bg-primary-600 border border-primary-500/50 flex items-center justify-center text-neutral-200 text-xs font-bold"
      >
        JD
      </div>
    </div>
  </aside>

  <!-- Mobile Hamburger (shown when sidebar closed on mobile) -->
  <button
    v-if="sidebar.isMobile"
    @click="sidebar.toggleSidebar"
    :class="[
      'fixed z-50 top-4 left-4',
      'w-9 h-9 rounded-lg flex items-center justify-center',
      'bg-primary-800 border border-primary-700/60 text-neutral-300',
      'shadow-lg shadow-primary-950/40',
      'transition-all duration-300',
      sidebar.isOpen ? 'opacity-0 pointer-events-none' : 'opacity-100',
    ]"
    aria-label="Open sidebar"
  >
    <svg width="16" height="16" viewBox="0 0 16 16" fill="none">
      <path
        d="M2 4h12M2 8h12M2 12h12"
        stroke="currentColor"
        stroke-width="1.5"
        stroke-linecap="round"
      />
    </svg>
  </button>
</template>

<style scoped>
.scrollbar-hide {
  scrollbar-width: none;
  -ms-overflow-style: none;
}
.scrollbar-hide::-webkit-scrollbar {
  display: none;
}

.label-enter-active,
.label-leave-active {
  transition:
    opacity 0.2s ease,
    transform 0.2s ease;
}
.label-enter-from,
.label-leave-to {
  opacity: 0;
  transform: translateX(-6px);
}

.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.25s ease;
}
.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}

.sim-card-enter-active,
.sim-card-leave-active {
  transition:
    opacity 0.2s ease,
    transform 0.2s ease;
}
.sim-card-enter-from,
.sim-card-leave-to {
  opacity: 0;
  transform: translateY(8px);
}
</style>
