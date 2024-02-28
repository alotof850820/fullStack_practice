<script setup lang="ts">
import { ref } from 'vue'
import { apiPostSearchFriends } from '@/api'
import { useUserStore } from '@/stores/user'
import { Icon } from '@iconify/vue'
import { go } from '@/router'

const user = useUserStore()

const friend = ref<any[]>([])
const tab = ref<number>(0)

const searchFriends = () => {
  const userId = user.userData.ID
  apiPostSearchFriends({
    userId
  }).then((res: any) => {
    friend.value = res.data.Data
  })
}
const setTab = (type: number) => {
  tab.value = type
  if (type === 0) {
    searchFriends()
  } else if (type === 1) {
    //
  } else if (type === 2) {
    //
  }
}
setTab(0)
</script>

<template>
  <div class="chat-room">
    <div class="title">聯絡人</div>
    <div class="contact_box">
      <div
        class="friend"
        v-for="item in friend"
        :key="item.ID"
        @click="go('ChatRoom', { id: item.ID, name: item.Name })"
      >
        <img class="avatar" src="https://cdn-icons-png.flaticon.com/512/149/149071.png" alt="" />
        <div class="name">{{ item.Name }}</div>
        <Icon icon="flowbite:angle-right-outline" />
      </div>
    </div>

    <div class="tabs">
      <div class="tab" :class="{ active: tab === 0 }" @click="setTab(0)">好友</div>
      <div class="tab" :class="{ active: tab === 1 }" @click="setTab(1)">群聊</div>
      <div class="tab" :class="{ active: tab === 2 }" @click="setTab(2)">我的</div>
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
    padding: 2vw;
    text-align: center;
  }
  .contact_box {
    flex: 1;
    border: 1px solid #ccc;
    padding: 2vw;
    overflow-y: auto;
    cursor: pointer;

    .friend {
      display: flex;
      align-items: center;
      justify-content: space-between;
      border-bottom: 1px solid #ccc;
      padding: 3vw;
      .avatar {
        width: 10vw;
        height: 10vw;
        object-fit: cover;
        border-radius: 50%;
      }
    }
  }

  .tabs {
    display: flex;
    justify-content: space-around;
    background-color: #f1f1f1;
    padding: 2vw;

    .tab {
      cursor: pointer;

      &.active {
        color: #4caf50;
        font-weight: bold;
      }
    }
  }
}
</style>
