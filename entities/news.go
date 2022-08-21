package entities

import (
	"gorm.io/gorm"
)

// News 新闻实体类
type News struct {
	gorm.Model
	Title          string `gorm:"not null"`     // 新闻标题
	Content        string `gorm:"not null"`     // 新闻内容, 富文本
	LastModifiedId uint   `gorm:"default:null"` // 最后修改者Id
}

func (News) TableName() string {
	return "n_news"
}
