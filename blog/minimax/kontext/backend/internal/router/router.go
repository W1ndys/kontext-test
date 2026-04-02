package router

import (
	"blog/internal/handler"
	"blog/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.New()
	r.Use(middleware.Recovery())
	r.Use(middleware.Logger())
	r.Use(middleware.CORS())

	r.Static("/uploads", "./uploads")

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	v1 := r.Group("/api/v1")
	{
		authHandler := handler.NewAuthHandler()
		v1.POST("/auth/register", authHandler.Register)
		v1.POST("/auth/login", authHandler.Login)
		v1.GET("/auth/current", authHandler.GetCurrentUser)

		articleHandler := handler.NewArticleHandler()
		v1.GET("/articles", articleHandler.List)
		v1.GET("/articles/:id", articleHandler.Get)
		v1.GET("/articles/slug/:slug", articleHandler.GetBySlug)

		categoryHandler := handler.NewCategoryHandler()
		v1.GET("/categories", categoryHandler.List)
		v1.GET("/categories/:id", categoryHandler.Get)

		tagHandler := handler.NewTagHandler()
		v1.GET("/tags", tagHandler.List)
		v1.GET("/tags/:id", tagHandler.Get)

		commentHandler := handler.NewCommentHandler()
		v1.GET("/articles/:id/comments", commentHandler.ListByArticle)
		v1.POST("/comments", commentHandler.Create)
	}

	admin := r.Group("/api/v1/admin")
	admin.Use(middleware.JWTAuth())
	{
		articleHandler := handler.NewArticleHandler()
		admin.GET("/articles", articleHandler.ListAll)
		admin.POST("/articles", articleHandler.Create)
		admin.PUT("/articles/:id", articleHandler.Update)
		admin.DELETE("/articles/:id", articleHandler.Delete)

		categoryHandler := handler.NewCategoryHandler()
		admin.POST("/categories", categoryHandler.Create)
		admin.PUT("/categories/:id", categoryHandler.Update)
		admin.DELETE("/categories/:id", categoryHandler.Delete)

		tagHandler := handler.NewTagHandler()
		admin.POST("/tags", tagHandler.Create)
		admin.PUT("/tags/:id", tagHandler.Update)
		admin.DELETE("/tags/:id", tagHandler.Delete)

		commentHandler := handler.NewCommentHandler()
		admin.GET("/comments", commentHandler.List)
		admin.PUT("/comments/:id/status", commentHandler.UpdateStatus)
		admin.DELETE("/comments/:id", commentHandler.Delete)
	}

	uploadHandler := handler.NewUploadHandler()
	admin.POST("/upload", uploadHandler.UploadImage)

	return r
}
