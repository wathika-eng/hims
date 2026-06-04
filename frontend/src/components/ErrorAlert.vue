<script setup lang="ts">
defineProps<{
  message: string | null
  title?: string
}>()
</script>

<template>
  <div v-if="message" class="flex items-start gap-3 px-4 py-3 rounded-2xl bg-cupertino-red/5 border border-cupertino-red/15 text-sm">
    <svg viewBox="0 0 24 24" class="w-5 h-5 shrink-0 mt-0.5 text-cupertino-red" fill="currentColor">
      <path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm1 15h-2v-2h2v2zm0-4h-2V7h2v6z"/>
    </svg>
    <div>
      <p v-if="title" class="font-semibold text-cupertino-red mb-0.5">{{ title }}</p>
      <p class="text-cupertino-red/80">{{ message }}</p>
    </div>
  </div>
</template>
