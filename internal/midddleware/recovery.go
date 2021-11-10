package midddleware

import (
	"github.com/gin-gonic/gin"
	"github.com/hayuzi/blogserver/global"
	"github.com/hayuzi/blogserver/pkg/app"
	"github.com/hayuzi/blogserver/pkg/errcode"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				global.Logger.WithCallersFrames().Panicf(c, "panic cover err: %v", err)
				app.NewResponse(c).ToResponseError(errcode.ServerError)
				c.Abort()
			}
		}()
		c.Next()
	}
}
