package router

import (
	"blog-backend/config"
	"blog-backend/controller"
	"blog-backend/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())
	r.Static("/uploads", config.AppConfig.UploadDir)

	public := r.Group("/api")
	{
		public.GET("/articles", controller.GetArticles)
		public.GET("/articles/:id", controller.GetArticleDetail)
		public.GET("/categories", controller.GetCategories)
		public.GET("/tags", controller.GetTags)
		public.GET("/search", controller.SearchArticles)
		public.GET("/timeline", controller.GetArticlesTimeline)
		public.POST("/auth/login", controller.Login)
		public.GET("/stats", controller.GetStats)
	}

	admin := r.Group("/api/admin")
	admin.Use(middleware.JWTAuth())
	{
		admin.GET("/profile", controller.GetProfile)
		admin.POST("/articles", controller.CreateArticle)
		admin.PUT("/articles/:id", controller.UpdateArticle)
		admin.DELETE("/articles/:id", controller.DeleteArticle)
		admin.POST("/categories", controller.CreateCategory)
		admin.PUT("/categories/:id", controller.UpdateCategory)
		admin.DELETE("/categories/:id", controller.DeleteCategory)
		admin.POST("/tags", controller.CreateTag)
		admin.PUT("/tags/:id", controller.UpdateTag)
		admin.DELETE("/tags/:id", controller.DeleteTag)
		admin.POST("/upload", controller.UploadImage)
	}

	return r
}
