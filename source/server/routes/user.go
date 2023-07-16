package routes

import (
	"basic/source/handler"
	"github.com/gin-gonic/gin"
)

func BindUserRoutes(router *gin.Engine, userHandler handler.UserHandler) {
	users := router.Group("/user")

	//users.GET("", userHandler.)
	users.GET("", userHandler.GetProfile)
	users.POST("/register", userHandler.Register)
	users.POST("/login", userHandler.Login)
	//users.PATCH("/:id", userHandler.DeleteUser)

	//users.Use(middlewares.JwtMiddleware().MiddlewareFunc())
	//users.GET("", middlewares.AuthorizationMiddleware(enumerations.NORMAL), controllers.FindUsers)
	//users.GET("/:id", middlewares.AuthorizationMiddleware(enumerations.NORMAL), controllers.FindUserById)
}
