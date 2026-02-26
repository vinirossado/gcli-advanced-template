package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

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
	userHandler *handler.UserHandler,
) *httpserver.Server {
	// Use release mode in production; debug mode otherwise
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
		middleware.ResponseLogMiddleware(logger),
		middleware.RequestLogMiddleware(logger),
	)

	// Health check — required by load balancers and Kubernetes probes
	s.GET("/health", func(ctx *gin.Context) {
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
		// No route group has permission
		noAuthRouter := v1.Group("/")
		{
			noAuthRouter.POST("/register", userHandler.Register)
			noAuthRouter.POST("/login", userHandler.Login)
		}
		// Non-strict permission routing group
		noStrictAuthRouter := v1.Group("/").Use(middleware.NoStrictAuth(jwt, logger))
		{
			noStrictAuthRouter.GET("/user", userHandler.GetProfile)
		}

		// Strict permission routing group
		strictAuthRouter := v1.Group("/")
		strictAuthRouter.Use(middleware.StrictAuth(jwt, logger))
		{
			strictAuthRouter.PUT("/user", userHandler.UpdateProfile)
		}
	}

	return s
}
