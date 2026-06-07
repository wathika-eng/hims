import axios from 'axios'
import router from '../router'
import type { ApiResponse, LoginResponse, Patient, Program } from '../types'

function extractError(err: any): string {
  if (axios.isCancel(err)) return ''
  const data = err.response?.data
  if (!data) return 'Network error — server may be unreachable'

  if (typeof data.error === 'string' && data.error !== 'false') return data.error
  if (typeof data.data === 'string') return data.data
  if (typeof data.message === 'string') return data.message

  if (Array.isArray(data.data)) {
    return data.data.map((e: any) => `${e.failedField}: ${e.tag}`).join(', ')
  }

  if (err.response?.status === 401) return 'Session expired. Please sign in again.'
  if (err.response?.status === 403) return 'You do not have permission to perform this action.'
  if (err.response?.status >= 500) return 'Server error. Please try again later.'

  return 'An unexpected error occurred'
}

const api = axios.create({
  baseURL: '/api/v1',
  headers: { 'Content-Type': 'application/json' },
  timeout: 15000,
})

api.interceptors.request.use((config) => {
  const token = localStorage.getItem('token')
  if (token) {
    config.headers.Authorization = `Bearer ${token}`
  }
  return config
})

api.interceptors.response.use(
  (res) => res,
  (err) => {
    if (err.response?.status === 401) {
      localStorage.removeItem('token')
      router.push('/login')
    }
    err.userMessage = extractError(err)
    return Promise.reject(err)
  },
)

export const authApi = {
  signup(body: { firstName: string; lastName: string; email: string; password: string; licenseNumber: string; specialization: string }) {
    return api.post<ApiResponse<string>>('/auth/signup', body)
  },
  login(body: { email: string; password: string }) {
    return api.post<LoginResponse>('/auth/login', body)
  },
}

export const patientApi = {
  list() {
    return api.get<ApiResponse<Patient[]>>('/protected/patients')
  },
  get(id: number) {
    return api.get<ApiResponse<Patient>>(`/protected/patients/${id}`)
  },
  create(body: { firstName: string; lastName: string; idNumber: string; phoneNumber: string; gender: string; age: number }) {
    return api.post<ApiResponse<string>>('/protected/patients', body)
  },
}

export const programApi = {
  list() {
    return api.get<ApiResponse<Program[]>>('/protected/programs')
  },
  create(body: { program: string; programCode: number }) {
    return api.post<ApiResponse<string>>('/protected/programs', body)
  },
  enroll(body: { programName: string; patientID: string }) {
    return api.post<ApiResponse<Patient>>('/protected/patients/enroll', body)
  },
}

export async function downloadPdf() {
  const token = localStorage.getItem('token')
  const res = await fetch('/api/v1/protected/programs/pdf', {
    headers: token ? { Authorization: `Bearer ${token}` } : {},
  })
  if (!res.ok) {
    const body = await res.json().catch(() => ({}))
    throw new Error(body.message || body.error || 'Failed to download PDF')
  }
  const blob = await res.blob()
  const url = URL.createObjectURL(blob)
  const a = document.createElement('a')
  a.href = url
  a.download = 'programs_summary.pdf'
  document.body.appendChild(a)
  a.click()
  document.body.removeChild(a)
  URL.revokeObjectURL(url)
}

export default api
