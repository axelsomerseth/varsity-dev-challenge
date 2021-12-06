package router

import (
	"path/filepath"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"

	"github.com/axelsomerseth/varsity-dev-challenge/backend/controllers"
)

func Setup() *gin.Engine {
	router := gin.Default()
	router.RedirectTrailingSlash = true
	router.ForwardedByClientIP = true

	// CORS policy
	router.Use(cors.AllowAll())

	// router.SetTrustedProxies([]string{""})

	router.StaticFile("/", filepath.Join("web", "index.html"))
	router.StaticFS("/public/css", gin.Dir(filepath.Join("web", "public", "css"), false))
	router.StaticFS("/public/js", gin.Dir(filepath.Join("web", "public", "js"), false))
	router.StaticFile("/auth_config.json", filepath.Join("web", "auth_config.json"))

	// Health check
	router.GET("/health", controllers.Status)

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			user := v1.Group("/users")
			{
				user.POST("/", controllers.CreateUser)
				user.GET("/:userId", controllers.GetUser)
			}

			post := v1.Group("posts")
			{
				post.GET("/", controllers.ListPosts)
				post.GET("/:postId", controllers.GetPost)
				post.POST("/", controllers.CreatePost)
			}
		}
	}

	return router
}
