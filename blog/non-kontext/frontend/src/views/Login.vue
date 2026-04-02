<template>
  <div class="flex justify-center items-center min-h-[60vh]">
    <form @submit.prevent="handleLogin" class="w-full max-w-sm space-y-4">
      <h1 class="text-2xl font-bold text-center mb-8">Login</h1>
      <div v-if="error" class="text-red-500 text-sm text-center">{{ error }}</div>
      <div>
        <label class="block text-sm text-gray-600 mb-1">Username</label>
        <input v-model="username" type="text" class="w-full border border-gray-300 rounded px-3 py-2 focus:outline-none focus:border-gray-500" required />
      </div>
      <div>
        <label class="block text-sm text-gray-600 mb-1">Password</label>
        <input v-model="password" type="password" class="w-full border border-gray-300 rounded px-3 py-2 focus:outline-none focus:border-gray-500" required />
      </div>
      <button type="submit" :disabled="loading" class="w-full bg-gray-900 text-white py-2 rounded hover:bg-gray-800 disabled:opacity-50">
        {{ loading ? 'Logging in...' : 'Login' }}
      </button>
    </form>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '../stores/auth'
import { login } from '../api'

const username = ref('')
const password = ref('')
const error = ref('')
const loading = ref(false)
const router = useRouter()
const auth = useAuthStore()

async function handleLogin() {
  error.value = ''
  loading.value = true
  try {
    const { data } = await login(username.value, password.value)
    auth.setToken(data.token)
    router.push('/admin/articles')
  } catch (e) {
    error.value = 'Invalid username or password'
  } finally {
    loading.value = false
  }
}
</script>
