<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { RouterLink } from 'vue-router'
import { getAllArticles, deleteArticle } from '@/api/article'
import { formatDateTime } from '@/utils/time'
import type { Article } from '@/types'

const articles = ref<Article[]>([])
const loading = ref(true)

onMounted(async () => {
  try {
    articles.value = await getAllArticles()
  } catch (e) {
    console.error('Failed to load articles:', e)
  } finally {
    loading.value = false
  }
})

async function handleDelete(id: number) {
  if (!confirm('确定要删除这篇文章吗？')) return
  try {
    await deleteArticle(id)
    articles.value = articles.value.filter(a => a.id !== id)
  } catch (e) {
    alert('删除失败')
  }
}
</script>

<template>
  <div>
    <div class="flex justify-between items-center mb-6">
      <h2 class="text-2xl font-bold">文章管理</h2>
      <RouterLink to="/admin/articles/new" class="btn btn-primary">
        新建文章
      </RouterLink>
    </div>

    <div class="bg-white rounded-xl shadow-md overflow-hidden">
      <table class="w-full">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">标题</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">分类</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">状态</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">时间</th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="article in articles" :key="article.id" class="hover:bg-gray-50">
            <td class="px-6 py-4">
              <RouterLink :to="`/article/${article.slug}`" target="_blank" class="text-primary-600 hover:underline">
                {{ article.title }}
              </RouterLink>
            </td>
            <td class="px-6 py-4 text-gray-600">{{ article.category?.name || '-' }}</td>
            <td class="px-6 py-4">
              <span
                :class="article.status === 'published' ? 'bg-green-100 text-green-800' : 'bg-gray-100 text-gray-800'"
                class="px-2 py-1 text-xs rounded-full"
              >
                {{ article.status === 'published' ? '已发布' : '草稿' }}
              </span>
            </td>
            <td class="px-6 py-4 text-gray-500 text-sm">{{ formatDateTime(article.created_at) }}</td>
            <td class="px-6 py-4 text-right">
              <RouterLink :to="`/admin/articles/${article.id}/edit`" class="text-primary-600 hover:underline mr-4">
                编辑
              </RouterLink>
              <button @click="handleDelete(article.id)" class="text-red-600 hover:underline">
                删除
              </button>
            </td>
          </tr>
        </tbody>
      </table>
      <div v-if="loading" class="text-center py-12 text-gray-500">加载中...</div>
      <div v-if="!loading && articles.length === 0" class="text-center py-12 text-gray-500">
        暂无文章
      </div>
    </div>
  </div>
</template>
