<template>
  <div class="login_container">
    <h1>註冊帳號</h1>

    <form class="login_form" @submit.prevent="login">
      <div class="input">
        <label for="username">帳號:</label>
        <input v-model="username" type="text" id="username" required />
      </div>

      <div class="input">
        <label for="password">密碼:</label>
        <input v-model="password" type="password" id="password" required />
      </div>

      <div class="input">
        <label for="password">確認密碼:</label>
        <input v-model="rePassword" type="password" id="password" required />
      </div>

      <button class="btn" type="submit">註冊</button>
    </form>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { apiPostRegister } from '@/api'
import { go } from '@/router'

const username = ref('我是8號')
const password = ref('111111')
const rePassword = ref('111111')

const login = () => {
  apiPostRegister({
    name: username.value,
    password: password.value,
    repassword: rePassword.value
  }).then((res: any) => {
    if (res.data.code === 200) {
      go('Login')
    }
  })
}
</script>

<style scoped lang="scss">
.login_container {
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
  background-color: rgba(21, 164, 154, 0.508);
  align-items: center;
  .login_form {
    display: flex;
    flex-direction: column;
    gap: 3vw;
    padding: 5vw;
    background-color: rgb(31, 92, 118);
    margin-bottom: 5vw;
    .input {
      display: flex;
      flex-direction: column;
      gap: 1vw;

      label {
        font-size: 2.5vw;
        color: aliceblue;
      }

      input {
        background: none;
        border: none;
        border-bottom: 1px solid aliceblue;
        color: aliceblue;
        font-size: 2vw;
        padding: 0.5vw;
        outline: none;
        &:focus {
          border-bottom: 2px solid aliceblue;
        }
      }
    }
    .btn {
      font-size: 2vw;
      background: rgb(15, 148, 189);
      border: none;
      color: aliceblue;
      padding: 0.5vw;
      border: 1px solid aliceblue;
      &:hover {
        background-color: aliceblue;
        color: rgb(31, 92, 118);
        cursor: pointer;
      }
    }
  }
}
</style>
