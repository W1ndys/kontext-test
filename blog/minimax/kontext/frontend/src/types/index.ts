export interface User {
  id: number
  username: string
  nickname: string
  avatar: string
  created_at: string
}

export interface Category {
  id: number
  name: string
  slug: string
  sort: number
}

export interface Tag {
  id: number
  name: string
  slug: string
}

export interface Article {
  id: number
  title: string
  slug: string
  content: string
  summary: string
  cover_image: string
  status: 'draft' | 'published'
  view_count: number
  category_id: number
  category: Category
  tags: Tag[]
  created_at: string
  updated_at: string
}

export interface Comment {
  id: number
  article_id: number
  nickname: string
  email: string
  content: string
  ip: string
  status: 'pending' | 'approved' | 'rejected'
  created_at: string
}

export interface LoginRequest {
  username: string
  password: string
}

export interface LoginResponse {
  token: string
  expires_at: number
  user: User
}

export interface RegisterRequest {
  username: string
  password: string
  nickname: string
}

export interface CreateArticleDTO {
  title: string
  slug?: string
  content: string
  summary?: string
  cover_image?: string
  status?: 'draft' | 'published'
  category_id?: number
  tag_ids?: number[]
}

export interface UpdateArticleDTO {
  title?: string
  slug?: string
  content?: string
  summary?: string
  cover_image?: string
  status?: 'draft' | 'published'
  category_id?: number
  tag_ids?: number[]
}

export interface CreateCommentDTO {
  article_id: number
  nickname: string
  email?: string
  content: string
}

export interface ArticleListQuery {
  page?: number
  page_size?: number
  category_id?: number
  tag_id?: number
  status?: string
}

export interface CommentListQuery {
  page?: number
  page_size?: number
  status?: string
}

export interface PageResult<T> {
  list: T[]
  total: number
  page: number
}
