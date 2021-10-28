package errcode

var (
	Success                   = NewError(0, "成功")
	ServerError               = NewError(10000000, "服务内部错误")
	InvalidParams             = NewError(10000001, "参数错误")
	RecordNotFound            = NewError(10000002, "记录不存在")
	TooManyRequests           = NewError(10000003, "请求过多")
	UnauthorizedAuthNotExists = NewError(10000100, "鉴权失败，找不到对应的AppKey和AppSecret")
	UnauthorizedTokenError    = NewError(10000101, "鉴权失败，Token错误")
	UnauthorizedTokenTimeout  = NewError(10000102, "鉴权失败，Token超时间")
	UnauthorizedTokenGenerate = NewError(10000103, "鉴权失败，Token生成失败")
)
