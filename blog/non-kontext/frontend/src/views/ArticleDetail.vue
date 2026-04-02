<template>
  <article v-if="article" class="max-w-3xl mx-auto">
    <header class="mb-8">
      <h1 class="text-3xl font-bold text-gray-900 mb-3">{{ article.title }}</h1>
      <div class="text-sm text-gray-400 flex gap-4">
        <span>{{ formatDate(article.created_at) }}</span>
        <router-link v-if="article.category" :to="`/category/${article.category}`" class="text-gray-500 hover:text-gray-700 no-underline">{{ article.category }}</router-link>
      </div>
      <div v-if="tagList.length" class="mt-3 flex gap-2 flex-wrap">
        <router-link v-for="tag in tagList" :key="tag" :to="`/tag/${tag}`" class="text-xs text-gray-500 bg-gray-100 px-2 py-1 rounded no-underline hover:bg-gray-200">{{ tag }}</router-link>
      </div>
    </header>
    <div class="prose prose-gray max-w-none" v-html="renderedContent"></div>
    <div class="mt-12">
      <router-link to="/" class="text-gray-500 hover:text-gray-700 no-underline text-sm">&larr; Back to Home</router-link>
    </div>
  </article>
  <div v-else-if="loading" class="text-gray-400 py-8 text-center">Loading...</div>
  <div v-else class="text-gray-400 py-8 text-center">Article not found.</div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { getArticle } from '../api'
import MarkdownIt from 'markdown-it'

const md = new MarkdownIt()
const route = useRoute()
const article = ref(null)
const loading = ref(true)

const renderedContent = computed(() => {
  if (!article.value) return ''
  return md.render(article.value.content)
})

const tagList = computed(() => {
  if (!article.value?.tags) return []
  return article.value.tags.split(',').map(t => t.trim()).filter(Boolean)
})

function formatDate(dateStr) {
  return new Date(dateStr).toLocaleDateString('zh-CN')
}

onMounted(async () => {
  try {
    const { data } = await getArticle(route.params.id)
    article.value = data
  } catch (e) {
    article.value = null
  } finally {
    loading.value = false
  }
})
</script>
