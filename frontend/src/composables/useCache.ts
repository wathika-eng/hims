import { ref } from 'vue'

interface CacheEntry<T> {
  data: T
  timestamp: number
  ttl: number
}

const DB_NAME = 'hims-cache'
const DB_VERSION = 1
const STORE_NAME = 'api-cache'

let dbPromise: Promise<IDBDatabase> | null = null

function openDB(): Promise<IDBDatabase> {
  if (dbPromise) return dbPromise
  dbPromise = new Promise((resolve, reject) => {
    const req = indexedDB.open(DB_NAME, DB_VERSION)
    req.onupgradeneeded = () => {
      req.result.createObjectStore(STORE_NAME)
    }
    req.onsuccess = () => resolve(req.result)
    req.onerror = () => reject(req.error)
  })
  return dbPromise
}

async function getCache<T>(key: string): Promise<T | null> {
  try {
    const db = await openDB()
    const tx = db.transaction(STORE_NAME, 'readonly')
    const store = tx.objectStore(STORE_NAME)
    const req = store.get(key)
    return new Promise((resolve) => {
      req.onsuccess = () => {
        const entry = req.result as CacheEntry<T> | undefined
        if (!entry) return resolve(null)
        if (Date.now() - entry.timestamp > entry.ttl) {
          tx.commit()
          deleteKey(key)
          return resolve(null)
        }
        resolve(entry.data)
      }
      req.onerror = () => resolve(null)
    })
  } catch {
    return null
  }
}

async function setCache<T>(key: string, data: T, ttl: number): Promise<void> {
  try {
    const db = await openDB()
    const tx = db.transaction(STORE_NAME, 'readwrite')
    const store = tx.objectStore(STORE_NAME)
    const entry: CacheEntry<T> = { data, timestamp: Date.now(), ttl }
    store.put(entry, key)
  } catch {
  }
}

async function deleteKey(key: string): Promise<void> {
  try {
    const db = await openDB()
    const tx = db.transaction(STORE_NAME, 'readwrite')
    const store = tx.objectStore(STORE_NAME)
    store.delete(key)
  } catch {
  }
}

export async function clearAllCache(): Promise<void> {
  try {
    const db = await openDB()
    const tx = db.transaction(STORE_NAME, 'readwrite')
    const store = tx.objectStore(STORE_NAME)
    store.clear()
  } catch {
  }
}

export function useCache() {
  const loading = ref(false)

  async function fetchWithCache<T>(
    key: string,
    fetcher: () => Promise<T>,
    ttl = 5 * 60 * 1000,
  ): Promise<{ data: T; fromCache: boolean }> {
    const cached = await getCache<T>(key)
    if (cached) {
      return { data: cached, fromCache: true }
    }

    loading.value = true
    try {
      const data = await fetcher()
      await setCache(key, data, ttl)
      return { data, fromCache: false }
    } finally {
      loading.value = false
    }
  }

  async function invalidate(key: string) {
    await deleteKey(key)
  }

  async function invalidateAll() {
    await clearAllCache()
  }

  return { loading, fetchWithCache, invalidate, invalidateAll }
}
