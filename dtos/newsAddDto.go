package dtos

// NewsAddDto 添加新闻Dto
type NewsAddDto struct {
	Title          string // 新闻标题
	Content        string // 新闻内容
	LastModifiedId uint   // 最后修改者Id
}
