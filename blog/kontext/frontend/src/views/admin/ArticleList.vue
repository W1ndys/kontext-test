<template>
  <AdminLayout>
    <!-- Header -->
    <div class="flex items-center justify-between mb-6">
      <h2 class="text-xl font-bold text-gray-800">文章列表</h2>
      <router-link
        to="/admin/articles/create"
        class="inline-flex items-center gap-2 px-4 py-2 bg-blue-600 text-white text-sm font-medium rounded-lg hover:bg-blue-700 transition-colors"
      >
        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
        </svg>
        新建文章
      </router-link>
    </div>

    <!-- Loading -->
    <div v-if="loading" class="bg-white rounded-xl border border-gray-100 shadow-sm overflow-hidden">
      <div class="p-6 space-y-4">
        <div v-for="i in 5" :key="i" class="flex items-center gap-4 animate-pulse">
          <div class="h-5 bg-gray-200 rounded flex-1"></div>
          <div class="h-5 bg-gray-200 rounded w-20"></div>
          <div class="h-5 bg-gray-200 rounded w-24"></div>
          <div class="h-5 bg-gray-200 rounded w-16"></div>
        </div>
      </div>
    </div>

    <!-- Table -->
    <div v-else class="bg-white rounded-xl border border-gray-100 shadow-sm overflow-hidden">
      <div class="overflow-x-auto">
        <table class="w-full">
          <thead>
            <tr class="bg-gray-50 border-b border-gray-100">
              <th class="text-left px-6 py-3 text-xs font-semibold text-gray-500 uppercase tracking-wider">标题</th>
              <th class="text-left px-6 py-3 text-xs font-semibold text-gray-500 uppercase tracking-wider hidden sm:table-cell">分类</th>
              <th class="text-left px-6 py-3 text-xs font-semibold text-gray-500 uppercase tracking-wider hidden md:table-cell">标签</th>
              <th class="text-left px-6 py-3 text-xs font-semibold text-gray-500 uppercase tracking-wider hidden lg:table-cell">日期</th>
              <th class="text-left px-6 py-3 text-xs font-semibold text-gray-500 uppercase tracking-wider hidden lg:table-cell">浏览</th>
              <th class="text-right px-6 py-3 text-xs font-semibold text-gray-500 uppercase tracking-wider">操作</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-gray-100">
            <tr v-if="articles.length === 0">
              <td colspan="6" class="px-6 py-12 text-center text-gray-400">
                暂无文章，点击右上角新建
              </td>
            </tr>
            <tr
              v-for="article in articles"
              :key="article.id"
              class="hover:bg-gray-50 transition-colors"
            >
              <td class="px-6 py-4">
                <span class="text-sm font-medium text-gray-800 line-clamp-1">{{ article.title }}</span>
              </td>
              <td class="px-6 py-4 hidden sm:table-cell">
                <span v-if="article.category" class="text-xs px-2 py-0.5 bg-blue-50 text-blue-600 rounded-full">
                  {{ article.category.name }}
                </span>
                <span v-else class="text-xs text-gray-400">-</span>
              </td>
              <td class="px-6 py-4 hidden md:table-cell">
                <div class="flex flex-wrap gap-1">
                  <span
                    v-for="tag in (article.tags || []).slice(0, 3)"
                    :key="tag.id"
                    class="text-xs px-2 py-0.5 bg-gray-100 text-gray-500 rounded-full"
                  >
                    {{ tag.name }}
                  </span>
                  <span v-if="(article.tags || []).length > 3" class="text-xs text-gray-400">
                    +{{ article.tags.length - 3 }}
                  </span>
                </div>
              </td>
              <td class="px-6 py-4 hidden lg:table-cell">
                <span class="text-sm text-gray-500">{{ formatDate(article.created_at) }}</span>
              </td>
              <td class="px-6 py-4 hidden lg:table-cell">
                <span class="text-sm text-gray-500">{{ article.views || 0 }}</span>
              </td>
              <td class="px-6 py-4 text-right">
                <div class="flex items-center justify-end gap-2">
                  <router-link
                    :to="`/admin/articles/edit/${article.id}`"
                    class="text-sm text-blue-600 hover:text-blue-700 transition-colors"
                  >
                    编辑
                  </router-link>
                  <button
                    @click="handleDelete(article)"
                    class="text-sm text-red-500 hover:text-red-600 transition-colors"
                  >
                    删除
                  </button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- Pagination -->
      <div class="px-6 py-4 border-t border-gray-100">
        <Pagination
          :current-page="currentPage"
          :total-pages="totalPages"
          :total-items="totalItems"
          @page-change="handlePageChange"
        />
      </div>
    </div>

    <!-- Delete Confirmation Modal -->
    <div
      v-if="showDeleteModal"
      class="fixed inset-0 bg-black/50 z-50 flex items-center justify-center px-4"
      @click.self="showDeleteModal = false"
    >
      <div class="bg-white rounded-xl shadow-xl p-6 w-full max-w-sm">
        <h3 class="text-lg font-semibold text-gray-800 mb-2">确认删除</h3>
        <p class="text-gray-500 text-sm mb-6">
          确定要删除文章「{{ deleteTarget?.title }}」吗？此操作不可撤销。
        </p>
        <div class="flex justify-end gap-3">
          <button
            @click="showDeleteModal = false"
            class="px-4 py-2 text-sm text-gray-600 bg-gray-100 rounded-lg hover:bg-gray-200 transition-colors"
          >
            取消
          </button>
          <button
            @click="confirmDelete"
            :disabled="deleting"
            class="px-4 py-2 text-sm text-white bg-red-500 rounded-lg hover:bg-red-600 transition-colors disabled:opacity-60"
          >
            {{ deleting ? '删除中...' : '确认删除' }}
          </button>
        </div>
      </div>
    </div>
  </AdminLayout>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import AdminLayout from '@/components/AdminLayout.vue'
import Pagination from '@/components/Pagination.vue'
import { getArticleList, deleteArticle } from '@/api/article'

const articles = ref([])
const loading = ref(true)
const currentPage = ref(1)
const totalPages = ref(1)
const totalItems = ref(0)
const pageSize = 15

const showDeleteModal = ref(false)
const deleteTarget = ref(null)
const deleting = ref(false)

function formatDate(dateStr) {
  if (!dateStr) return ''
  const date = new Date(dateStr)
  return date.toLocaleDateString('zh-CN')
}

async function fetchArticles(page = 1) {
  loading.value = true
  try {
    const res = await getArticleList({ page, page_size: pageSize })
    articles.value = res.data || res.articles || []
    totalItems.value = res.total || 0
    totalPages.value = Math.ceil(totalItems.value / pageSize) || 1
    currentPage.value = page
  } catch (err) {
    console.error('Failed to fetch articles:', err)
  } finally {
    loading.value = false
  }
}

function handlePageChange(page) {
  fetchArticles(page)
}

function handleDelete(article) {
  deleteTarget.value = article
  showDeleteModal.value = true
}

async function confirmDelete() {
  if (!deleteTarget.value) return
  deleting.value = true
  try {
    await deleteArticle(deleteTarget.value.id)
    showDeleteModal.value = false
    deleteTarget.value = null
    fetchArticles(currentPage.value)
  } catch (err) {
    console.error('Failed to delete article:', err)
  } finally {
    deleting.value = false
  }
}

onMounted(() => {
  fetchArticles()
})
</script>
