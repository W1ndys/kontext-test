<template>
  <div class="flex gap-12">
    <div class="flex-1 min-w-0">
      <h1 v-if="filterLabel" class="text-lg font-semibold text-gray-700 mb-6">{{ filterLabel }}</h1>
      <div v-if="loading" class="text-gray-400 py-8 text-center">Loading...</div>
      <div v-else-if="articles.length === 0" class="text-gray-400 py-8 text-center">No articles yet.</div>
      <div v-else>
        <ArticleCard v-for="article in articles" :key="article.id" :article="article" />
        <Pagination :page="page" :total="total" :size="size" @change="onPageChange" />
      </div>
    </div>
    <div class="hidden lg:block w-56 shrink-0">
      <Sidebar />
    </div>
  </div>
</template>

<script setup>
import { ref, watch, computed } from 'vue'
import { useRoute } from 'vue-router'
import { getArticles } from '../api'
import ArticleCard from '../components/ArticleCard.vue'
import Sidebar from '../components/Sidebar.vue'
import Pagination from '../components/Pagination.vue'

const route = useRoute()
const articles = ref([])
const total = ref(0)
const page = ref(1)
const size = 10
const loading = ref(false)

const filterLabel = computed(() => {
  if (route.params.name && route.path.startsWith('/category')) return `Category: ${route.params.name}`
  if (route.params.name && route.path.startsWith('/tag')) return `Tag: ${route.params.name}`
  return ''
})

async function fetchArticles() {
  loading.value = true
  try {
    const params = { page: page.value, size }
    if (route.params.name && route.path.startsWith('/category')) params.category = route.params.name
    if (route.params.name && route.path.startsWith('/tag')) params.tag = route.params.name
    const { data } = await getArticles(params)
    articles.value = data.articles || []
    total.value = data.total
  } catch (e) {
    articles.value = []
  } finally {
    loading.value = false
  }
}

function onPageChange(newPage) {
  page.value = newPage
}

watch(() => route.fullPath, () => {
  page.value = 1
  fetchArticles()
}, { immediate: true })

watch(page, fetchArticles)
</script>
