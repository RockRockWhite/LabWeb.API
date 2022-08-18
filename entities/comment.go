package entities

import "gorm.io/gorm"

// Comment 评论实体类
type Comment struct {
	gorm.Model
	UserId    uint   // 发布者Id
	Content   string // 评论内容
	ArticleId uint   // 博文Id
	ParentId  uint   // 父评论Id
}

func (Comment) TableName() string {
	return "a_comments"
}
