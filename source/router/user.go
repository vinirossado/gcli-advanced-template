package routes

import (
	"basic/pkg/jwt"
	"basic/pkg/logger"
	"basic/source/handler"
	"net/http"
)

func BindUserRoutes(mux *http.ServeMux, jwt *jwt.JWT, userHandler handler.UserHandler, log *logger.Logger) {

	mux.HandleFunc("/user/login", userHandler.Login)
	mux.HandleFunc("/user/register", userHandler.Register)

	// mux.Handle("/user/profile", middleware.StrictAuth(jwt, log, http.HandlerFunc(userHandler.GetProfile)))
}


