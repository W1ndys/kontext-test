import request from '@/utils/request'
import type { Article, ArticleListQuery, CreateArticleDTO, UpdateArticleDTO, PageResult } from '@/types'

export const getArticleList = async (params: ArticleListQuery = {}) => {
  const response = await request.get<PageResult<Article>>('/articles', { params })
  return response.data.data
}

export const getArticleById = async (id: number) => {
  const response = await request.get<Article>(`/articles/${id}`)
  return response.data.data
}

export const getArticleBySlug = async (slug: string) => {
  const response = await request.get<Article>(`/articles/slug/${slug}`)
  return response.data.data
}

export const getAllArticles = async () => {
  const response = await request.get<Article[]>('/articles')
  return response.data.data
}

export const createArticle = async (data: CreateArticleDTO) => {
  const response = await request.post<Article>('/articles', data)
  return response.data.data
}

export const updateArticle = async (id: number, data: UpdateArticleDTO) => {
  const response = await request.put<Article>(`/articles/${id}`, data)
  return response.data.data
}

export const deleteArticle = async (id: number) => {
  await request.delete(`/articles/${id}`)
}
