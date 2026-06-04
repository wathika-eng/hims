export function maskPhone(phone: string): string {
  if (!phone || phone.length < 7) return phone
  const visible = 3
  const prefix = phone.slice(0, visible)
  const suffix = phone.slice(-visible)
  return `${prefix}****${suffix}`
}

export function maskID(id: string): string {
  if (!id || id.length < 5) return id
  const prefix = id.slice(0, 2)
  const suffix = id.slice(-2)
  return `${prefix}****${suffix}`
}
