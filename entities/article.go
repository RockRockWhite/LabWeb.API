package entities

import "gorm.io/gorm"

// Article 博文实体类
type Article struct {
	gorm.Model
	UserId  uint   // 发布者Id
	Title   string // 博文标题
	Content string // 博文内容
	Views   uint   // 博文浏览量

	Tags     []Tag     `gorm:"-"` // 博文标签
	Comments []Comment `gorm:"-"` // 博文评论
	Stars    []Star    `gorm:"-"` // 博文点赞
}

func (Article) TableName() string {
	return "a_articles"
}
