package gormplugins

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"gorm.io/gorm"
	"strings"
)

const (
	parentSpanGormKey = "opentracing:parent.span"
	spanGormKey       = "opentracing:span"
)

// WithContext SetSpanToGorm sets span to gorm settings, returns cloned DB
func WithContext(ctx context.Context, db *gorm.DB) *gorm.DB {
	if ctx == nil {
		return db
	}
	parentSpan := opentracing.SpanFromContext(ctx)
	if parentSpan == nil {
		return db
	}
	return db.Set(parentSpanGormKey, parentSpan)
}

// AddGormCallbacks adds callbacks for tracing, you should call SetSpanToGorm to make them work
func AddGormCallbacks(db *gorm.DB) {
	callbacks := newCallbacks()
	registerCallbacks(db, "create", callbacks)
	registerCallbacks(db, "query", callbacks)
	registerCallbacks(db, "update", callbacks)
	registerCallbacks(db, "delete", callbacks)
	registerCallbacks(db, "row_query", callbacks)
}

type callbacks struct{}

func newCallbacks() *callbacks {
	return &callbacks{}
}

func (c *callbacks) beforeCreate(db *gorm.DB)   { c.before(db) }
func (c *callbacks) afterCreate(db *gorm.DB)    { c.after(db, "INSERT") }
func (c *callbacks) beforeQuery(db *gorm.DB)    { c.before(db) }
func (c *callbacks) afterQuery(db *gorm.DB)     { c.after(db, "SELECT") }
func (c *callbacks) beforeUpdate(db *gorm.DB)   { c.before(db) }
func (c *callbacks) afterUpdate(db *gorm.DB)    { c.after(db, "UPDATE") }
func (c *callbacks) beforeDelete(db *gorm.DB)   { c.before(db) }
func (c *callbacks) afterDelete(db *gorm.DB)    { c.after(db, "DELETE") }
func (c *callbacks) beforeRowQuery(db *gorm.DB) { c.before(db) }
func (c *callbacks) afterRowQuery(db *gorm.DB)  { c.after(db, "") }

func (c *callbacks) before(db *gorm.DB) {
	// TODO@yuzi
	span, _ := opentracing.StartSpanFromContext(db.Statement.Context, parentSpanGormKey)
	// span = opentracing.StartSpan("sql", opentracing.ChildOf(span.Context()))
	// 利用db实例去传递span
	db.InstanceSet(spanGormKey, span)
	//db.InstanceSet(startTime, time.Now())

	//parentSpan := val.(opentracing.Span)
	//tr := span.Tracer()

	//ext.DBType.Set(sp, "sql")
	//db.Set(spanGormKey, sp)
}

func (c *callbacks) after(db *gorm.DB, operation string) {
	_span, ok := db.InstanceGet(spanGormKey)
	if !ok {
		return
	}
	span := _span.(opentracing.Span)
	defer span.Finish()
	if operation == "" {
		operation = strings.ToUpper(strings.Split(db.Dialector.Explain(db.Statement.SQL.String(), db.Statement.Vars...), " ")[0])
	}
	span.SetTag("db.instance", db.Statement.Table)
	span.SetTag("db.table", db.Statement.Table)
	span.SetTag("db.method", operation)
	span.SetTag("db.err", db.Statement.Error)
	span.SetTag("db.count", db.Statement.RowsAffected)

}

func registerCallbacks(db *gorm.DB, name string, c *callbacks) {
	beforeName := fmt.Sprintf("tracing:%v_before", name)
	afterName := fmt.Sprintf("tracing:%v_after", name)
	gormCallbackName := fmt.Sprintf("gorm:%v", name)
	// gorm does some magic, if you pass CallbackProcessor here - nothing works
	switch name {
	case "create":
		_ = db.Callback().Create().Before(gormCallbackName).Register(beforeName, c.beforeCreate)
		_ = db.Callback().Create().After(gormCallbackName).Register(afterName, c.afterCreate)
	case "query":
		_ = db.Callback().Query().Before(gormCallbackName).Register(beforeName, c.beforeQuery)
		_ = db.Callback().Query().After(gormCallbackName).Register(afterName, c.afterQuery)
	case "update":
		_ = db.Callback().Update().Before(gormCallbackName).Register(beforeName, c.beforeUpdate)
		_ = db.Callback().Update().After(gormCallbackName).Register(afterName, c.afterUpdate)
	case "delete":
		_ = db.Callback().Delete().Before(gormCallbackName).Register(beforeName, c.beforeDelete)
		_ = db.Callback().Delete().After(gormCallbackName).Register(afterName, c.afterDelete)
	case "row":
		_ = db.Callback().Row().Before(gormCallbackName).Register(beforeName, c.beforeRowQuery)
		_ = db.Callback().Row().After(gormCallbackName).Register(afterName, c.afterRowQuery)
	}
}
