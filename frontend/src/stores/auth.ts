import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { authApi } from '../api'
import { useRouter } from 'vue-router'
import { clearAllCache } from '../composables/useCache'

export const useAuthStore = defineStore('auth', () => {
  const token = ref<string | null>(localStorage.getItem('token'))
  const error = ref<string | null>(null)
  const loading = ref(false)
  const router = useRouter()

  const isAuthenticated = computed(() => !!token.value)

  async function login(email: string, password: string) {
    loading.value = true
    error.value = null
    try {
      const res = await authApi.login({ email, password })
      token.value = res.data.accessToken
      localStorage.setItem('token', res.data.accessToken)
      await router.push('/')
    } catch (e: any) {
      error.value = e.userMessage || 'Login failed'
    } finally {
      loading.value = false
    }
  }

  async function signup(body: { firstName: string; lastName: string; email: string; password: string; licenseNumber: string; specialization: string }) {
    loading.value = true
    error.value = null
    try {
      await authApi.signup(body)
      await login(body.email, body.password)
    } catch (e: any) {
      error.value = e.userMessage || 'Signup failed'
    } finally {
      loading.value = false
    }
  }

  function clearError() {
    error.value = null
  }

  function logout() {
    token.value = null
    localStorage.removeItem('token')
    clearAllCache()
    router.push('/login')
  }

  return { token, error, loading, isAuthenticated, login, signup, logout, clearError }
})
