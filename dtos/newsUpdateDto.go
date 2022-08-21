package dtos

import (
	"github.com/RockRockWhite/LabWeb.API/entities"
)

// NewsUpdateDto 修改新闻Dto
type NewsUpdateDto struct {
	Title   string // 新闻标题
	Content string // 新闻内容
}

// NewsDtoFromEntity 从entity转换UpdateDto
func NewsDtoFromEntity(n *entities.News) *NewsUpdateDto {
	return &NewsUpdateDto{
		Title:   n.Title,
		Content: n.Content,
	}
}

// ApplyUpdateToEntity 将Update应用到Entity
func (dto *NewsUpdateDto) ApplyUpdateToEntity(entity *entities.News, lastModifiedId uint) {
	entity.Title = dto.Title
	entity.Content = dto.Content
	entity.LastModifiedId = lastModifiedId
}
