<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { RouterLink } from 'vue-router'
import { useArticleStore } from '@/stores/article'
import ArticleCard from '@/components/ArticleCard.vue'
import Pagination from '@/components/Pagination.vue'
import Loading from '@/components/common/Loading.vue'

const articleStore = useArticleStore()
const loading = ref(true)
const currentPage = ref(1)
const pageSize = 9

onMounted(async () => {
  loading.value = false
  await fetchArticles()
})

async function fetchArticles(page = 1) {
  loading.value = true
  try {
    await articleStore.fetchArticles({
      page,
      page_size: pageSize,
      status: 'published'
    })
    currentPage.value = page
  } finally {
    loading.value = false
  }
}

function onPageChange(page: number) {
  fetchArticles(page)
  window.scrollTo({ top: 0, behavior: 'smooth' })
}
</script>

<template>
  <div class="max-w-6xl mx-auto px-4 py-8">
    <h1 class="text-3xl font-bold mb-8">所有文章</h1>
    <Loading v-if="loading" />
    <div v-else class="grid md:grid-cols-2 lg:grid-cols-3 gap-6">
      <ArticleCard
        v-for="article in articleStore.articles"
        :key="article.id"
        :article="article"
      />
    </div>
    <div v-if="!loading && articleStore.articles.length === 0" class="text-center text-gray-500 py-12">
      暂无文章
    </div>
    <Pagination
      v-if="articleStore.total > pageSize"
      :current="currentPage"
      :total="articleStore.total"
      :page-size="pageSize"
      @change="onPageChange"
    />
  </div>
</template>
