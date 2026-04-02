import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import type { User } from '@/types'
import { login as apiLogin, getCurrentUser as apiGetCurrentUser } from '@/api/auth'

export const useAuthStore = defineStore('auth', () => {
  const token = ref<string | null>(localStorage.getItem('token'))
  const user = ref<User | null>(JSON.parse(localStorage.getItem('user') || 'null'))

  const isAuthenticated = computed(() => !!token.value)

  async function login(username: string, password: string) {
    const response = await apiLogin({ username, password })
    token.value = response.token
    user.value = response.user
    localStorage.setItem('token', response.token)
    localStorage.setItem('user', JSON.stringify(response.user))
    return response
  }

  async function fetchCurrentUser() {
    if (!token.value) return null
    try {
      const response = await apiGetCurrentUser()
      user.value = response
      localStorage.setItem('user', JSON.stringify(response))
      return response
    } catch {
      logout()
      return null
    }
  }

  function logout() {
    token.value = null
    user.value = null
    localStorage.removeItem('token')
    localStorage.removeItem('user')
  }

  return {
    token,
    user,
    isAuthenticated,
    login,
    fetchCurrentUser,
    logout
  }
})
