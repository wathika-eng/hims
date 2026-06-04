import { createRouter, createWebHistory } from 'vue-router'

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
        { path: 'enroll', name: 'enroll', component: () => import('../views/EnrollView.vue') },
      ],
    },
  ],
})

router.beforeEach((to, _from, next) => {
  const token = localStorage.getItem('token')
  if (to.meta.requiresAuth && !token) {
    next('/login')
  } else if (to.meta.guest && token) {
    next('/')
  } else {
    next()
  }
})

export default router
