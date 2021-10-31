package model

type Tag struct {
	Model
	TagName   string `json:"tagName"`
	Weight    int    `json:"weight"`
	TagStatus int    `json:"tagStatus"`
}

// TableName 会将 的表名重写
func (t Tag) TableName() string {
	return "blog_tag"
}

const TagStatusNormal = 1 // 正常
const TagStatusHidden = 2 // 禁用
