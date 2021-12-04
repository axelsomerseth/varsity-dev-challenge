package router

import (
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

	router.SetTrustedProxies([]string{""})

	// Health check
	router.GET("/health", controllers.Status)

	return router
}
