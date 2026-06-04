<script setup lang="ts">
import { ref } from 'vue'
import { useAuthStore } from '../stores/auth'
import FormField from '../components/FormField.vue'

const auth = useAuthStore()

const form = ref({
  firstName: '',
  lastName: '',
  email: '',
  password: '',
  licenseNumber: '',
  specialization: '',
})

async function handleSignup() {
  auth.clearError()
  await auth.signup(form.value)
}
</script>

<template>
  <div class="min-h-screen bg-cupertino-gray-50 flex items-center justify-center p-4">
    <div class="w-full max-w-sm">
      <div class="text-center mb-8">
        <div class="w-16 h-16 rounded-2xl bg-cupertino-green mx-auto flex items-center justify-center shadow-lg shadow-cupertino-green/20 mb-4">
          <svg viewBox="0 0 24 24" class="w-9 h-9 text-white" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round">
            <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"/><circle cx="12" cy="7" r="4"/>
          </svg>
        </div>
        <h1 class="text-2xl font-bold text-cupertino-gray-900 tracking-tight">Create Account</h1>
        <p class="mt-1 text-sm text-cupertino-gray-400">Register as a doctor</p>
      </div>

      <form @submit.prevent="handleSignup" class="bg-white/70 backdrop-blur-2xl rounded-3xl border border-cupertino-gray-100/60 p-6 shadow-sm space-y-3.5">
        <div class="grid grid-cols-2 gap-3">
          <FormField label="First">
            <input v-model="form.firstName" placeholder="John" class="w-full px-4 py-3 rounded-xl bg-cupertino-gray-50 border border-cupertino-gray-100 text-sm text-cupertino-gray-900 placeholder:text-cupertino-gray-300 focus:outline-none focus:ring-2 focus:ring-cupertino-blue/30 focus:border-cupertino-blue transition-all" />
          </FormField>
          <FormField label="Last">
            <input v-model="form.lastName" placeholder="Doe" class="w-full px-4 py-3 rounded-xl bg-cupertino-gray-50 border border-cupertino-gray-100 text-sm text-cupertino-gray-900 placeholder:text-cupertino-gray-300 focus:outline-none focus:ring-2 focus:ring-cupertino-blue/30 focus:border-cupertino-blue transition-all" />
          </FormField>
        </div>

        <FormField label="Email">
          <input v-model="form.email" type="email" placeholder="doctor@clinic.com" class="w-full px-4 py-3 rounded-xl bg-cupertino-gray-50 border border-cupertino-gray-100 text-sm text-cupertino-gray-900 placeholder:text-cupertino-gray-300 focus:outline-none focus:ring-2 focus:ring-cupertino-blue/30 focus:border-cupertino-blue transition-all" />
        </FormField>

        <FormField label="Password">
          <input v-model="form.password" type="password" placeholder="Min 4 characters" class="w-full px-4 py-3 rounded-xl bg-cupertino-gray-50 border border-cupertino-gray-100 text-sm text-cupertino-gray-900 placeholder:text-cupertino-gray-300 focus:outline-none focus:ring-2 focus:ring-cupertino-blue/30 focus:border-cupertino-blue transition-all" />
        </FormField>

        <FormField label="License No.">
          <input v-model="form.licenseNumber" placeholder="License number" class="w-full px-4 py-3 rounded-xl bg-cupertino-gray-50 border border-cupertino-gray-100 text-sm text-cupertino-gray-900 placeholder:text-cupertino-gray-300 focus:outline-none focus:ring-2 focus:ring-cupertino-blue/30 focus:border-cupertino-blue transition-all" />
        </FormField>

        <FormField label="Specialization">
          <input v-model="form.specialization" placeholder="e.g. Cardiology" class="w-full px-4 py-3 rounded-xl bg-cupertino-gray-50 border border-cupertino-gray-100 text-sm text-cupertino-gray-900 placeholder:text-cupertino-gray-300 focus:outline-none focus:ring-2 focus:ring-cupertino-blue/30 focus:border-cupertino-blue transition-all" />
        </FormField>

        <div v-if="auth.error" class="flex items-center gap-2 text-xs text-cupertino-red bg-cupertino-red/5 rounded-lg px-3 py-2">
          <svg viewBox="0 0 24 24" class="w-4 h-4 shrink-0" fill="currentColor"><path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm1 15h-2v-2h2v2zm0-4h-2V7h2v6z"/></svg>
          <span>{{ auth.error }}</span>
        </div>

        <button
          type="submit"
          :disabled="auth.loading"
          class="w-full py-3 rounded-xl bg-cupertino-green text-white text-sm font-semibold shadow-sm shadow-cupertino-green/20 hover:bg-cupertino-green/90 active:scale-[0.98] transition-all disabled:opacity-50 disabled:cursor-not-allowed"
        >
          <span v-if="auth.loading" class="inline-flex items-center gap-2">
            <svg class="animate-spin h-4 w-4" viewBox="0 0 24 24" fill="none"><circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"/><path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z"/></svg>
            Creating account...
          </span>
          <span v-else>Create Account</span>
        </button>
      </form>

      <p class="mt-6 text-center text-sm text-cupertino-gray-400">
        Already have an account?
        <router-link to="/login" class="text-cupertino-blue font-semibold hover:underline">Sign In</router-link>
      </p>
    </div>
  </div>
</template>
