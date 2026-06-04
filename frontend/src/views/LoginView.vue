<script setup lang="ts">
import { ref } from 'vue'
import { useAuthStore } from '../stores/auth'
import FormField from '../components/FormField.vue'

const auth = useAuthStore()

const email = ref('')
const password = ref('')
const showPassword = ref(false)
const showDemo = ref(false)

const demoAccounts = [
  { email: 'john.doe@hospital.com', password: 'password123', name: 'Dr. John Doe', specialization: 'Cardiology' },
  { email: 'jane.smith@hospital.com', password: 'password123', name: 'Dr. Jane Smith', specialization: 'Pediatrics' },
  { email: 'samuel.mwangi@hospital.com', password: 'password123', name: 'Dr. Samuel Mwangi', specialization: 'Internal Medicine' },
] as const

async function handleLogin() {
  auth.clearError()
  await auth.login(email.value, password.value)
}

function selectDemo(acc: typeof demoAccounts[number]) {
  email.value = acc.email
  password.value = acc.password
  handleLogin()
}
</script>

<template>
  <div class="min-h-screen bg-cupertino-gray-50 flex items-center justify-center p-4">
    <div class="w-full max-w-sm">
      <div class="text-center mb-8">
        <div class="w-16 h-16 rounded-2xl bg-cupertino-blue mx-auto flex items-center justify-center shadow-lg shadow-cupertino-blue/20 mb-4">
          <svg viewBox="0 0 24 24" class="w-9 h-9 text-white" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round">
            <path d="M12 2L2 7l10 5 10-5-10-5zM2 17l10 5 10-5M2 12l10 5 10-5"/>
          </svg>
        </div>
        <h1 class="text-2xl font-bold text-cupertino-gray-900 tracking-tight">Welcome to HIMS</h1>
        <p class="mt-1 text-sm text-cupertino-gray-400">Sign in to your account</p>
      </div>

      <form @submit.prevent="handleLogin" class="bg-white/70 backdrop-blur-2xl rounded-3xl border border-cupertino-gray-100/60 p-6 shadow-sm space-y-4">
        <FormField label="Email">
          <input
            v-model="email"
            type="email"
            placeholder="doctor@clinic.com"
            class="w-full px-4 py-3 rounded-xl bg-cupertino-gray-50 border border-cupertino-gray-100 text-sm text-cupertino-gray-900 placeholder:text-cupertino-gray-300 focus:outline-none focus:ring-2 focus:ring-cupertino-blue/30 focus:border-cupertino-blue transition-all"
          />
        </FormField>

        <FormField label="Password">
          <div class="relative">
            <input
              v-model="password"
              :type="showPassword ? 'text' : 'password'"
              placeholder="Your password"
              class="w-full px-4 py-3 rounded-xl bg-cupertino-gray-50 border border-cupertino-gray-100 text-sm text-cupertino-gray-900 placeholder:text-cupertino-gray-300 focus:outline-none focus:ring-2 focus:ring-cupertino-blue/30 focus:border-cupertino-blue transition-all pr-10"
            />
            <button type="button" @click="showPassword = !showPassword" class="absolute right-3 top-1/2 -translate-y-1/2 text-cupertino-gray-400 hover:text-cupertino-gray-600">
              <svg v-if="!showPassword" viewBox="0 0 24 24" class="w-5 h-5" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <path d="M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z"/><circle cx="12" cy="12" r="3"/>
              </svg>
              <svg v-else viewBox="0 0 24 24" class="w-5 h-5" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <path d="M17.94 17.94A10.07 10.07 0 0 1 12 20c-7 0-11-8-11-8a18.45 18.45 0 0 1 5.06-5.94M9.9 4.24A9.12 9.12 0 0 1 12 4c7 0 11 8 11 8a18.5 18.5 0 0 1-2.16 3.19m-6.72-1.07a3 3 0 1 1-4.24-4.24"/>
                <line x1="1" y1="1" x2="23" y2="23"/>
              </svg>
            </button>
          </div>
        </FormField>

        <div v-if="auth.error" class="flex items-center gap-2 text-xs text-cupertino-red bg-cupertino-red/5 rounded-lg px-3 py-2">
          <svg viewBox="0 0 24 24" class="w-4 h-4 shrink-0" fill="currentColor"><path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm1 15h-2v-2h2v2zm0-4h-2V7h2v6z"/></svg>
          <span>{{ auth.error }}</span>
        </div>

        <button
          type="submit"
          :disabled="auth.loading"
          class="w-full py-3 rounded-xl bg-cupertino-blue text-white text-sm font-semibold shadow-sm shadow-cupertino-blue/20 hover:bg-cupertino-blue/90 active:scale-[0.98] transition-all disabled:opacity-50 disabled:cursor-not-allowed"
        >
          <span v-if="auth.loading" class="inline-flex items-center gap-2">
            <svg class="animate-spin h-4 w-4" viewBox="0 0 24 24" fill="none"><circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"/><path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z"/></svg>
            Signing in...
          </span>
          <span v-else>Sign In</span>
        </button>

        <div class="border-t border-cupertino-gray-100/60 pt-4">
          <button type="button" @click="showDemo = !showDemo" class="flex items-center gap-2 text-xs font-semibold text-cupertino-gray-400 hover:text-cupertino-gray-600 transition-colors w-full">
            <svg viewBox="0 0 24 24" class="w-4 h-4" :class="showDemo ? 'rotate-90' : ''" fill="currentColor"><path d="M10 6L8.59 7.41 13.17 12l-4.58 4.59L10 18l6-6z"/></svg>
            Quick login with demo accounts
          </button>
          <Transition name="demo">
            <div v-if="showDemo" class="mt-3 space-y-2">
              <button
                v-for="acc in demoAccounts"
                :key="acc.email"
                type="button"
                @click="selectDemo(acc)"
                class="w-full text-left px-3.5 py-2.5 rounded-xl bg-cupertino-gray-50 border border-cupertino-gray-100 hover:bg-cupertino-blue/5 hover:border-cupertino-blue/20 transition-all active:scale-[0.99]"
              >
                <p class="text-sm font-medium text-cupertino-gray-900">{{ acc.name }}</p>
                <p class="text-xs text-cupertino-gray-400 mt-0.5">{{ acc.email }} · {{ acc.specialization }}</p>
              </button>
            </div>
          </Transition>
        </div>
      </form>

      <p class="mt-6 text-center text-sm text-cupertino-gray-400">
        Don't have an account?
        <router-link to="/signup" class="text-cupertino-blue font-semibold hover:underline">Sign Up</router-link>
      </p>
    </div>
  </div>
</template>

<style scoped>
.demo-enter-active { transition: all 0.25s ease-out; }
.demo-leave-active { transition: all 0.2s ease-in; }
.demo-enter-from, .demo-leave-to { opacity: 0; transform: translateY(-8px); }
</style>
