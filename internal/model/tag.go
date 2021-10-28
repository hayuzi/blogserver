package model

type Tag struct {
	Model
	TagName   string `json:"tag_name"`
	Weight    string `json:"weight"`
	TagStatus int    `json:"tag_status"`
}

// TableName 会将 的表名重写
func (t Tag) TableName() string {
	return "blog_tag"
}
