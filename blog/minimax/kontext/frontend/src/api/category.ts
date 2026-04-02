import request from '@/utils/request'
import type { Category } from '@/types'

export const getCategoryList = async () => {
  const response = await request.get<Category[]>('/categories')
  return response.data.data
}

export const getCategoryById = async (id: number) => {
  const response = await request.get<Category>(`/categories/${id}`)
  return response.data.data
}

export const createCategory = async (data: { name: string; slug?: string; sort?: number }) => {
  const response = await request.post<Category>('/categories', data)
  return response.data.data
}

export const updateCategory = async (id: number, data: { name?: string; slug?: string; sort?: number }) => {
  const response = await request.put<Category>(`/categories/${id}`, data)
  return response.data.data
}

export const deleteCategory = async (id: number) => {
  await request.delete(`/categories/${id}`)
}

export const getCategoryBySlug = async (slug: string) => {
  const response = await request.get<Category[]>(`/categories`, { params: { slug } })
  return response.data.data?.[0] || null
}
