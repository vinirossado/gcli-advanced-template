package middleware

import (
	"basic/pkg/helper/resp"
	"basic/pkg/logger"
	"net/http"
	"sort"
	"strings"

	"github.com/duke-git/lancet/v2/cryptor"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// SignMiddleware is a middleware that validates the request signature.
// It checks for the presence of required headers and verifies the signature
// using a combination of header values and a secret key.
//
// This middleware ensures that the request is coming from a trusted source
// by validating the signature, which is a hash of the request data and a secret key.
//
// Parameters:
// - logger: Logger instance for logging purposes.
// - conf: Configuration instance to retrieve security settings.
//
// Returns:
// - gin.HandlerFunc: The middleware handler function.
func SignMiddleware(conf *viper.Viper, logger *logger.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requiredHeaders := []string{"Timestamp", "Nonce", "Sign", "App-Version"}

		// Check for the presence of required headers
		for _, header := range requiredHeaders {
			if value := ctx.Request.Header.Get(header); value == "" {
				resp.HandleError(ctx, http.StatusBadRequest, "Missing required header: "+header, nil)
				ctx.Abort()
				return
			}
		}

		// Collect data for signature verification
		data := map[string]string{
			"AppKey":     conf.GetString("security.api_sign.app_key"),
			"Timestamp":  ctx.Request.Header.Get("Timestamp"),
			"Nonce":      ctx.Request.Header.Get("Nonce"),
			"AppVersion": ctx.Request.Header.Get("App-Version"),
		}

		// Sort keys to ensure consistent ordering
		keys := make([]string, 0, len(data))
		for k := range data {
			keys = append(keys, k)
		}
		sort.Slice(keys, func(i, j int) bool { return strings.ToLower(keys[i]) < strings.ToLower(keys[j]) })

		// Concatenate sorted key-value pairs
		var strBuilder strings.Builder
		for _, k := range keys {
			strBuilder.WriteString(k + data[k])
		}
		strBuilder.WriteString(conf.GetString("security.api_sign.app_security"))

		// Verify the signature
		expectedSign := strings.ToUpper(cryptor.Md5String(strBuilder.String()))
		if ctx.Request.Header.Get("Sign") != expectedSign {
			resp.HandleError(ctx, http.StatusBadRequest, "Invalid signature", nil)
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
