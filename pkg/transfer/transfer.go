package transfer

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en2 "github.com/go-playground/validator/v10/translations/en"
	zh2 "github.com/go-playground/validator/v10/translations/zh"
)

type Transfers struct {
	TransZh ut.Translator
	TransEn ut.Translator
}

func NewGinValidateTrans() (tsf Transfers, err error) {
	// 获取翻译器
	uni := ut.New(en.New(), zh.New())
	transZh, _ := uni.GetTranslator("zh")
	transEn, _ := uni.GetTranslator("en")
	// 绑定到gin的validator上 ( 注意: trans 我们最后还要用来作为获取 翻译信息的 键，所以要存起来，并在后面使用 )
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		err = errors.New("RegisterTransfer binding.Validator.Engine().(*validator.Validate) err")
		return
	}
	err = zh2.RegisterDefaultTranslations(v, transZh)
	if err != nil {
		err = errors.New(fmt.Sprintf("RegisterTransfer zh2.RegisterDefaultTranslations err:%v", err))
		return
	}
	err = en2.RegisterDefaultTranslations(v, transEn)
	if err != nil {
		err = errors.New(fmt.Sprintf("RegisterTransfer en2.RegisterDefaultTranslations err:%v", err))
		return
	}
	tsf.TransZh = transZh
	tsf.TransEn = transEn
	return
}
