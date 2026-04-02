<script setup lang="ts">
import { RouterLink } from 'vue-router'
import { ref, onMounted } from 'vue'
import { getCategoryList } from '@/api/category'
import type { Category } from '@/types'

const categories = ref<Category[]>([])
const isMenuOpen = ref(false)

onMounted(async () => {
  try {
    categories.value = await getCategoryList()
  } catch (e) {
    console.error('Failed to load categories:', e)
  }
})
</script>

<template>
  <header class="bg-white shadow-sm sticky top-0 z-50">
    <nav class="max-w-6xl mx-auto px-4">
      <div class="flex justify-between items-center h-16">
        <RouterLink to="/" class="text-xl font-bold text-primary-600">
          我的博客
        </RouterLink>

        <div class="hidden md:flex items-center space-x-6">
          <RouterLink to="/" class="text-gray-600 hover:text-primary-600">首页</RouterLink>
          <RouterLink to="/articles" class="text-gray-600 hover:text-primary-600">文章</RouterLink>
          <RouterLink to="/about" class="text-gray-600 hover:text-primary-600">关于</RouterLink>
          <RouterLink to="/admin/login" class="text-gray-600 hover:text-primary-600">登录</RouterLink>
        </div>

        <button class="md:hidden" @click="isMenuOpen = !isMenuOpen">
          <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16"/>
          </svg>
        </button>
      </div>

      <div v-if="isMenuOpen" class="md:hidden pb-4">
        <RouterLink to="/" class="block py-2 text-gray-600">首页</RouterLink>
        <RouterLink to="/articles" class="block py-2 text-gray-600">文章</RouterLink>
        <RouterLink to="/about" class="block py-2 text-gray-600">关于</RouterLink>
        <RouterLink to="/admin/login" class="block py-2 text-gray-600">登录</RouterLink>
      </div>
    </nav>
  </header>
</template>
