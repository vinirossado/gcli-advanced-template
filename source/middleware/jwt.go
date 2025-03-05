package middleware

import (
	"basic/pkg/helper/resp"
	"basic/pkg/jwt"
	"basic/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// StrictAuth is a middleware that requires a valid JWT token to access the route.
// If the token is missing or invalid, it responds with an unauthorized error.
func StrictAuth(j *jwt.JWT, logger *logger.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := getTokenString(ctx)
		if tokenString == "" {
			handleUnauthorized(ctx, logger, "No token")
			return
		}

		claims, err := j.ParseToken(tokenString)
		if err != nil {
			handleUnauthorized(ctx, logger, "Token error", err)
			return
		}

		ctx.Set("claims", claims)
		recoveryLoggerFunc(ctx, logger)
		ctx.Next()
	}
}

// NoStrictAuth is a middleware that allows access to the route with or without a valid JWT token.
// If the token is present and valid, it sets the claims in the context. Otherwise, it proceeds without authentication.
func NoStrictAuth(j *jwt.JWT, logger *logger.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := getTokenString(ctx)
		if tokenString == "" {
			ctx.Next()
			return
		}

		claims, err := j.ParseToken(tokenString)
		if err != nil {
			ctx.Next()
			return
		}

		ctx.Set("claims", claims)
		recoveryLoggerFunc(ctx, logger)
		ctx.Next()
	}
}

// getTokenString retrieves the JWT token from the request header, cookie, or query parameters.
func getTokenString(ctx *gin.Context) string {
	tokenString := ctx.Request.Header.Get("Authorization")
	if tokenString == "" {
		tokenString, _ = ctx.Cookie("accessToken")
	}
	if tokenString == "" {
		tokenString = ctx.Query("accessToken")
	}
	return tokenString
}

// handleUnauthorized handles unauthorized access by logging the event and responding with an error.
func handleUnauthorized(ctx *gin.Context, logger *logger.Logger, message string, err ...error) {
	logger.WithContext(ctx).Warn(message, zap.Any("data", map[string]interface{}{
		"url":    ctx.Request.URL,
		"params": ctx.Params,
	}), zap.Errors("errors", err))
	resp.HandleError(ctx, http.StatusUnauthorized, "Unauthorized", nil)
	ctx.Abort()
}

// recoveryLoggerFunc logs the user ID from the JWT claims if available.
func recoveryLoggerFunc(ctx *gin.Context, logger *logger.Logger) {
	if userInfo, ok := ctx.MustGet("claims").(*jwt.MyCustomClaims); ok {
		logger.WithValue(ctx, zap.String("UserId", userInfo.UserID))
	}
}
