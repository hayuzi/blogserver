package v1

// UserChangePwdReq 删除
type UserChangePwdReq struct {
	Pwd      string `json:"pwd" form:"pwd"`           // 旧密码
	Password string `json:"password" form:"password"` // 新密码
	Confirm  string `json:"confirm" form:"confirm"`   // 新密码确认
}
type UserChangePwdRes struct {
	Id int `json:"id"`
}
