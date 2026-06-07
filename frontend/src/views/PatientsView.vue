<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import { patientApi } from '../api'
import { useCache } from '../composables/useCache'
import type { Patient } from '../types'
import { maskPhone, maskID, maskName } from '../utils/mask'
import PageHeader from '../components/PageHeader.vue'
import ErrorAlert from '../components/ErrorAlert.vue'

const patients = ref<Patient[]>([])
const loading = ref(true)
const fetchError = ref<string | null>(null)
const search = ref('')
const { fetchWithCache } = useCache()

const filtered = computed(() => {
  if (!search.value) return patients.value
  const q = search.value.toLowerCase()
  return patients.value.filter(
    p =>
      p.firstName.toLowerCase().includes(q) ||
      p.lastName.toLowerCase().includes(q) ||
      p.idNumber.includes(q) ||
      p.phoneNumber.includes(q),
  )
})

onMounted(async () => {
  try {
    const { data } = await fetchWithCache(
      'patients',
      () => patientApi.list().then(r => r.data.data),
      2 * 60 * 1000,
    )
    patients.value = data
  } catch (e: any) {
    fetchError.value = e.userMessage || 'Failed to load patients'
  } finally {
    loading.value = false
  }
})
</script>

<template>
  <div>
    <div class="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-4 mb-6 sm:mb-8">
      <PageHeader title="Patients" subtitle="Manage all registered patients" />
      <router-link
        to="/patients/new"
        class="self-stretch sm:self-auto inline-flex items-center justify-center gap-2 px-4 py-2.5 rounded-xl bg-cupertino-blue text-white text-sm font-semibold shadow-sm hover:bg-cupertino-blue/90 active:scale-[0.98] transition-all"
      >
        <svg viewBox="0 0 24 24" class="w-4 h-4" fill="currentColor"><path d="M19 13h-6v6h-2v-6H5v-2h6V5h2v6h6v2z"/></svg>
        New Patient
      </router-link>
    </div>

    <ErrorAlert :message="fetchError" title="Could not load patients" />

    <div class="mb-5">
      <div class="relative max-w-full sm:max-w-xs">
        <svg viewBox="0 0 24 24" class="w-4 h-4 text-cupertino-gray-400 absolute left-3.5 top-1/2 -translate-y-1/2" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
          <circle cx="11" cy="11" r="8"/><line x1="21" y1="21" x2="16.65" y2="16.65"/>
        </svg>
        <input v-model="search" placeholder="Search patients..." class="w-full pl-10 pr-4 py-2.5 rounded-xl bg-white/70 backdrop-blur-2xl border border-cupertino-gray-100/60 text-sm text-cupertino-gray-900 placeholder:text-cupertino-gray-300 focus:outline-none focus:ring-2 focus:ring-cupertino-blue/30 transition-all shadow-sm" />
      </div>
    </div>

    <div v-if="loading" class="space-y-2">
      <div v-for="i in 6" :key="i" class="flex items-center gap-4 px-5 py-4 bg-white/70 rounded-2xl border border-cupertino-gray-100/60 animate-pulse">
        <div class="w-8 h-8 rounded-full bg-cupertino-gray-100" />
        <div class="flex-1 space-y-1.5">
          <div class="h-3 w-40 bg-cupertino-gray-100 rounded" />
          <div class="h-2 w-24 bg-cupertino-gray-50 rounded" />
        </div>
        <div class="h-3 w-16 bg-cupertino-gray-100 rounded" />
      </div>
    </div>

    <div v-else-if="!fetchError && filtered.length === 0" class="text-center py-20 text-cupertino-gray-400">
      <svg viewBox="0 0 24 24" class="w-12 h-12 mx-auto mb-3 text-cupertino-gray-200" fill="currentColor"><path d="M12 12c2.21 0 4-1.79 4-4s-1.79-4-4-4-4 1.79-4 4 1.79 4 4 4zm-8 6c0-2.67 5.33-4 8-4s8 1.33 8 4v2H4v-2z"/></svg>
      <p class="text-sm">{{ search ? 'No patients match your search' : 'No patients registered yet' }}</p>
    </div>

    <!-- Desktop table -->
    <div v-else-if="!fetchError" class="bg-white/70 backdrop-blur-2xl rounded-2xl border border-cupertino-gray-100/60 shadow-sm overflow-hidden hidden sm:block">
      <div class="overflow-x-auto">
        <table class="w-full">
          <thead>
            <tr class="border-b border-cupertino-gray-100/60">
              <th class="text-left px-5 py-3 text-xs font-semibold text-cupertino-gray-400 uppercase tracking-wide">Patient</th>
              <th class="text-left px-5 py-3 text-xs font-semibold text-cupertino-gray-400 uppercase tracking-wide">ID No.</th>
              <th class="text-left px-5 py-3 text-xs font-semibold text-cupertino-gray-400 uppercase tracking-wide">Phone</th>
              <th class="text-left px-5 py-3 text-xs font-semibold text-cupertino-gray-400 uppercase tracking-wide">Gender</th>
              <th class="text-left px-5 py-3 text-xs font-semibold text-cupertino-gray-400 uppercase tracking-wide">Age</th>
              <th class="text-left px-5 py-3 text-xs font-semibold text-cupertino-gray-400 uppercase tracking-wide">Programs</th>
              <th class="text-right px-5 py-3"></th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="p in filtered" :key="p.ID" class="border-b border-cupertino-gray-100/40 hover:bg-cupertino-gray-50/50 transition-colors">
              <td class="px-5 py-3">
                <router-link :to="`/patients/${p.ID}`" class="flex items-center gap-3">
                  <div class="w-8 h-8 rounded-full bg-cupertino-blue/10 text-cupertino-blue flex items-center justify-center text-xs font-semibold">
                    {{ p.firstName[0] }}{{ p.lastName[0] }}
                  </div>
                  <span class="text-sm font-medium text-cupertino-gray-900">{{ maskName(p.firstName, p.lastName) }}</span>
                </router-link>
              </td>
              <td class="px-5 py-3 text-sm text-cupertino-gray-500 font-mono">{{ maskID(p.idNumber) }}</td>
              <td class="px-5 py-3 text-sm text-cupertino-gray-500 font-mono">{{ maskPhone(p.phoneNumber) }}</td>
              <td class="px-5 py-3">
                <span class="inline-flex text-xs font-medium px-2.5 py-1 rounded-lg" :class="p.gender === 'male' ? 'bg-cupertino-blue/10 text-cupertino-blue' : 'bg-cupertino-purple/10 text-cupertino-purple'">
                  {{ p.gender }}
                </span>
              </td>
              <td class="px-5 py-3 text-sm text-cupertino-gray-500">{{ p.age }}</td>
              <td class="px-5 py-3">
                <span class="inline-flex items-center gap-1 text-xs text-cupertino-gray-400">
                  <svg viewBox="0 0 24 24" class="w-3.5 h-3.5" fill="currentColor"><path d="M4 4h12v12H4V4zm14 6h2v8H8v-2h10v-6z"/></svg>
                  {{ p.patientPrograms?.length || 0 }}
                </span>
              </td>
              <td class="px-5 py-3 text-right">
                <router-link :to="`/patients/${p.ID}`" class="text-xs font-semibold text-cupertino-blue hover:underline">View</router-link>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- Mobile cards -->
    <div v-else-if="!fetchError" class="sm:hidden space-y-3">
      <router-link
        v-for="p in filtered"
        :key="p.ID"
        :to="`/patients/${p.ID}`"
        class="block bg-white/70 backdrop-blur-2xl rounded-2xl border border-cupertino-gray-100/60 p-4 shadow-sm hover:shadow-md hover:border-cupertino-gray-200/60 transition-all active:scale-[0.99]"
      >
        <div class="flex items-center gap-3 mb-3">
          <div class="w-10 h-10 rounded-full bg-cupertino-blue/10 text-cupertino-blue flex items-center justify-center text-sm font-semibold">
            {{ p.firstName[0] }}{{ p.lastName[0] }}
          </div>
          <div class="min-w-0 flex-1">
            <p class="text-sm font-semibold text-cupertino-gray-900 truncate">{{ maskName(p.firstName, p.lastName) }}</p>
            <p class="text-xs text-cupertino-gray-400 font-mono truncate">{{ maskID(p.idNumber) }}</p>
          </div>
          <svg viewBox="0 0 24 24" class="w-4 h-4 text-cupertino-gray-300 shrink-0" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polyline points="9 18 15 12 9 6"/></svg>
        </div>
        <div class="flex items-center gap-3 text-xs text-cupertino-gray-500">
          <span class="inline-flex items-center gap-1">
            <svg viewBox="0 0 24 24" class="w-3.5 h-3.5" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M22 16.92v3a2 2 0 0 1-2.18 2 19.79 19.79 0 0 1-8.63-3.07 19.5 19.5 0 0 1-6-6 19.79 19.79 0 0 1-3.07-8.67A2 2 0 0 1 4.11 2h3a2 2 0 0 1 2 1.72 12.84 12.84 0 0 0 .7 2.81 2 2 0 0 1-.45 2.11L8.09 9.91a16 16 0 0 0 6 6l1.27-1.27a2 2 0 0 1 2.11-.45 12.84 12.84 0 0 0 2.81.7A2 2 0 0 1 22 16.92z"/></svg>
            {{ maskPhone(p.phoneNumber) }}
          </span>
          <span class="inline-flex items-center gap-1 px-2 py-0.5 rounded-md text-[11px] font-medium" :class="p.gender === 'male' ? 'bg-cupertino-blue/10 text-cupertino-blue' : 'bg-cupertino-purple/10 text-cupertino-purple'">
            {{ p.gender }}, {{ p.age }}
          </span>
          <span class="inline-flex items-center gap-1 ml-auto">
            <svg viewBox="0 0 24 24" class="w-3.5 h-3.5" fill="currentColor"><path d="M4 4h12v12H4V4zm14 6h2v8H8v-2h10v-6z"/></svg>
            {{ p.patientPrograms?.length || 0 }}
          </span>
        </div>
      </router-link>
    </div>
  </div>
</template>
