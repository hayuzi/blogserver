package errcode

var (
	Success                  = NewError(200, "成功")
	ServerError              = NewError(10000000, "服务内部错误")
	InvalidParams            = NewError(10000001, "参数错误")
	RecordNotFound           = NewError(10000002, "记录不存在")
	TooManyRequests          = NewError(10000003, "请求过多")
	UnauthorizedTokenError   = NewError(10000010, "鉴权失败，Token错误")
	UnauthorizedTokenTimeout = NewError(10000011, "鉴权失败，Token超时间")
	UnauthorizedUserNotAdmin = NewError(10000012, "鉴权失败，非管理员账号")
	AuthTokenGenerateFail    = NewError(10000013, "登陆授权失败，Token生成失败")
	AuthPwdNotCorrect        = NewError(10000014, "登陆授权失败，用户名或密码错误")
	AuthUserNotExists        = NewError(10000015, "登陆授权失败，找不到对应用户")
	AuthRegisterFail         = NewError(10000016, "注册失败")
	AuthUsernameExists       = NewError(10000017, "注册失败，用户名已被使用")
)
