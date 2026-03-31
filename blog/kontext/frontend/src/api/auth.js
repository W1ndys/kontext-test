import request from './request'

export function login(data) {
  return request.post('/api/auth/login', data)
}

export function getProfile() {
  return request.get('/api/admin/profile')
}
