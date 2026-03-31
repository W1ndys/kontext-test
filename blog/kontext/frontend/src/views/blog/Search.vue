<template>
  <BlogLayout>
    <div class="max-w-3xl mx-auto">
      <h1 class="text-3xl font-bold text-gray-800 mb-8">搜索文章</h1>

      <!-- Search Input -->
      <div class="relative mb-8">
        <svg class="w-5 h-5 absolute left-4 top-1/2 -translate-y-1/2 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
        </svg>
        <input
          v-model="keyword"
          type="text"
          placeholder="输入关键词搜索..."
          class="w-full pl-12 pr-4 py-3 bg-white border border-gray-200 rounded-xl text-gray-800 placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all"
          @keyup.enter="handleSearch"
        />
        <button
          @click="handleSearch"
          class="absolute right-2 top-1/2 -translate-y-1/2 px-4 py-1.5 bg-blue-600 text-white text-sm rounded-lg hover:bg-blue-700 transition-colors"
        >
          搜索
        </button>
      </div>

      <!-- Loading -->
      <div v-if="loading" class="space-y-4">
        <div v-for="i in 3" :key="i" class="bg-white rounded-xl p-6 animate-pulse">
          <div class="h-5 bg-gray-200 rounded w-3/4 mb-3"></div>
          <div class="h-4 bg-gray-200 rounded w-full mb-2"></div>
          <div class="h-4 bg-gray-200 rounded w-1/2"></div>
        </div>
      </div>

      <!-- Results -->
      <div v-else-if="searched">
        <p class="text-sm text-gray-400 mb-4">
          找到 {{ results.length }} 条结果
        </p>

        <div v-if="results.length > 0" class="space-y-4">
          <ArticleCard
            v-for="article in results"
            :key="article.id"
            :article="article"
          />
        </div>

        <div v-else class="text-center py-16">
          <svg class="w-16 h-16 mx-auto text-gray-300 mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
          </svg>
          <p class="text-gray-400 text-lg">未找到相关文章</p>
          <p class="text-gray-300 text-sm mt-2">换个关键词试试吧</p>
        </div>
      </div>

      <!-- Initial State -->
      <div v-else class="text-center py-16">
        <svg class="w-16 h-16 mx-auto text-gray-300 mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
        </svg>
        <p class="text-gray-400">输入关键词搜索文章</p>
      </div>
    </div>
  </BlogLayout>
</template>

<script setup>
import { ref } from 'vue'
import BlogLayout from '@/components/BlogLayout.vue'
import ArticleCard from '@/components/ArticleCard.vue'
import { searchArticles } from '@/api/article'

const keyword = ref('')
const results = ref([])
const loading = ref(false)
const searched = ref(false)

async function handleSearch() {
  const q = keyword.value.trim()
  if (!q) return

  loading.value = true
  searched.value = true
  try {
    const res = await searchArticles({ q })
    results.value = res.data || res.articles || []
  } catch (err) {
    console.error('Search failed:', err)
    results.value = []
  } finally {
    loading.value = false
  }
}
</script>
