import { defineStore } from 'pinia'
import { ref } from 'vue'
import type { Article, ArticleListQuery } from '@/types'
import { getArticleList as apiGetList, getArticleById as apiGetById, getArticleBySlug as apiGetBySlug } from '@/api/article'

export const useArticleStore = defineStore('article', () => {
  const articles = ref<Article[]>([])
  const currentArticle = ref<Article | null>(null)
  const total = ref(0)
  const loading = ref(false)

  async function fetchArticles(params: ArticleListQuery = {}) {
    loading.value = true
    try {
      const response = await apiGetList(params)
      articles.value = response.list
      total.value = response.total
      return response
    } finally {
      loading.value = false
    }
  }

  async function fetchArticleById(id: number) {
    loading.value = true
    try {
      currentArticle.value = await apiGetById(id)
      return currentArticle.value
    } finally {
      loading.value = false
    }
  }

  async function fetchArticleBySlug(slug: string) {
    loading.value = true
    try {
      currentArticle.value = await apiGetBySlug(slug)
      return currentArticle.value
    } finally {
      loading.value = false
    }
  }

  function clearCurrent() {
    currentArticle.value = null
  }

  return {
    articles,
    currentArticle,
    total,
    loading,
    fetchArticles,
    fetchArticleById,
    fetchArticleBySlug,
    clearCurrent
  }
})
