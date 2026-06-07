<script setup lang="ts">
import { ref } from 'vue'
import { useRoute } from 'vue-router'
import AppSidebar from '../components/AppSidebar.vue'
import VerificationBanner from '../components/VerificationBanner.vue'
import { useAuthStore } from '../stores/auth'

const route = useRoute()
const auth = useAuthStore()
const sidebarOpen = ref(false)

function closeSidebar() {
  sidebarOpen.value = false
}

const bottomLinks = [
  { to: '/', label: 'Home', icon: 'M3 9l9-7 9 7v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V9z', iconActive: 'M3 9l9-7 9 7v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V9z M9 22V12h6v10' },
  { to: '/patients', label: 'Patients', icon: 'M12 12c2.21 0 4-1.79 4-4s-1.79-4-4-4-4 1.79-4 4 1.79 4 4 4zm-8 6c0-2.67 5.33-4 8-4s8 1.33 8 4v2H4v-2z' },
  { to: '/programs', label: 'Programs', icon: 'M4 4h12v12H4V4zm14 6h2v8H8v-2h10v-6z' },
  { to: '/enroll', label: 'Enroll', icon: 'M7 17c-1.1 0-2.04-.4-2.82-1.18C3.4 15.04 3 14.1 3 13s.4-2.04 1.18-2.82C5.04 9.4 5.97 9 7 9h4v2H7c-.55 0-1.02.2-1.41.59S5 12.45 5 13s.2 1.02.59 1.41S6.45 15 7 15h4v2H7zm4-4v-2h6v2h-6zm3 4v-2h3c.55 0 1.02-.2 1.41-.59S19 13.55 19 13s-.2-1.02-.59-1.41S18.55 11 18 11h-3V9h3c1.1 0 2.04.4 2.82 1.18C20.6 10.96 21 11.9 21 13s-.4 2.04-1.18 2.82C19.04 16.6 18.1 17 17 17h-3z' },
]

function isActive(path: string) {
  if (path === '/') return route.path === '/'
  return route.path.startsWith(path)
}
</script>

<template>
  <div class="flex min-h-screen bg-cupertino-gray-50">
    <AppSidebar :open="sidebarOpen" @close="closeSidebar" />

    <main class="flex-1 flex flex-col overflow-auto pb-16 lg:pb-0">
      <!-- Mobile top bar -->
      <div class="sticky top-0 z-30 lg:hidden bg-white/80 backdrop-blur-2xl border-b border-cupertino-gray-100/50">
        <div class="flex items-center justify-between px-4 h-12">
          <button
            @click="sidebarOpen = true"
            class="p-1.5 -ml-1.5 rounded-lg text-cupertino-gray-500 hover:text-cupertino-gray-700 hover:bg-cupertino-gray-100 transition-colors"
            aria-label="Open menu"
          >
            <svg viewBox="0 0 24 24" class="w-5 h-5" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round">
              <line x1="3" y1="6" x2="21" y2="6"/><line x1="3" y1="12" x2="21" y2="12"/><line x1="3" y1="18" x2="21" y2="18"/>
            </svg>
          </button>
          <div class="flex items-center gap-2">
            <div class="w-6 h-6 rounded-lg bg-cupertino-blue flex items-center justify-center">
              <svg viewBox="0 0 24 24" class="w-3.5 h-3.5 text-white" fill="none" stroke="currentColor" stroke-width="3" stroke-linecap="round">
                <path d="M12 2L2 7l10 5 10-5-10-5zM2 17l10 5 10-5M2 12l10 5 10-5"/>
              </svg>
            </div>
            <span class="text-sm font-semibold text-cupertino-gray-900">HIMS</span>
            <span class="text-[10px] font-semibold px-1.5 py-[1px] rounded-md bg-cupertino-orange/10 text-cupertino-orange">DEMO</span>
          </div>
          <button
            @click="auth.logout()"
            class="p-1.5 -mr-1.5 rounded-lg text-cupertino-gray-400 hover:text-cupertino-red/80 hover:bg-cupertino-red/5 transition-colors"
            aria-label="Sign out"
          >
            <svg viewBox="0 0 24 24" class="w-5 h-5" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <path d="M9 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h4" />
              <polyline points="16 17 21 12 16 7" />
              <line x1="21" y1="12" x2="9" y2="12" />
            </svg>
          </button>
        </div>
      </div>

      <VerificationBanner />
      <div class="flex-1 p-4 sm:p-6 lg:p-8">
        <router-view />
      </div>
    </main>

    <!-- Mobile bottom navigation -->
    <nav class="fixed bottom-0 left-0 right-0 z-30 lg:hidden bg-white/90 backdrop-blur-2xl border-t border-cupertino-gray-100/50 safe-area-bottom">
      <div class="flex items-center justify-around px-2 py-1">
        <router-link
          v-for="link in bottomLinks"
          :key="link.to"
          :to="link.to"
          class="flex flex-col items-center gap-0.5 px-3 py-1.5 rounded-xl text-[10px] font-medium transition-colors min-w-0"
          :class="isActive(link.to) ? 'text-cupertino-blue' : 'text-cupertino-gray-400'"
        >
          <svg viewBox="0 0 24 24" class="w-5 h-5" :fill="isActive(link.to) ? 'currentColor' : 'none'" :stroke="isActive(link.to) ? 'none' : 'currentColor'" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round">
            <path :d="link.icon" />
          </svg>
          <span class="truncate max-w-full">{{ link.label }}</span>
        </router-link>
      </div>
    </nav>
  </div>
</template>

<style>
.safe-area-bottom {
  padding-bottom: env(safe-area-inset-bottom, 0px);
}
</style>
