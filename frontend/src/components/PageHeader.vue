<script setup lang="ts">
defineProps<{
  title: string
  subtitle?: string
}>()
</script>

<template>
  <div class="mb-8">
    <h1 class="text-3xl font-bold text-cupertino-gray-900 tracking-tight">{{ title }}</h1>
    <p v-if="subtitle" class="mt-1 text-base text-cupertino-gray-400">{{ subtitle }}</p>
  </div>
</template>
