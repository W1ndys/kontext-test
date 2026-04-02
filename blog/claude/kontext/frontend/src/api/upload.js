import request from './request'

export function uploadImage(file) {
  const formData = new FormData()
  formData.append('image', file)
  return request.post('/api/admin/upload', formData, {
    headers: { 'Content-Type': 'multipart/form-data' }
  })
}
