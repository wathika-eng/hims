export function required(value: unknown, label: string): string | null {
  if (value === null || value === undefined) return `${label} is required`
  if (typeof value === 'string' && !value.trim()) return `${label} is required`
  if (typeof value === 'number' && value === 0) return `${label} is required`
  return null
}

export function minLength(value: string, min: number, label: string): string | null {
  if (!value || value.trim().length < min) return `${label} must be at least ${min} characters`
  return null
}

export function lettersOnly(value: string, label: string): string | null {
  if (!value) return null
  if (!/^[A-Za-z\s]+$/.test(value.trim())) return `${label} must contain only letters`
  return null
}

export function alphanumeric(value: string, label: string): string | null {
  if (!value) return null
  if (!/^[A-Za-z0-9]+$/.test(value.trim())) return `${label} must contain only letters and numbers`
  return null
}

export function validEmail(value: string): string | null {
  if (!value) return null
  if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(value.trim())) return 'Invalid email format'
  return null
}

export function validKenyanPhone(value: string): string | null {
  if (!value) return null
  const cleaned = value.trim().replace(/[\s-]/g, '')
  if (!/^(07\d{8}|\+?2547\d{8})$/.test(cleaned)) return 'Must be a valid Kenyan phone number (e.g. 0712345678 or +254712345678)'
  return null
}

export function inRange(value: number, min: number, max: number, label: string): string | null {
  if (value === undefined || value === null || value < min || value > max) return `${label} must be between ${min} and ${max}`
  return null
}

type ValidationRule = () => string | null

export function validateAll(rules: Record<string, ValidationRule>): Record<string, string> {
  const errors: Record<string, string> = {}
  for (const [field, rule] of Object.entries(rules)) {
    const err = rule()
    if (err) errors[field] = err
  }
  return errors
}
