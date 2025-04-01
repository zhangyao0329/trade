/**
 * userStore.js
 * 用户个人信息
 */

import { defineStore } from 'pinia'
import { loginApi } from '@/api/user'
import { getUserInfoAPI } from '@/api/profiles'
import { ref } from 'vue'
import { ElMessage } from 'element-plus'
// import useASE from '@/hooks/useASE'
// const { decrypt } = useASE()

export const useUserStore = defineStore(
  'user',
  () => {
    const userInfo = ref({})
    const getUserInfo = async ({ mail, password }) => {
      const res = await loginApi({ mail, password })
      if (res.data.code === 1) {
        ElMessage.success('登录成功')
        userInfo.value = res.data.data
        // console.log('登录数据：', userInfo.value)
      } else ElMessage.error('用户名或密码错误')
    }

    // 清除用户信息
    const clearUserInfo = () => {
      userInfo.value = {}
    }

    // 获取用户信息
    const fetchUserInfo = async () => {
      // console.log('pinia原数据：', userInfo.value)
      const res = await getUserInfoAPI()
      // res.data.data.password = decrypt(res.data.data.password)
      userInfo.value = {
        ...userInfo.value, // 保留现有字段
        ...res.data.data // 更新新的字段
      }
      // console.log('pinia更新数据：', userInfo.value)
    }

    return {
      userInfo,
      getUserInfo,
      clearUserInfo,
      fetchUserInfo
    }
  },
  {
    persist: true
  }
)
