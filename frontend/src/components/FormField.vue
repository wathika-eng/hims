<script setup lang="ts">
defineProps<{
  label: string
  error?: string | null
}>()
</script>

<template>
  <div>
    <label class="block text-xs font-semibold text-cupertino-gray-500 mb-1.5 tracking-wide uppercase">{{ label }}</label>
    <slot />
    <p v-if="error" class="mt-1 text-xs text-cupertino-red/80 flex items-center gap-1">
      <svg viewBox="0 0 24 24" class="w-3.5 h-3.5 shrink-0" fill="currentColor"><path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm1 15h-2v-2h2v2zm0-4h-2V7h2v6z"/></svg>
      {{ error }}
    </p>
  </div>
</template>
