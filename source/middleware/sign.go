package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"basic/pkg/helper/resp"
	"basic/pkg/logger"
)

func SignMiddleware(logger *logger.Logger, conf *viper.Viper) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requiredHeaders := []string{"Timestamp", "Nonce", "Sign", "App-Version"}

		for _, header := range requiredHeaders {
			value, ok := ctx.Request.Header[header]
			if !ok || len(value) == 0 {
				resp.HandleError(ctx, http.StatusBadRequest, "BadRequest", nil)
				ctx.Abort()
				return
			}
		}

		data := map[string]string{
			"AppKey":     conf.GetString("security.api_sign.app_key"),
			"Timestamp":  ctx.Request.Header.Get("Timestamp"),
			"Nonce":      ctx.Request.Header.Get("Nonce"),
			"AppVersion": ctx.Request.Header.Get("App-Version"),
		}

		var keys []string
		for k := range data {
			keys = append(keys, k)
		}
		sort.Slice(keys, func(i, j int) bool { return strings.ToLower(keys[i]) < strings.ToLower(keys[j]) })

		var msg strings.Builder
		for _, k := range keys {
			msg.WriteString(k)
			msg.WriteString(data[k])
		}

		secret := conf.GetString("security.api_sign.app_security")
		mac := hmac.New(sha256.New, []byte(secret))
		mac.Write([]byte(msg.String()))
		expected := strings.ToUpper(hex.EncodeToString(mac.Sum(nil)))

		if ctx.Request.Header.Get("Sign") != expected {
			resp.HandleError(ctx, http.StatusBadRequest, "BadRequest", nil)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
