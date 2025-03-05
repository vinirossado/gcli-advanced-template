package resp

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// HandleSuccess sends a successful response with the given code, message, and data.
func HandleSuccess(w http.ResponseWriter, code int, message string, data interface{}) {
	if data == nil {
		data = map[string]interface{}{}
	}
	resp := response{Code: code, Message: message, Data: data}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(resp)
}

// HandleError sends an error response with the given HTTP code, message, and data.
func HandleError(w http.ResponseWriter, httpCode int, message string, data interface{}) {
	if data == nil {
		data = map[string]string{}
	}
	resp := response{Code: httpCode, Message: message, Data: data}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	json.NewEncoder(w).Encode(resp)
}
