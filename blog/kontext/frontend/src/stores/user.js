import { ref, computed } from 'vue'
import { defineStore } from 'pinia'

export const useUserStore = defineStore('user', () => {
  const token = ref(localStorage.getItem('blog_token') || '')
  const userInfo = ref(JSON.parse(localStorage.getItem('blog_user') || 'null'))

  const isLoggedIn = computed(() => !!token.value)

  function setAuth(newToken, user) {
    token.value = newToken
    userInfo.value = user
    localStorage.setItem('blog_token', newToken)
    localStorage.setItem('blog_user', JSON.stringify(user))
  }

  function clearAuth() {
    token.value = ''
    userInfo.value = null
    localStorage.removeItem('blog_token')
    localStorage.removeItem('blog_user')
  }

  return {
    token,
    userInfo,
    isLoggedIn,
    setAuth,
    clearAuth
  }
})
