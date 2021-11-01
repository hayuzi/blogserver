package errcode

var (
	Success         = NewError(0, "成功")
	ServerError     = NewError(10000000, "服务内部错误")
	InvalidParams   = NewError(10000001, "参数错误")
	RecordNotFound  = NewError(10000002, "记录不存在")
	TooManyRequests = NewError(10000003, "请求过多")

	UnauthorizedTokenError    = NewError(210001, "鉴权失败，Token错误")
	UnauthorizedTokenTimeout  = NewError(210002, "鉴权失败，Token超时间")
	UnauthorizedTokenGenerate = NewError(210003, "鉴权失败，Token生成失败")
	UnauthorizedAuthError     = NewError(210004, "鉴权失败，用户名或密码错误")
	UnauthorizedAuthNotExists = NewError(210005, "鉴权失败，找不到对应用户")
	UnauthorizedUserNotAdmin  = NewError(210009, "鉴权失败，非管理员账号")
)
