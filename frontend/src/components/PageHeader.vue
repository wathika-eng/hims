<script setup lang="ts">
defineProps<{
  title: string
  subtitle?: string
}>()
</script>

<template>
  <div class="mb-6 sm:mb-8">
    <h1 class="text-2xl sm:text-3xl font-bold text-cupertino-gray-900 tracking-tight">{{ title }}</h1>
    <p v-if="subtitle" class="mt-0.5 sm:mt-1 text-sm sm:text-base text-cupertino-gray-500">{{ subtitle }}</p>
  </div>
</template>
