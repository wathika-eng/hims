<script setup lang="ts">
import { useRoute } from 'vue-router'
import { useAuthStore } from '../stores/auth'

const route = useRoute()
const auth = useAuthStore()

const links = [
  { to: '/', label: 'Dashboard', icon: 'grid' },
  { to: '/patients', label: 'Patients', icon: 'person.2' },
  { to: '/patients/new', label: 'New Patient', icon: 'person.badge.plus' },
  { to: '/programs', label: 'Programs', icon: 'square.on.square' },
  { to: '/programs/new', label: 'New Program', icon: 'plus.square' },
  { to: '/enroll', label: 'Enroll Patient', icon: 'link' },
]

function isActive(path: string) {
  if (path === '/') return route.path === '/'
  return route.path.startsWith(path)
}

function iconSvg(name: string) {
  const icons: Record<string, string> = {
    grid: 'M3 9h6V3H3v6zm0 12h6v-6H3v6zm9-12h6V3h-6v6zm0 12h6v-6h-6v6zM21 3v6h-6V3h6zm0 12h-6v6h6v-6z',
    'person.2': 'M12 12c2.21 0 4-1.79 4-4s-1.79-4-4-4-4 1.79-4 4 1.79 4 4 4zm-8 6c0-2.67 5.33-4 8-4s8 1.33 8 4v2H4v-2z',
    'person.badge.plus': 'M12 12c2.21 0 4-1.79 4-4s-1.79-4-4-4-4 1.79-4 4 1.79 4 4 4zm-4 5h3v3h2v-3h3v-2h-3v-3h-2v3H8v2z',
    'square.on.square': 'M4 4h12v12H4V4zm14 6h2v8H8v-2h10v-6z',
    'plus.square': 'M6 3h12a3 3 0 013 3v12a3 3 0 01-3 3H6a3 3 0 01-3-3V6a3 3 0 013-3zm6 4v4h4v2h-4v4h-2v-4H6v-2h4V7h2z',
    link: 'M7 17c-1.1 0-2.04-.4-2.82-1.18C3.4 15.04 3 14.1 3 13s.4-2.04 1.18-2.82C5.04 9.4 5.97 9 7 9h4v2H7c-.55 0-1.02.2-1.41.59S5 12.45 5 13s.2 1.02.59 1.41S6.45 15 7 15h4v2H7zm4-4v-2h6v2h-6zm3 4v-2h3c.55 0 1.02-.2 1.41-.59S19 13.55 19 13s-.2-1.02-.59-1.41S18.55 11 18 11h-3V9h3c1.1 0 2.04.4 2.82 1.18C20.6 10.96 21 11.9 21 13s-.4 2.04-1.18 2.82C19.04 16.6 18.1 17 17 17h-3z',
  }
  return icons[name] || ''
}
</script>

<template>
  <aside class="w-64 min-h-screen bg-white/80 backdrop-blur-2xl border-r border-cupertino-gray-100/50 flex flex-col">
    <div class="px-5 pt-6 pb-4">
      <div class="flex items-center gap-3">
        <div class="w-9 h-9 rounded-xl bg-cupertino-blue flex items-center justify-center shadow-sm">
          <svg viewBox="0 0 24 24" class="w-5 h-5 text-white" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round">
            <path d="M12 2L2 7l10 5 10-5-10-5zM2 17l10 5 10-5M2 12l10 5 10-5"/>
          </svg>
        </div>
        <div>
          <h1 class="text-sm font-semibold text-cupertino-gray-900 leading-tight">HIMS</h1>
          <p class="text-[11px] text-cupertino-gray-400 leading-tight">Health Information</p>
        </div>
      </div>
    </div>

    <nav class="flex-1 px-3 space-y-0.5">
      <router-link
        v-for="link in links"
        :key="link.to"
        :to="link.to"
        :class="[
          'flex items-center gap-3 px-3 py-2 rounded-xl text-sm font-medium transition-all duration-200',
          isActive(link.to)
            ? 'bg-cupertino-blue/10 text-cupertino-blue'
            : 'text-cupertino-gray-500 hover:text-cupertino-gray-700 hover:bg-cupertino-gray-50',
        ]"
      >
        <svg viewBox="0 0 24 24" class="w-5 h-5 shrink-0" fill="currentColor">
          <path :d="iconSvg(link.icon)" />
        </svg>
        <span>{{ link.label }}</span>
      </router-link>
    </nav>

    <div class="px-3 pb-4 pt-3 border-t border-cupertino-gray-100/50">
      <button
        @click="auth.logout()"
        class="flex items-center gap-3 px-3 py-2 rounded-xl text-sm font-medium text-cupertino-red/80 hover:bg-cupertino-red/5 w-full transition-colors"
      >
        <svg viewBox="0 0 24 24" class="w-5 h-5" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <path d="M9 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h4" />
          <polyline points="16 17 21 12 16 7" />
          <line x1="21" y1="12" x2="9" y2="12" />
        </svg>
        <span>Sign Out</span>
      </button>
    </div>
  </aside>
</template>
