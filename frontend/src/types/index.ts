export interface Doctor {
  id: number
  firstName: string
  lastName: string
  email: string
  role: string
  licenseNumber: string
  specialization: string
  createdAt: string
}

export interface Patient {
  ID: number
  firstName: string
  lastName: string
  idNumber: string
  phoneNumber: string
  gender: string
  age: number
  role: string
  dateOfBirth: string | null
  createdAt: string
  UpdatedAt: string | null
  DeletedAt: string | null
  patientPrograms: Program[] | null
}

export interface Program {
  ID: number
  program: string
  programCode: number
  patientsEnrolled?: Patient[]
}

export interface ApiResponse<T> {
  data: T
  error: string | boolean
  time?: string
}

export interface LoginResponse {
  accessToken: string
}

export interface AuthState {
  token: string | null
  user: { email: string; role: string } | null
}
