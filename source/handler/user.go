package handler

import (
	"basic/pkg/helper/convert"
	"basic/pkg/helper/resp"
	"basic/source/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserHandler interface {
	GetUserById(ctx *gin.Context)
	GetUsers(ctx *gin.Context)
	CreateUser(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
}

type userHandler struct {
	*Handler
	userService service.UserService
}

func NewUserHandler(handler *Handler, userService service.UserService) UserHandler {
	return &userHandler{
		Handler:     handler,
		userService: userService,
	}
}

func (h *userHandler) GetUsers(ctx *gin.Context) {
	users, err := h.userService.GetUsers()
	if err != nil {
		h.logger.Info("GetUserByID", zap.Any("user", users))
		resp.HandleError(ctx, http.StatusInternalServerError, 1, err.Error(), nil)
		return
	}
	resp.HandleSuccess(ctx, users)
}

func (h *userHandler) GetUserById(ctx *gin.Context) {
	id := convert.ConvertToInt(ctx.Params.ByName("id"))

	user, err := h.userService.GetUserById(id)
	if err != nil {
		h.logger.Info("GetUserByID", zap.Any("user", user))
		resp.HandleError(ctx, http.StatusInternalServerError, 1, err.Error(), nil)
		return
	}
	resp.HandleSuccess(ctx, user)
}

func (h *userHandler) CreateUser(ctx *gin.Context) {
	resp.HandleSuccess(ctx, nil)
}

func (h *userHandler) UpdateUser(ctx *gin.Context) {
	resp.HandleSuccess(ctx, nil)
}

func (h *userHandler) DeleteUser(ctx *gin.Context) {
	resp.HandleSuccess(ctx, nil)
}
