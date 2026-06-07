<script setup lang="ts">
import { ref, onMounted } from 'vue'

const isDemo = ref(false)
const unverified = ref(false)
const dismissed = ref(false)

onMounted(async () => {
  try {
    const token = localStorage.getItem('token')
    if (!token) return
    const payload = JSON.parse(atob(token.split('.')[1]))
    if (payload.demo) {
      isDemo.value = true
    }
    if (payload.aud === '' || payload.aud === 'patient') {
      unverified.value = true
    }
  } catch {
  }
})

function dismiss() {
  dismissed.value = true
}
</script>

<template>
  <div v-if="!dismissed" class="mx-4 sm:mx-6 lg:mx-8 mt-4 sm:mt-6 mb-0 flex flex-col gap-2">
    <!-- Demo mode banner -->
    <div v-if="isDemo" class="flex items-start gap-3 px-4 py-3 rounded-2xl bg-cupertino-orange/15 border-l-4 border-cupertino-orange shadow-sm">
      <svg viewBox="0 0 24 24" class="w-5 h-5 shrink-0 mt-0.5 text-cupertino-orange" fill="currentColor">
        <path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm-2 15l-5-5 1.41-1.41L10 14.17l7.59-7.59L19 8l-9 9z"/>
      </svg>
      <div class="flex-1 min-w-0">
        <p class="text-sm font-semibold text-cupertino-orange">Demo Environment</p>
        <p class="text-xs text-cupertino-orange/90 mt-0.5">All data shown is simulated test data. Names, IDs, and phone numbers are masked — nothing is real.</p>
      </div>
      <button @click="dismiss" aria-label="Dismiss banner" class="shrink-0 p-1 rounded-lg text-cupertino-orange/50 hover:text-cupertino-orange hover:bg-cupertino-orange/10 transition-colors">
        <svg viewBox="0 0 24 24" class="w-4 h-4" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round"><line x1="18" y1="6" x2="6" y2="18"/><line x1="6" y1="6" x2="18" y2="18"/></svg>
      </button>
    </div>

    <!-- Unverified account banner -->
    <div v-if="unverified" class="flex items-start gap-3 px-4 py-3 rounded-2xl bg-cupertino-blue/10 border border-cupertino-blue/20">
      <svg viewBox="0 0 24 24" class="w-5 h-5 shrink-0 mt-0.5 text-cupertino-blue" fill="currentColor">
        <path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm-2 15l-5-5 1.41-1.41L10 14.17l7.59-7.59L19 8l-9 9z"/>
      </svg>
      <div class="flex-1 min-w-0">
        <p class="text-sm font-semibold text-cupertino-blue">Account not verified</p>
        <p class="text-xs text-cupertino-blue/80 mt-0.5">Please verify your email address to access all features. Check your inbox for a verification link.</p>
      </div>
    </div>
  </div>
</template>
