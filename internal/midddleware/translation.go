package midddleware

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/hayuzi/blogserver/global"
)

func Translations() gin.HandlerFunc {
	return func(c *gin.Context) {
		// RegisterDefaultTranslations 注册方法访问全局的 gin/binding.StructValidator 并循环注册方法到 每个trans对应的map上
		// 在 gin/binding.StructValidator.validate[trans] 这个map上做了写入
		// 在实时请求的goroutine中注册翻译方法，可能发生并发map读写导致程序panic（ 实际发生 ）
		// 所以迁移到应用启动的时候生成，并在middle中获取对应的trans并注入到 context中

		var trans ut.Translator
		locale := c.GetHeader("locale")
		//if locale == "" {
		//	// locale 通常取决于 http 请求头的 'Accept-Language'()
		//	acceptLanguage := c.GetHeader("Accept-Language")
		//	if strings.Contains(acceptLanguage, "zh") {
		//		locale = "zh"
		//	}
		//}

		switch locale {
		case "zh":
			trans = global.Trans.TransZh
		case "en":
			trans = global.Trans.TransEn
		default:
			trans = global.Trans.TransZh
		}
		c.Set("trans", trans)
		c.Next()
	}
}
