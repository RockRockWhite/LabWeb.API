package dtos

import (
	"github.com/RockRockWhite/LabWeb.API/entities"
	"time"
)

// PaperAddDto 添加论文Dto
type PaperAddDto struct {
	Title       string              // 论文标题
	Abstract    string              // 论文简介 富文本
	Content     string              // 论文内容 富文本
	Authors     string              // 论文作者
	State       entities.PaperState // 论文状态 枚举
	PublishedAt time.Time           // 发布时间
	PublishedIn string              // 发表单位
}

// ToEntity 转换成Entity
func (dto *PaperAddDto) ToEntity(lastModifiedId uint) *entities.Paper {
	entity := entities.Paper{
		Title:          dto.Title,
		Abstract:       dto.Abstract,
		Content:        dto.Content,
		Authors:        dto.Authors,
		State:          dto.State,
		PublishedAt:    dto.PublishedAt,
		PublishedIn:    dto.PublishedIn,
		LastModifiedId: lastModifiedId,
	}

	return &entity
}
