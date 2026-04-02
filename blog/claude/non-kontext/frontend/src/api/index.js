import axios from 'axios'

const api = axios.create({
  baseURL: '/api',
})

api.interceptors.request.use((config) => {
  const token = localStorage.getItem('token')
  if (token) {
    config.headers.Authorization = `Bearer ${token}`
  }
  return config
})

api.interceptors.response.use(
  (response) => response,
  (error) => {
    if (error.response?.status === 401) {
      localStorage.removeItem('token')
      window.location.href = '/login'
    }
    return Promise.reject(error)
  }
)

export function login(username, password) {
  return api.post('/login', { username, password })
}

export function getArticles(params = {}) {
  return api.get('/articles', { params })
}

export function getArticle(id) {
  return api.get(`/articles/${id}`)
}

export function getCategories() {
  return api.get('/categories')
}

export function getTags() {
  return api.get('/tags')
}

export function createArticle(data) {
  return api.post('/admin/articles', data)
}

export function updateArticle(id, data) {
  return api.put(`/admin/articles/${id}`, data)
}

export function deleteArticle(id) {
  return api.delete(`/admin/articles/${id}`)
}

export function uploadImage(file) {
  const formData = new FormData()
  formData.append('image', file)
  return api.post('/admin/upload', formData, {
    headers: { 'Content-Type': 'multipart/form-data' },
  })
}

export default api
