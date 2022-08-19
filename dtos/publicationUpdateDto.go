package dtos

import (
	"github.com/RockRockWhite/LabWeb.API/entities"
	"time"
)

// PublicationUpdateDto 修改论文Dto
type PublicationUpdateDto struct {
	Title       string                    // 论文标题
	Abstract    string                    // 论文简介 富文本
	Content     string                    // 论文内容 富文本
	Authors     string                    // 论文作者
	State       entities.PublicationState // 论文状态 枚举
	PublishedAt time.Time                 // 发布时间
	PublishedIn string                    // 发表单位
}

// PublicationDtoFromEntity 从entity转换UpdateDto
func PublicationDtoFromEntity(p *entities.Publication) *PublicationUpdateDto {
	return &PublicationUpdateDto{
		Title:       p.Title,
		Abstract:    p.Abstract,
		Content:     p.Content,
		Authors:     p.Authors,
		State:       p.State,
		PublishedAt: p.PublishedAt,
		PublishedIn: p.PublishedIn,
	}
}

// ApplyUpdateToEntity 将Update应用到Entity
func (dto *PublicationUpdateDto) ApplyUpdateToEntity(entity *entities.Publication, lastModifiedId uint) {
	entity.Title = dto.Title
	entity.Abstract = dto.Abstract
	entity.Content = dto.Content
	entity.Authors = dto.Authors
	entity.State = dto.State
	entity.PublishedAt = dto.PublishedAt
	entity.PublishedIn = dto.PublishedIn
	entity.LastModifiedId = lastModifiedId
}
