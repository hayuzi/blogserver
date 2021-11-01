package model

type Article struct {
	Model
	Title         string `json:"title"`
	Sketch        string `json:"sketch"`
	Content       string `json:"content"`
	TagId         int    `json:"tagId"`
	Weight        int    `json:"weight"`
	ArticleStatus int    `json:"articleStatus"`
	Tag           Tag    `json:"tag" gorm:"foreignKey:TagId"`
}

// TableName 会将 的表名重写
func (t Article) TableName() string {
	return "blog_article"
}

const ArticleStatusDraft = 1     // 草稿
const ArticleStatusPublished = 2 // 发布
