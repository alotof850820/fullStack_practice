import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useToastStore = defineStore('toast', () => {
  const isVisible = ref(false)
  const message = ref('')

  const showToast = (msg: string) => {
    message.value = msg
    isVisible.value = true

    setTimeout(() => {
      isVisible.value = false
    }, 3000) // 3秒後隱藏 Toast，你可以調整這個時間
  }

  return { isVisible, message, showToast }
})
