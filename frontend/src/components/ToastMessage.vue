<template>
  <div v-if="toast.isVisible" class="toast">{{ toast.message }}</div>
</template>

<script setup lang="ts">
import { useToastStore } from '@/stores/toast'

const toast = useToastStore()
</script>

<style scoped lang="scss">
.toast {
  position: absolute;
  bottom: 3vw;
  left: 50%;
  transform: translateX(-50%);
  padding: 1vw 3vw;
  background-color: #333;
  color: #fff;
  border-radius: 5px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.2);
  z-index: 99999999;
  font-size: 3vw;
}
</style>
