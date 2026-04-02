import request from './request'

export function getTags() {
  return request.get('/api/tags')
}

export function createTag(data) {
  return request.post('/api/admin/tags', data)
}

export function updateTag(id, data) {
  return request.put(`/api/admin/tags/${id}`, data)
}

export function deleteTag(id) {
  return request.delete(`/api/admin/tags/${id}`)
}
