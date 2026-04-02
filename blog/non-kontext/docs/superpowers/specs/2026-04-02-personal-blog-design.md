# 个人博客系统设计文档

## 概述

单用户个人博客系统，支持 Markdown 写作和粘贴上传图片。前后端分离架构，Go 服务同时托管 API 和前端静态文件，Docker 部署。

## 技术选型

| 层级 | 技术 |
|------|------|
| 前端框架 | Vue 3 + Vite |
| 前端 UI | TailwindCSS |
| 前端路由 | Vue Router |
| 状态管理 | Pinia |
| Markdown 编辑器 | md-editor-v3 |
| Markdown 渲染 | markdown-it（md-editor-v3 内置） |
| 后端 | Go + net/http + gorilla/mux |
| 数据库 | SQLite（modernc.org/sqlite，纯 Go 无 CGO） |
| ORM | gorm.io/gorm + gorm.io/driver/sqlite |
| 认证 | JWT（golang-jwt/jwt） |
| 密码哈希 | bcrypt |
| 部署 | Docker + docker-compose |

## 项目结构

```
blog/
├── frontend/
│   ├── src/
│   │   ├── views/         # 页面组件
│   │   ├── components/    # 通用组件
│   │   ├── router/        # Vue Router 配置
│   │   ├── api/           # Axios API 封装
│   │   └── stores/        # Pinia 状态管理
│   ├── package.json
│   └── vite.config.js
├── backend/
│   ├── main.go            # 入口
│   ├── handler/           # HTTP 处理函数
│   ├── model/             # 数据模型
│   ├── middleware/         # JWT 中间件
│   └── go.mod
├── data/                  # 运行时数据（Docker volume）
│   ├── blog.db
│   └── uploads/
├── Dockerfile
└── docker-compose.yml
```

## 数据模型

### articles 表

| 字段 | 类型 | 说明 |
|------|------|------|
| id | INTEGER PK | 自增主键 |
| title | TEXT | 文章标题 |
| content | TEXT | Markdown 原文 |
| category | TEXT | 分类 |
| tags | TEXT | 标签（逗号分隔） |
| created_at | DATETIME | 创建时间 |
| updated_at | DATETIME | 更新时间 |

### admin 表

| 字段 | 类型 | 说明 |
|------|------|------|
| id | INTEGER PK | 固定为 1（单用户） |
| username | TEXT | 用户名 |
| password | TEXT | bcrypt 哈希 |

## API 设计

### 公开接口（前台）

- `GET /api/articles` — 文章列表，支持 `?category=&tag=&page=&size=` 分页查询
- `GET /api/articles/:id` — 文章详情
- `GET /api/categories` — 分类列表
- `GET /api/tags` — 标签列表

### 认证接口

- `POST /api/login` — 登录，请求体 `{username, password}`，返回 `{token}`

### 管理接口（需 JWT，Authorization: Bearer）

- `POST /api/admin/articles` — 新建文章
- `PUT /api/admin/articles/:id` — 编辑文章
- `DELETE /api/admin/articles/:id` — 删除文章
- `POST /api/admin/upload` — 上传图片（multipart/form-data），返回 `{url}`

### 静态资源

- `GET /uploads/:filename` — 上传的图片
- 其余路径 — 前端 SPA（index.html）

## 前端页面

### 前台（公开）

- **首页** `/` — 文章列表按时间倒序分页，右侧分类/标签侧边栏
- **文章详情** `/article/:id` — Markdown 渲染，标题/时间/分类/标签
- **分类筛选** `/category/:name` — 按分类过滤
- **标签筛选** `/tag/:name` — 按标签过滤

### 后台（需登录）

- **登录页** `/login` — 用户名/密码表单
- **文章管理** `/admin/articles` — 文章列表表格，编辑/删除
- **文章编辑** `/admin/articles/new` 和 `/admin/articles/edit/:id` — md-editor-v3 编辑器，标题/分类/标签输入，粘贴图片自动上传

### UI 风格

- 极简设计，黑白灰为主
- 大量留白，内容区最大宽度 720px 居中
- 系统默认字体

## 认证流程

1. 管理员提交用户名密码到 `POST /api/login`
2. 后端 bcrypt 校验，成功返回 JWT（有效期 24 小时）
3. 前端存 JWT 到 localStorage，请求通过 `Authorization: Bearer <token>` 携带
4. Go 中间件校验 JWT，失败返回 401
5. 前端路由守卫：访问 `/admin/*` 检查 JWT，无效跳转登录页
6. 前端收到 401 响应自动跳转登录页

## 初始化

- 首次启动自动建表
- 内置默认管理员 `admin / admin123`
- SQLite 文件 `data/blog.db`，上传图片 `data/uploads/`

## Docker 部署

多阶段构建：

1. **Stage 1（frontend）**：node:18-alpine，npm install + npm run build
2. **Stage 2（backend）**：golang:1.22-alpine，go build
3. **Stage 3（runtime）**：alpine，复制二进制和前端 dist

docker-compose 挂载 `./data:/app/data` 持久化数据库和图片。

端口：8080。
