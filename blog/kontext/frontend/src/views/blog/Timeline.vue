<template>
  <BlogLayout>
    <h1 class="text-3xl font-bold text-gray-800 mb-8">时间轴</h1>

    <!-- Loading -->
    <div v-if="loading" class="space-y-8">
      <div v-for="i in 3" :key="i" class="animate-pulse">
        <div class="h-6 bg-gray-200 rounded w-32 mb-4"></div>
        <div class="ml-6 space-y-3">
          <div class="h-4 bg-gray-200 rounded w-3/4"></div>
          <div class="h-4 bg-gray-200 rounded w-2/3"></div>
        </div>
      </div>
    </div>

    <!-- Timeline -->
    <div v-else-if="Object.keys(groupedArticles).length > 0" class="relative">
      <!-- Timeline line -->
      <div class="absolute left-3 top-2 bottom-2 w-0.5 bg-blue-200"></div>

      <div v-for="(articles, yearMonth) in groupedArticles" :key="yearMonth" class="mb-10">
        <!-- Year-Month Header -->
        <div class="flex items-center gap-4 mb-4 relative">
          <div class="w-6 h-6 bg-blue-600 rounded-full border-4 border-blue-100 z-10 flex-shrink-0"></div>
          <h2 class="text-xl font-bold text-gray-800">{{ yearMonth }}</h2>
        </div>

        <!-- Articles -->
        <div class="ml-10 space-y-3">
          <router-link
            v-for="article in articles"
            :key="article.id"
            :to="`/article/${article.id}`"
            class="flex items-center gap-3 group"
          >
            <div class="w-2 h-2 bg-gray-300 rounded-full group-hover:bg-blue-500 transition-colors flex-shrink-0"></div>
            <span class="text-sm text-gray-400 flex-shrink-0">{{ formatDay(article.created_at) }}</span>
            <span class="text-gray-700 group-hover:text-blue-600 transition-colors truncate">
              {{ article.title }}
            </span>
          </router-link>
        </div>
      </div>
    </div>

    <!-- Empty -->
    <div v-else class="text-center py-20">
      <svg class="w-16 h-16 mx-auto text-gray-300 mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
      </svg>
      <p class="text-gray-400 text-lg">暂无时间轴数据</p>
    </div>
  </BlogLayout>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import BlogLayout from '@/components/BlogLayout.vue'
import { getTimeline } from '@/api/article'

const timelineData = ref([])
const loading = ref(true)

const groupedArticles = computed(() => {
  const groups = {}
  const articles = timelineData.value
  for (const article of articles) {
    const date = new Date(article.created_at)
    const key = `${date.getFullYear()}年${String(date.getMonth() + 1).padStart(2, '0')}月`
    if (!groups[key]) {
      groups[key] = []
    }
    groups[key].push(article)
  }
  return groups
})

function formatDay(dateStr) {
  if (!dateStr) return ''
  const date = new Date(dateStr)
  return `${String(date.getMonth() + 1).padStart(2, '0')}-${String(date.getDate()).padStart(2, '0')}`
}

onMounted(async () => {
  try {
    const res = await getTimeline()
    timelineData.value = res.data || res || []
  } catch (err) {
    console.error('Failed to fetch timeline:', err)
  } finally {
    loading.value = false
  }
})
</script>
