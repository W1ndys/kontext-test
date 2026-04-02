<script setup lang="ts">
import { ref } from 'vue'
import { useRouter, RouterLink } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const router = useRouter()
const authStore = useAuthStore()
const username = ref('')
const password = ref('')
const loading = ref(false)
const error = ref('')

async function handleLogin() {
  if (!username.value || !password.value) {
    error.value = '请输入用户名和密码'
    return
  }

  loading.value = true
  error.value = ''

  try {
    await authStore.login(username.value, password.value)
    router.push('/admin')
  } catch (e: any) {
    error.value = e.message || '登录失败'
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div class="min-h-screen flex items-center justify-center bg-gray-100">
    <div class="bg-white p-8 rounded-xl shadow-md w-full max-w-md">
      <h1 class="text-2xl font-bold text-center mb-6">博客后台登录</h1>

      <div v-if="error" class="bg-red-50 text-red-600 p-3 rounded-lg mb-4">
        {{ error }}
      </div>

      <form @submit.prevent="handleLogin">
        <div class="mb-4">
          <label class="block text-gray-700 mb-2">用户名</label>
          <input
            v-model="username"
            type="text"
            class="input"
            placeholder="请输入用户名"
          />
        </div>
        <div class="mb-6">
          <label class="block text-gray-700 mb-2">密码</label>
          <input
            v-model="password"
            type="password"
            class="input"
            placeholder="请输入密码"
          />
        </div>
        <button
          type="submit"
          :disabled="loading"
          class="btn btn-primary w-full"
        >
          {{ loading ? '登录中...' : '登录' }}
        </button>
      </form>
    </div>
  </div>
</template>
