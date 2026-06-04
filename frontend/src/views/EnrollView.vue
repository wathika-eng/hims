<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { patientApi, programApi } from '../api'
import type { Patient, Program } from '../types'
import { useToast } from '../composables/useToast'
import { maskID } from '../utils/mask'
import PageHeader from '../components/PageHeader.vue'
import ErrorAlert from '../components/ErrorAlert.vue'
import FormField from '../components/FormField.vue'

const toast = useToast()
const loading = ref(false)
const serverError = ref<string | null>(null)

const patients = ref<Patient[]>([])
const programs = ref<Program[]>([])
const loadError = ref<string | null>(null)
const loadingData = ref(true)

const form = ref({
  patientID: '',
  programName: '',
})

onMounted(async () => {
  try {
    const [pats, progs] = await Promise.all([
      patientApi.list(),
      programApi.list(),
    ])
    patients.value = pats.data.data
    programs.value = progs.data.data
  } catch (e: any) {
    loadError.value = e.userMessage || 'Failed to load form data'
  } finally {
    loadingData.value = false
  }
})

async function handleEnroll() {
  if (!form.value.patientID || !form.value.programName) {
    serverError.value = 'Please select both a patient and a program'
    return
  }
  loading.value = true
  serverError.value = null
  try {
    const res = await programApi.enroll(form.value)
    const name = `${res.data.data.firstName} ${res.data.data.lastName}`
    toast.success(`Enrolled ${name} successfully`)
    form.value = { patientID: '', programName: '' }
  } catch (e: any) {
    serverError.value = e.userMessage || 'Failed to enroll patient'
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div class="max-w-lg">
    <PageHeader title="Enroll Patient" subtitle="Assign a patient to a healthcare program" />

    <ErrorAlert :message="loadError" title="Could not load form data" />

    <div v-if="loadingData" class="bg-white/70 rounded-3xl border border-cupertino-gray-100/60 p-6 space-y-4 animate-pulse">
      <div class="h-12 bg-cupertino-gray-100 rounded-xl" />
      <div class="h-12 bg-cupertino-gray-100 rounded-xl" />
    </div>

    <form v-else-if="!loadError" @submit.prevent="handleEnroll" class="bg-white/70 backdrop-blur-2xl rounded-3xl border border-cupertino-gray-100/60 p-6 shadow-sm space-y-4">
      <FormField label="Patient">
        <select v-model="form.patientID" required class="w-full px-4 py-3 rounded-xl bg-cupertino-gray-50 border border-cupertino-gray-100 text-sm text-cupertino-gray-900 focus:outline-none focus:ring-2 focus:ring-cupertino-blue/30 transition-all appearance-none">
          <option value="" disabled>Select a patient</option>
          <option v-for="p in patients" :key="p.ID" :value="p.idNumber">{{ p.firstName }} {{ p.lastName }} ({{ maskID(p.idNumber) }})</option>
        </select>
        <p v-if="!patients.length && !loadingData" class="mt-1 text-xs text-cupertino-orange flex items-center gap-1">
          <svg viewBox="0 0 24 24" class="w-3.5 h-3.5 shrink-0" fill="currentColor"><path d="M1 21h22L12 2 1 21zm12-3h-2v-2h2v2zm0-4h-2v-4h2v4z"/></svg>
          No patients available. Create one first.
        </p>
      </FormField>

      <FormField label="Program">
        <select v-model="form.programName" required class="w-full px-4 py-3 rounded-xl bg-cupertino-gray-50 border border-cupertino-gray-100 text-sm text-cupertino-gray-900 focus:outline-none focus:ring-2 focus:ring-cupertino-blue/30 transition-all appearance-none">
          <option value="" disabled>Select a program</option>
          <option v-for="prog in programs" :key="prog.ID" :value="prog.program">{{ prog.program }} (#{{ prog.programCode }})</option>
        </select>
        <p v-if="!programs.length" class="mt-1 text-xs text-cupertino-orange flex items-center gap-1">
          <svg viewBox="0 0 24 24" class="w-3.5 h-3.5 shrink-0" fill="currentColor"><path d="M1 21h22L12 2 1 21zm12-3h-2v-2h2v2zm0-4h-2v-4h2v4z"/></svg>
          No programs available. Create one first.
        </p>
      </FormField>

      <p v-if="serverError" class="text-xs text-cupertino-red bg-cupertino-red/5 rounded-lg px-3 py-2">{{ serverError }}</p>

      <button type="submit" :disabled="loading" class="w-full py-3 rounded-xl bg-cupertino-blue text-white text-sm font-semibold shadow-sm hover:bg-cupertino-blue/90 active:scale-[0.98] transition-all disabled:opacity-50">
        <span v-if="loading" class="inline-flex items-center gap-2">
          <svg class="animate-spin h-4 w-4" viewBox="0 0 24 24" fill="none"><circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"/><path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z"/></svg>
          Enrolling...
        </span>
        <span v-else>Enroll Patient</span>
      </button>
    </form>
  </div>
</template>
