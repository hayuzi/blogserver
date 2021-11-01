package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/hayuzi/blogserver/pkg/app"
)

type Common struct{}

func NewCommon() Common {
	return Common{}
}

func (t Common) Dashboard(c *gin.Context) {
	response := app.NewResponse(c)
	response.ToResponse(nil)
	return
}
