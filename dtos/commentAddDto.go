package dtos

import (
	"github.com/RockRockWhite/LabWeb.API/entities"
)

// CommentAddDto 博创建文评论Dto
type CommentAddDto struct {
	Content  string // 评论内容
	ParentId uint   // 父评论Id
}

// ToEntity 转换成Entity
func (dto *CommentAddDto) ToEntity(articleId uint, userId uint) *entities.Comment {
	return &entities.Comment{
		UserId:    userId,
		Content:   dto.Content,
		ArticleId: articleId,
		ParentId:  dto.ParentId,
	}
}
