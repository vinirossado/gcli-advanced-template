package routes

import (
	"basic/pkg/helper/resp"
	"basic/pkg/logger"
	"basic/source/handler"
	"basic/source/middleware"
	"github.com/gin-gonic/gin"
)

func NewServerHTTP(logger *logger.Logger,
	jwt *middleware.JWT,
	userHandler handler.UserHandler) *gin.Engine {

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(
		middleware.CORSMiddleware(),
	)

	r.GET("/", func(ctx *gin.Context) {
		resp.HandleSuccess(ctx, map[string]interface{}{
			"say": "Hi sua-mae!",
		})
	})

	BindUserRoutes(r, jwt, userHandler, logger)

	return r
}
