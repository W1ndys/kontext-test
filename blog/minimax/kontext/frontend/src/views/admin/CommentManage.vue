<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useCommentStore } from '@/stores/comment'
import { formatDateTime } from '@/utils/time'
import Pagination from '@/components/Pagination.vue'

const commentStore = useCommentStore()
const loading = ref(true)
const currentPage = ref(1)
const filterStatus = ref('')

onMounted(async () => {
  await fetchComments()
})

async function fetchComments(page = 1) {
  loading.value = true
  try {
    await commentStore.fetchComments({
      page,
      page_size: 10,
      status: filterStatus.value || undefined
    })
    currentPage.value = page
  } finally {
    loading.value = false
  }
}

function onPageChange(page: number) {
  fetchComments(page)
}

async function handleApprove(id: number) {
  try {
    await commentStore.updateStatus(id, 'approved')
  } catch (e) {
    alert('操作失败')
  }
}

async function handleReject(id: number) {
  try {
    await commentStore.updateStatus(id, 'rejected')
  } catch (e) {
    alert('操作失败')
  }
}

async function handleDelete(id: number) {
  if (!confirm('确定要删除这条评论吗？')) return
  try {
    await commentStore.removeComment(id)
  } catch (e) {
    alert('删除失败')
  }
}

function getStatusClass(status: string) {
  switch (status) {
    case 'approved': return 'bg-green-100 text-green-800'
    case 'rejected': return 'bg-red-100 text-red-800'
    default: return 'bg-yellow-100 text-yellow-800'
  }
}

function getStatusText(status: string) {
  switch (status) {
    case 'approved': return '已通过'
    case 'rejected': return '已拒绝'
    default: return '待审核'
  }
}
</script>

<template>
  <div>
    <div class="flex justify-between items-center mb-6">
      <h2 class="text-2xl font-bold">评论管理</h2>
      <select v-model="filterStatus" @change="fetchComments()" class="input w-40">
        <option value="">全部</option>
        <option value="pending">待审核</option>
        <option value="approved">已通过</option>
        <option value="rejected">已拒绝</option>
      </select>
    </div>

    <div class="bg-white rounded-xl shadow-md overflow-hidden">
      <table class="w-full">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">昵称</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">内容</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">文章</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">状态</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">时间</th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="comment in commentStore.comments" :key="comment.id" class="hover:bg-gray-50">
            <td class="px-6 py-4">
              <div>{{ comment.nickname }}</div>
              <div class="text-gray-400 text-xs">{{ comment.email || '-' }}</div>
            </td>
            <td class="px-6 py-4">
              <div class="max-w-xs truncate">{{ comment.content }}</div>
            </td>
            <td class="px-6 py-4 text-gray-500 text-sm">#{{ comment.article_id }}</td>
            <td class="px-6 py-4">
              <span :class="['px-2 py-1 text-xs rounded-full', getStatusClass(comment.status)]">
                {{ getStatusText(comment.status) }}
              </span>
            </td>
            <td class="px-6 py-4 text-gray-500 text-sm">{{ formatDateTime(comment.created_at) }}</td>
            <td class="px-6 py-4 text-right">
              <template v-if="comment.status === 'pending'">
                <button @click="handleApprove(comment.id)" class="text-green-600 hover:underline mr-4">通过</button>
                <button @click="handleReject(comment.id)" class="text-yellow-600 hover:underline mr-4">拒绝</button>
              </template>
              <button @click="handleDelete(comment.id)" class="text-red-600 hover:underline">删除</button>
            </td>
          </tr>
        </tbody>
      </table>
      <div v-if="loading" class="text-center py-12 text-gray-500">加载中...</div>
      <div v-if="!loading && commentStore.comments.length === 0" class="text-center py-12 text-gray-500">
        暂无评论
      </div>
    </div>

    <Pagination
      v-if="commentStore.total > 10"
      :current="currentPage"
      :total="commentStore.total"
      :page-size="10"
      @change="onPageChange"
    />
  </div>
</template>
