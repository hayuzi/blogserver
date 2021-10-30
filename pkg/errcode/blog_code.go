package errcode

var (
	TagListFail   = NewError(20001000, "标签列表获取失败")
	TagDetailFail = NewError(20001001, "标签详情获取失败")
	TagCreateFail = NewError(20001002, "标签创建失败")
	TagUpdateFail = NewError(20001003, "标签更新失败")
	TagDeleteFail = NewError(20001004, "标签删除失败")
	TagCountFail  = NewError(20001005, "标签统计失败")
)
