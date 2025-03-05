package handler

import (
	"encoding/json"
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

// Register godoc
// @Summary Register new user
// @Description Register new user
// @Tags user
// @Accept json
// @Produce json
// @Param body body RegisterRequest true "Register Request"
// @Success 201 {string} string "New account has been created."
// @Failure 400 {string} string "invalid request"
// @Router /register [post]
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

// Login processa a autenticação do usuário
func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req service.LoginRequest

	// Decodifica o JSON da requisição
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		resp.HandleError(w, http.StatusBadRequest, errors.Wrap(err, "invalid request").Error(), nil)
		return
	}

	// Executa o login no serviço
	token, err := h.userService.Login(r.Context(), &req)
	if err != nil {
		resp.HandleError(w, http.StatusUnauthorized, err.Error(), nil)
		return
	}

	// Envia resposta JSON de sucesso
	resp.HandleSuccess(w, http.StatusOK, "Login Success", map[string]string{
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
