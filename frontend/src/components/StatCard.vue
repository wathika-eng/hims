<script setup lang="ts">
defineProps<{
  label: string
  value: string | number
  icon: string
  color?: string
}>()
</script>

<template>
  <div class="bg-white/70 backdrop-blur-2xl rounded-2xl border border-cupertino-gray-100/60 p-5 shadow-sm transition-all duration-200 hover:shadow-md hover:border-cupertino-gray-200/60">
    <div class="flex items-start justify-between">
      <div>
        <p class="text-sm font-medium text-cupertino-gray-500">{{ label }}</p>
        <p class="mt-1 text-3xl font-bold text-cupertino-gray-900 tracking-tight">{{ value }}</p>
      </div>
      <div
        class="w-10 h-10 rounded-xl flex items-center justify-center shrink-0"
        :style="{ backgroundColor: (color || '#007aff') + '15', color: color || '#007aff' }"
      >
        <svg viewBox="0 0 24 24" class="w-5 h-5" fill="currentColor">
          <path :d="icon" />
        </svg>
      </div>
    </div>
  </div>
</template>
