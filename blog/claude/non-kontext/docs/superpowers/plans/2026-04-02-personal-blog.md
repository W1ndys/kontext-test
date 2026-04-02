# Personal Blog System Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Build a single-user personal blog with Vue 3 frontend + Go backend, supporting Markdown writing and paste-to-upload images, deployed via Docker.

**Architecture:** Frontend Vue 3 SPA communicates with Go REST API. Go serves both the API and the built frontend static files. SQLite stores data, JWT handles auth. Docker multi-stage build produces a single container.

**Tech Stack:** Vue 3, Vite, TailwindCSS, Pinia, Vue Router, md-editor-v3, Go, gorilla/mux, GORM, modernc.org/sqlite, golang-jwt/jwt, bcrypt, Docker

---

## File Structure

```
frontend/
├── index.html
├── package.json
├── vite.config.js
├── tailwind.config.js
├── postcss.config.js
├── src/
│   ├── main.js
│   ├── App.vue
│   ├── api/
│   │   └── index.js              # Axios instance + API functions
│   ├── stores/
│   │   └── auth.js               # Pinia auth store
│   ├── router/
│   │   └── index.js              # Vue Router config with guards
│   ├── views/
│   │   ├── Home.vue              # Article list + sidebar
│   │   ├── ArticleDetail.vue     # Single article view
│   │   ├── Login.vue             # Admin login form
│   │   ├── AdminArticles.vue     # Article management table
│   │   └── ArticleEditor.vue     # Create/edit article with md-editor-v3
│   └── components/
│       ├── AppHeader.vue         # Site header/nav
│       ├── AppFooter.vue         # Site footer
│       ├── ArticleCard.vue       # Article preview card
│       ├── Sidebar.vue           # Categories + tags sidebar
│       └── Pagination.vue        # Pagination component

backend/
├── go.mod
├── go.sum
├── main.go                       # Entry point, server setup, static file serving
├── model/
│   ├── db.go                     # Database init, connection, auto-migrate
│   └── article.go                # Article + Admin structs and GORM models
├── handler/
│   ├── auth.go                   # Login handler
│   ├── article.go                # Article CRUD handlers
│   └── upload.go                 # Image upload handler
└── middleware/
    └── jwt.go                    # JWT auth middleware

Dockerfile
docker-compose.yml
```

---

### Task 1: Initialize Go Backend Project

**Files:**
- Create: `backend/go.mod`
- Create: `backend/main.go`

- [ ] **Step 1: Initialize Go module**

```bash
cd backend
go mod init blog-backend
```

- [ ] **Step 2: Create main.go with minimal server**

Create `backend/main.go`:

```go
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"status":"ok"}`)
	})

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

- [ ] **Step 3: Verify server starts**

```bash
cd backend
go run main.go
```

In another terminal:
```bash
curl http://localhost:8080/api/health
```

Expected: `{"status":"ok"}`

- [ ] **Step 4: Commit**

```bash
git add backend/go.mod backend/main.go
git commit -m "feat: initialize Go backend with health check endpoint"
```

---

### Task 2: Database Models and Initialization

**Files:**
- Create: `backend/model/db.go`
- Create: `backend/model/article.go`
- Modify: `backend/go.mod` (new dependencies)

- [ ] **Step 1: Install dependencies**

```bash
cd backend
go get gorm.io/gorm
go get gorm.io/driver/sqlite
go get golang.org/x/crypto/bcrypt
```

Note: `gorm.io/driver/sqlite` with modernc backend — GORM's sqlite driver automatically uses `modernc.org/sqlite` when built without CGO. Verify by checking that the import resolves without CGO.

- [ ] **Step 2: Create model/db.go**

Create `backend/model/db.go`:

```go
package model

import (
	"log"
	"os"
	"path/filepath"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dataDir := "data"
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		log.Fatalf("failed to create data directory: %v", err)
	}

	uploadsDir := filepath.Join(dataDir, "uploads")
	if err := os.MkdirAll(uploadsDir, 0755); err != nil {
		log.Fatalf("failed to create uploads directory: %v", err)
	}

	dbPath := filepath.Join(dataDir, "blog.db")
	var err error
	DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	DB.AutoMigrate(&Article{}, &Admin{})

	var count int64
	DB.Model(&Admin{}).Count(&count)
	if count == 0 {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
		DB.Create(&Admin{
			Username: "admin",
			Password: string(hashedPassword),
		})
		log.Println("Default admin account created: admin / admin123")
	}
}
```

- [ ] **Step 3: Create model/article.go**

Create `backend/model/article.go`:

```go
package model

import "time"

type Article struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Category  string    `json:"category"`
	Tags      string    `json:"tags"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Admin struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username"`
	Password string `json:"-"`
}
```

- [ ] **Step 4: Update main.go to init DB**

Replace `backend/main.go`:

```go
package main

import (
	"fmt"
	"log"
	"net/http"

	"blog-backend/model"
)

