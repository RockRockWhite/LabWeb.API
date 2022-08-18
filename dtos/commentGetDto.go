package dtos

import (
	"github.com/RockRockWhite/LabWeb.API/entities"
	"time"
)

// CommentGetDto 博文评论GetDto
type CommentGetDto struct {
	Id        uint      // 博文评论Id
	CreatedAt time.Time // 博文评论创建时间
	UpdatedAt time.Time // 博文评论修改时间

	UserId    uint   // 发布者Id
	Content   string // 评论内容
	ArticleId uint   // 博文Id
	ParentId  uint   // 父评论Id
}

func ParseCommentEntity(comment *entities.Comment) *CommentGetDto {
	return &CommentGetDto{
		Id:        comment.ID,
		CreatedAt: comment.CreatedAt,
		UpdatedAt: comment.UpdatedAt,
		UserId:    comment.UserId,
		Content:   comment.Content,
		ArticleId: comment.ArticleId,
		ParentId:  comment.ParentId,
	}
}
