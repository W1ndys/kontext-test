<template>
  <BlogLayout>
    <!-- Loading -->
    <div v-if="loading" class="max-w-3xl mx-auto">
      <div class="animate-pulse">
        <div class="h-8 bg-gray-200 rounded w-3/4 mb-4"></div>
        <div class="flex gap-3 mb-8">
          <div class="h-5 bg-gray-200 rounded-full w-24"></div>
          <div class="h-5 bg-gray-200 rounded-full w-20"></div>
        </div>
        <div class="space-y-3">
          <div class="h-4 bg-gray-200 rounded w-full"></div>
          <div class="h-4 bg-gray-200 rounded w-full"></div>
          <div class="h-4 bg-gray-200 rounded w-5/6"></div>
          <div class="h-4 bg-gray-200 rounded w-full"></div>
          <div class="h-4 bg-gray-200 rounded w-4/6"></div>
        </div>
      </div>
    </div>

    <!-- Article -->
    <div v-else-if="article" class="max-w-3xl mx-auto">
      <!-- Back -->
      <button
        @click="$router.back()"
        class="flex items-center gap-1 text-sm text-gray-500 hover:text-blue-600 transition-colors mb-6"
      >
        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
        </svg>
        返回
      </button>

      <!-- Header -->
      <header class="mb-8">
        <h1 class="text-3xl sm:text-4xl font-bold text-gray-800 mb-4 leading-tight">
          {{ article.title }}
        </h1>
        <div class="flex flex-wrap items-center gap-3 text-sm text-gray-400">
          <span class="flex items-center gap-1">
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" />
            </svg>
            {{ formatDate(article.created_at) }}
          </span>
          <span v-if="article.views !== undefined" class="flex items-center gap-1">
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
            </svg>
            {{ article.views }} 次阅读
          </span>
          <span
            v-if="article.category"
            class="inline-flex items-center px-2.5 py-0.5 rounded-full bg-blue-50 text-blue-600 text-xs font-medium"
          >
            {{ article.category.name }}
          </span>
          <span
            v-for="tag in (article.tags || [])"
            :key="tag.id"
            class="inline-flex items-center px-2.5 py-0.5 rounded-full bg-gray-100 text-gray-500 text-xs font-medium"
          >
            #{{ tag.name }}
          </span>
        </div>
      </header>

      <!-- Content -->
      <article class="bg-white rounded-xl shadow-sm border border-gray-100 p-6 sm:p-10">
        <MarkdownRenderer :content="article.content" />
      </article>
    </div>

    <!-- Not Found -->
    <div v-else class="text-center py-20">
      <p class="text-gray-400 text-lg">文章不存在或已被删除</p>
      <router-link to="/" class="mt-4 inline-block text-blue-600 hover:text-blue-700">返回首页</router-link>
    </div>
  </BlogLayout>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import BlogLayout from '@/components/BlogLayout.vue'
import MarkdownRenderer from '@/components/MarkdownRenderer.vue'
import { getArticleDetail } from '@/api/article'

const route = useRoute()
const article = ref(null)
const loading = ref(true)

function formatDate(dateStr) {
  if (!dateStr) return ''
  const date = new Date(dateStr)
  return date.toLocaleDateString('zh-CN', {
    year: 'numeric',
    month: 'long',
    day: 'numeric'
  })
}

onMounted(async () => {
  try {
    const res = await getArticleDetail(route.params.id)
    article.value = res.data || res
  } catch (err) {
    console.error('Failed to fetch article:', err)
    article.value = null
  } finally {
    loading.value = false
  }
})
</script>
