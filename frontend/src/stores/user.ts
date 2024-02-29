import { ref } from 'vue'
import { defineStore } from 'pinia'

export const useUserStore = defineStore('user', () => {
  const userData = ref(JSON.parse(localStorage.getItem('user') || '{}'))
  const friends = ref<any>([])

  return { userData, friends }
})
