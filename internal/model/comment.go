package model

type Comment struct {
	Model
	ArticleId     int    `json:"articleId"`
	UserId        int    `json:"userId"`
	MentionUserId int    `json:"mentionUserId"`
	Content       string `json:"content"`
	CommentStatus int    `json:"commentStatus"`
	User          User   `json:"user" gorm:"foreignKey:UserId"`
	MentionUser   User   `json:"mentionUser" gorm:"foreignKey:MentionUserId"`
}

// TableName 会将 的表名重写
func (t Comment) TableName() string {
	return "blog_comment"
}

const CommentStatusNormal = 1 // 展示
const CommentStatusHidden = 2 // 隐藏
