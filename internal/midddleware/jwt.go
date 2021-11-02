package midddleware

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/hayuzi/blogserver/internal/model"
	"github.com/hayuzi/blogserver/pkg/app"
	"github.com/hayuzi/blogserver/pkg/consts"
	"github.com/hayuzi/blogserver/pkg/errcode"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		ecode := errcode.Success
		token := ""
		if s, ok := c.GetQuery("token"); ok {
			token = s
		} else {
			token = c.GetHeader("token")
		}

		if token == "" {
			ecode = errcode.InvalidParams
		} else {
			_, err := app.ParseToken(token)
			if err != nil {
				ecode = errcode.UnauthorizedTokenError
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					ecode = errcode.UnauthorizedTokenTimeout
				default:
					ecode = errcode.UnauthorizedTokenError
				}
			}
		}

		if ecode != errcode.Success {
			response := app.NewResponse(c)
			response.ToResponseError(ecode)
			c.Abort()
			return
		}
		c.Next()
	}
}

func JWTAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		ecode := errcode.Success
		token := ""
		if s, ok := c.GetQuery("token"); ok {
			token = s
		} else {
			token = c.GetHeader("token")
		}

		if token == "" {
			ecode = errcode.InvalidParams
		} else {
			claims, err := app.ParseToken(token)
			if err != nil {
				ecode = errcode.UnauthorizedTokenError
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					ecode = errcode.UnauthorizedTokenTimeout
				default:
					ecode = errcode.UnauthorizedTokenError
				}
			} else if claims.UserType != model.UserTypeAdmin {
				ecode = errcode.UnauthorizedUserNotAdmin
			}
		}

		if ecode != errcode.Success {
			response := app.NewResponse(c)
			response.ToResponseError(ecode)
			c.Abort()
			return
		}
		c.Next()
	}
}

func JWTInjectClaims() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, _ := app.GetJwtClaims(c)
		ctx := context.WithValue(c.Request.Context(), consts.ContextLoginUserKey, claims)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
