package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hayuzi/blogserver/pkg/app"
	"runtime"
)

type Common struct{}

func NewCommon() Common {
	return Common{}
}

func (t Common) Dashboard(c *gin.Context) {
	response := app.NewResponse(c)
	data := make(map[string]string)
	data["os"] = runtime.GOOS
	data["arch"] = runtime.GOARCH
	data["version"] = runtime.Version()
	data["cups"] = fmt.Sprintf("%d", runtime.NumCPU())
	response.ToResponse(data)
	return
}
