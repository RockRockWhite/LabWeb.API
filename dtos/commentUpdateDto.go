package dtos

import "github.com/RockRockWhite/LabWeb.API/entities"

// CommentUpdateDto 更新博文评论Dto
type CommentUpdateDto struct {
	Content string // 评论内容
}

// CommentUpdateDtoFromEntity 从entity转换UpdateDto
func CommentUpdateDtoFromEntity(comment *entities.Comment) *CommentUpdateDto {
	return &CommentUpdateDto{Content: comment.Content}
}

// ApplyUpdateToEntity 将Update应用到Entity
func (dto *CommentUpdateDto) ApplyUpdateToEntity(entity *entities.Comment) {
	entity.Content = dto.Content
}
