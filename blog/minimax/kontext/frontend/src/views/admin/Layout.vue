<script setup lang="ts">
import { RouterLink, useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const authStore = useAuthStore()
const router = useRouter()

const menuItems = [
  { path: '/admin/articles', name: '文章管理', icon: 'M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z' },
  { path: '/admin/categories', name: '分类管理', icon: 'M3 7v10a2 2 0 002 2h14a2 2 0 002-2V9a2 2 0 00-2-2h-6l-2-2H5a2 2 0 00-2 2z' },
  { path: '/admin/tags', name: '标签管理', icon: 'M7 7h.01M7 3h5c.512 0 1.024.195 1.414.586l7 7a2 2 0 010 2.828l-7 7a2 2 0 01-2.828 0l-7-7A2 2 0 013 12V7a4 4 0 014-4z' },
  { path: '/admin/comments', name: '评论管理', icon: 'M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z' },
]

function handleLogout() {
  authStore.logout()
  router.push('/admin/login')
}
</script>

<template>
  <div class="min-h-screen bg-gray-100">
    <header class="bg-white shadow-sm">
      <div class="max-w-7xl mx-auto px-4 py-4 flex justify-between items-center">
        <h1 class="text-xl font-bold text-primary-600">博客后台</h1>
        <div class="flex items-center gap-4">
          <span class="text-gray-600">{{ authStore.user?.nickname || authStore.user?.username }}</span>
          <RouterLink to="/" target="_blank" class="text-gray-500 hover:text-primary-600">
            查看博客
          </RouterLink>
          <button @click="handleLogout" class="text-gray-500 hover:text-red-600">
            退出登录
          </button>
        </div>
      </div>
    </header>

    <div class="flex">
      <aside class="w-64 bg-white shadow-sm min-h-screen">
        <nav class="p-4">
          <RouterLink
            v-for="item in menuItems"
            :key="item.path"
            :to="item.path"
            class="flex items-center gap-3 px-4 py-3 text-gray-600 hover:bg-primary-50 hover:text-primary-600 rounded-lg mb-1"
            active-class="bg-primary-50 text-primary-600"
          >
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" :d="item.icon" />
            </svg>
            {{ item.name }}
          </RouterLink>
        </nav>
      </aside>

      <main class="flex-1 p-6">
        <RouterView />
      </main>
    </div>
  </div>
</template>
