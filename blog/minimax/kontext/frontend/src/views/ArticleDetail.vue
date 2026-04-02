<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import { useRoute, RouterLink } from 'vue-router'
import { useArticleStore } from '@/stores/article'
import { useCommentStore } from '@/stores/comment'
import { renderMarkdown, formatDate } from '@/utils/time'
import CommentSection from '@/components/CommentSection.vue'
import Loading from '@/components/common/Loading.vue'

const route = useRoute()
const articleStore = useArticleStore()
const commentStore = useCommentStore()
const loading = ref(true)

const article = computed(() => articleStore.currentArticle)
const renderedContent = computed(() => article.value ? renderMarkdown(article.value.content) : '')

onMounted(async () => {
  try {
    const slug = route.params.slug as string
    await articleStore.fetchArticleBySlug(slug)
    if (article.value) {
      await commentStore.fetchCommentsByArticle(article.value.id)
    }
  } finally {
    loading.value = false
  }
})
</script>

<template>
  <div class="max-w-4xl mx-auto px-4 py-8">
    <Loading v-if="loading" />
    <article v-else-if="article" class="bg-white rounded-xl shadow-md p-8">
      <div class="mb-6">
        <RouterLink to="/articles" class="text-primary-600 hover:text-primary-800 mb-4 inline-block">
          &larr; 返回文章列表
        </RouterLink>
        <h1 class="text-4xl font-bold mb-4">{{ article.title }}</h1>
        <div class="flex items-center text-gray-500 text-sm">
          <span>{{ formatDate(article.created_at) }}</span>
          <span class="mx-2">|</span>
          <span>{{ article.view_count }} 阅读</span>
          <span class="mx-2">|</span>
          <RouterLink
            v-if="article.category"
            :to="`/category/${article.category.slug}`"
            class="text-primary-600 hover:text-primary-800"
          >
            {{ article.category.name }}
          </RouterLink>
        </div>
        <div v-if="article.tags?.length" class="flex gap-2 mt-3">
          <RouterLink
            v-for="tag in article.tags"
            :key="tag.id"
            :to="`/tag/${tag.slug}`"
            class="px-2 py-1 bg-gray-100 text-gray-600 text-sm rounded hover:bg-gray-200"
          >
            {{ tag.name }}
          </RouterLink>
        </div>
      </div>

      <div v-if="article.cover_image" class="mb-6">
        <img :src="article.cover_image" :alt="article.title" class="w-full h-64 object-cover rounded-lg" />
      </div>

      <div class="prose max-w-none" v-html="renderedContent"></div>
    </article>
    <div v-else class="text-center text-gray-500 py-12">
      文章不存在
    </div>

    <CommentSection v-if="article" :article-id="article.id" />
  </div>
</template>
