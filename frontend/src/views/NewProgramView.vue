<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { programApi } from '../api'
import { useToast } from '../composables/useToast'
import { useCache } from '../composables/useCache'
import { required, minLength, inRange, validateAll } from '../utils/validators'
import PageHeader from '../components/PageHeader.vue'
import FormField from '../components/FormField.vue'

const router = useRouter()
const toast = useToast()
const { invalidate } = useCache()
const loading = ref(false)
const serverError = ref<string | null>(null)

const form = ref({
  program: '',
  programCode: 0,
})

const fieldErrors = ref<Record<string, string>>({})

function validate() {
  fieldErrors.value = validateAll({
    program: () => required(form.value.program, 'Program name') ?? minLength(form.value.program, 3, 'Program name'),
    programCode: () => required(form.value.programCode, 'Program code') ?? inRange(form.value.programCode, 1, 9999, 'Program code'),
  })
  return Object.keys(fieldErrors.value).length === 0
}

async function handleSubmit() {
  if (!validate()) return
  loading.value = true
  serverError.value = null
  try {
    await programApi.create(form.value)
    await invalidate('programs')
    toast.success(`Program "${form.value.program}" created`)
    router.push('/programs')
  } catch (e: any) {
    serverError.value = e.userMessage || 'Failed to create program'
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div class="max-w-lg">
    <PageHeader title="New Program" subtitle="Add a new healthcare program" />

    <form @submit.prevent="handleSubmit" class="bg-white/70 backdrop-blur-2xl rounded-3xl border border-cupertino-gray-100/60 p-6 shadow-sm space-y-4">
      <FormField label="Program Name" :error="fieldErrors.program">
        <input v-model="form.program" placeholder="e.g. Hypertension Screening" class="w-full px-4 py-3 rounded-xl bg-cupertino-gray-50 border text-sm focus:outline-none focus:ring-2 focus:ring-cupertino-blue/30 transition-all" :class="fieldErrors.program ? 'border-cupertino-red/40' : 'border-cupertino-gray-100'" />
      </FormField>

      <FormField label="Program Code" :error="fieldErrors.programCode">
        <input v-model.number="form.programCode" type="number" min="1" max="9999" placeholder="e.g. 101" class="w-full px-4 py-3 rounded-xl bg-cupertino-gray-50 border text-sm focus:outline-none focus:ring-2 focus:ring-cupertino-blue/30 transition-all" :class="fieldErrors.programCode ? 'border-cupertino-red/40' : 'border-cupertino-gray-100'" />
      </FormField>

      <p v-if="serverError" class="text-xs text-cupertino-red bg-cupertino-red/5 rounded-lg px-3 py-2">{{ serverError }}</p>

      <button type="submit" :disabled="loading" class="w-full py-3 rounded-xl bg-cupertino-blue text-white text-sm font-semibold shadow-sm hover:bg-cupertino-blue/90 active:scale-[0.98] transition-all disabled:opacity-50">
        <span v-if="loading" class="inline-flex items-center gap-2">
          <svg class="animate-spin h-4 w-4" viewBox="0 0 24 24" fill="none"><circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"/><path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z"/></svg>
          Creating...
        </span>
        <span v-else>Create Program</span>
      </button>
    </form>
  </div>
</template>
