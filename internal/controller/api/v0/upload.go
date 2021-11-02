package v0

import (
	"github.com/gin-gonic/gin"
	"github.com/hayuzi/blogserver/internal/service"
	"github.com/hayuzi/blogserver/pkg/app"
	"github.com/hayuzi/blogserver/pkg/convert"
	"github.com/hayuzi/blogserver/pkg/errcode"
	"github.com/hayuzi/blogserver/pkg/upload"
)

type Upload struct{}

func NewUpload() Upload {
	return Upload{}
}

func (t Upload) UploadFile(c *gin.Context) {
	response := app.NewResponse(c)
	// 两个参数，file数据和 type类型 1=图片 2=excel
	file, fileHeader, err := c.Request.FormFile("file")
	fileType := convert.StrTo(c.PostForm("type")).MustInt()
	if err != nil {
		response.ToResponseError(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}
	if fileHeader == nil || fileType <= 0 {
		response.ToResponseError(errcode.InvalidParams)
		return
	}
	svc := service.New(c.Request.Context())
	resFileInfo, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader)
	if err != nil {
		response.ToResponseError(errcode.UploadFileFail.WithDetails(err.Error()))
		return
	}
	response.ToResponse(resFileInfo)
	return
}
