package app

import (
	"github.com/gin-gonic/gin"
	"github.com/hayuzi/blogserver/global"
	"github.com/hayuzi/blogserver/pkg/errcode"
	"net/http"
	"runtime"
)

type Response struct {
	Ctx *gin.Context
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{
		Ctx: ctx,
	}
}

func (r *Response) ToResponse(data interface{}) {
	if data == nil {
		data = make(map[string]interface{})
	}
	r.Ctx.JSON(http.StatusOK, gin.H{
		"code": errcode.Success.Code(),
		"msg":  "",
		"data": data,
	})
}

func (r *Response) ToResponseList(list interface{}, totalRows int64) {
	data := make(map[string]interface{})
	data["lists"] = list
	data["pageNum"] = GetPageNum(r.Ctx)
	data["pageSize"] = GetPageSize(r.Ctx)
	data["total"] = totalRows

	r.Ctx.JSON(http.StatusOK, gin.H{
		"code": errcode.Success.Code(),
		"msg":  "",
		"data": data,
	})
}

func (r *Response) ToResponseError(err *errcode.Error) {
	_, file, line, _ := runtime.Caller(1)
	global.Logger.Errorf("%s:%d %v", file, line, err.Error())
	response := gin.H{
		"code": err.Code(),
		"msg":  err.Msg(),
		"data": make(map[string]interface{}),
	}
	details := err.Details()
	if len(details) > 0 {
		response["details"] = details
		response["msg"] = err.Msg()
	}
	r.Ctx.JSON(err.StatusCode(), response)
}
