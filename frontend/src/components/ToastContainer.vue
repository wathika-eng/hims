<script setup lang="ts">
import { useToast } from '../composables/useToast'

const { toasts } = useToast()

function icon(type: string) {
  if (type === 'success') return 'M9 16.17L4.83 12l-1.42 1.41L9 19 21 7l-1.41-1.41L9 16.17z'
  if (type === 'error') return 'M18 6L6 18M6 6l12 12'
  return 'M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm-2 15l-5-5 1.41-1.41L10 14.17l7.59-7.59L19 8l-9 9z'
}

function bg(type: string) {
  if (type === 'success') return 'bg-cupertino-green/10 border-cupertino-green/30 text-cupertino-green'
  if (type === 'error') return 'bg-cupertino-red/10 border-cupertino-red/30 text-cupertino-red'
  return 'bg-cupertino-blue/10 border-cupertino-blue/30 text-cupertino-blue'
}
</script>

<template>
  <div class="fixed top-4 right-4 z-50 flex flex-col gap-2 max-w-sm">
    <TransitionGroup name="toast">
      <div
        v-for="t in toasts"
        :key="t.id"
        :class="['flex items-center gap-3 px-4 py-3 rounded-2xl border shadow-lg backdrop-blur-2xl text-sm font-medium transition-all', bg(t.type)]"
      >
        <svg viewBox="0 0 24 24" class="w-5 h-5 shrink-0" fill="currentColor">
          <path :d="icon(t.type)" />
        </svg>
        <span>{{ t.message }}</span>
      </div>
    </TransitionGroup>
  </div>
</template>

<style scoped>
.toast-enter-active { transition: all 0.3s ease-out; }
.toast-leave-active { transition: all 0.25s ease-in; }
.toast-enter-from { opacity: 0; transform: translateX(30px) scale(0.95); }
.toast-leave-to { opacity: 0; transform: translateX(30px) scale(0.95); }
</style>
