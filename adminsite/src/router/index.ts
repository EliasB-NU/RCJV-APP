import { createRouter, createWebHashHistory } from 'vue-router'

const router = createRouter({
  history: createWebHashHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: () => import('@/views/HomeView.vue'),
    },
    {
      path: '/login',
      name: 'login',
      component: () => import('@/views/LoginView.vue'),
    },
    {
      path: '/teams',
      name: 'Teams & Institutions',
      component: () => import('@/views/Teams&InstitutionsView.vue'),
    },
    {
      path: '/fields',
      name: 'Fields & ODS',
      component: () => import('@/views/Fields&ODSView.vue')
    }
  ],
})

export default router
