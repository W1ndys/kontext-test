<template>
  <div>
    <div class="flex justify-between items-center mb-6">
      <h1 class="text-2xl font-bold">Article Management</h1>
      <router-link to="/admin/articles/new" class="bg-gray-900 text-white px-4 py-2 rounded text-sm hover:bg-gray-800 no-underline">New Article</router-link>
    </div>
    <div v-if="loading" class="text-gray-400 py-8 text-center">Loading...</div>
    <table v-else class="w-full text-left">
      <thead>
        <tr class="border-b border-gray-200">
          <th class="py-2 text-sm font-semibold text-gray-600">Title</th>
          <th class="py-2 text-sm font-semibold text-gray-600 w-28">Category</th>
          <th class="py-2 text-sm font-semibold text-gray-600 w-28">Date</th>
          <th class="py-2 text-sm font-semibold text-gray-600 w-32 text-right">Actions</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="article in articles" :key="article.id" class="border-b border-gray-100">
          <td class="py-3 text-sm">{{ article.title }}</td>
          <td class="py-3 text-sm text-gray-500">{{ article.category }}</td>
          <td class="py-3 text-sm text-gray-400">{{ formatDate(article.created_at) }}</td>
          <td class="py-3 text-sm text-right space-x-3">
            <router-link :to="`/admin/articles/edit/${article.id}`" class="text-gray-600 hover:text-gray-900 no-underline">Edit</router-link>
            <button @click="handleDelete(article.id)" class="text-red-500 hover:text-red-700">Delete</button>
          </td>
        </tr>
        <tr v-if="articles.length === 0">
          <td colspan="4" class="py-8 text-center text-gray-400">No articles yet.</td>
        </tr>
      </tbody>
    </table>
    <Pagination :page="page" :total="total" :size="size" @change="onPageChange" />
  </div>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue'
import { getArticles, deleteArticle } from '../api'
import Pagination from '../components/Pagination.vue'

const articles = ref([])
const total = ref(0)
const page = ref(1)
const size = 20
const loading = ref(false)

async function fetchArticles() {
  loading.value = true
  try {
    const { data } = await getArticles({ page: page.value, size })
    articles.value = data.articles || []
    total.value = data.total
  } finally {
    loading.value = false
  }
}

async function handleDelete(id) {
  if (!confirm('Are you sure you want to delete this article?')) return
  await deleteArticle(id)
  fetchArticles()
}

function onPageChange(newPage) {
  page.value = newPage
}

function formatDate(dateStr) {
  return new Date(dateStr).toLocaleDateString('zh-CN')
}

watch(page, fetchArticles)
onMounted(fetchArticles)
</script>
