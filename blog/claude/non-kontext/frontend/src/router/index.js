import { createRouter, createWebHistory } from 'vue-router'
import { useAuthStore } from '../stores/auth'

import Home from '../views/Home.vue'
import ArticleDetail from '../views/ArticleDetail.vue'
import Login from '../views/Login.vue'
import AdminArticles from '../views/AdminArticles.vue'
import ArticleEditor from '../views/ArticleEditor.vue'

const routes = [
  { path: '/', component: Home },
  { path: '/article/:id', component: ArticleDetail },
  { path: '/category/:name', component: Home },
  { path: '/tag/:name', component: Home },
  { path: '/login', component: Login },
  { path: '/admin/articles', component: AdminArticles, meta: { requiresAuth: true } },
  { path: '/admin/articles/new', component: ArticleEditor, meta: { requiresAuth: true } },
  { path: '/admin/articles/edit/:id', component: ArticleEditor, meta: { requiresAuth: true } },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

router.beforeEach((to) => {
  if (to.meta.requiresAuth) {
    const auth = useAuthStore()
    if (!auth.isLoggedIn) {
      return '/login'
    }
  }
})

export default router
