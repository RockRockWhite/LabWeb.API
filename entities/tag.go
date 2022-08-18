package entities

import "gorm.io/gorm"

// Tag 博文标签实体类
type Tag struct {
	gorm.Model
	Name      string // 标签名称
	ArticleId uint   // 博文Id
}

func (Tag) TableName() string {
	return "a_tags"
}
