<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { RouterLink } from 'vue-router'
import { useArticleStore } from '@/stores/article'
import { getCategoryList } from '@/api/category'
import ArticleCard from '@/components/ArticleCard.vue'
import Loading from '@/components/common/Loading.vue'
import type { Category } from '@/types'

const articleStore = useArticleStore()
const categories = ref<Category[]>([])
const loading = ref(true)

onMounted(async () => {
  try {
    await Promise.all([
      articleStore.fetchArticles({ page: 1, page_size: 6, status: 'published' }),
      getCategoryList().then(data => categories.value = data)
    ])
  } finally {
    loading.value = false
  }
})
</script>

<template>
  <div class="max-w-6xl mx-auto px-4 py-8">
    <section class="mb-12">
      <h1 class="text-4xl font-bold text-center mb-4">欢迎来到我的博客</h1>
      <p class="text-gray-600 text-center text-lg">分享技术，记录生活</p>
    </section>

    <section class="mb-12">
      <h2 class="text-2xl font-bold mb-6">分类</h2>
      <div class="flex flex-wrap gap-3">
        <RouterLink
          v-for="category in categories"
          :key="category.id"
          :to="`/category/${category.slug}`"
          class="px-4 py-2 bg-primary-50 text-primary-700 rounded-full hover:bg-primary-100 transition-colors"
        >
          {{ category.name }}
        </RouterLink>
      </div>
    </section>

    <section>
      <h2 class="text-2xl font-bold mb-6">最新文章</h2>
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
      <div class="text-center mt-8">
        <RouterLink to="/articles" class="btn btn-primary">
          查看更多文章
        </RouterLink>
      </div>
    </section>
  </div>
</template>
