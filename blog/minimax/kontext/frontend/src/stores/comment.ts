import { defineStore } from 'pinia'
import { ref } from 'vue'
import type { Comment, CommentListQuery } from '@/types'
import { getCommentListByArticle as apiGetByArticle, createComment as apiCreate, getCommentList as apiGetList, updateCommentStatus as apiUpdateStatus, deleteComment as apiDelete } from '@/api/comment'

export const useCommentStore = defineStore('comment', () => {
  const comments = ref<Comment[]>([])
  const total = ref(0)
  const loading = ref(false)

  async function fetchCommentsByArticle(articleId: number) {
    loading.value = true
    try {
      comments.value = await apiGetByArticle(articleId)
      return comments.value
    } finally {
      loading.value = false
    }
  }

  async function createComment(data: { article_id: number; nickname: string; email?: string; content: string }) {
    const comment = await apiCreate(data)
    return comment
  }

  async function fetchComments(params: CommentListQuery = {}) {
    loading.value = true
    try {
      const response = await apiGetList(params)
      comments.value = response.list
      total.value = response.total
      return response
    } finally {
      loading.value = false
    }
  }

  async function updateStatus(id: number, status: string) {
    const comment = await apiUpdateStatus(id, status)
    const index = comments.value.findIndex(c => c.id === id)
    if (index !== -1) {
      comments.value[index] = comment
    }
    return comment
  }

  async function removeComment(id: number) {
    await apiDelete(id)
    comments.value = comments.value.filter(c => c.id !== id)
  }

  return {
    comments,
    total,
    loading,
    fetchCommentsByArticle,
    createComment,
    fetchComments,
    updateStatus,
    removeComment
  }
})
