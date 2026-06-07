import { createRouter, createWebHistory } from 'vue-router'

function isTokenExpired(token: string): boolean {
  try {
    const payload = JSON.parse(atob(token.split('.')[1]))
    if (!payload.exp) return false
    return Date.now() >= payload.exp * 1000
  } catch {
    return true
  }
}

function clearSession() {
  localStorage.removeItem('token')
}

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/login',
      name: 'login',
      component: () => import('../views/LoginView.vue'),
      meta: { guest: true },
    },
    {
      path: '/signup',
      name: 'signup',
      component: () => import('../views/SignupView.vue'),
      meta: { guest: true },
    },
    {
      path: '/',
      component: () => import('../layouts/DefaultLayout.vue'),
      meta: { requiresAuth: true },
      children: [
        { path: '', name: 'dashboard', component: () => import('../views/DashboardView.vue') },
        { path: 'patients', name: 'patients', component: () => import('../views/PatientsView.vue') },
        { path: 'patients/new', name: 'new-patient', component: () => import('../views/NewPatientView.vue') },
        { path: 'patients/:id', name: 'patient-profile', component: () => import('../views/PatientProfileView.vue') },
        { path: 'programs', name: 'programs', component: () => import('../views/ProgramsView.vue') },
        { path: 'programs/new', name: 'new-program', component: () => import('../views/NewProgramView.vue') },
        { path: 'programs/:id', name: 'program-detail', component: () => import('../views/ProgramDetailView.vue') },
        { path: 'enroll', name: 'enroll', component: () => import('../views/EnrollView.vue') },
      ],
    },
  ],
})

router.beforeEach((to, _from, next) => {
  const token = localStorage.getItem('token')

  if (token && isTokenExpired(token)) {
    clearSession()
    return next('/login')
  }

  if (to.meta.requiresAuth && !token) {
    return next('/login')
  }

  if (to.meta.guest && token) {
    return next('/')
  }

  next()
})

export default router
