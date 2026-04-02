<template>
  <div class="min-h-screen flex flex-col bg-gray-50">
    <!-- Top Navigation -->
    <header class="bg-white shadow-sm sticky top-0 z-50">
      <div class="max-w-5xl mx-auto px-4 sm:px-6">
        <div class="flex items-center justify-between h-16">
          <!-- Logo -->
          <router-link to="/" class="text-xl font-bold text-blue-600 hover:text-blue-700 transition-colors">
            我的博客
          </router-link>

          <!-- Desktop Nav -->
          <nav class="hidden md:flex items-center space-x-8">
            <router-link
              v-for="link in navLinks"
              :key="link.path"
              :to="link.path"
              class="text-gray-600 hover:text-blue-600 transition-colors text-sm font-medium"
              active-class="text-blue-600"
            >
              {{ link.name }}
            </router-link>
            <router-link
              to="/search"
              class="text-gray-600 hover:text-blue-600 transition-colors"
              active-class="text-blue-600"
            >
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
              </svg>
            </router-link>
          </nav>

          <!-- Mobile Hamburger -->
          <button
            class="md:hidden text-gray-600 hover:text-blue-600 transition-colors"
            @click="mobileMenuOpen = !mobileMenuOpen"
          >
            <svg v-if="!mobileMenuOpen" class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
            </svg>
            <svg v-else class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>
      </div>

      <!-- Mobile Menu -->
      <div v-if="mobileMenuOpen" class="md:hidden border-t border-gray-100">
        <nav class="px-4 py-3 space-y-2">
          <router-link
            v-for="link in navLinks"
            :key="link.path"
            :to="link.path"
            class="block px-3 py-2 text-gray-600 hover:text-blue-600 hover:bg-blue-50 rounded-lg transition-colors text-sm font-medium"
            active-class="text-blue-600 bg-blue-50"
            @click="mobileMenuOpen = false"
          >
            {{ link.name }}
          </router-link>
          <router-link
            to="/search"
            class="block px-3 py-2 text-gray-600 hover:text-blue-600 hover:bg-blue-50 rounded-lg transition-colors text-sm font-medium"
            @click="mobileMenuOpen = false"
          >
            搜索
          </router-link>
        </nav>
      </div>
    </header>

    <!-- Main Content -->
    <main class="flex-1 max-w-5xl mx-auto w-full px-4 sm:px-6 py-8">
      <slot />
    </main>

    <!-- Footer -->
    <footer class="bg-white border-t border-gray-200">
      <div class="max-w-5xl mx-auto px-4 sm:px-6 py-6">
        <p class="text-center text-gray-400 text-sm">
          &copy; {{ currentYear }} 我的博客. All rights reserved.
        </p>
      </div>
    </footer>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'

const mobileMenuOpen = ref(false)

const currentYear = computed(() => new Date().getFullYear())

const navLinks = [
  { name: '首页', path: '/' },
  { name: '分类', path: '/categories' },
  { name: '标签', path: '/tags' },
  { name: '时间轴', path: '/timeline' },
  { name: '关于', path: '/about' }
]
</script>
