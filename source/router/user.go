package routes

import (
	"basic/pkg/jwt"
	"basic/pkg/logger"
	"basic/source/handler"
	"basic/source/middleware"
	"github.com/gin-gonic/gin"
)

func BindUserRoutes(router *gin.Engine, jwt *jwt.JWT, userHandler handler.UserHandler, log *logger.Logger) {
	users := router.Group("/user")
	users.POST("/login", userHandler.Login)
	users.POST("/register", userHandler.Register)

	users.Use(middleware.StrictAuth(jwt, log))
	users.GET("", userHandler.GetProfile)
}
