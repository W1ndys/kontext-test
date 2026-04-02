import request from '@/utils/request'
import type { Tag } from '@/types'

export const getTagList = async () => {
  const response = await request.get<Tag[]>('/tags')
  return response.data.data
}

export const getTagById = async (id: number) => {
  const response = await request.get<Tag>(`/tags/${id}`)
  return response.data.data
}

export const createTag = async (data: { name: string; slug?: string }) => {
  const response = await request.post<Tag>('/tags', data)
  return response.data.data
}

export const updateTag = async (id: number, data: { name?: string; slug?: string }) => {
  const response = await request.put<Tag>(`/tags/${id}`, data)
  return response.data.data
}

export const deleteTag = async (id: number) => {
  await request.delete(`/tags/${id}`)
}

export const getTagBySlug = async (slug: string) => {
  const response = await request.get<Tag[]>(`/tags`, { params: { slug } })
  return response.data.data?.[0] || null
}
