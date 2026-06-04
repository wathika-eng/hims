<script setup lang="ts">
import ToastContainer from './components/ToastContainer.vue'
</script>

<template>
  <ToastContainer />
  <router-view v-slot="{ Component }">
    <Transition name="page" mode="out-in">
      <component :is="Component" />
    </Transition>
  </router-view>
</template>

<style>
.page-enter-active { transition: opacity 0.2s ease-out, transform 0.2s ease-out; }
.page-leave-active { transition: opacity 0.15s ease-in, transform 0.15s ease-in; }
.page-enter-from { opacity: 0; transform: translateY(8px); }
.page-leave-to { opacity: 0; transform: translateY(-8px); }
</style>
