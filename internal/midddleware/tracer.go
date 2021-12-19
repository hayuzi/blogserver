package midddleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/hayuzi/blogserver/global"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
)

func Tracing() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx context.Context
		var traceId string
		var spanId string
		//Header中字段  Uber-Trace-Id: 5a4a0384f7facfaa:5a4a0384f7facfaa:0000000000000000:1
		parentSpanContext, _ := opentracing.GlobalTracer().Extract(opentracing.TextMap, opentracing.HTTPHeadersCarrier(c.Request.Header))
		span, ctx := opentracing.StartSpanFromContextWithTracer(c.Request.Context(), global.Tracer, c.Request.URL.Path, opentracing.ChildOf(parentSpanContext))
		//span := opentracing.SpanFromContext(c.Request.Context())
		//if parentSpanContext != nil {
		//	span, ctx = opentracing.StartSpanFromContextWithTracer(c.Request.Context(), global.Tracer, c.Request.URL.Path, opentracing.ChildOf(parentSpanContext))
		//} else {
		//	span, ctx = opentracing.StartSpanFromContextWithTracer(c.Request.Context(), global.Tracer, c.Request.URL.Path)
		//}
		defer span.Finish()

		spanContext := span.Context()
		switch spanContext.(type) {
		case jaeger.SpanContext:
			traceId = spanContext.(jaeger.SpanContext).TraceID().String()
			spanId = spanContext.(jaeger.SpanContext).SpanID().String()
		}
		c.Set("X-Trace-ID", traceId)
		c.Set("X-Span-ID", spanId)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
