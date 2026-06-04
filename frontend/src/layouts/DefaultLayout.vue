<script setup lang="ts">
import AppSidebar from '../components/AppSidebar.vue'
import VerificationBanner from '../components/VerificationBanner.vue'
</script>

<template>
  <div class="flex min-h-screen bg-cupertino-gray-50">
    <AppSidebar />
    <main class="flex-1 flex flex-col overflow-auto">
      <VerificationBanner />
      <div class="flex-1 p-8">
        <router-view />
      </div>
    </main>
  </div>
</template>
