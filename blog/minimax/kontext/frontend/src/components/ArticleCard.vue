<script setup lang="ts">
import { RouterLink } from 'vue-router'
import type { Article } from '@/types'
import { formatDate } from '@/utils/time'

defineProps<{
  article: Article
}>()
</script>

<template>
  <RouterLink :to="`/article/${article.slug}`" class="card hover:shadow-lg transition-shadow">
    <div v-if="article.cover_image" class="h-48 overflow-hidden">
      <img :src="article.cover_image" :alt="article.title" class="w-full h-full object-cover" />
    </div>
    <div class="p-6">
      <h3 class="text-xl font-bold mb-2 line-clamp-2 hover:text-primary-600">{{ article.title }}</h3>
      <p v-if="article.summary" class="text-gray-600 text-sm mb-4 line-clamp-2">{{ article.summary }}</p>
      <div class="flex items-center text-gray-500 text-xs">
        <span>{{ formatDate(article.created_at) }}</span>
        <span class="mx-2">|</span>
        <span>{{ article.view_count }} 阅读</span>
      </div>
      <div v-if="article.tags?.length" class="flex flex-wrap gap-2 mt-3">
        <span
          v-for="tag in article.tags.slice(0, 3)"
          :key="tag.id"
          class="px-2 py-1 bg-gray-100 text-gray-600 text-xs rounded"
        >
          {{ tag.name }}
        </span>
      </div>
    </div>
  </RouterLink>
</template>

<style scoped>
.line-clamp-2 {
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}
</style>
