package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gincom/pkg/errors"
	"go.uber.org/zap"

	"basic/pkg/helper/resp"
	"basic/source/service"
)

type PersonHandler struct {
	*Handler
	personService service.PersonService
}

func NewPersonHandler(handler *Handler, personService service.PersonService) *PersonHandler {
	return &PersonHandler{
		Handler:       handler,
		personService: personService,
	}
}

func (h *PersonHandler) GetPersonById(ctx *gin.Context) {
	var params struct {
		Id uint `uri:"id" binding:"required"`
	}

	if err := ctx.ShouldBindUri(&params); err != nil {
		resp.HandleError(ctx, http.StatusBadRequest, 1, err.Error(), nil)
		return
	}

	person, err := h.personService.GetPersonById(ctx, params.Id)
	h.logger.Info("GetPersonByID", zap.Any("person", person))

	if err != nil {
		resp.HandleError(ctx, http.StatusInternalServerError, 1, err.Error(), nil)
		return
	}

	resp.HandleSuccess(ctx, http.StatusOK, "Success", person)
}

func (h *PersonHandler) GetAllPerson(ctx *gin.Context) {
	person, err := h.personService.GetAllPerson(ctx)
	h.logger.Info("GetAllPerson", zap.Any("person", person))

	if err != nil {
		resp.HandleError(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	resp.HandleSuccess(ctx, http.StatusOK, "Success", person)
}

func (h *PersonHandler) CreatePerson(ctx *gin.Context) {
	req := new(service.CreatePersonRequest)

	if err := ctx.ShouldBindJSON(req); err != nil {
		resp.HandleError(ctx, http.StatusBadRequest, 1, errors.Wrap(err, "invalid request").Error(), nil)
		return
	}

	if _, err := h.personService.CreatePerson(ctx, req); err != nil {
		resp.HandleError(ctx, http.StatusBadRequest, 1, errors.Wrap(err, "invalid request").Error(), nil)
		return
	}

	resp.HandleSuccess(ctx, http.StatusOK, "Success", nil)
}

func (h *PersonHandler) UpdatePerson(ctx *gin.Context) {
	req := new(service.UpdatePersonRequest)

	if err := ctx.ShouldBindJSON(req); err != nil {
		resp.HandleError(ctx, http.StatusBadRequest, 1, errors.Wrap(err, "invalid request").Error(), nil)
		return
	}

	if _, err := h.personService.UpdatePerson(ctx, req); err != nil {
		resp.HandleError(ctx, http.StatusBadRequest, 1, errors.Wrap(err, "invalid request").Error(), nil)
		return
	}

	resp.HandleSuccess(ctx, http.StatusOK, "Success", nil)
}

func (h *PersonHandler) DeletePerson(ctx *gin.Context) {
	resp.HandleSuccess(ctx, http.StatusOK, "Success", nil)
}
