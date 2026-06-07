<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { programApi } from '../api'
import { useCache } from '../composables/useCache'
import type { Program, Patient } from '../types'
import { maskID, maskPhone, maskName } from '../utils/mask'
import ErrorAlert from '../components/ErrorAlert.vue'

const route = useRoute()
const router = useRouter()
const { fetchWithCache, loading } = useCache()

const program = ref<Program | null>(null)
const fetchError = ref<string | null>(null)
const initialLoading = ref(true)

const enrolledPatients = computed<Patient[]>(() => program.value?.patientsEnrolled || [])

onMounted(async () => {
  const id = Number(route.params.id)
  try {
    const { data: programs } = await fetchWithCache(
      'programs',
      () => programApi.list().then(r => r.data.data),
      2 * 60 * 1000,
    )
    program.value = programs.find(p => p.ID === id) || null
    if (!program.value) {
      fetchError.value = 'Program not found'
    }
  } catch (e: any) {
    fetchError.value = e.userMessage || 'Failed to load program'
  } finally {
    initialLoading.value = false
  }
})
</script>

<template>
  <div>
    <button
      @click="router.back()"
      class="inline-flex items-center gap-1.5 text-sm text-cupertino-blue font-medium mb-4 hover:underline"
    >
      <svg viewBox="0 0 24 24" class="w-4 h-4" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"><polyline points="15 18 9 12 15 6"/></svg>
      Back
    </button>

    <ErrorAlert :message="fetchError" title="Error" />

    <div v-if="initialLoading" class="animate-pulse space-y-4">
      <div class="h-8 w-48 bg-cupertino-gray-100 rounded-lg" />
      <div class="h-4 w-32 bg-cupertino-gray-50 rounded" />
      <div class="mt-6 space-y-3">
        <div v-for="i in 4" :key="i" class="flex items-center gap-4 px-4 py-3 bg-white/70 rounded-2xl border border-cupertino-gray-100/60">
          <div class="w-10 h-10 rounded-full bg-cupertino-gray-100" />
          <div class="flex-1 space-y-1.5">
            <div class="h-3 w-36 bg-cupertino-gray-100 rounded" />
            <div class="h-2 w-20 bg-cupertino-gray-50 rounded" />
          </div>
        </div>
      </div>
    </div>

    <template v-else-if="program">
      <div class="flex items-center gap-4 mb-6 sm:mb-8">
        <div class="w-12 h-12 rounded-2xl bg-cupertino-green/10 text-cupertino-green flex items-center justify-center text-lg font-bold">
          {{ program.programCode }}
        </div>
        <div>
          <h1 class="text-2xl font-bold text-cupertino-gray-900">{{ program.program }}</h1>
          <p class="text-sm text-cupertino-gray-400">Code: #{{ program.programCode }} &middot; {{ enrolledPatients.length }} enrolled</p>
        </div>
      </div>

      <div class="bg-white/70 backdrop-blur-2xl rounded-2xl border border-cupertino-gray-100/60 shadow-sm overflow-hidden">
        <div v-if="enrolledPatients.length === 0" class="text-center py-16 text-cupertino-gray-400">
          <svg viewBox="0 0 24 24" class="w-12 h-12 mx-auto mb-3 text-cupertino-gray-200" fill="currentColor"><path d="M12 12c2.21 0 4-1.79 4-4s-1.79-4-4-4-4 1.79-4 4 1.79 4 4 4zm-8 6c0-2.67 5.33-4 8-4s8 1.33 8 4v2H4v-2z"/></svg>
          <p class="text-sm">No patients enrolled in this program</p>
        </div>

        <template v-else>
          <!-- Desktop table -->
          <div class="hidden sm:block overflow-x-auto">
            <table class="w-full">
              <thead>
                <tr class="border-b border-cupertino-gray-100/60">
                  <th class="text-left px-5 py-3 text-xs font-semibold text-cupertino-gray-400 uppercase tracking-wide">Patient</th>
                  <th class="text-left px-5 py-3 text-xs font-semibold text-cupertino-gray-400 uppercase tracking-wide">ID No.</th>
                  <th class="text-left px-5 py-3 text-xs font-semibold text-cupertino-gray-400 uppercase tracking-wide">Phone</th>
                  <th class="text-left px-5 py-3 text-xs font-semibold text-cupertino-gray-400 uppercase tracking-wide">Gender</th>
                  <th class="text-left px-5 py-3 text-xs font-semibold text-cupertino-gray-400 uppercase tracking-wide">Age</th>
                  <th class="text-right px-5 py-3"></th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="p in enrolledPatients" :key="p.ID" class="border-b border-cupertino-gray-100/40 hover:bg-cupertino-gray-50/50 transition-colors">
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
                  <td class="px-5 py-3 text-right">
                    <router-link :to="`/patients/${p.ID}`" class="text-xs font-semibold text-cupertino-blue hover:underline">View</router-link>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>

          <!-- Mobile cards -->
          <div class="sm:hidden space-y-2 p-4">
            <router-link
              v-for="p in enrolledPatients"
              :key="p.ID"
              :to="`/patients/${p.ID}`"
              class="flex items-center gap-3 px-3 py-3 rounded-xl hover:bg-cupertino-gray-50 transition-colors active:scale-[0.99]"
            >
              <div class="w-9 h-9 rounded-full bg-cupertino-blue/10 text-cupertino-blue flex items-center justify-center text-xs font-semibold">
                {{ p.firstName[0] }}{{ p.lastName[0] }}
              </div>
              <div class="min-w-0 flex-1">
                <p class="text-sm font-medium text-cupertino-gray-900 truncate">{{ maskName(p.firstName, p.lastName) }}</p>
                <p class="text-xs text-cupertino-gray-400 font-mono">{{ maskID(p.idNumber) }}</p>
              </div>
              <div class="flex items-center gap-2 text-xs text-cupertino-gray-500">
                <span class="px-2 py-0.5 rounded-md font-medium" :class="p.gender === 'male' ? 'bg-cupertino-blue/10 text-cupertino-blue' : 'bg-cupertino-purple/10 text-cupertino-purple'">
                  {{ p.gender }}, {{ p.age }}
                </span>
                <svg viewBox="0 0 24 24" class="w-4 h-4 text-cupertino-gray-300" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polyline points="9 18 15 12 9 6"/></svg>
              </div>
            </router-link>
          </div>
        </template>
      </div>

      <p v-if="loading && !initialLoading" class="mt-4 text-xs text-cupertino-gray-400 flex items-center gap-2">
        <svg class="animate-spin h-3 w-3" viewBox="0 0 24 24" fill="none"><circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"/><path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z"/></svg>
        Refreshing data...
      </p>
    </template>
  </div>
</template>
