import request from './request'

export function getArticleList(params) {
  return request.get('/api/articles', { params })
}

export function getArticleDetail(id) {
  return request.get(`/api/articles/${id}`)
}

export function createArticle(data) {
  return request.post('/api/admin/articles', data)
}

export function updateArticle(id, data) {
  return request.put(`/api/admin/articles/${id}`, data)
}

export function deleteArticle(id) {
  return request.delete(`/api/admin/articles/${id}`)
}

export function searchArticles(params) {
  return request.get('/api/search', { params })
}

export function getTimeline() {
  return request.get('/api/timeline')
}
