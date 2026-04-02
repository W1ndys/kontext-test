<template>
  <article class="py-6 border-b border-gray-100">
    <router-link :to="`/article/${article.id}`" class="no-underline">
      <h2 class="text-xl font-semibold text-gray-900 hover:text-gray-600 mb-2">{{ article.title }}</h2>
    </router-link>
    <div class="text-sm text-gray-400 mb-3 flex gap-4">
      <span>{{ formatDate(article.created_at) }}</span>
      <router-link v-if="article.category" :to="`/category/${article.category}`" class="text-gray-500 hover:text-gray-700 no-underline">{{ article.category }}</router-link>
    </div>
    <p class="text-gray-600 leading-relaxed">{{ excerpt }}</p>
    <div v-if="tagList.length" class="mt-3 flex gap-2 flex-wrap">
      <router-link v-for="tag in tagList" :key="tag" :to="`/tag/${tag}`" class="text-xs text-gray-500 bg-gray-100 px-2 py-1 rounded no-underline hover:bg-gray-200">{{ tag }}</router-link>
    </div>
  </article>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps({
  article: { type: Object, required: true },
})

const excerpt = computed(() => {
  const text = props.article.content.replace(/[#*`>\-\[\]!()]/g, '')
  return text.length > 150 ? text.slice(0, 150) + '...' : text
})

const tagList = computed(() => {
  if (!props.article.tags) return []
  return props.article.tags.split(',').map(t => t.trim()).filter(Boolean)
})

function formatDate(dateStr) {
  return new Date(dateStr).toLocaleDateString('zh-CN')
}
</script>