func main() {
	model.InitDB()

	http.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"status":"ok"}`)
	})

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

- [ ] **Step 5: Verify DB initializes**

```bash
cd backend
go run main.go
```

Expected output includes: `Default admin account created: admin / admin123`
Verify `data/blog.db` file was created.

- [ ] **Step 6: Commit**

```bash
git add backend/
git commit -m "feat: add database models and auto-initialization with default admin"
```

---

### Task 3: JWT Middleware

**Files:**
- Create: `backend/middleware/jwt.go`

- [ ] **Step 1: Install JWT dependency**

```bash
cd backend
go get github.com/golang-jwt/jwt/v5
```

- [ ] **Step 2: Create middleware/jwt.go**

Create `backend/middleware/jwt.go`:

```go
package middleware

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JWTSecret = []byte("blog-secret-key-change-in-production")

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateToken(username string) (string, error) {
	claims := Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JWTSecret)
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, `{"error":"invalid token format"}`, http.StatusUnauthorized)
			return
		}

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(parts[1], claims, func(token *jwt.Token) (interface{}, error) {
			return JWTSecret, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, `{"error":"invalid token"}`, http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "username", claims.Username)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
```

- [ ] **Step 3: Verify it compiles**

```bash
cd backend
go build ./...
```

Expected: no errors.

- [ ] **Step 4: Commit**

```bash
git add backend/middleware/ backend/go.mod backend/go.sum
git commit -m "feat: add JWT authentication middleware"
```

---

### Task 4: Auth Handler (Login)

**Files:**
- Create: `backend/handler/auth.go`

- [ ] **Step 1: Create handler/auth.go**

Create `backend/handler/auth.go`:

```go
package handler

import (
	"encoding/json"
	"net/http"

	"blog-backend/middleware"
	"blog-backend/model"

	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, `{"error":"method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"invalid request"}`, http.StatusBadRequest)
		return
	}

	var admin model.Admin
	if err := model.DB.Where("username = ?", req.Username).First(&admin).Error; err != nil {
		http.Error(w, `{"error":"invalid credentials"}`, http.StatusUnauthorized)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(req.Password)); err != nil {
		http.Error(w, `{"error":"invalid credentials"}`, http.StatusUnauthorized)
		return
	}

	token, err := middleware.GenerateToken(admin.Username)
	if err != nil {
		http.Error(w, `{"error":"failed to generate token"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(LoginResponse{Token: token})
}
```

- [ ] **Step 2: Verify it compiles**

```bash
cd backend
go build ./...
```

- [ ] **Step 3: Commit**

```bash
git add backend/handler/auth.go
git commit -m "feat: add login handler with bcrypt password verification"
```

---

### Task 5: Article CRUD Handlers

**Files:**
- Create: `backend/handler/article.go`

- [ ] **Step 1: Create handler/article.go**

Create `backend/handler/article.go`:

```go
package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"blog-backend/model"

	"github.com/gorilla/mux"
)

type ArticleListResponse struct {
	Articles []model.Article `json:"articles"`
	Total    int64           `json:"total"`
	Page     int             `json:"page"`
	Size     int             `json:"size"`
}

func GetArticles(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	size, _ := strconv.Atoi(r.URL.Query().Get("size"))
	category := r.URL.Query().Get("category")
	tag := r.URL.Query().Get("tag")

	if page < 1 {
		page = 1
	}
	if size < 1 || size > 50 {
		size = 10
	}

	query := model.DB.Model(&model.Article{})
	if category != "" {
		query = query.Where("category = ?", category)
	}
	if tag != "" {
		query = query.Where("tags LIKE ?", "%"+tag+"%")
	}

	var total int64
	query.Count(&total)

	var articles []model.Article
	query.Order("created_at DESC").Offset((page - 1) * size).Limit(size).Find(&articles)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ArticleListResponse{
		Articles: articles,
		Total:    total,
		Page:     page,
		Size:     size,
	})
}

func GetArticle(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var article model.Article
	if err := model.DB.First(&article, id).Error; err != nil {
		http.Error(w, `{"error":"article not found"}`, http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(article)
}

func GetCategories(w http.ResponseWriter, r *http.Request) {
	var categories []string
	model.DB.Model(&model.Article{}).Distinct().Pluck("category", &categories)

	// Filter out empty strings
	var result []string
	for _, c := range categories {
		if c != "" {
			result = append(result, c)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func GetTags(w http.ResponseWriter, r *http.Request) {
	var tagStrings []string
	model.DB.Model(&model.Article{}).Pluck("tags", &tagStrings)

	tagSet := make(map[string]bool)
	for _, ts := range tagStrings {
		for _, t := range strings.Split(ts, ",") {
			t = strings.TrimSpace(t)
			if t != "" {
				tagSet[t] = true
			}
		}
	}

	var result []string
	for t := range tagSet {
		result = append(result, t)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func CreateArticle(w http.ResponseWriter, r *http.Request) {
	var article model.Article
	if err := json.NewDecoder(r.Body).Decode(&article); err != nil {
		http.Error(w, `{"error":"invalid request"}`, http.StatusBadRequest)
		return
	}

	if err := model.DB.Create(&article).Error; err != nil {
		http.Error(w, `{"error":"failed to create article"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(article)
}

func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var existing model.Article
	if err := model.DB.First(&existing, id).Error; err != nil {
		http.Error(w, `{"error":"article not found"}`, http.StatusNotFound)
		return
	}

	var input model.Article
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, `{"error":"invalid request"}`, http.StatusBadRequest)
		return
	}

	model.DB.Model(&existing).Updates(model.Article{
		Title:    input.Title,
		Content:  input.Content,
		Category: input.Category,
		Tags:     input.Tags,
	})

	model.DB.First(&existing, id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(existing)
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	if err := model.DB.Delete(&model.Article{}, id).Error; err != nil {
		http.Error(w, `{"error":"failed to delete article"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "deleted"})
}
```

- [ ] **Step 2: Install gorilla/mux**

```bash
cd backend
go get github.com/gorilla/mux
```

- [ ] **Step 3: Verify it compiles**

```bash
cd backend
go build ./...
```

- [ ] **Step 4: Commit**

```bash
git add backend/handler/article.go backend/go.mod backend/go.sum
git commit -m "feat: add article CRUD handlers with pagination and filtering"
```

---

### Task 6: Image Upload Handler

**Files:**
- Create: `backend/handler/upload.go`

- [ ] **Step 1: Create handler/upload.go**

Create `backend/handler/upload.go`:

```go
package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func UploadImage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, `{"error":"method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	r.ParseMultipartForm(10 << 20) // 10MB max

	file, header, err := r.FormFile("image")
	if err != nil {
		http.Error(w, `{"error":"failed to read file"}`, http.StatusBadRequest)
		return
	}
	defer file.Close()

	ext := filepath.Ext(header.Filename)
	filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	savePath := filepath.Join("data", "uploads", filename)

	dst, err := os.Create(savePath)
	if err != nil {
		http.Error(w, `{"error":"failed to save file"}`, http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, `{"error":"failed to save file"}`, http.StatusInternalServerError)
		return
	}

	url := fmt.Sprintf("/uploads/%s", filename)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"url": url})
}
```

- [ ] **Step 2: Verify it compiles**

```bash
cd backend
go build ./...
```

- [ ] **Step 3: Commit**

```bash
git add backend/handler/upload.go
git commit -m "feat: add image upload handler with timestamp-based naming"
```

---

### Task 7: Wire Up Router and Static File Serving in main.go

**Files:**
- Modify: `backend/main.go`

- [ ] **Step 1: Replace main.go with full router setup**

Replace `backend/main.go`:

```go
package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"blog-backend/handler"
	"blog-backend/middleware"
	"blog-backend/model"

	"github.com/gorilla/mux"
)

