<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import { patientApi, programApi } from '../api'
import { useCache } from '../composables/useCache'
import type { Patient, Program } from '../types'
import { maskID, maskName } from '../utils/mask'
import PageHeader from '../components/PageHeader.vue'
import StatCard from '../components/StatCard.vue'
import ErrorAlert from '../components/ErrorAlert.vue'

const patients = ref<Patient[]>([])
const programs = ref<Program[]>([])
const loading = ref(true)
const fetchError = ref<string | null>(null)
const { fetchWithCache } = useCache()

const enrolledCount = computed(() => patients.value.filter(p => p.patientPrograms?.length).length)
const avgAge = computed(() =>
  patients.value.length
    ? Math.round(patients.value.reduce((a, p) => a + p.age, 0) / patients.value.length)
    : 0,
)

const icons = {
  patients: 'M12 12c2.21 0 4-1.79 4-4s-1.79-4-4-4-4 1.79-4 4 1.79 4 4 4zm-8 6c0-2.67 5.33-4 8-4s8 1.33 8 4v2H4v-2z',
  programs: 'M4 4h12v12H4V4zm14 6h2v8H8v-2h10v-6z',
  enrolled: 'M9 16.17L4.83 12l-1.42 1.41L9 19 21 7l-1.41-1.41L9 16.17z',
  age: 'M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm0 3c1.66 0 3 1.34 3 3s-1.34 3-3 3-3-1.34-3-3 1.34-3 3-3zm0 14.2c-2.5 0-4.71-1.28-6-3.22.03-1.99 4-3.08 6-3.08 1.99 0 5.97 1.09 6 3.08-1.29 1.94-3.5 3.22-6 3.22z',
}

onMounted(async () => {
  try {
    const [pats, progs] = await Promise.all([
      fetchWithCache('patients', () => patientApi.list().then(r => r.data.data), 2 * 60 * 1000),
      fetchWithCache('programs', () => programApi.list().then(r => r.data.data), 2 * 60 * 1000),
    ])
    patients.value = pats.data
    programs.value = progs.data
  } catch (e: any) {
    fetchError.value = e.userMessage || 'Failed to load dashboard data'
  } finally {
    loading.value = false
  }
})
</script>

<template>
  <div>
    <PageHeader title="Dashboard" subtitle="Overview of your health information system" />

    <ErrorAlert :message="fetchError" title="Failed to load" class="mb-4 sm:mb-6" />

    <div v-if="loading" class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4">
      <div v-for="i in 4" :key="i" class="bg-white/70 rounded-2xl border border-cupertino-gray-100/60 p-5 shadow-sm animate-pulse">
        <div class="h-3 w-20 bg-cupertino-gray-100 rounded mb-3" />
        <div class="h-8 w-16 bg-cupertino-gray-100 rounded" />
      </div>
    </div>

    <div v-else class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4">
      <StatCard label="Total Patients" :value="patients.length" :icon="icons.patients" color="#007aff" />
      <StatCard label="Programs" :value="programs.length" :icon="icons.programs" color="#34c759" />
      <StatCard label="Enrolled" :value="enrolledCount" :icon="icons.enrolled" color="#af52de" />
      <StatCard label="Avg. Age" :value="avgAge" :icon="icons.age" color="#ff9500" />
    </div>

    <div class="mt-8 grid grid-cols-1 lg:grid-cols-2 gap-6">
      <div class="bg-white/70 backdrop-blur-2xl rounded-2xl border border-cupertino-gray-100/60 p-5 shadow-sm">
        <h2 class="text-base font-semibold text-cupertino-gray-900 mb-4">Recent Patients</h2>
        <div v-if="loading" class="space-y-2">
          <div v-for="i in 5" :key="i" class="flex items-center gap-3 px-3 py-2.5 animate-pulse">
            <div class="w-8 h-8 rounded-full bg-cupertino-gray-100" />
            <div class="flex-1"><div class="h-3 w-32 bg-cupertino-gray-100 rounded mb-1" /><div class="h-2 w-20 bg-cupertino-gray-50 rounded" /></div>
          </div>
        </div>
        <div v-else-if="patients.length === 0" class="text-sm text-cupertino-gray-400 py-8 text-center">No patients registered yet</div>
        <div v-else class="space-y-2">
          <router-link
            v-for="p in patients.slice(0, 5)"
            :key="p.ID"
            :to="`/patients/${p.ID}`"
            class="flex items-center justify-between px-3 py-2.5 rounded-xl hover:bg-cupertino-gray-50 transition-colors"
          >
            <div class="flex items-center gap-3">
              <div class="w-8 h-8 rounded-full bg-cupertino-blue/10 text-cupertino-blue flex items-center justify-center text-xs font-semibold">
                {{ p.firstName[0] }}{{ p.lastName[0] }}
              </div>
              <div>
                <p class="text-sm font-medium text-cupertino-gray-900">{{ maskName(p.firstName, p.lastName) }}</p>
                <p class="text-xs text-cupertino-gray-400 font-mono">{{ maskID(p.idNumber) }}</p>
              </div>
            </div>
            <span class="text-xs text-cupertino-gray-400">{{ p.gender }}, {{ p.age }}</span>
          </router-link>
        </div>
      </div>

      <div class="bg-white/70 backdrop-blur-2xl rounded-2xl border border-cupertino-gray-100/60 p-5 shadow-sm">
        <h2 class="text-base font-semibold text-cupertino-gray-900 mb-4">Programs</h2>
        <div v-if="loading" class="space-y-2">
          <div v-for="i in 4" :key="i" class="flex items-center gap-3 px-3 py-2.5 animate-pulse">
            <div class="w-8 h-8 rounded-lg bg-cupertino-gray-100" />
            <div class="h-3 w-40 bg-cupertino-gray-100 rounded" />
          </div>
        </div>
        <div v-else-if="programs.length === 0" class="text-sm text-cupertino-gray-400 py-8 text-center">No programs yet</div>
        <div v-else class="space-y-2">
          <div v-for="prog in programs" :key="prog.ID" class="flex items-center justify-between px-3 py-2.5 rounded-xl hover:bg-cupertino-gray-50 transition-colors">
            <div class="flex items-center gap-3">
              <div class="w-8 h-8 rounded-lg bg-cupertino-green/10 text-cupertino-green flex items-center justify-center text-xs font-semibold">
                {{ prog.programCode }}
              </div>
              <p class="text-sm font-medium text-cupertino-gray-900">{{ prog.program }}</p>
            </div>
            <span class="text-xs text-cupertino-gray-400">#{{ prog.programCode }}</span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
