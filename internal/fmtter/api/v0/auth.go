package v0

// AuthRegisterReq 注册
type AuthRegisterReq struct {
	Username string `json:"username" form:"username" binding:"required,max=64"`
	Email    string `json:"email" form:"email" binding:"required,max=64"`
	Pwd      string `json:"pwd" form:"pwd" binding:"required,max=64"`
}
type AuthRegisterRes struct {
	Token    string `json:"token"`
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	UserType int    `json:"userType"`
}

// AuthLoginReq 登陆
type AuthLoginReq struct {
	Username string `json:"username" form:"username" binding:"required,max=64"`
	Pwd      string `json:"pwd" form:"pwd" binding:"required,max=64"`
}
type AuthLoginRes struct {
	Token    string `json:"token"`
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	UserType int    `json:"userType"`
}
