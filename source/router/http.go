//package routes
//
//import (
//	"basic/pkg/helper/resp"
//	"basic/pkg/logger"
//	"basic/source/handler"
//	"basic/source/middleware"
//	"github.com/gin-gonic/gin"
//	"net/http"
//)
//
//func NewServerHTTP(logger *logger.Logger,
//	jwt *middleware.JWT,
//	userHandler handler.UserHandler) *gin.Engine {
//
//	gin.SetMode(gin.ReleaseMode)
//	r := gin.Default()
//	r.Use(
//		middleware.CORSMiddleware(),
//	)
//
//	r.GET("/", func(ctx *gin.Context) {
//		resp.HandleSuccess(ctx, http.StatusOK, "Connected", map[string]interface{}{
//			"say": "Hi Welcome to your new API!",
//		})
//	})
//
//	BindUserRoutes(r, jwt, userHandler, logger)
//
//	return r
//}

package routes

import (
	"basic/pkg/helper/resp"
	"basic/pkg/jwt"
	"basic/pkg/logger"
	"basic/pkg/server/http"
	"basic/source/handler"
	"basic/source/middleware"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	// swagger doc
	//docs.SwaggerInfo.BasePath = "/v1"
	s.GET("/swagger/*any", ginSwagger.WrapHandler(
		swaggerfiles.Handler,
		//ginSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger/doc.json", conf.GetInt("app.http.port"))),
		ginSwagger.DefaultModelsExpandDepth(-1),
		ginSwagger.PersistAuthorization(true),
	))

	s.Use(
		middleware.CORSMiddleware(),
		middleware.ResponseLogMiddleware(logger),
		middleware.RequestLogMiddleware(logger),
		//middleware.SignMiddleware(log),
	)
	s.GET("/", func(ctx *gin.Context) {
		logger.WithContext(ctx).Info("hello")
		resp.HandleSuccess(ctx, 200, "", map[string]interface{}{
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
		strictAuthRouter := v1.Group("/").Use(middleware.StrictAuth(jwt, logger))
		{
			strictAuthRouter.PUT("/user", userHandler.UpdateProfile)
		}
	}

	return s
}
