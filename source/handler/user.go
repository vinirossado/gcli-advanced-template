package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"basic/pkg/helper/resp"
	"basic/source/service"
)

type UserHandler struct {
	*Handler
	userService service.UserService
}

func NewUserHandler(handler *Handler, userService service.UserService) *UserHandler {
	return &UserHandler{
		Handler:     handler,
		userService: userService,
	}
}

func (h *UserHandler) Register(ctx *gin.Context) {
	req := new(service.RegisterRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		resp.HandleError(ctx, http.StatusBadRequest, errors.Wrap(err, "invalid request").Error(), nil)
		return
	}

	if err := h.userService.Register(ctx, req); err != nil {
		h.logger.WithContext(ctx).Error("userService.Register error", zap.Error(err))
		resp.HandleError(ctx, http.StatusBadRequest, errors.Wrap(err, "invalid request").Error(), nil)
		return
	}

	resp.HandleSuccess(ctx, http.StatusCreated, "New account has been created.", nil)
}

func (h *UserHandler) Login(ctx *gin.Context) {
	var req service.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		resp.HandleError(ctx, http.StatusBadRequest, errors.Wrap(err, "invalid request").Error(), nil)
		return
	}

	token, err := h.userService.Login(ctx, &req)
	if err != nil {
		resp.HandleError(ctx, http.StatusUnauthorized, err.Error(), nil)
		return
	}

	resp.HandleSuccess(ctx, http.StatusOK, "Login Success", gin.H{
		"accessToken": token,
	})
}

func (h *UserHandler) GetProfile(ctx *gin.Context) {
	userID := GetUserIDFromCtx(ctx)
	if userID == "" {
		resp.HandleError(ctx, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}

	user, err := h.userService.GetProfile(ctx, userID)
	if err != nil {
		resp.HandleError(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	resp.HandleSuccess(ctx, http.StatusOK, "Get Profile", gin.H{
		"user": user,
	})
}

func (h *UserHandler) UpdateProfile(ctx *gin.Context) {
	userID := GetUserIDFromCtx(ctx)

	var req service.UpdateProfileRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		resp.HandleError(ctx, http.StatusBadRequest, errors.Wrap(err, "invalid request").Error(), nil)
		return
	}

	if err := h.userService.UpdateProfile(ctx, userID, &req); err != nil {
		resp.HandleError(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	resp.HandleSuccess(ctx, http.StatusOK, "Updated User", gin.H{
		"updated": "true",
	})
}
