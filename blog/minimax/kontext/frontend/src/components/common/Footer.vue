<script setup lang="ts">
import { RouterLink } from 'vue-router'
</script>

<template>
  <footer class="bg-gray-800 text-gray-300 py-8 mt-12">
    <div class="max-w-6xl mx-auto px-4">
      <div class="flex flex-col md:flex-row justify-between items-center">
        <div class="mb-4 md:mb-0">
          <RouterLink to="/" class="text-xl font-bold text-white">我的博客</RouterLink>
          <p class="text-sm mt-2">分享技术，记录生活</p>
        </div>
        <div class="flex space-x-6">
          <RouterLink to="/about" class="hover:text-white">关于</RouterLink>
          <a href="https://github.com" target="_blank" class="hover:text-white">GitHub</a>
        </div>
      </div>
      <div class="text-center mt-6 text-sm text-gray-500">
        &copy; {{ new Date().getFullYear() }} 我的博客. All rights reserved.
      </div>
    </div>
  </footer>
</template>
