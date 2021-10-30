package v1

type TagListReq struct {
	Id        int    `json:"id" form:"id" binding:"min:0"`
	TagName   string `json:"tagName" form:"tagName" binding:"max=100"`
	TagStatus string `json:"tagStatus" form:"tagStatus" binding:"min=0"`
}
