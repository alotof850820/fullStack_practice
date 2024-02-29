<script setup lang="ts">
import { ref } from 'vue'
import { apiPostSearchFriends } from '@/api'
import { useUserStore } from '@/stores/user'
import { Icon } from '@iconify/vue'
import { go } from '@/router'
import { apiPostaddFriend } from '@/api'
import { useToastStore } from '@/stores/toast'

const user = useUserStore()
const toast = useToastStore()

const tab = ref<number>(0)
const showModal = ref<boolean>(false)
const friendId = ref<number>(0)

const searchFriends = () => {
  const userId = user.userData.ID
  apiPostSearchFriends({
    userId
  }).then((res: any) => {
    user.friends = res.data.Data
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
const addFriend = () => {
  if (+friendId.value === user.userData.ID) {
    toast.showToast('不能添加自己')
    return
  }
  if (user.friends.map((item: any) => item.ID).includes(+friendId.value)) {
    toast.showToast('已經是好友了')
    return
  }
  apiPostaddFriend({
    userId: user.userData.ID,
    targetId: friendId.value
  })
    .then((res: any) => {
      toast.showToast(res.data.Msg)
      showModal.value = false
    })
    .catch((err: any) => {
      toast.showToast(err)
    })
}
setTab(0)
</script>

<template>
  <div class="chat-room">
    <div class="title">聯絡人</div>
    <div v-if="tab === 0" class="contact_box">
      <div
        class="friend"
        v-for="item in user.friends"
        :key="item.ID"
        @click="go('ChatRoom', { id: item.ID, name: item.Name })"
      >
        <img class="avatar" src="https://cdn-icons-png.flaticon.com/512/149/149071.png" alt="" />
        <div class="name">{{ item.Name }}</div>
        <Icon icon="flowbite:angle-right-outline" />
      </div>
    </div>
    <div v-if="tab === 1" class="contact_box"></div>
    <div v-if="tab === 2" class="contact_box" @click="showModal = true">
      <div class="list_box">添加好友</div>
    </div>

    <div class="tabs">
      <div class="tab" :class="{ active: tab === 0 }" @click="setTab(0)">好友</div>
      <div class="tab" :class="{ active: tab === 1 }" @click="setTab(1)">群聊</div>
      <div class="tab" :class="{ active: tab === 2 }" @click="setTab(2)">我的</div>
    </div>

    <!-- modal -->
    <div v-if="showModal" class="modal">
      <div class="modal-content">
        <span class="close" @click="showModal = false">&times;</span>
        <label for="friendId">好友ID:</label>
        <input v-model="friendId" type="text" id="friendId" placeholder="请输入好友ID" />
        <button @click="addFriend">添加</button>
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
    padding: 2vw;
    text-align: center;
  }
  .contact_box {
    flex: 1;
    border: 1px solid #ccc;
    padding: 2vw;
    overflow-y: auto;

    .friend {
      display: flex;
      align-items: center;
      justify-content: space-between;
      border-bottom: 1px solid #ccc;
      padding: 3vw;
      cursor: pointer;
      .avatar {
        width: 10vw;
        height: 10vw;
        object-fit: cover;
        border-radius: 50%;
      }
    }
    .list_box {
      padding: 3vw;
      border-bottom: 1px solid #ccc;
      cursor: pointer;
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

  .modal {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background-color: rgba(0, 0, 0, 0.7);
    display: flex;
    align-items: center;

    .modal-content {
      background-color: #fefefe;
      margin: 15% auto;
      padding: 20px;
      border: 1px solid #888;
      width: 40%;

      .close {
        color: #aaa;
        float: right;
        font-size: 28px;
        font-weight: bold;
        cursor: pointer;

        &:hover,
        &:focus {
          color: black;
          text-decoration: none;
          cursor: pointer;
        }
      }
    }
  }
}
</style>
