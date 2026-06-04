<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { programApi, downloadPdf } from '../api'
import { useToast } from '../composables/useToast'
import type { Program } from '../types'
import PageHeader from '../components/PageHeader.vue'
import ErrorAlert from '../components/ErrorAlert.vue'

const programs = ref<Program[]>([])
const loading = ref(true)
const downloading = ref(false)
const fetchError = ref<string | null>(null)
const toast = useToast()

onMounted(async () => {
  try {
    const res = await programApi.list()
    programs.value = res.data.data
  } catch (e: any) {
    fetchError.value = e.userMessage || 'Failed to load programs'
  } finally {
    loading.value = false
  }
})

async function handleDownload() {
  downloading.value = true
  try {
    await downloadPdf()
    toast.success('PDF downloaded')
  } catch (e: any) {
    toast.error(e.message || 'Failed to download PDF')
  } finally {
    downloading.value = false
  }
}
</script>

<template>
  <div>
    <div class="flex items-center justify-between mb-8">
      <PageHeader title="Programs" subtitle="Manage healthcare programs" />
      <div class="flex gap-2">
        <button
          @click="handleDownload"
          :disabled="downloading"
          class="inline-flex items-center gap-2 px-4 py-2.5 rounded-xl bg-white/70 backdrop-blur-2xl border border-cupertino-gray-100/60 text-cupertino-gray-700 text-sm font-semibold shadow-sm hover:bg-cupertino-gray-50 active:scale-[0.98] transition-all disabled:opacity-50"
        >
          <svg v-if="!downloading" viewBox="0 0 24 24" class="w-4 h-4" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"/><polyline points="14 2 14 8 20 8"/><line x1="16" y1="13" x2="8" y2="13"/><line x1="16" y1="17" x2="8" y2="17"/></svg>
          <svg v-else class="animate-spin h-4 w-4" viewBox="0 0 24 24" fill="none"><circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"/><path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z"/></svg>
          {{ downloading ? 'Downloading...' : 'Export PDF' }}
        </button>
        <router-link
          to="/programs/new"
          class="inline-flex items-center gap-2 px-4 py-2.5 rounded-xl bg-cupertino-blue text-white text-sm font-semibold shadow-sm hover:bg-cupertino-blue/90 active:scale-[0.98] transition-all"
        >
          <svg viewBox="0 0 24 24" class="w-4 h-4" fill="currentColor"><path d="M19 13h-6v6h-2v-6H5v-2h6V5h2v6h6v2z"/></svg>
          New Program
        </router-link>
      </div>
    </div>

    <ErrorAlert :message="fetchError" title="Could not load programs" />

    <div v-if="loading" class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4">
      <div v-for="i in 6" :key="i" class="bg-white/70 rounded-2xl border border-cupertino-gray-100/60 p-5 animate-pulse">
        <div class="w-10 h-10 rounded-xl bg-cupertino-gray-100 mb-3" />
        <div class="h-4 w-36 bg-cupertino-gray-100 rounded mb-2" />
        <div class="h-3 w-20 bg-cupertino-gray-50 rounded" />
      </div>
    </div>

    <div v-else-if="!fetchError && programs.length === 0" class="text-center py-20 text-cupertino-gray-400">
      <svg viewBox="0 0 24 24" class="w-12 h-12 mx-auto mb-3 text-cupertino-gray-200" fill="currentColor"><path d="M4 4h12v12H4V4zm14 6h2v8H8v-2h10v-6z"/></svg>
      <p class="text-sm">No programs yet</p>
    </div>

    <div v-else-if="!fetchError" class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4">
      <div v-for="prog in programs" :key="prog.ID" class="bg-white/70 backdrop-blur-2xl rounded-2xl border border-cupertino-gray-100/60 p-5 shadow-sm hover:shadow-md hover:border-cupertino-gray-200/60 transition-all">
        <div class="w-10 h-10 rounded-xl bg-cupertino-green/10 text-cupertino-green flex items-center justify-center text-sm font-bold mb-3">
          {{ prog.programCode }}
        </div>
        <h3 class="text-base font-semibold text-cupertino-gray-900 mb-1">{{ prog.program }}</h3>
        <p class="text-xs text-cupertino-gray-400">Code: #{{ prog.programCode }}</p>
      </div>
    </div>
  </div>
</template>
