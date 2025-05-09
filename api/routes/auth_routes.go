package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/t2469/attendance-system.git/controllers"
	"github.com/t2469/attendance-system.git/middleware"
)

func addAuthRoutes(router *gin.Engine) {
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)

	auth := router.Group("/", middleware.AuthMiddleware())
	{
		auth.POST("/logout", controllers.Logout)
		auth.GET("/current_account", controllers.CurrentAccount)
	}
}
