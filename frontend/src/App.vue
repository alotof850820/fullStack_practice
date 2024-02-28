<script setup lang="ts">
import { RouterView } from 'vue-router'
import ToastMessage from '@/components/ToastMessage.vue'
</script>

<template>
  <main class="main">
    <RouterView />
    <ToastMessage />
  </main>
</template>

<style>
body {
  box-sizing: border-box;
  margin: 0;
  padding: 0;
  font-size: 5vw;
}
main {
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
  align-items: center;
  position: relative;
}
</style>
