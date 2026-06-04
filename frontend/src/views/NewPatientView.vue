<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { patientApi } from '../api'
import { useToast } from '../composables/useToast'
import PageHeader from '../components/PageHeader.vue'
import FormField from '../components/FormField.vue'

const router = useRouter()
const toast = useToast()
const loading = ref(false)
const serverError = ref<string | null>(null)

const form = ref({
  firstName: '',
  lastName: '',
  idNumber: '',
  phoneNumber: '',
  gender: 'male',
  age: 0,
})

const fieldErrors = ref<Record<string, string>>({})

function validate() {
  fieldErrors.value = {}
  if (!form.value.firstName || form.value.firstName.length < 3) fieldErrors.value.firstName = 'At least 3 characters'
  if (!form.value.lastName || form.value.lastName.length < 3) fieldErrors.value.lastName = 'At least 3 characters'
  if (!form.value.idNumber || form.value.idNumber.length < 8) fieldErrors.value.idNumber = 'At least 8 characters'
  if (!form.value.phoneNumber || form.value.phoneNumber.length < 7) fieldErrors.value.phoneNumber = 'At least 7 characters'
  if (!form.value.age || form.value.age < 1 || form.value.age > 255) fieldErrors.value.age = 'Must be 1–255'
  return Object.keys(fieldErrors.value).length === 0
}

async function handleSubmit() {
  if (!validate()) return
  loading.value = true
  serverError.value = null
  try {
    await patientApi.create(form.value)
    toast.success(`Patient ${form.value.firstName} ${form.value.lastName} created`)
    router.push('/patients')
  } catch (e: any) {
    serverError.value = e.userMessage || 'Failed to create patient'
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div class="max-w-lg">
    <PageHeader title="New Patient" subtitle="Register a new patient in the system" />

    <form @submit.prevent="handleSubmit" class="bg-white/70 backdrop-blur-2xl rounded-3xl border border-cupertino-gray-100/60 p-6 shadow-sm space-y-4">
      <div class="grid grid-cols-2 gap-3">
        <FormField label="First Name" :error="fieldErrors.firstName">
          <input v-model="form.firstName" class="w-full px-4 py-3 rounded-xl bg-cupertino-gray-50 border text-sm focus:outline-none focus:ring-2 focus:ring-cupertino-blue/30 transition-all" :class="fieldErrors.firstName ? 'border-cupertino-red/40' : 'border-cupertino-gray-100'" />
        </FormField>
        <FormField label="Last Name" :error="fieldErrors.lastName">
          <input v-model="form.lastName" class="w-full px-4 py-3 rounded-xl bg-cupertino-gray-50 border text-sm focus:outline-none focus:ring-2 focus:ring-cupertino-blue/30 transition-all" :class="fieldErrors.lastName ? 'border-cupertino-red/40' : 'border-cupertino-gray-100'" />
        </FormField>
      </div>

      <FormField label="ID Number" :error="fieldErrors.idNumber">
        <input v-model="form.idNumber" placeholder="National ID" class="w-full px-4 py-3 rounded-xl bg-cupertino-gray-50 border text-sm focus:outline-none focus:ring-2 focus:ring-cupertino-blue/30 transition-all" :class="fieldErrors.idNumber ? 'border-cupertino-red/40' : 'border-cupertino-gray-100'" />
      </FormField>

      <FormField label="Phone Number" :error="fieldErrors.phoneNumber">
        <input v-model="form.phoneNumber" placeholder="+2547XX XXX XXX" class="w-full px-4 py-3 rounded-xl bg-cupertino-gray-50 border text-sm focus:outline-none focus:ring-2 focus:ring-cupertino-blue/30 transition-all" :class="fieldErrors.phoneNumber ? 'border-cupertino-red/40' : 'border-cupertino-gray-100'" />
      </FormField>

      <div class="grid grid-cols-2 gap-3">
        <FormField label="Gender">
          <select v-model="form.gender" class="w-full px-4 py-3 rounded-xl bg-cupertino-gray-50 border border-cupertino-gray-100 text-sm text-cupertino-gray-900 focus:outline-none focus:ring-2 focus:ring-cupertino-blue/30 transition-all appearance-none">
            <option value="male">Male</option>
            <option value="female">Female</option>
          </select>
        </FormField>
        <FormField label="Age" :error="fieldErrors.age">
          <input v-model.number="form.age" type="number" min="0" max="255" class="w-full px-4 py-3 rounded-xl bg-cupertino-gray-50 border text-sm focus:outline-none focus:ring-2 focus:ring-cupertino-blue/30 transition-all" :class="fieldErrors.age ? 'border-cupertino-red/40' : 'border-cupertino-gray-100'" />
        </FormField>
      </div>

      <p v-if="serverError" class="text-xs text-cupertino-red bg-cupertino-red/5 rounded-lg px-3 py-2">{{ serverError }}</p>

      <button type="submit" :disabled="loading" class="w-full py-3 rounded-xl bg-cupertino-blue text-white text-sm font-semibold shadow-sm hover:bg-cupertino-blue/90 active:scale-[0.98] transition-all disabled:opacity-50">
        <span v-if="loading" class="inline-flex items-center gap-2">
          <svg class="animate-spin h-4 w-4" viewBox="0 0 24 24" fill="none"><circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"/><path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z"/></svg>
          Creating...
        </span>
        <span v-else>Create Patient</span>
      </button>
    </form>
  </div>
</template>
