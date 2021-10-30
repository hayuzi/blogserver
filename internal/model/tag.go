package model

type Tag struct {
	Model
	TagName   string `json:"tagName"`
	Weight    string `json:"weight"`
	TagStatus int    `json:"tagStatus"`
}

// TableName 会将 的表名重写
func (t Tag) TableName() string {
	return "blog_tag"
}
