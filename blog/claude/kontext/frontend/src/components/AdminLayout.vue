<template>
  <div class="min-h-screen flex bg-gray-100">
    <!-- Sidebar Overlay (mobile) -->
    <div
      v-if="sidebarOpen"
      class="fixed inset-0 bg-black/50 z-40 lg:hidden"
      @click="sidebarOpen = false"
    ></div>

    <!-- Sidebar -->
    <aside
      :class="[
        'fixed inset-y-0 left-0 z-50 w-64 bg-gray-900 text-white flex flex-col transition-transform duration-300 lg:translate-x-0 lg:static lg:z-auto',
        sidebarOpen ? 'translate-x-0' : '-translate-x-full'
      ]"
    >
      <!-- Logo -->
      <div class="h-16 flex items-center px-6 border-b border-gray-800">
        <router-link to="/admin/dashboard" class="text-lg font-bold text-white hover:text-blue-400 transition-colors">
          博客管理
        </router-link>
      </div>

      <!-- Nav Links -->
      <nav class="flex-1 px-4 py-6 space-y-1">
        <router-link
          v-for="link in sidebarLinks"
          :key="link.path"
          :to="link.path"
          class="flex items-center gap-3 px-3 py-2.5 rounded-lg text-gray-300 hover:text-white hover:bg-gray-800 transition-colors text-sm font-medium"
          active-class="!text-white !bg-blue-600"
          @click="sidebarOpen = false"
        >
          <span v-html="link.icon" class="w-5 h-5 flex-shrink-0"></span>
          {{ link.name }}
        </router-link>
      </nav>

      <!-- Logout -->
      <div class="px-4 py-4 border-t border-gray-800">
        <button
          @click="handleLogout"
          class="flex items-center gap-3 w-full px-3 py-2.5 rounded-lg text-gray-300 hover:text-white hover:bg-red-600/20 transition-colors text-sm font-medium"
        >
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1" />
          </svg>
          退出登录
        </button>
      </div>
    </aside>

    <!-- Main Area -->
    <div class="flex-1 flex flex-col min-w-0">
      <!-- Top Header -->
      <header class="h-16 bg-white shadow-sm flex items-center justify-between px-4 sm:px-6 sticky top-0 z-30">
        <div class="flex items-center gap-4">
          <!-- Mobile toggle -->
          <button
            class="lg:hidden text-gray-600 hover:text-gray-900"
            @click="sidebarOpen = !sidebarOpen"
          >
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
            </svg>
          </button>
          <h1 class="text-lg font-semibold text-gray-800">{{ pageTitle }}</h1>
        </div>
        <div class="flex items-center gap-4">
          <router-link to="/" target="_blank" class="text-sm text-gray-500 hover:text-blue-600 transition-colors">
            查看博客
          </router-link>
          <span class="text-sm text-gray-400">{{ userStore.userInfo?.username || '管理员' }}</span>
        </div>
      </header>

      <!-- Page Content -->
      <main class="flex-1 p-4 sm:p-6">
        <slot />
      </main>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useUserStore } from '@/stores/user'

const props = defineProps({
  title: {
    type: String,
    default: ''
  }
})

const route = useRoute()
const router = useRouter()
const userStore = useUserStore()
const sidebarOpen = ref(false)

const pageTitles = {
  AdminDashboard: '仪表盘',
  AdminArticleList: '文章管理',
  AdminArticleCreate: '新建文章',
  AdminArticleEdit: '编辑文章',
  AdminCategoryManage: '分类管理',
  AdminTagManage: '标签管理'
}

const pageTitle = computed(() => {
  return props.title || pageTitles[route.name] || '管理后台'
})

const sidebarLinks = [
  {
    name: '仪表盘',
    path: '/admin/dashboard',
    icon: '<svg fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6"/></svg>'
  },
  {
    name: '文章管理',
    path: '/admin/articles',
    icon: '<svg fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"/></svg>'
  },
  {
    name: '分类管理',
    path: '/admin/categories',
    icon: '<svg fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 7v10a2 2 0 002 2h14a2 2 0 002-2V9a2 2 0 00-2-2h-6l-2-2H5a2 2 0 00-2 2z"/></svg>'
  },
  {
    name: '标签管理',
    path: '/admin/tags',
    icon: '<svg fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 7h.01M7 3h5c.512 0 1.024.195 1.414.586l7 7a2 2 0 010 2.828l-7 7a2 2 0 01-2.828 0l-7-7A1.994 1.994 0 013 12V7a4 4 0 014-4z"/></svg>'
  }
]

function handleLogout() {
  userStore.clearAuth()
  router.push('/admin/login')
}
</script>
