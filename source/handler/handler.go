package handler

import (
	"basic/pkg/logger"
	"basic/source/middleware"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	logger *logger.Logger
}

func NewHandler(logger *logger.Logger) *Handler {
	return &Handler{
		logger: logger,
	}
}
func GetUserIdFromCtx(ctx *gin.Context) string {
	v, exists := ctx.Get("claims")

	if !exists {
		return ""
	}
	return v.(*middleware.MyCustomClaims).UserId
}
