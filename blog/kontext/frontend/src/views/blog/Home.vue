<template>
  <BlogLayout>
    <!-- Hero Section -->
    <div class="mb-10">
      <h1 class="text-3xl sm:text-4xl font-bold text-gray-800 mb-2">欢迎来到我的博客</h1>
      <p class="text-gray-500">记录技术、分享生活</p>
    </div>

    <!-- Loading -->
    <div v-if="loading" class="space-y-6">
      <div v-for="i in 3" :key="i" class="bg-white rounded-xl p-6 animate-pulse">
        <div class="h-6 bg-gray-200 rounded w-3/4 mb-4"></div>
        <div class="h-4 bg-gray-200 rounded w-full mb-2"></div>
        <div class="h-4 bg-gray-200 rounded w-2/3 mb-4"></div>
        <div class="flex gap-2">
          <div class="h-5 bg-gray-200 rounded-full w-20"></div>
          <div class="h-5 bg-gray-200 rounded-full w-16"></div>
        </div>
      </div>
    </div>

    <!-- Article List -->
    <div v-else-if="articles.length > 0" class="space-y-6">
      <ArticleCard
        v-for="article in articles"
        :key="article.id"
        :article="article"
      />

      <Pagination
        :current-page="currentPage"
        :total-pages="totalPages"
        :total-items="totalItems"
        @page-change="handlePageChange"
      />
    </div>

    <!-- Empty State -->
    <div v-else class="text-center py-20">
      <svg class="w-16 h-16 mx-auto text-gray-300 mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
      </svg>
      <p class="text-gray-400 text-lg">暂无文章</p>
    </div>
  </BlogLayout>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import BlogLayout from '@/components/BlogLayout.vue'
import ArticleCard from '@/components/ArticleCard.vue'
import Pagination from '@/components/Pagination.vue'
import { getArticleList } from '@/api/article'

const articles = ref([])
const loading = ref(true)
const currentPage = ref(1)
const totalPages = ref(1)
const totalItems = ref(0)
const pageSize = 10

async function fetchArticles(page = 1) {
  loading.value = true
  try {
    const res = await getArticleList({ page, page_size: pageSize })
    articles.value = res.data || res.articles || []
    totalItems.value = res.total || 0
    totalPages.value = Math.ceil(totalItems.value / pageSize) || 1
    currentPage.value = page
  } catch (err) {
    console.error('Failed to fetch articles:', err)
    articles.value = []
  } finally {
    loading.value = false
  }
}

function handlePageChange(page) {
  fetchArticles(page)
  window.scrollTo({ top: 0, behavior: 'smooth' })
}

onMounted(() => {
  fetchArticles()
})
</script>
