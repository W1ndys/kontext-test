<template>
  <BlogLayout>
    <h1 class="text-3xl font-bold text-gray-800 mb-8">文章分类</h1>

    <!-- Loading -->
    <div v-if="loading" class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4">
      <div v-for="i in 6" :key="i" class="bg-white rounded-xl p-6 animate-pulse">
        <div class="h-5 bg-gray-200 rounded w-1/2 mb-2"></div>
        <div class="h-4 bg-gray-200 rounded w-1/4"></div>
      </div>
    </div>

    <div v-else>
      <!-- Category Cards -->
      <div v-if="!selectedCategory" class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4">
        <div
          v-for="cat in categories"
          :key="cat.id"
          class="bg-white rounded-xl p-6 border border-gray-100 hover:shadow-md hover:border-blue-100 transition-all cursor-pointer group"
          @click="selectCategory(cat)"
        >
          <div class="flex items-center justify-between">
            <h2 class="text-lg font-semibold text-gray-700 group-hover:text-blue-600 transition-colors">
              {{ cat.name }}
            </h2>
            <span class="text-sm text-gray-400 bg-gray-100 px-2.5 py-0.5 rounded-full">
              {{ cat.article_count || 0 }} 篇
            </span>
          </div>
        </div>
      </div>

      <!-- Empty -->
      <div v-if="!selectedCategory && categories.length === 0" class="text-center py-20">
        <p class="text-gray-400 text-lg">暂无分类</p>
      </div>

      <!-- Category Articles -->
      <div v-if="selectedCategory">
        <div class="flex items-center gap-3 mb-6">
          <button
            @click="selectedCategory = null; categoryArticles = []"
            class="flex items-center gap-1 text-sm text-gray-500 hover:text-blue-600 transition-colors"
          >
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
            </svg>
            返回分类
          </button>
          <h2 class="text-xl font-bold text-gray-800">
            {{ selectedCategory.name }}
            <span class="text-sm text-gray-400 font-normal ml-2">{{ selectedCategory.article_count || 0 }} 篇文章</span>
          </h2>
        </div>

        <!-- Articles -->
        <div v-if="categoryLoading" class="space-y-4">
          <div v-for="i in 3" :key="i" class="bg-white rounded-xl p-6 animate-pulse">
            <div class="h-5 bg-gray-200 rounded w-3/4 mb-3"></div>
            <div class="h-4 bg-gray-200 rounded w-1/2"></div>
          </div>
        </div>

        <div v-else-if="categoryArticles.length > 0" class="space-y-4">
          <ArticleCard
            v-for="article in categoryArticles"
            :key="article.id"
            :article="article"
          />
        </div>

        <div v-else class="text-center py-12">
          <p class="text-gray-400">该分类下暂无文章</p>
        </div>
      </div>
    </div>
  </BlogLayout>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import BlogLayout from '@/components/BlogLayout.vue'
import ArticleCard from '@/components/ArticleCard.vue'
import { getCategories } from '@/api/category'
import { getArticleList } from '@/api/article'

const categories = ref([])
const loading = ref(true)
const selectedCategory = ref(null)
const categoryArticles = ref([])
const categoryLoading = ref(false)

async function fetchCategories() {
  loading.value = true
  try {
    const res = await getCategories()
    categories.value = res.data || res || []
  } catch (err) {
    console.error('Failed to fetch categories:', err)
  } finally {
    loading.value = false
  }
}

async function selectCategory(cat) {
  selectedCategory.value = cat
  categoryLoading.value = true
  try {
    const res = await getArticleList({ category_id: cat.id, page_size: 100 })
    categoryArticles.value = res.data || res.articles || []
  } catch (err) {
    console.error('Failed to fetch category articles:', err)
    categoryArticles.value = []
  } finally {
    categoryLoading.value = false
  }
}

onMounted(() => {
  fetchCategories()
})
</script>
