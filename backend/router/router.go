package router

import (
	"github.com/axelsomerseth/varsity-dev-challenge/backend/controllers"
	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	router := gin.Default()
	router.RedirectTrailingSlash = true
	router.ForwardedByClientIP = true

	router.SetTrustedProxies([]string{""})

	// Health check
	router.GET("/health", controllers.Status)

	return router
}
