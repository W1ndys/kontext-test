<template>
  <router-link
    :to="`/article/${article.id}`"
    class="block bg-white rounded-xl shadow-sm border border-gray-100 hover:shadow-md hover:border-blue-100 transition-all duration-300 overflow-hidden group"
  >
    <div class="p-6">
      <!-- Title -->
      <h2 class="text-xl font-bold text-gray-800 group-hover:text-blue-600 transition-colors mb-3 line-clamp-2">
        {{ article.title }}
      </h2>

      <!-- Summary -->
      <p v-if="article.summary" class="text-gray-500 text-sm leading-relaxed mb-4 line-clamp-3">
        {{ article.summary }}
      </p>

      <!-- Meta -->
      <div class="flex flex-wrap items-center gap-3 text-xs text-gray-400">
        <!-- Date -->
        <span class="flex items-center gap-1">
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" />
          </svg>
          {{ formatDate(article.created_at) }}
        </span>

        <!-- Views -->
        <span v-if="article.views !== undefined" class="flex items-center gap-1">
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
          </svg>
          {{ article.views }}
        </span>

        <!-- Category -->
        <span
          v-if="article.category"
          class="inline-flex items-center px-2.5 py-0.5 rounded-full bg-blue-50 text-blue-600 font-medium"
        >
          {{ article.category.name }}
        </span>

        <!-- Tags -->
        <span
          v-for="tag in (article.tags || [])"
          :key="tag.id"
          class="inline-flex items-center px-2.5 py-0.5 rounded-full bg-gray-100 text-gray-500 font-medium"
        >
          #{{ tag.name }}
        </span>
      </div>
    </div>
  </router-link>
</template>

<script setup>
defineProps({
  article: {
    type: Object,
    required: true
  }
})

function formatDate(dateStr) {
  if (!dateStr) return ''
  const date = new Date(dateStr)
  return date.toLocaleDateString('zh-CN', {
    year: 'numeric',
    month: 'long',
    day: 'numeric'
  })
}
</script>
