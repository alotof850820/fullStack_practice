import { createRouter, createWebHistory } from 'vue-router'
import Login from '@/views/LoginView.vue'
import Register from '@/views/RegisterView.vue'
import ToChat from '@/views/ToChatView.vue'
import chatRoom from '@/views/ChatRoom.vue'

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
    },
    {
      path: '/toChat',
      name: 'ToChat',
      component: ToChat
    },
    {
      path: '/chatRoom/:id/:name',
      name: 'ChatRoom',
      component: chatRoom
    }
  ]
})

export const go = (name: string, params: any) => {
  router.push({ name, params })
}
export const goBack = () => {
  router.go(-1)
}

export default router
