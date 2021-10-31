package model

type User struct {
	Model
	Username string `json:"username"`
	Email    string `json:"email"`
	Pwd      string `json:"pwd"`
	UserType int    `json:"userType"`
}

// TableName 会将 的表名重写
func (t User) TableName() string {
	return "blog_user"
}

const UserTypeAdmin = 1 // 管理员
const UserTypeUser = 2  // 前台用户
