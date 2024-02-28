<script setup lang="ts">
import { useUserStore } from '@/stores/user'
import { onBeforeUnmount, onMounted, ref } from 'vue'
import { useRoute } from 'vue-router'
import { Icon } from '@iconify/vue'
import { goBack } from '@/router'

const route = useRoute()

const user = useUserStore()

const allMessage = ref<any[]>([])
const newMessage = ref<string>('')
const socket = ref<WebSocket | null>(null)
const enterMsg = ref<string>('')
const socketUrl = `ws://127.0.0.1:8081/user/sendUserMsg?userId=${user.userData.ID}&token=${user.userData.Identity}`

const sendMessage = () => {
  if (socket.value && socket.value.readyState === WebSocket.OPEN) {
    const data = {
      TargetId: +route.params.id, //接收者
      Type: 1, //消息類型 群聊 私聊 廣播
      Media: 1, //消息類型 文字 圖片 音訊
      userId: user.userData.ID, //发送者
      Content: newMessage.value //消息內容
    }
    socket.value.send(JSON.stringify(data))
  }
  newMessage.value = ''
}

// 打开 WebSocket 连接
const openWebSocket = () => {
  socket.value = new WebSocket(socketUrl)
  socket.value?.addEventListener('message', (event) => {
    try {
      const data = JSON.parse(event.data)
      allMessage.value.push(data)
    } catch (error) {
      enterMsg.value = event.data
    }
  })
}

// 关闭 WebSocket 连接
const closeWebSocket = () => {
  if (socket.value) {
    socket.value.close()
    socket.value = null
  }
}

// 在组件挂载时打开 WebSocket 连接
onMounted(() => {
  openWebSocket()
})

// 在组件销毁前关闭 WebSocket 连接
onBeforeUnmount(() => {
  closeWebSocket()
})
</script>

<template>
  <div class="chat-room">
    <div class="title">
      <Icon class="goBack" icon="ep:back" @click="goBack" />
      <div class="name">與"{{ route.params.name }}"聊天</div>
      <Icon class="goBack" icon="" />
    </div>
    <div class="contact_box">
      <div class="enter">{{ enterMsg }}</div>
      <div
        v-for="(msg, index) in allMessage"
        :key="index"
        class="message"
        :class="{ right: msg.userId !== user.userData.ID, left: msg.userId === user.userData.ID }"
      >
        <span>{{ msg.userId === user.userData.ID ? '我' : route.params.name }}:</span>
        {{ msg.Content }}
      </div>
    </div>

    <div class="input-box">
      <input v-model="newMessage" @keyup.enter="sendMessage" placeholder="Type your message..." />
      <button @click="sendMessage">Send</button>
    </div>
  </div>
</template>

<style scoped lang="scss">
.chat-room {
  display: flex;
  flex-direction: column;
  height: 100vh;
  width: 100vw;

  .title {
    background-color: #4caf50;
    color: white;
    text-align: center;

    display: flex;
    align-items: center;
    justify-content: space-between;
    .name {
      flex: 1;
    }
    .goBack {
      cursor: pointer;
      margin-right: auto;
      padding: 4vw;
    }
  }
  .contact_box {
    flex: 1;
    border: 1px solid #ccc;
    padding: 2vw;
    overflow-y: auto;
    .enter {
      text-align: center;
    }

    .message {
      padding: 5px;
      margin-bottom: 5px;

      &.left {
        text-align: end;
      }

      &.right {
        text-align: start;
      }
    }
  }

  .input-box {
    display: flex;
    align-items: center;
    padding: 2vw;
    background-color: #f1f1f1;

    input {
      flex: 1;
      padding: 1vw;
      margin-right: 1.6vw;
    }

    button {
      padding: 1.2vw 2vw;
      background-color: #4caf50;
      color: white;
      border: none;
      cursor: pointer;
    }
  }
}
</style>
