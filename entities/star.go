package entities

import "gorm.io/gorm"

// Star 博文点赞实体类
type Star struct {
	gorm.Model
	UserId    uint // 点赞者Id
	ArticleId uint // 点赞博文Id
}

func (Star) TableName() string {
	return "a_stars"
}
