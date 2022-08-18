package dtos

import (
	"github.com/RockRockWhite/LabWeb.API/entities"
)

// TagAddDto 创建博文标签Dto
type TagAddDto struct {
	Name string // 标签名称
}

// ToEntity 转换成Entity
func (dto *TagAddDto) ToEntity(articleId uint) *entities.Tag {
	return &entities.Tag{
		ArticleId: articleId,
		Name:      dto.Name,
	}
}