func main() {
	model.InitDB()

	r := mux.NewRouter()

	// Public API
	r.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"ok"}`))
	}).Methods("GET")
	r.HandleFunc("/api/login", handler.Login).Methods("POST")
	r.HandleFunc("/api/articles", handler.GetArticles).Methods("GET")
	r.HandleFunc("/api/articles/{id}", handler.GetArticle).Methods("GET")
	r.HandleFunc("/api/categories", handler.GetCategories).Methods("GET")
	r.HandleFunc("/api/tags", handler.GetTags).Methods("GET")

	// Admin API (JWT protected)
	admin := r.PathPrefix("/api/admin").Subrouter()
	admin.Use(middleware.AuthMiddleware)
	admin.HandleFunc("/articles", handler.CreateArticle).Methods("POST")
	admin.HandleFunc("/articles/{id}", handler.UpdateArticle).Methods("PUT")
	admin.HandleFunc("/articles/{id}", handler.DeleteArticle).Methods("DELETE")
	admin.HandleFunc("/upload", handler.UploadImage).Methods("POST")

	// Serve uploaded images
	uploadsDir := filepath.Join("data", "uploads")
	r.PathPrefix("/uploads/").Handler(
		http.StripPrefix("/uploads/", http.FileServer(http.Dir(uploadsDir))),
	)

	// Serve frontend static files
	staticDir := "static"
	if _, err := os.Stat(staticDir); err == nil {
		r.PathPrefix("/assets/").Handler(
			http.StripPrefix("/", http.FileServer(http.Dir(staticDir))),
		)
		// SPA fallback: serve index.html for all other routes
		r.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, filepath.Join(staticDir, "index.html"))
		})
	}

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
```

- [ ] **Step 2: Test the server starts and routes work**

```bash
cd backend
go run main.go
```

Test in another terminal:
```bash
# Health check
curl http://localhost:8080/api/health
# Expected: {"status":"ok"}

# Login
curl -X POST http://localhost:8080/api/login -H "Content-Type: application/json" -d '{"username":"admin","password":"admin123"}'
# Expected: {"token":"eyJ..."}

# Get articles (empty)
curl http://localhost:8080/api/articles
# Expected: {"articles":null,"total":0,"page":1,"size":10}
```

- [ ] **Step 3: Commit**

```bash
git add backend/main.go
git commit -m "feat: wire up all routes with JWT-protected admin endpoints and static serving"
```

---

### Task 8: Initialize Vue Frontend Project

**Files:**
- Create: `frontend/` (entire scaffolded project)

- [ ] **Step 1: Create Vue project with Vite**

```bash
cd "D:/Github-projects/W1ndys/kontext-test/blog/non-kontext"
npm create vite@latest frontend -- --template vue
```

- [ ] **Step 2: Install dependencies**

```bash
cd frontend
npm install
npm install vue-router@4 pinia axios md-editor-v3
npm install -D tailwindcss @tailwindcss/vite
```

- [ ] **Step 3: Configure TailwindCSS**

Replace `frontend/src/style.css`:

```css
@import "tailwindcss";
```

Add TailwindCSS plugin to `frontend/vite.config.js`:

```js
import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import tailwindcss from '@tailwindcss/vite'

export default defineConfig({
  plugins: [
    vue(),
    tailwindcss(),
  ],
  server: {
    proxy: {
      '/api': 'http://localhost:8080',
      '/uploads': 'http://localhost:8080',
    }
  }
})
```

- [ ] **Step 4: Verify frontend starts**

```bash
cd frontend
npm run dev
```

Expected: Vite dev server starts, browser shows default Vue page.

- [ ] **Step 5: Commit**

