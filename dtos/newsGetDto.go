package dtos

type NewsGetDto struct {
	Id             uint   // 新闻Id
	Title          string // 新闻标题
	Content        string // 新闻内容
	LastModifiedId uint   // 最后修改者Id
}
