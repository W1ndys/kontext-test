import request from '@/utils/request'
import type { Comment, CreateCommentDTO, CommentListQuery, PageResult } from '@/types'

export const getCommentListByArticle = async (articleId: number) => {
  const response = await request.get<Comment[]>(`/articles/${articleId}/comments`)
  return response.data.data
}

export const createComment = async (data: CreateCommentDTO) => {
  const response = await request.post<Comment>('/comments', data)
  return response.data.data
}

export const getCommentList = async (params: CommentListQuery = {}) => {
  const response = await request.get<PageResult<Comment>>('/comments', { params })
  return response.data.data
}

export const updateCommentStatus = async (id: number, status: string) => {
  const response = await request.put<Comment>(`/comments/${id}/status`, { status })
  return response.data.data
}

export const deleteComment = async (id: number) => {
  await request.delete(`/comments/${id}`)
}