```bash
git add frontend/
git commit -m "feat: initialize Vue frontend with Vite, TailwindCSS, and dev proxy"
```

---

### Task 9: Frontend API Layer and Auth Store

**Files:**
- Create: `frontend/src/api/index.js`
- Create: `frontend/src/stores/auth.js`

- [ ] **Step 1: Create api/index.js**

Create `frontend/src/api/index.js`:

```js
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
```

- [ ] **Step 2: Create stores/auth.js**

Create `frontend/src/stores/auth.js`:

```js
import { defineStore } from 'pinia'
import { ref, computed } from 'vue'

export const useAuthStore = defineStore('auth', () => {
  const token = ref(localStorage.getItem('token') || '')

  const isLoggedIn = computed(() => !!token.value)

  function setToken(newToken) {
    token.value = newToken
    localStorage.setItem('token', newToken)
  }

  function logout() {
    token.value = ''
    localStorage.removeItem('token')
  }

  return { token, isLoggedIn, setToken, logout }
})
```

- [ ] **Step 3: Commit**

```bash
git add frontend/src/api/ frontend/src/stores/
git commit -m "feat: add API layer with axios interceptors and Pinia auth store"
```

---

### Task 10: Vue Router with Auth Guards

**Files:**
- Create: `frontend/src/router/index.js`
- Modify: `frontend/src/main.js`
- Modify: `frontend/src/App.vue`

- [ ] **Step 1: Create router/index.js**

Create `frontend/src/router/index.js`:

```js
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
  {
    path: '/admin/articles',
    component: AdminArticles,
    meta: { requiresAuth: true },
  },
  {
    path: '/admin/articles/new',
    component: ArticleEditor,
    meta: { requiresAuth: true },
  },
  {
    path: '/admin/articles/edit/:id',
    component: ArticleEditor,
    meta: { requiresAuth: true },
  },
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
```

- [ ] **Step 2: Update main.js**

Replace `frontend/src/main.js`:

```js
import { createApp } from 'vue'
import { createPinia } from 'pinia'
import App from './App.vue'
import router from './router'
import './style.css'

const app = createApp(App)
app.use(createPinia())
app.use(router)
app.mount('#app')
```

- [ ] **Step 3: Update App.vue**

Replace `frontend/src/App.vue`:

```vue
<template>
  <div class="min-h-screen bg-white text-gray-900">
    <AppHeader />
    <main class="max-w-4xl mx-auto px-4 py-8">
      <router-view />
    </main>
    <AppFooter />
  </div>
</template>

<script setup>
import AppHeader from './components/AppHeader.vue'
import AppFooter from './components/AppFooter.vue'
</script>
```

- [ ] **Step 4: Create placeholder view files**

Create these placeholder files so the router doesn't error. Each one will be implemented in subsequent tasks.

`frontend/src/views/Home.vue`:
```vue
<template><div>Home</div></template>
```

`frontend/src/views/ArticleDetail.vue`:
```vue
<template><div>Article Detail</div></template>
```

`frontend/src/views/Login.vue`:
```vue
<template><div>Login</div></template>
```

`frontend/src/views/AdminArticles.vue`:
```vue
<template><div>Admin Articles</div></template>
```

`frontend/src/views/ArticleEditor.vue`:
```vue
<template><div>Article Editor</div></template>
```

- [ ] **Step 5: Create AppHeader.vue and AppFooter.vue**

Create `frontend/src/components/AppHeader.vue`:

```vue
<template>
  <header class="border-b border-gray-200">
    <nav class="max-w-4xl mx-auto px-4 py-4 flex justify-between items-center">
      <router-link to="/" class="text-xl font-bold text-gray-900 no-underline">
        My Blog
      </router-link>
      <div class="flex gap-4 items-center">
        <router-link to="/" class="text-gray-600 hover:text-gray-900 no-underline">
          Home
        </router-link>
        <template v-if="auth.isLoggedIn">
          <router-link to="/admin/articles" class="text-gray-600 hover:text-gray-900 no-underline">
            Admin
          </router-link>
          <button @click="handleLogout" class="text-gray-600 hover:text-gray-900">
            Logout
          </button>
        </template>
        <router-link v-else to="/login" class="text-gray-600 hover:text-gray-900 no-underline">
          Login
        </router-link>
      </div>
    </nav>
  </header>
</template>

<script setup>
import { useAuthStore } from '../stores/auth'
import { useRouter } from 'vue-router'

const auth = useAuthStore()
const router = useRouter()

function handleLogout() {
  auth.logout()
  router.push('/')
}
</script>
```

Create `frontend/src/components/AppFooter.vue`:

```vue
<template>
  <footer class="border-t border-gray-200 mt-16">
    <div class="max-w-4xl mx-auto px-4 py-6 text-center text-gray-400 text-sm">
      &copy; {{ new Date().getFullYear() }} My Blog. All rights reserved.
    </div>
  </footer>
</template>
```

- [ ] **Step 6: Verify frontend compiles**

```bash
cd frontend
npm run dev
```

Expected: no compilation errors, page shows "Home" with header and footer.

- [ ] **Step 7: Commit**

```bash
git add frontend/src/
git commit -m "feat: add Vue Router with auth guards, header, footer, and placeholder views"
```

---

### Task 11: Login Page

**Files:**
- Modify: `frontend/src/views/Login.vue`

- [ ] **Step 1: Implement Login.vue**

Replace `frontend/src/views/Login.vue`:

