<template>
  <header class="border-b border-gray-200">
    <nav class="max-w-4xl mx-auto px-4 py-4 flex justify-between items-center">
      <router-link to="/" class="text-xl font-bold text-gray-900 no-underline">
        My Blog
      </router-link>
      <div class="flex gap-4 items-center">
        <router-link to="/" class="text-gray-600 hover:text-gray-900 no-underline">Home</router-link>
        <template v-if="auth.isLoggedIn">
          <router-link to="/admin/articles" class="text-gray-600 hover:text-gray-900 no-underline">Admin</router-link>
          <button @click="handleLogout" class="text-gray-600 hover:text-gray-900">Logout</button>
        </template>
        <router-link v-else to="/login" class="text-gray-600 hover:text-gray-900 no-underline">Login</router-link>
      </div>
    </nav>
  </header>
</template>

<script setup>
import { useAuthStore } from '../stores/auth'
import { useRouter } from 'vue-router'

const auth = useAuthStore()
const router = useRouter()

function handleLogout() {
  auth.logout()
  router.push('/')
}
</script>
