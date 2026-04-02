# 博客系统

一个使用 Vue 3 + Go + SQLite 构建的轻量级个人博客系统。

## 功能特性

- 前台展示：首页、文章列表、文章详情、分类/标签筛选、评论功能
- 后台管理：文章编辑（支持 Markdown）、分类管理、标签管理、评论审核
- 图片上传：支持复制粘贴上传图片
- 响应式设计：适配移动端和桌面端

## 技术栈

### 前端
- Vue 3 + TypeScript
- Vite
- TailwindCSS
- Vue Router
- Pinia
- marked + DOMPurify

### 后端
- Go + Gin
- GORM + SQLite
- JWT 认证

## 快速开始

### 后端

```bash
cd backend
go mod tidy
go run main.go
```

### 前端

```bash
cd frontend
npm install
npm run dev
```

### Docker 部署

```bash
docker-compose up -d
```

## 默认账号

- 用户名：admin
- 密码：admin123

## API 接口

### 公开接口
- `GET /api/v1/articles` - 获取文章列表
- `GET /api/v1/articles/:id` - 获取文章详情
- `GET /api/v1/categories` - 获取分类列表
- `GET /api/v1/tags` - 获取标签列表
- `GET /api/v1/articles/:id/comments` - 获取评论列表
- `POST /api/v1/comments` - 提交评论

### 管理接口（需认证）
- `POST /api/v1/auth/login` - 登录
- `POST /api/v1/articles` - 创建文章
- `PUT /api/v1/articles/:id` - 更新文章
- `DELETE /api/v1/articles/:id` - 删除文章
- `POST /api/v1/admin/upload` - 上传图片
