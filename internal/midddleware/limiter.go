package midddleware

import (
	"github.com/gin-gonic/gin"
	"github.com/hayuzi/blogserver/pkg/app"
	"github.com/hayuzi/blogserver/pkg/errcode"
	"github.com/hayuzi/blogserver/pkg/ratelimiter"
)

func RateLimiter(l ratelimiter.LimiterIface) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := l.Key(c)
		if bucket, ok := l.GetBucket(key); ok {
			count := bucket.TakeAvailable(1)
			if count == 0 {
				response := app.NewResponse(c)
				response.ToResponseError(errcode.TooManyRequests)
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
