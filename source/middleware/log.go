package middleware

import (
	"bytes"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"

	"basic/pkg/logger"
)

func RequestLogMiddleware(logger *logger.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		trace := uuid.New().String()
		logger.WithValue(ctx, zap.String("trace", trace))
		logger.WithValue(ctx, zap.String("request_method", ctx.Request.Method))
		logger.WithValue(ctx, zap.String("request_url", ctx.Request.URL.Path))

		// Read and restore body so downstream handlers can still read it,
		// but do NOT log it — it may contain passwords or other sensitive data
		if ctx.Request.Body != nil {
			bodyBytes, _ := ctx.GetRawData()
			ctx.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}

		logger.WithContext(ctx).Info("Request")
		ctx.Next()
	}
}

func ResponseLogMiddleware(logger *logger.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		startTime := time.Now()
		ctx.Next()
		duration := time.Since(startTime).String()
		// Log status code and duration only — not the response body (may contain PII)
		logger.WithContext(ctx).Info("Response",
			zap.Int("status", ctx.Writer.Status()),
			zap.String("duration", duration),
		)
	}
}
