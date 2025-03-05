package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"basic/pkg/helper/resp"
	"basic/pkg/jwt"
	"basic/pkg/logger"
	"basic/pkg/server/http"
	"basic/source/handler"
	"basic/source/middleware"
)

func NewHTTPServer(
	logger *logger.Logger,
	conf *viper.Viper,
	jwt *jwt.JWT,
	userHandler *handler.UserHandler,

) *http.Server {
	gin.SetMode(gin.DebugMode)

	s := http.NewServer(
		gin.Default(),
		logger,
		http.WithServerHost(conf.GetString("http.host")),
		http.WithServerPort(conf.GetInt("http.port")),
	)

	healthCheck(s, logger)

	setMiddlewares(s, logger)

	BindUserRoutes(s, jwt, userHandler, logger)

	middleware.SignMiddleware(conf, logger)
	middleware.StrictAuth(jwt, logger)

	// v1 := s.Group("/v1")
	// {
	// 	// No route group has permission
	// 	noAuthRouter := v1.Group("/")
	// 	{
	// 		noAuthRouter.POST("/register", userHandler.Register)
	// 		noAuthRouter.POST("/login", userHandler.Login)
	// 	}
	// 	// Non-strict permission routing group
	// 	noStrictAuthRouter := v1.Group("/").Use(middleware.NoStrictAuth(jwt, logger))
	// 	{
	// 		noStrictAuthRouter.GET("/user", userHandler.GetProfile)
	// 		// BindUserRoutes()
	// 	}

	// 	// Strict permission routing group
	// 	strictAuthRouter := v1.Group("/").Use(middleware.StrictAuth(jwt, logger))
	// 	{
	// 		strictAuthRouter.PUT("/user", userHandler.UpdateProfile)
	// 	}
	// }

	return s
}

func setMiddlewares(s *http.Server, logger *logger.Logger) {
	s.Use(
		middleware.CORSMiddleware(),
		middleware.ResponseLogMiddleware(logger),
		middleware.RequestLogMiddleware(logger),
	)
}

func healthCheck(s *http.Server, logger *logger.Logger) {
	s.GET("/", func(ctx *gin.Context) {
		logger.WithContext(ctx).Info("hello")
		resp.HandleSuccess(ctx, 200, "Welcome to your new Golang API", map[string]interface{}{
			":)": "Thank you for using Gcli!",
		})
	})
}
