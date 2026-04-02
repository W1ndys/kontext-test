<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { useArticleStore } from '@/stores/article'
import { getTagBySlug } from '@/api/tag'
import ArticleCard from '@/components/ArticleCard.vue'
import Loading from '@/components/common/Loading.vue'
import type { Tag } from '@/types'

const route = useRoute()
const articleStore = useArticleStore()
const tag = ref<Tag | null>(null)
const loading = ref(true)

onMounted(async () => {
  try {
    const slug = route.params.slug as string
    tag.value = await getTagBySlug(slug)
    await articleStore.fetchArticles({
      page: 1,
      page_size: 10,
      tag_id: tag.value.id,
      status: 'published'
    })
  } finally {
    loading.value = false
  }
})
</script>

<template>
  <div class="max-w-6xl mx-auto px-4 py-8">
    <h1 v-if="tag" class="text-3xl font-bold mb-8">标签: {{ tag.name }}</h1>
    <Loading v-if="loading" />
    <div v-else class="grid md:grid-cols-2 lg:grid-cols-3 gap-6">
      <ArticleCard
        v-for="article in articleStore.articles"
        :key="article.id"
        :article="article"
      />
    </div>
    <div v-if="!loading && articleStore.articles.length === 0" class="text-center text-gray-500 py-12">
      该标签下暂无文章
    </div>
  </div>
</template>
