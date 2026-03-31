<template>
  <div class="min-h-screen flex items-center justify-center bg-gradient-to-br from-blue-500 via-blue-600 to-indigo-700 px-4">
    <div class="w-full max-w-md">
      <!-- Login Card -->
      <div class="bg-white rounded-2xl shadow-2xl p-8">
        <!-- Header -->
        <div class="text-center mb-8">
          <div class="w-16 h-16 bg-blue-100 rounded-2xl flex items-center justify-center mx-auto mb-4">
            <svg class="w-8 h-8 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z" />
            </svg>
          </div>
          <h1 class="text-2xl font-bold text-gray-800">博客管理</h1>
          <p class="text-gray-400 text-sm mt-1">请登录以继续</p>
        </div>

        <!-- Error Message -->
        <div v-if="errorMsg" class="mb-4 p-3 bg-red-50 border border-red-200 rounded-lg text-red-600 text-sm">
          {{ errorMsg }}
        </div>

        <!-- Form -->
        <form @submit.prevent="handleLogin" class="space-y-5">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1.5">用户名</label>
            <input
              v-model="form.username"
              type="text"
              placeholder="请输入用户名"
              required
              class="w-full px-4 py-2.5 border border-gray-300 rounded-lg text-gray-800 placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all"
            />
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1.5">密码</label>
            <input
              v-model="form.password"
              type="password"
              placeholder="请输入密码"
              required
              class="w-full px-4 py-2.5 border border-gray-300 rounded-lg text-gray-800 placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all"
            />
          </div>

          <button
            type="submit"
            :disabled="submitting"
            class="w-full py-2.5 bg-blue-600 text-white font-medium rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 transition-all disabled:opacity-60 disabled:cursor-not-allowed"
          >
            <span v-if="submitting" class="flex items-center justify-center gap-2">
              <svg class="animate-spin h-4 w-4" viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4" fill="none"></circle>
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z"></path>
              </svg>
              登录中...
            </span>
            <span v-else>登录</span>
          </button>
        </form>

        <!-- Back Link -->
        <div class="mt-6 text-center">
          <router-link to="/" class="text-sm text-gray-400 hover:text-blue-600 transition-colors">
            返回博客首页
          </router-link>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useUserStore } from '@/stores/user'
import { login } from '@/api/auth'

const router = useRouter()
const route = useRoute()
const userStore = useUserStore()

const form = reactive({
  username: '',
  password: ''
})

const submitting = ref(false)
const errorMsg = ref('')

async function handleLogin() {
  errorMsg.value = ''
  submitting.value = true

  try {
    const res = await login({
      username: form.username,
      password: form.password
    })
    const token = res.token || res.data?.token
    const user = res.user || res.data?.user || { username: form.username }
    userStore.setAuth(token, user)

    const redirect = route.query.redirect || '/admin/dashboard'
    router.push(redirect)
  } catch (err) {
    errorMsg.value = err.response?.data?.error || '登录失败，请检查用户名和密码'
  } finally {
    submitting.value = false
  }
}
</script>
