import { createRouter, createWebHistory } from 'vue-router'
import type { RouteRecordRaw } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    name: 'Home',
    component: () => import('@/views/Home.vue')
  },
  {
    path: '/articles',
    name: 'ArticleList',
    component: () => import('@/views/ArticleList.vue')
  },
  {
    path: '/article/:slug',
    name: 'ArticleDetail',
    component: () => import('@/views/ArticleDetail.vue')
  },
  {
    path: '/category/:slug',
    name: 'Category',
    component: () => import('@/views/Category.vue')
  },
  {
    path: '/tag/:slug',
    name: 'Tag',
    component: () => import('@/views/Tag.vue')
  },
  {
    path: '/about',
    name: 'About',
    component: () => import('@/views/About.vue')
  },
  {
    path: '/admin/login',
    name: 'Login',
    component: () => import('@/views/admin/Login.vue')
  },
  {
    path: '/admin',
    component: () => import('@/views/admin/Layout.vue'),
    meta: { requiresAuth: true },
    children: [
      {
        path: '',
        redirect: '/admin/articles'
      },
      {
        path: 'articles',
        name: 'ArticleManage',
        component: () => import('@/views/admin/ArticleManage.vue')
      },
      {
        path: 'articles/new',
        name: 'ArticleCreate',
        component: () => import('@/views/admin/ArticleEdit.vue')
      },
      {
        path: 'articles/:id/edit',
        name: 'ArticleEdit',
        component: () => import('@/views/admin/ArticleEdit.vue')
      },
      {
        path: 'categories',
        name: 'CategoryManage',
        component: () => import('@/views/admin/CategoryManage.vue')
      },
      {
        path: 'tags',
        name: 'TagManage',
        component: () => import('@/views/admin/TagManage.vue')
      },
      {
        path: 'comments',
        name: 'CommentManage',
        component: () => import('@/views/admin/CommentManage.vue')
      }
    ]
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes,
  scrollBehavior(to, from, savedPosition) {
    if (savedPosition) {
      return savedPosition
    } else {
      return { top: 0 }
    }
  }
})

router.beforeEach((to, from, next) => {
  if (to.meta.requiresAuth) {
    const authStore = useAuthStore()
    if (!authStore.isAuthenticated) {
      next({ name: 'Login', query: { redirect: to.fullPath } })
      return
    }
  }
  next()
})

export default router
