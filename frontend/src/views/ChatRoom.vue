<script setup lang="ts">
import { useUserStore } from '@/stores/user'
import { onBeforeUnmount, onMounted, ref } from 'vue'
import { useRoute } from 'vue-router'
import { Icon } from '@iconify/vue'
import { goBack } from '@/router'
import { GIF } from '@/assets/gif'
import { apiPostUploadImg } from '@/api'

const route = useRoute()

const user = useUserStore()

const showGif = ref<boolean>(false)
const showUpload = ref<boolean>(false)
const imgInput = ref<HTMLInputElement | null>(null)
const allMessage = ref<any[]>([])
const newMessage = ref<string>('')
const socket = ref<WebSocket | null>(null)
const enterMsg = ref<string>('')
const socketUrl = `ws://127.0.0.1:8081/user/sendUserMsg?userId=${user.userData.ID}&token=`

const sendMessage = () => {
  if (socket.value && socket.value.readyState === WebSocket.OPEN) {
    const data = {
      TargetId: +route.params.id, //接收者
      Type: 1, //消息類型 群聊 私聊 廣播
      Media: 1, //消息類型 文字 圖片 音訊
      userId: user.userData.ID, //发送者
      Content: newMessage.value, //消息內容
      Url: ''
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
const sendGif = (url: string) => {
  if (socket.value && socket.value.readyState === WebSocket.OPEN) {
    const data = {
      TargetId: +route.params.id, //接收者
      Type: 1, //消息類型 群聊 私聊 廣播
      Media: 4, //消息類型 文字 圖片 音訊
      userId: user.userData.ID, //发送者
      Content: newMessage.value, //消息內容
      Url: url
    }
    socket.value.send(JSON.stringify(data))
    showGif.value = false
  }
}
const uploadImg = () => {
  // showUpload.value = !showUpload.value
  imgInput.value?.click()
}
const handleFileChange = (event: any) => {
  const formData = new FormData()

  // 将选择的文件添加到 FormData 中
  const file = event.target?.files[0]
  formData.append('file', file)
  apiPostUploadImg(formData)
    .then((res) => {
      // ./assets/upload/17091291611076632673.png
      const data = {
        TargetId: +route.params.id, //接收者
        Type: 1, //消息類型 群聊 私聊 廣播
        Media: 2, //消息類型 文字 圖片 音訊 gif
        userId: user.userData.ID, //发送者
        Content: newMessage.value, //消息內容
        url: res.data.Data
      }
      socket.value?.send(JSON.stringify(data))
      showUpload.value = false
    })
    .catch((error) => {
      console.error(error)
    })
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
        <span> {{ msg.userId === user.userData.ID ? '我' : route.params.name }}: </span>
        <span> {{ msg.Content }}</span>
        <img v-show="msg.Media === 4" :src="`${msg.url}`" alt="" />
        <img class="img" v-show="msg.Media === 2" :src="`${msg.url}`" alt="" />
      </div>
    </div>
    <div>
      <div v-show="showGif" class="gif_box">
        <img
          @click="sendGif(gif)"
          class="img"
          v-for="(gif, index) in GIF"
          :key="index"
          :src="gif"
          alt=""
        />
      </div>
      <div v-show="showUpload" class="upload_box">
        <Icon @click="uploadImg" class="upload" icon="material-symbols:image-outline" />
        <input ref="imgInput" type="file" style="display: none" @change="handleFileChange" />
      </div>
      <div class="input-box">
        <Icon @click="showGif = !showGif" class="gif" icon="ant-design:smile-outlined" />
        <Icon @click="showUpload = !showUpload" class="gif" icon="icons8:plus" />
        <input v-model="newMessage" @keyup.enter="sendMessage" placeholder="Type your message..." />
        <button @click="sendMessage">Send</button>
      </div>
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
      padding: 1vw;
      margin-bottom: 1vw;
      font-size: 4vw;
      display: flex;
      align-items: center;

      &.left {
        justify-content: flex-end;
      }

      &.right {
        justify-content: flex-start;
      }

      img {
        width: 4vw;
        height: 4vw;
      }
      .img {
        width: 40vw;
        height: auto;
        object-fit: cover;
      }
    }
  }

  .input-box {
    display: flex;
    align-items: center;
    padding: 2vw;
    background-color: #f1f1f1;
    gap: 2vw;

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
    .gif {
      cursor: pointer;
    }
  }
  .gif_box {
    display: flex;
    flex-wrap: wrap;
    padding: 1vw;
    gap: 1vw;
    .img {
      width: 4vw;
      height: 4vw;
      cursor: pointer;
    }
  }
  .upload_box {
    display: flex;
    flex-wrap: wrap;
    padding: 1vw;
    gap: 1vw;
    .upload {
      cursor: pointer;
      width: 4vw;
      height: 4vw;
    }
  }
}
</style>
