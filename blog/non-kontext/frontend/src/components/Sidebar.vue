<template>
  <aside class="space-y-8">
    <div>
      <h3 class="text-sm font-semibold text-gray-900 uppercase tracking-wide mb-3">Categories</h3>
      <ul class="space-y-1">
        <li v-for="cat in categories" :key="cat">
          <router-link :to="`/category/${cat}`" class="text-gray-600 hover:text-gray-900 text-sm no-underline">{{ cat }}</router-link>
        </li>
        <li v-if="!categories.length" class="text-gray-400 text-sm">No categories yet</li>
      </ul>
    </div>
    <div>
      <h3 class="text-sm font-semibold text-gray-900 uppercase tracking-wide mb-3">Tags</h3>
      <div class="flex flex-wrap gap-2">
        <router-link v-for="tag in tags" :key="tag" :to="`/tag/${tag}`" class="text-xs text-gray-500 bg-gray-100 px-2 py-1 rounded no-underline hover:bg-gray-200">{{ tag }}</router-link>
        <span v-if="!tags.length" class="text-gray-400 text-sm">No tags yet</span>
      </div>
    </div>
  </aside>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { getCategories, getTags } from '../api'

const categories = ref([])
const tags = ref([])

onMounted(async () => {
  try {
    const [catRes, tagRes] = await Promise.all([getCategories(), getTags()])
    categories.value = catRes.data || []
    tags.value = tagRes.data || []
  } catch (e) {
    // silently fail
  }
})
</script>
