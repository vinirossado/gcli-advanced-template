package resp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func HandleSuccess(ctx *gin.Context, code int, message string, data interface{}) {
	if data == nil {
		data = map[string]string{}
	}
	resp := response{Code: code, Message: message, Data: data}
	ctx.JSON(http.StatusOK, resp)
}

func HandleError(ctx *gin.Context, httpCode int, message string, data interface{}) {
	if data == nil {
		data = map[string]string{}
	}
	resp := response{Code: httpCode, Message: message, Data: data}
	ctx.JSON(httpCode, resp)
}
