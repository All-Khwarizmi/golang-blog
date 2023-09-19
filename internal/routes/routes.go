package routes

import (
	"github.com/All-Khwarizmi/golang-blog/internal/controllers"
	"github.com/All-Khwarizmi/golang-blog/internal/middlewares"
	"github.com/All-Khwarizmi/golang-blog/internal/web/pages"
	"github.com/gin-gonic/gin"
)

func Auth(sm *gin.Engine) {
	// Protected routes
	// Set environement variable

	protected := sm.Group("/protected")
	protected.Use(middlewares.Auth())
	protected.GET("/home-page", pages.ProtectedHome)
}

func Public(sm *gin.Engine) {
	sm.LoadHTMLGlob("../../internal/web/templates/*")
	sm.GET("/index", pages.Login)
	sm.POST("/login", controllers.LoginHandler)
}
