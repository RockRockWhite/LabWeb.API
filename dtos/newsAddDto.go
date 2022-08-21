package dtos

import (
	"github.com/RockRockWhite/LabWeb.API/entities"
)

// NewsAddDto 添加新闻Dto
type NewsAddDto struct {
	Title          string // 新闻标题
	Content        string // 新闻内容
	LastModifiedId uint   // 最后修改者Id
}

// ToEntity 转换成Entity
func (dto *NewsAddDto) ToEntity(lastModifiedId uint) *entities.News {
	entity := entities.News{
		Title:          dto.Title,
		Content:        dto.Content,
		LastModifiedId: lastModifiedId,
	}

	return &entity
}
