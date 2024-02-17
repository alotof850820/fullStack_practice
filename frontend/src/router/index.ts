import { createRouter, createWebHistory } from 'vue-router'
import Login from '@/views/LoginView.vue'
import Register from '@/views/RegisterView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'Login',
      component: Login
    },
    {
      path: '/register',
      name: 'Register',
      component: Register
    }
  ]
})

export const go = (name: string) => {
  router.push({ name })
}

export default router
