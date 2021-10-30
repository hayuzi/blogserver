package app

import (
	"github.com/gin-gonic/gin"
	"github.com/hayuzi/blogserver/global"
	"github.com/hayuzi/blogserver/pkg/convert"
)

func GetPageNum(c *gin.Context) int {
	pageNum := convert.StrTo(c.Query("pageNum")).MustInt()
	if pageNum <= 0 {
		return 1
	}
	return pageNum
}

func GetPageSize(c *gin.Context) int {
	pageSize := convert.StrTo(c.Query("pageSize")).MustInt()
	if pageSize <= 0 {
		return global.AppSetting.DefaultPageSize
	}
	if pageSize > global.AppSetting.MaxPageSize {
		return global.AppSetting.MaxPageSize
	}
	return pageSize
}

func GetPageOffset(page, pageSize int) int {
	result := 0
	if page > 0 {
		result = (page - 1) * pageSize
	}
	return result
}

func InitPagination(pageNum, pageSize int) (int, int) {
	if pageNum <= 0 {
		pageNum = 1
	}
	if pageSize <= 0 {
		pageSize = global.AppSetting.DefaultPageSize
	}
	if pageSize > global.AppSetting.MaxPageSize {
		pageSize = global.AppSetting.MaxPageSize
	}
	return pageNum, pageSize
}