```vue
<template>
  <div class="flex justify-center items-center min-h-[60vh]">
    <form @submit.prevent="handleLogin" class="w-full max-w-sm space-y-4">
      <h1 class="text-2xl font-bold text-center mb-8">Login</h1>
      <div v-if="error" class="text-red-500 text-sm text-center">{{ error }}</div>
      <div>
        <label class="block text-sm text-gray-600 mb-1">Username</label>
        <input
          v-model="username"
          type="text"
          class="w-full border border-gray-300 rounded px-3 py-2 focus:outline-none focus:border-gray-500"
          required
        />
      </div>
      <div>
        <label class="block text-sm text-gray-600 mb-1">Password</label>
        <input
          v-model="password"
          type="password"
          class="w-full border border-gray-300 rounded px-3 py-2 focus:outline-none focus:border-gray-500"
          required
        />
      </div>
      <button
        type="submit"
        :disabled="loading"
        class="w-full bg-gray-900 text-white py-2 rounded hover:bg-gray-800 disabled:opacity-50"
      >
        {{ loading ? 'Logging in...' : 'Login' }}
      </button>
    </form>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '../stores/auth'
import { login } from '../api'

const username = ref('')
const password = ref('')
const error = ref('')
const loading = ref(false)
const router = useRouter()
const auth = useAuthStore()

async function handleLogin() {
  error.value = ''
  loading.value = true
  try {
    const { data } = await login(username.value, password.value)
    auth.setToken(data.token)
    router.push('/admin/articles')
  } catch (e) {
    error.value = 'Invalid username or password'
  } finally {
    loading.value = false
  }
}
</script>
```

- [ ] **Step 2: Verify login page renders**

```bash
cd frontend
npm run dev
```

Navigate to `http://localhost:5173/login`. Expected: login form renders.

- [ ] **Step 3: Commit**

```bash
git add frontend/src/views/Login.vue
git commit -m "feat: implement login page with form validation"
```

---

### Task 12: Shared Components (ArticleCard, Sidebar, Pagination)

**Files:**
- Create: `frontend/src/components/ArticleCard.vue`
- Create: `frontend/src/components/Sidebar.vue`
- Create: `frontend/src/components/Pagination.vue`

- [ ] **Step 1: Create ArticleCard.vue**

Create `frontend/src/components/ArticleCard.vue`:

```vue
<template>
  <article class="py-6 border-b border-gray-100">
    <router-link :to="`/article/${article.id}`" class="no-underline">
      <h2 class="text-xl font-semibold text-gray-900 hover:text-gray-600 mb-2">
        {{ article.title }}
      </h2>
    </router-link>
    <div class="text-sm text-gray-400 mb-3 flex gap-4">
      <span>{{ formatDate(article.created_at) }}</span>
      <router-link
        v-if="article.category"
        :to="`/category/${article.category}`"
        class="text-gray-500 hover:text-gray-700 no-underline"
      >
        {{ article.category }}
      </router-link>
    </div>
    <p class="text-gray-600 leading-relaxed">
      {{ excerpt }}
    </p>
    <div v-if="tagList.length" class="mt-3 flex gap-2 flex-wrap">
      <router-link
        v-for="tag in tagList"
        :key="tag"
        :to="`/tag/${tag}`"
        class="text-xs text-gray-500 bg-gray-100 px-2 py-1 rounded no-underline hover:bg-gray-200"
      >
        {{ tag }}
      </router-link>
    </div>
  </article>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps({
  article: { type: Object, required: true },
})

const excerpt = computed(() => {
  const text = props.article.content.replace(/[#*`>\-\[\]!()]/g, '')
  return text.length > 150 ? text.slice(0, 150) + '...' : text
})

const tagList = computed(() => {
  if (!props.article.tags) return []
  return props.article.tags.split(',').map(t => t.trim()).filter(Boolean)
})

function formatDate(dateStr) {
  return new Date(dateStr).toLocaleDateString('zh-CN')
}
</script>
```

- [ ] **Step 2: Create Sidebar.vue**

Create `frontend/src/components/Sidebar.vue`:

```vue
<template>
  <aside class="space-y-8">
    <div>
      <h3 class="text-sm font-semibold text-gray-900 uppercase tracking-wide mb-3">Categories</h3>
      <ul class="space-y-1">
        <li v-for="cat in categories" :key="cat">
          <router-link
            :to="`/category/${cat}`"
            class="text-gray-600 hover:text-gray-900 text-sm no-underline"
          >
            {{ cat }}
          </router-link>
        </li>
        <li v-if="!categories.length" class="text-gray-400 text-sm">No categories yet</li>
      </ul>
    </div>
    <div>
      <h3 class="text-sm font-semibold text-gray-900 uppercase tracking-wide mb-3">Tags</h3>
      <div class="flex flex-wrap gap-2">
        <router-link
          v-for="tag in tags"
          :key="tag"
          :to="`/tag/${tag}`"
          class="text-xs text-gray-500 bg-gray-100 px-2 py-1 rounded no-underline hover:bg-gray-200"
        >
          {{ tag }}
        </router-link>
        <span v-if="!tags.length" class="text-gray-400 text-sm">No tags yet</span>
      </div>
    </div>
  </aside>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { getCategories, getTags } from '../api'

const categories = ref([])
const tags = ref([])

