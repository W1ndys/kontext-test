import { createRouter, createWebHistory } from 'vue-router'
import { useUserStore } from '@/stores/user'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: () => import('@/views/blog/Home.vue')
  },
  {
    path: '/article/:id',
    name: 'ArticleDetail',
    component: () => import('@/views/blog/ArticleDetail.vue')
  },
  {
    path: '/categories',
    name: 'CategoryArchive',
    component: () => import('@/views/blog/CategoryArchive.vue')
  },
  {
    path: '/tags',
    name: 'TagArchive',
    component: () => import('@/views/blog/TagArchive.vue')
  },
  {
    path: '/timeline',
    name: 'Timeline',
    component: () => import('@/views/blog/Timeline.vue')
  },
  {
    path: '/about',
    name: 'About',
    component: () => import('@/views/blog/About.vue')
  },
  {
    path: '/search',
    name: 'Search',
    component: () => import('@/views/blog/Search.vue')
  },
  {
    path: '/admin/login',
    name: 'AdminLogin',
    component: () => import('@/views/admin/Login.vue')
  },
  {
    path: '/admin',
    redirect: '/admin/dashboard',
    meta: { requiresAuth: true },
    children: [
      {
        path: 'dashboard',
        name: 'AdminDashboard',
        component: () => import('@/views/admin/Dashboard.vue'),
        meta: { requiresAuth: true }
      },
      {
        path: 'articles',
        name: 'AdminArticleList',
        component: () => import('@/views/admin/ArticleList.vue'),
        meta: { requiresAuth: true }
      },
      {
        path: 'articles/create',
        name: 'AdminArticleCreate',
        component: () => import('@/views/admin/ArticleEdit.vue'),
        meta: { requiresAuth: true }
      },
      {
        path: 'articles/edit/:id',
        name: 'AdminArticleEdit',
        component: () => import('@/views/admin/ArticleEdit.vue'),
        meta: { requiresAuth: true }
      },
      {
        path: 'categories',
        name: 'AdminCategoryManage',
        component: () => import('@/views/admin/CategoryManage.vue'),
        meta: { requiresAuth: true }
      },
      {
        path: 'tags',
        name: 'AdminTagManage',
        component: () => import('@/views/admin/TagManage.vue'),
        meta: { requiresAuth: true }
      }
    ]
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes,
  scrollBehavior() {
    return { top: 0 }
  }
})

router.beforeEach((to, from, next) => {
  if (to.meta.requiresAuth) {
    const userStore = useUserStore()
    if (!userStore.isLoggedIn) {
      next({ path: '/admin/login', query: { redirect: to.fullPath } })
    } else {
      next()
    }
  } else {
    next()
  }
})

export default router
