package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"golang.org/x/time/rate"

	"basic/pkg/helper/resp"
)

type ipLimiter struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

var (
	limiters sync.Map
)

func getLimiter(ip string, rps float64, burst int) *rate.Limiter {
	val, ok := limiters.Load(ip)
	if !ok {
		l := &ipLimiter{
			limiter:  rate.NewLimiter(rate.Limit(rps), burst),
			lastSeen: time.Now(),
		}
		limiters.Store(ip, l)
		return l.limiter
	}
	entry := val.(*ipLimiter)
	entry.lastSeen = time.Now()
	return entry.limiter
}

// cleanupLimiters removes entries that haven't been seen in the last 5 minutes.
// Called once at startup as a background goroutine.
func cleanupLimiters() {
	for {
		time.Sleep(5 * time.Minute)
		limiters.Range(func(key, value any) bool {
			entry := value.(*ipLimiter)
			if time.Since(entry.lastSeen) > 5*time.Minute {
				limiters.Delete(key)
			}
			return true
		})
	}
}

func init() {
	go cleanupLimiters()
}

// RateLimitMiddleware limits requests per client IP.
// Config keys (with defaults):
//
//	http.rate_limit.rps   — requests per second (default 10)
//	http.rate_limit.burst — burst size          (default 30)
func RateLimitMiddleware(conf *viper.Viper) gin.HandlerFunc {
	rps := conf.GetFloat64("http.rate_limit.rps")
	if rps <= 0 {
		rps = 10
	}
	burst := conf.GetInt("http.rate_limit.burst")
	if burst <= 0 {
		burst = 30
	}

	return func(ctx *gin.Context) {
		ip := ctx.ClientIP()
		if !getLimiter(ip, rps, burst).Allow() {
			resp.HandleError(ctx, http.StatusTooManyRequests, "TooManyRequests", nil)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