onMounted(async () => {
  try {
    const [catRes, tagRes] = await Promise.all([getCategories(), getTags()])
    categories.value = catRes.data || []
    tags.value = tagRes.data || []
  } catch (e) {
    // silently fail
  }
})
</script>
```

- [ ] **Step 3: Create Pagination.vue**

Create `frontend/src/components/Pagination.vue`:

```vue
<template>
  <div v-if="totalPages > 1" class="flex justify-center gap-2 mt-8">
    <button
      @click="$emit('change', page - 1)"
      :disabled="page <= 1"
      class="px-3 py-1 border border-gray-300 rounded text-sm disabled:opacity-30 hover:bg-gray-50"
    >
      Prev
    </button>
    <span class="px-3 py-1 text-sm text-gray-500">
      {{ page }} / {{ totalPages }}
    </span>
    <button
      @click="$emit('change', page + 1)"
      :disabled="page >= totalPages"
      class="px-3 py-1 border border-gray-300 rounded text-sm disabled:opacity-30 hover:bg-gray-50"
    >
      Next
    </button>
  </div>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps({
  page: { type: Number, required: true },
  total: { type: Number, required: true },
  size: { type: Number, default: 10 },
})

defineEmits(['change'])

const totalPages = computed(() => Math.ceil(props.total / props.size))
</script>
```

- [ ] **Step 4: Commit**

```bash
git add frontend/src/components/ArticleCard.vue frontend/src/components/Sidebar.vue frontend/src/components/Pagination.vue
git commit -m "feat: add ArticleCard, Sidebar, and Pagination components"
```

---

### Task 13: Home Page (Article List)

**Files:**
- Modify: `frontend/src/views/Home.vue`

- [ ] **Step 1: Implement Home.vue**

Replace `frontend/src/views/Home.vue`:

```vue
<template>
  <div class="flex gap-12">
    <div class="flex-1 min-w-0">
      <h1 v-if="filterLabel" class="text-lg font-semibold text-gray-700 mb-6">
        {{ filterLabel }}
      </h1>
      <div v-if="loading" class="text-gray-400 py-8 text-center">Loading...</div>
      <div v-else-if="articles.length === 0" class="text-gray-400 py-8 text-center">
        No articles yet.
      </div>
      <div v-else>
        <ArticleCard v-for="article in articles" :key="article.id" :article="article" />
        <Pagination :page="page" :total="total" :size="size" @change="onPageChange" />
      </div>
    </div>
    <div class="hidden lg:block w-56 shrink-0">
      <Sidebar />
    </div>
  </div>
</template>

