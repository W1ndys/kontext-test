import request from '@/utils/request'
import type { LoginRequest, LoginResponse, RegisterRequest, User } from '@/types'

export const login = async (data: LoginRequest) => {
  const response = await request.post<LoginResponse>('/auth/login', data)
  return response.data.data
}

export const register = async (data: RegisterRequest) => {
  const response = await request.post<User>('/auth/register', data)
  return response.data.data
}

export const getCurrentUser = async () => {
  const response = await request.get<User>('/auth/current')
  return response.data.data
}
