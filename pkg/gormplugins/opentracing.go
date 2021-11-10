package gormplugins

import (
	"github.com/opentracing/opentracing-go"
	"gorm.io/gorm"
)

const (
	callBackBeforeName = "opentracing:before"
	callBackAfterName  = "opentracing:after"

	gormParentSpanKey = "gorm"
	gormSpanKey       = "gorm:span"
)

func before(db *gorm.DB) {
	// 先从父级spans生成子span ---> 这里命名为gorm，但实际上可以自定义
	// 自己喜欢的operationName
	span, _ := opentracing.StartSpanFromContext(db.Statement.Context, gormParentSpanKey)

	// 利用db实例去传递span
	db.InstanceSet(gormSpanKey, span)
	return
}

func after(db *gorm.DB) {
	// 从GORM的DB实例中取出span
	_span, isExist := db.InstanceGet(gormSpanKey)
	if !isExist {
		// 不存在就直接抛弃掉
		return
	}

	// 断言进行类型转换
	span, ok := _span.(opentracing.Span)
	if !ok {
		return
	}
	// <---- 一定一定一定要Finsih掉！！！
	defer span.Finish()

	span.SetTag("db.err", db.Error)
	span.SetTag("db.instance", db.Statement.Table)
	span.SetTag("db.table", db.Statement.Table)
	span.SetTag("db.sql", db.Statement.SQL.String())
	span.SetTag("db.count", db.Statement.RowsAffected)
	return
}

const ()

type OpentracingPlugin struct{}

func (op *OpentracingPlugin) Name() string {
	return "opentracingPlugin"
}

func (op *OpentracingPlugin) Initialize(db *gorm.DB) (err error) {
	// 开始前 - 并不是都用相同的方法，可以自己自定义
	_ = db.Callback().Create().Before("gorm:before_create").Register(callBackBeforeName, before)
	_ = db.Callback().Query().Before("gorm:query").Register(callBackBeforeName, before)
	_ = db.Callback().Delete().Before("gorm:before_delete").Register(callBackBeforeName, before)
	_ = db.Callback().Update().Before("gorm:setup_reflect_value").Register(callBackBeforeName, before)
	_ = db.Callback().Row().Before("gorm:row").Register(callBackBeforeName, before)
	_ = db.Callback().Raw().Before("gorm:raw").Register(callBackBeforeName, before)

	// 结束后 - 并不是都用相同的方法，可以自己自定义
	_ = db.Callback().Create().After("gorm:after_create").Register(callBackAfterName, after)
	_ = db.Callback().Query().After("gorm:after_query").Register(callBackAfterName, after)
	_ = db.Callback().Delete().After("gorm:after_delete").Register(callBackAfterName, after)
	_ = db.Callback().Update().After("gorm:after_update").Register(callBackAfterName, after)
	_ = db.Callback().Row().After("gorm:row").Register(callBackAfterName, after)
	_ = db.Callback().Raw().After("gorm:raw").Register(callBackAfterName, after)
	return
}

// 告诉编译器这个结构体实现了gorm.Plugin接口
var _ gorm.Plugin = &OpentracingPlugin{}