<script setup>
import { ref, watch, computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { getArticles } from '../api'
import ArticleCard from '../components/ArticleCard.vue'
import Sidebar from '../components/Sidebar.vue'
import Pagination from '../components/Pagination.vue'

const route = useRoute()
const router = useRouter()
const articles = ref([])
const total = ref(0)
const page = ref(1)
const size = 10
const loading = ref(false)

const filterLabel = computed(() => {
  if (route.params.name && route.path.startsWith('/category')) {
    return `Category: ${route.params.name}`
  }
  if (route.params.name && route.path.startsWith('/tag')) {
    return `Tag: ${route.params.name}`
  }
  return ''
})

async function fetchArticles() {
  loading.value = true
  try {
    const params = { page: page.value, size }
    if (route.params.name && route.path.startsWith('/category')) {
      params.category = route.params.name
    }
    if (route.params.name && route.path.startsWith('/tag')) {
      params.tag = route.params.name
    }
    const { data } = await getArticles(params)
    articles.value = data.articles || []
    total.value = data.total
  } catch (e) {
    articles.value = []
  } finally {
    loading.value = false
  }
}

function onPageChange(newPage) {
  page.value = newPage
}

watch(() => route.fullPath, () => {
  page.value = 1
  fetchArticles()
}, { immediate: true })

watch(page, fetchArticles)
</script>
```

- [ ] **Step 2: Verify home page renders**

Start both backend (`cd backend && go run main.go`) and frontend (`cd frontend && npm run dev`). Navigate to `http://localhost:5173/`.

Expected: page shows "No articles yet." with sidebar.

- [ ] **Step 3: Commit**

```bash
git add frontend/src/views/Home.vue
git commit -m "feat: implement home page with article list, filtering, and pagination"
```

---

### Task 14: Article Detail Page

**Files:**
- Modify: `frontend/src/views/ArticleDetail.vue`

- [ ] **Step 1: Implement ArticleDetail.vue**

Replace `frontend/src/views/ArticleDetail.vue`:

```vue
<template>
  <article v-if="article" class="max-w-3xl mx-auto">
    <header class="mb-8">
      <h1 class="text-3xl font-bold text-gray-900 mb-3">{{ article.title }}</h1>
      <div class="text-sm text-gray-400 flex gap-4">
        <span>{{ formatDate(article.created_at) }}</span>
        <router-link
          v-if="article.category"
          :to="`/category/${article.category}`"
          class="text-gray-500 hover:text-gray-700 no-underline"
        >
          {{ article.category }}
        </router-link>
      </div>
      <div v-if="tagList.length" class="mt-3 flex gap-2 flex-wrap">
        <router-link
          v-for="tag in tagList"
          :key="tag"
          :to="`/tag/${tag}`"
          class="text-xs text-gray-500 bg-gray-100 px-2 py-1 rounded no-underline hover:bg-gray-200"
        >
          {{ tag }}
        </router-link>
      </div>
    </header>
    <div class="prose prose-gray max-w-none" v-html="renderedContent"></div>
    <div class="mt-12">
      <router-link to="/" class="text-gray-500 hover:text-gray-700 no-underline text-sm">
        &larr; Back to Home
      </router-link>
    </div>
  </article>
  <div v-else-if="loading" class="text-gray-400 py-8 text-center">Loading...</div>
  <div v-else class="text-gray-400 py-8 text-center">Article not found.</div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { getArticle } from '../api'
import MarkdownIt from 'markdown-it'

const md = new MarkdownIt()
const route = useRoute()
const article = ref(null)
const loading = ref(true)

const renderedContent = computed(() => {
  if (!article.value) return ''
  return md.render(article.value.content)
})

const tagList = computed(() => {
  if (!article.value?.tags) return []
  return article.value.tags.split(',').map(t => t.trim()).filter(Boolean)
})

function formatDate(dateStr) {
  return new Date(dateStr).toLocaleDateString('zh-CN')
}

onMounted(async () => {
  try {
    const { data } = await getArticle(route.params.id)
    article.value = data
  } catch (e) {
    article.value = null
  } finally {
    loading.value = false
  }
})
</script>
```

- [ ] **Step 2: Install markdown-it**

```bash
cd frontend
npm install markdown-it
```

- [ ] **Step 3: Commit**

```bash
git add frontend/src/views/ArticleDetail.vue frontend/package.json frontend/package-lock.json
git commit -m "feat: implement article detail page with markdown rendering"
```

---

### Task 15: Admin Article Management Page

**Files:**
- Modify: `frontend/src/views/AdminArticles.vue`

- [ ] **Step 1: Implement AdminArticles.vue**

Replace `frontend/src/views/AdminArticles.vue`:

```vue
<template>
  <div>
    <div class="flex justify-between items-center mb-6">
      <h1 class="text-2xl font-bold">Article Management</h1>
      <router-link
        to="/admin/articles/new"
        class="bg-gray-900 text-white px-4 py-2 rounded text-sm hover:bg-gray-800 no-underline"
      >
        New Article
      </router-link>
    </div>
    <div v-if="loading" class="text-gray-400 py-8 text-center">Loading...</div>
    <table v-else class="w-full text-left">
      <thead>
        <tr class="border-b border-gray-200">
          <th class="py-2 text-sm font-semibold text-gray-600">Title</th>
          <th class="py-2 text-sm font-semibold text-gray-600 w-28">Category</th>
          <th class="py-2 text-sm font-semibold text-gray-600 w-28">Date</th>
          <th class="py-2 text-sm font-semibold text-gray-600 w-32 text-right">Actions</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="article in articles" :key="article.id" class="border-b border-gray-100">
          <td class="py-3 text-sm">{{ article.title }}</td>
          <td class="py-3 text-sm text-gray-500">{{ article.category }}</td>
          <td class="py-3 text-sm text-gray-400">{{ formatDate(article.created_at) }}</td>
          <td class="py-3 text-sm text-right space-x-3">
            <router-link
              :to="`/admin/articles/edit/${article.id}`"
              class="text-gray-600 hover:text-gray-900 no-underline"
            >
              Edit
            </router-link>
            <button
              @click="handleDelete(article.id)"
              class="text-red-500 hover:text-red-700"
            >
              Delete
            </button>
          </td>
        </tr>
        <tr v-if="articles.length === 0">
          <td colspan="4" class="py-8 text-center text-gray-400">No articles yet.</td>
        </tr>
      </tbody>
    </table>
    <Pagination :page="page" :total="total" :size="size" @change="onPageChange" />
  </div>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue'
import { getArticles, deleteArticle } from '../api'
import Pagination from '../components/Pagination.vue'

const articles = ref([])
const total = ref(0)
const page = ref(1)
const size = 20
const loading = ref(false)

async function fetchArticles() {
  loading.value = true
  try {
    const { data } = await getArticles({ page: page.value, size })
    articles.value = data.articles || []
    total.value = data.total
  } finally {
    loading.value = false
  }
}

async function handleDelete(id) {
  if (!confirm('Are you sure you want to delete this article?')) return
  await deleteArticle(id)
  fetchArticles()
}

function onPageChange(newPage) {
  page.value = newPage
}

function formatDate(dateStr) {
  return new Date(dateStr).toLocaleDateString('zh-CN')
}

watch(page, fetchArticles)
onMounted(fetchArticles)
</script>
```

- [ ] **Step 2: Commit**

```bash
git add frontend/src/views/AdminArticles.vue
git commit -m "feat: implement admin article management page with delete"
```

---

### Task 16: Article Editor with Markdown and Paste Upload

**Files:**
- Modify: `frontend/src/views/ArticleEditor.vue`

- [ ] **Step 1: Implement ArticleEditor.vue**

Replace `frontend/src/views/ArticleEditor.vue`:

```vue
<template>
  <div>
    <h1 class="text-2xl font-bold mb-6">{{ isEdit ? 'Edit Article' : 'New Article' }}</h1>
    <form @submit.prevent="handleSubmit" class="space-y-4">
      <div>
        <label class="block text-sm text-gray-600 mb-1">Title</label>
        <input
          v-model="form.title"
          type="text"
          class="w-full border border-gray-300 rounded px-3 py-2 focus:outline-none focus:border-gray-500"
          required
        />
      </div>
      <div class="flex gap-4">
        <div class="flex-1">
          <label class="block text-sm text-gray-600 mb-1">Category</label>
          <input
            v-model="form.category"
            type="text"
            class="w-full border border-gray-300 rounded px-3 py-2 focus:outline-none focus:border-gray-500"
          />
        </div>
        <div class="flex-1">
          <label class="block text-sm text-gray-600 mb-1">Tags (comma separated)</label>
          <input
            v-model="form.tags"
            type="text"
            class="w-full border border-gray-300 rounded px-3 py-2 focus:outline-none focus:border-gray-500"
            placeholder="tag1, tag2"
          />
        </div>
      </div>
      <div>
        <label class="block text-sm text-gray-600 mb-1">Content</label>
        <MdEditor
          v-model="form.content"
          language="en-US"
          :style="{ height: '500px' }"
          @onUploadImg="handleUploadImg"
        />
      </div>
      <div class="flex gap-3">
        <button
          type="submit"
          :disabled="saving"
          class="bg-gray-900 text-white px-6 py-2 rounded text-sm hover:bg-gray-800 disabled:opacity-50"
        >
          {{ saving ? 'Saving...' : 'Save' }}
        </button>
        <router-link
          to="/admin/articles"
          class="border border-gray-300 px-6 py-2 rounded text-sm text-gray-600 hover:bg-gray-50 no-underline"
        >
          Cancel
        </router-link>
      </div>
    </form>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { MdEditor } from 'md-editor-v3'
import 'md-editor-v3/lib/style.css'
import { getArticle, createArticle, updateArticle, uploadImage } from '../api'

const route = useRoute()
const router = useRouter()
const saving = ref(false)

const isEdit = computed(() => !!route.params.id)

const form = ref({
  title: '',
  content: '',
  category: '',
  tags: '',
})

async function handleUploadImg(files, callback) {
  const urls = []
  for (const file of files) {
    try {
      const { data } = await uploadImage(file)
      urls.push(data.url)
    } catch (e) {
      console.error('Upload failed:', e)
    }
  }
  callback(urls)
}

async function handleSubmit() {
  saving.value = true
  try {
    if (isEdit.value) {
      await updateArticle(route.params.id, form.value)
    } else {
      await createArticle(form.value)
    }
    router.push('/admin/articles')
  } catch (e) {
    alert('Failed to save article')
  } finally {
    saving.value = false
  }
}

onMounted(async () => {
  if (isEdit.value) {
    try {
      const { data } = await getArticle(route.params.id)
      form.value = {
        title: data.title,
        content: data.content,
        category: data.category,
        tags: data.tags,
      }
    } catch (e) {
      alert('Failed to load article')
      router.push('/admin/articles')
    }
  }
})
</script>
```

- [ ] **Step 2: Verify editor renders with paste upload**

Start backend and frontend. Login at `/login` with `admin / admin123`. Navigate to `/admin/articles/new`.

Expected: Markdown editor renders. Pasting an image in the editor should trigger upload and insert the image URL.

- [ ] **Step 3: Commit**

```bash
git add frontend/src/views/ArticleEditor.vue
git commit -m "feat: implement article editor with md-editor-v3 and paste image upload"
```

---

### Task 17: Dockerfile and docker-compose.yml

**Files:**
- Create: `Dockerfile`
- Create: `docker-compose.yml`

- [ ] **Step 1: Create Dockerfile**

Create `Dockerfile` in project root (`non-kontext/`):

```dockerfile
# Stage 1: Build frontend
FROM node:18-alpine AS frontend
WORKDIR /app/frontend
COPY frontend/package*.json ./
RUN npm install
COPY frontend/ ./
RUN npm run build

# Stage 2: Build backend
FROM golang:1.22-alpine AS backend
WORKDIR /app
COPY backend/go.mod backend/go.sum ./
RUN go mod download
COPY backend/ ./
RUN CGO_ENABLED=0 go build -o server .

# Stage 3: Runtime
FROM alpine:latest
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=backend /app/server ./server
COPY --from=frontend /app/frontend/dist ./static
RUN mkdir -p data/uploads
EXPOSE 8080
CMD ["./server"]
```

- [ ] **Step 2: Create docker-compose.yml**

Create `docker-compose.yml` in project root:

```yaml
services:
  blog:
    build: .
    ports:
      - "8080:8080"
    volumes:
      - ./data:/app/data
    restart: unless-stopped
```

- [ ] **Step 3: Create .dockerignore**

Create `.dockerignore` in project root:

```
data/
frontend/node_modules/
backend/data/
.git/
```

- [ ] **Step 4: Commit**

```bash
git add Dockerfile docker-compose.yml .dockerignore
git commit -m "feat: add Docker multi-stage build and docker-compose config"
```

---

### Task 18: End-to-End Verification

- [ ] **Step 1: Build and run with Docker**

```bash
cd "D:/Github-projects/W1ndys/kontext-test/blog/non-kontext"
docker compose up --build
```

Expected: Container builds and starts, logs show `Server starting on :8080`.

- [ ] **Step 2: Test full workflow**

Open browser at `http://localhost:8080`:

1. Home page loads with "No articles yet" message
2. Navigate to `/login`, login with `admin / admin123`
3. Redirected to `/admin/articles`
4. Click "New Article", fill in title/category/tags, write Markdown content
5. Paste an image in the editor — should upload and show
6. Save article
7. Navigate to home page — article appears in list
8. Click article — detail page with rendered Markdown
9. Test category and tag filtering from sidebar
10. Test edit and delete from admin page

- [ ] **Step 3: Final commit**

```bash
git add -A
git commit -m "feat: personal blog system complete - Vue+Go+SQLite+Docker"
```
