package handler

import (
	"github.com/gin-gonic/gin"

	"basic/pkg/jwt"
	"basic/pkg/logger"
)

type Handler struct {
	logger *logger.Logger
}

func NewHandler(logger *logger.Logger) *Handler {
	return &Handler{
		logger: logger,
	}
}
func GetUserIDFromCtx(ctx *gin.Context) string {
	v, exists := ctx.Get("claims")

	if !exists {
		return ""
	}
	return v.(*jwt.MyCustomClaims).UserID
}
