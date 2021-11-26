package app

import (
	"context"
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/hayuzi/blogserver/global"
)

type ValidError struct {
	Key     string
	Message string
}

type ValidErrors []*ValidError

func (v *ValidError) Error() string {
	return v.Message
}

func (v ValidErrors) Errors() []string {
	errs := make([]string, 0)
	for _, err := range v {
		errs = append(errs, err.Error())
	}
	return errs
}

func BindAndValid(c *gin.Context, v interface{}) (bool, ValidErrors) {
	var errs ValidErrors = make([]*ValidError, 0)
	err := c.ShouldBind(v)
	if err != nil {
		t := c.Value("trans")
		trans, _ := t.(ut.Translator)
		verrs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 使用其他方法
			global.Logger.Error(context.Background(), err.Error())
			errs = append(errs, &ValidError{
				Key:     "service",
				Message: "参数数据类型不符合要求",
			})
			return true, errs
		}
		for key, value := range verrs.Translate(trans) {
			errs = append(errs, &ValidError{
				Key:     key,
				Message: value,
			})
		}
		return true, errs
	}
	return false, nil
}
