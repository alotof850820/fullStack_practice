<template>
  <div class="login_container">
    <div class="form_box">
      <h3>登入</h3>
      <form class="login_form" @submit.prevent="login">
        <div class="input">
          <label for="username">帳號:</label>
          <input v-model="username" type="text" id="username" required />
        </div>

        <div class="input">
          <label for="password">密碼:</label>
          <input v-model="password" type="password" id="password" required />
        </div>

        <button class="btn" type="submit">登入</button>
      </form>

      <div class="footer">
        <p class="register" @click="go('Register')">註冊帳號</p>

        <p class="register">忘記密碼</p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { apiPostlogin } from '@/api'
import { go } from '@/router'
import { useToastStore } from '@/stores/toast'

const username = ref('我是9號')
const password = ref('111111')

const toast = useToastStore()

const login = () => {
  apiPostlogin({
    name: username.value,
    password: password.value
  })
    .then((res: any) => {
      localStorage.setItem('user', JSON.stringify(res.data.data))
      toast.showToast(res.data.message)
      go('ToChat')
    })
    .catch((err: any) => {
      toast.showToast(err)
    })
}
</script>

<style scoped lang="scss">
.login_container {
  width: 100%;
  height: 100vh;
  display: flex;
  flex-direction: column;
  background-color: rgba(21, 164, 154, 0.508);
  align-items: center;

  .form_box {
    flex: 1;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    .login_form {
      display: flex;
      flex-direction: column;
      gap: 3vw;
      padding: 5vw;
      background-color: rgb(31, 92, 118);

      justify-content: center;

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

  .footer {
    display: flex;
    gap: 3vw;

    .register {
      font-size: 2vw;
      color: rgb(31, 92, 118);
      margin-top: 2vw;
      background-color: aliceblue;
      padding: 1vw;
      border-radius: 0.5vw;
      cursor: pointer;

      &:hover {
        background-color: rgb(31, 92, 118);
        color: aliceblue;

        transition: 0.5s;
        transform: scale(1.1);
        box-shadow: 0 0 10px rgb(31, 92, 118);
        text-shadow: 0 0 10px aliceblue;
      }
    }
  }
}
</style>
