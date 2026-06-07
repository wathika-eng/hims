<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { patientApi } from '../api'
import type { Patient } from '../types'
import { maskPhone, maskID, maskName } from '../utils/mask'
import ErrorAlert from '../components/ErrorAlert.vue'

const route = useRoute()
const patient = ref<Patient | null>(null)
const loading = ref(true)
const fetchError = ref<string | null>(null)

onMounted(async () => {
  try {
    const res = await patientApi.get(Number(route.params.id))
    patient.value = res.data.data
  } catch (e: any) {
    if (e.response?.status === 404) {
      fetchError.value = 'Patient not found'
    } else {
      fetchError.value = e.userMessage || 'Failed to load patient'
    }
  } finally {
    loading.value = false
  }
})

function formatDate(dateStr: string | null | undefined) {
  if (!dateStr) return '—'
  return new Date(dateStr).toLocaleDateString('en-US', {
    year: 'numeric', month: 'short', day: 'numeric', hour: '2-digit', minute: '2-digit',
  })
}
</script>

<template>
  <div class="max-w-2xl">
    <router-link to="/patients" class="inline-flex items-center gap-1.5 text-sm text-cupertino-blue font-medium mb-4 hover:underline">
      <svg viewBox="0 0 24 24" class="w-4 h-4" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"><polyline points="15 18 9 12 15 6"/></svg>
      Back to Patients
    </router-link>

    <ErrorAlert :message="fetchError" title="Error" />

    <div v-if="loading" class="space-y-6 animate-pulse">
      <div class="flex items-center gap-4 mb-8">
        <div class="w-16 h-16 rounded-full bg-cupertino-gray-100" />
        <div class="space-y-2">
          <div class="h-6 w-48 bg-cupertino-gray-100 rounded" />
          <div class="h-3 w-32 bg-cupertino-gray-50 rounded" />
        </div>
      </div>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
        <div v-for="i in 2" :key="i" class="bg-white/70 rounded-2xl border border-cupertino-gray-100/60 p-5">
          <div class="h-3 w-24 bg-cupertino-gray-100 rounded mb-4" />
          <div class="space-y-3">
            <div v-for="j in 4" :key="j" class="h-3 w-full bg-cupertino-gray-50 rounded" />
          </div>
        </div>
      </div>
    </div>

    <template v-else-if="patient">
      <div class="flex items-center gap-4 mb-6 sm:mb-8">
        <div class="w-16 h-16 rounded-full bg-cupertino-blue/10 text-cupertino-blue flex items-center justify-center text-xl font-bold">
          {{ patient.firstName[0] }}{{ patient.lastName[0] }}
        </div>
        <div>
          <h1 class="text-2xl font-bold text-cupertino-gray-900">{{ maskName(patient.firstName, patient.lastName) }}</h1>
          <p class="text-sm text-cupertino-gray-400">Patient ID: {{ patient.ID }} &middot; {{ patient.role }}</p>
        </div>
      </div>

      <div class="grid grid-cols-1 md:grid-cols-2 gap-4 mb-6">
        <div class="bg-white/70 backdrop-blur-2xl rounded-2xl border border-cupertino-gray-100/60 p-5 shadow-sm">
          <h3 class="text-xs font-semibold text-cupertino-gray-400 uppercase tracking-wide mb-3">Personal Info</h3>
          <dl class="space-y-2.5">
            <div class="flex justify-between"><dt class="text-sm text-cupertino-gray-400">ID Number</dt><dd class="text-sm font-medium text-cupertino-gray-900 font-mono">{{ maskID(patient.idNumber) }}</dd></div>
            <div class="flex justify-between"><dt class="text-sm text-cupertino-gray-400">Gender</dt><dd class="text-sm font-medium text-cupertino-gray-900 capitalize">{{ patient.gender }}</dd></div>
            <div class="flex justify-between"><dt class="text-sm text-cupertino-gray-400">Age</dt><dd class="text-sm font-medium text-cupertino-gray-900">{{ patient.age }}</dd></div>
            <div class="flex justify-between"><dt class="text-sm text-cupertino-gray-400">Phone</dt><dd class="text-sm font-medium text-cupertino-gray-900 font-mono">{{ maskPhone(patient.phoneNumber) }}</dd></div>
          </dl>
        </div>
        <div class="bg-white/70 backdrop-blur-2xl rounded-2xl border border-cupertino-gray-100/60 p-5 shadow-sm">
          <h3 class="text-xs font-semibold text-cupertino-gray-400 uppercase tracking-wide mb-3">Timeline</h3>
          <dl class="space-y-2.5">
            <div class="flex justify-between"><dt class="text-sm text-cupertino-gray-400">Created</dt><dd class="text-sm font-medium text-cupertino-gray-900">{{ formatDate(patient.createdAt) }}</dd></div>
            <div class="flex justify-between"><dt class="text-sm text-cupertino-gray-400">Updated</dt><dd class="text-sm font-medium text-cupertino-gray-900">{{ formatDate(patient.UpdatedAt) }}</dd></div>
          </dl>
        </div>
      </div>

      <div class="bg-white/70 backdrop-blur-2xl rounded-2xl border border-cupertino-gray-100/60 p-5 shadow-sm">
        <h3 class="text-xs font-semibold text-cupertino-gray-400 uppercase tracking-wide mb-3">Enrolled Programs</h3>
        <div v-if="!patient.patientPrograms?.length" class="text-sm text-cupertino-gray-400 py-4 text-center">Not enrolled in any programs</div>
        <div v-else class="space-y-2">
          <div v-for="prog in patient.patientPrograms" :key="prog.ID" class="flex items-center gap-3 px-3 py-2.5 rounded-xl bg-cupertino-green/5">
            <div class="w-7 h-7 rounded-lg bg-cupertino-green/10 text-cupertino-green flex items-center justify-center text-xs font-semibold">{{ prog.programCode }}</div>
            <span class="text-sm font-medium text-cupertino-gray-900">{{ prog.program }}</span>
          </div>
        </div>
      </div>
    </template>

    <div v-else-if="!fetchError" class="text-center py-20 text-cupertino-gray-400">
      <svg viewBox="0 0 24 24" class="w-12 h-12 mx-auto mb-3 text-cupertino-gray-200" fill="currentColor"><path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm1 15h-2v-2h2v2zm0-4h-2V7h2v6z"/></svg>
      <p class="text-sm">Patient not found</p>
    </div>
  </div>
</template>
