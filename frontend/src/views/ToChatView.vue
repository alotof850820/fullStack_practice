<script setup lang="ts">
import { ref } from 'vue'
import { useUserStore } from '@/stores/user'
import { Icon } from '@iconify/vue'
import { go } from '@/router'
import { useToastStore } from '@/stores/toast'
import {
  apiPostCreateGroup,
  apiPostLoadGroup,
  apiPostaddFriend,
  apiPostSearchFriends,
  apiPostJoinGroup
} from '@/api'

const user = useUserStore()
const toast = useToastStore()

const tab = ref<number>(0)
const showModal = ref<boolean>(false)
const showCreateGroupModal = ref<boolean>(false)
const showAddGroup = ref<boolean>(false)
const friendId = ref<number>(0)
const groupId = ref<number>(0)
const groupName = ref<string>('')
const groupDescription = ref<string>('')

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
    loadGroups()
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
const createGroup = () => {
  apiPostCreateGroup({
    name: groupName.value,
    description: groupDescription.value,
    ownerId: user.userData.ID
  })
    .then((res: any) => {
      toast.showToast(res.data.Msg)
      showCreateGroupModal.value = false
    })
    .catch((err: any) => {
      toast.showToast(err)
    })
}
const loadGroups = () => {
  apiPostLoadGroup({
    ownerId: user.userData.ID
  })
    .then((res: any) => {
      user.groups = res.data.Data
    })
    .catch((err: any) => {
      toast.showToast(err)
    })
}
const addGroup = () => {
  apiPostJoinGroup({
    userId: user.userData.ID,
    groupId: groupId.value
  })
    .then((res: any) => {
      toast.showToast(res.data.Msg)
      showAddGroup.value = false
    })
    .catch((err: any) => {
      toast.showToast(err)
    })
}
setTab(0)
</script>

<template>
  <div class="chat-room">
    <div class="title">{{ tab === 0 ? '好友' : tab === 1 ? '群組' : '個人中心' }}</div>
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
    <div v-if="tab === 1" class="contact_box">
      <div
        class="friend"
        v-for="item in user.groups"
        :key="item.ID"
        @click="go('ChatRoom', { id: item.ID, name: item.Name })"
      >
        <img class="avatar" src="https://cdn-icons-png.flaticon.com/512/149/149071.png" alt="" />
        <div class="name">{{ item.Name }}({{ item.ID }})</div>
        <Icon icon="flowbite:angle-right-outline" />
      </div>
    </div>
    <div v-if="tab === 2" class="contact_box">
      <div class="list_box" @click="showModal = true">添加好友</div>
      <div class="list_box" @click="showCreateGroupModal = true">建立群組</div>
      <div class="list_box" @click="showAddGroup = true">加入群組</div>
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

    <div v-if="showCreateGroupModal" class="modal">
      <div class="modal-content">
        <span @click="showCreateGroupModal = false" class="close">&times;</span>
        <h2>建立群组</h2>
        <form @submit.prevent="createGroup">
          <label for="groupName">群组名称:</label>
          <input v-model="groupName" type="text" id="groupName" required />

          <label for="groupDescription">介绍:</label>
          <textarea v-model="groupDescription" id="groupDescription" rows="4"></textarea>

          <button type="submit">建立</button>
        </form>
      </div>
    </div>

    <div v-if="showAddGroup" class="modal">
      <div class="modal-content">
        <span class="close" @click="showAddGroup = false">&times;</span>
        <label for="friendId">群號ID:</label>
        <input v-model="groupId" type="text" id="friendId" placeholder="请输入群ID" />
        <button @click="addGroup">添加</button>
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
