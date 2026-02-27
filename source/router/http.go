package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/gorm"

	"basic/pkg/helper/resp"
	"basic/pkg/jwt"
	"basic/pkg/logger"
	httpserver "basic/pkg/server/http"
	"basic/source/handler"
	"basic/source/middleware"
)

func NewHTTPServer(
	logger *logger.Logger,
	conf *viper.Viper,
	jwt *jwt.JWT,
	db *gorm.DB,
	userHandler *handler.UserHandler,
) *httpserver.Server {
	if env := conf.GetString("env"); env == "prod" || env == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	s := httpserver.NewServer(
		gin.Default(),
		logger,
		httpserver.WithServerHost(conf.GetString("http.host")),
		httpserver.WithServerPort(conf.GetInt("http.port")),
	)

	s.Use(
		middleware.CORSMiddleware(conf.GetStringSlice("http.cors.allowed_origins")),
		middleware.RateLimitMiddleware(conf),
		middleware.ResponseLogMiddleware(logger),
		middleware.RequestLogMiddleware(logger),
	)

	// Health check — verifies DB connectivity for load balancers and k8s probes
	s.GET("/health", func(ctx *gin.Context) {
		sqlDB, err := db.DB()
		if err != nil {
			ctx.JSON(http.StatusServiceUnavailable, gin.H{"status": "unhealthy", "error": "db unavailable"})
			return
		}
		if err = sqlDB.PingContext(ctx.Request.Context()); err != nil {
			ctx.JSON(http.StatusServiceUnavailable, gin.H{"status": "unhealthy", "error": "db unreachable"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	s.GET("/", func(ctx *gin.Context) {
		logger.WithContext(ctx).Info("hello")
		resp.HandleSuccess(ctx, 200, "Welcome to your new Golang API", map[string]interface{}{
			":)": "Thank you for using Gcli!",
		})
	})

	v1 := s.Group("/v1")
	{
		noAuthRouter := v1.Group("/")
		{
			noAuthRouter.POST("/register", userHandler.Register)
			noAuthRouter.POST("/login", userHandler.Login)
		}
		noStrictAuthRouter := v1.Group("/").Use(middleware.NoStrictAuth(jwt, logger))
		{
			noStrictAuthRouter.GET("/user", userHandler.GetProfile)
		}

		strictAuthRouter := v1.Group("/")
		strictAuthRouter.Use(middleware.StrictAuth(jwt, logger))
		{
			strictAuthRouter.PUT("/user", userHandler.UpdateProfile)
		}
	}

	return s
}
