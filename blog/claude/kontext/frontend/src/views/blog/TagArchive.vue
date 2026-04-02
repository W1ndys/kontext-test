<template>
  <BlogLayout>
    <h1 class="text-3xl font-bold text-gray-800 mb-8">标签云</h1>

    <!-- Loading -->
    <div v-if="loading" class="flex flex-wrap gap-3">
      <div v-for="i in 12" :key="i" class="h-8 bg-gray-200 rounded-full animate-pulse" :style="{ width: (40 + Math.random() * 60) + 'px' }"></div>
    </div>

    <div v-else>
      <!-- Tag Cloud -->
      <div v-if="!selectedTag && tags.length > 0" class="flex flex-wrap gap-3 mb-8">
        <button
          v-for="tag in tags"
          :key="tag.id"
          class="inline-flex items-center gap-1 px-4 py-2 rounded-full border border-gray-200 hover:border-blue-300 hover:bg-blue-50 hover:text-blue-600 transition-all cursor-pointer"
          :style="{ fontSize: getTagSize(tag) + 'rem' }"
          @click="selectTag(tag)"
        >
          <span>#{{ tag.name }}</span>
          <span class="text-xs text-gray-400 ml-1">({{ tag.article_count || 0 }})</span>
        </button>
      </div>

      <!-- Empty -->
      <div v-if="!selectedTag && tags.length === 0" class="text-center py-20">
        <p class="text-gray-400 text-lg">暂无标签</p>
      </div>

      <!-- Tag Articles -->
      <div v-if="selectedTag">
        <div class="flex items-center gap-3 mb-6">
          <button
            @click="selectedTag = null; tagArticles = []"
            class="flex items-center gap-1 text-sm text-gray-500 hover:text-blue-600 transition-colors"
          >
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
            </svg>
            返回标签
          </button>
          <h2 class="text-xl font-bold text-gray-800">
            #{{ selectedTag.name }}
            <span class="text-sm text-gray-400 font-normal ml-2">{{ selectedTag.article_count || 0 }} 篇文章</span>
          </h2>
        </div>

        <div v-if="tagLoading" class="space-y-4">
          <div v-for="i in 3" :key="i" class="bg-white rounded-xl p-6 animate-pulse">
            <div class="h-5 bg-gray-200 rounded w-3/4 mb-3"></div>
            <div class="h-4 bg-gray-200 rounded w-1/2"></div>
          </div>
        </div>

        <div v-else-if="tagArticles.length > 0" class="space-y-4">
          <ArticleCard
            v-for="article in tagArticles"
            :key="article.id"
            :article="article"
          />
        </div>

        <div v-else class="text-center py-12">
          <p class="text-gray-400">该标签下暂无文章</p>
        </div>
      </div>
    </div>
  </BlogLayout>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import BlogLayout from '@/components/BlogLayout.vue'
import ArticleCard from '@/components/ArticleCard.vue'
import { getTags } from '@/api/tag'
import { getArticleList } from '@/api/article'

const tags = ref([])
const loading = ref(true)
const selectedTag = ref(null)
const tagArticles = ref([])
const tagLoading = ref(false)

function getTagSize(tag) {
  const count = tag.article_count || 0
  const maxCount = Math.max(...tags.value.map(t => t.article_count || 0), 1)
  const minSize = 0.85
  const maxSize = 1.4
  return minSize + ((count / maxCount) * (maxSize - minSize))
}

async function fetchTags() {
  loading.value = true
  try {
    const res = await getTags()
    tags.value = res.data || res || []
  } catch (err) {
    console.error('Failed to fetch tags:', err)
  } finally {
    loading.value = false
  }
}

async function selectTag(tag) {
  selectedTag.value = tag
  tagLoading.value = true
  try {
    const res = await getArticleList({ tag_id: tag.id, page_size: 100 })
    tagArticles.value = res.data || res.articles || []
  } catch (err) {
    console.error('Failed to fetch tag articles:', err)
    tagArticles.value = []
  } finally {
    tagLoading.value = false
  }
}

onMounted(() => {
  fetchTags()
})
</script>
