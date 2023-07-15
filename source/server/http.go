package server

import (
	"basic/pkg/helper/resp"
	logger "basic/pkg/logger"
	"basic/source/handler"
	"basic/source/middleware"
	"basic/source/routes"
	"github.com/gin-gonic/gin"
)

func NewServerHTTP(
	logger *logger.Logger,
	userHandler handler.UserHandler,
) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(
		middleware.CORSMiddleware(),
	)

	routes.BindUserRoutes(r, userHandler)

	r.GET("/", func(ctx *gin.Context) {
		resp.HandleSuccess(ctx, map[string]interface{}{
			"say": "Hi sua-mae!",
		})
	})

	return r
}
