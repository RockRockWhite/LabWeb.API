package dtos

import (
	"github.com/RockRockWhite/LabWeb.API/entities"
)

// ArticleAddDto 添加博文Dto
type ArticleAddDto struct {
	Titile  string // 博文标题
	Content string // 博文内容

	Tags []TagAddDto // 博文标签
}

// ToEntity 转换成Entity
func (dto *ArticleAddDto) ToEntity(userId uint) *entities.Article {
	entity := entities.Article{
		UserId:   userId,
		Title:    dto.Titile,
		Content:  dto.Content,
		Views:    0,
		Tags:     nil,
		Comments: nil,
		Stars:    nil,
	}

	// 转换Tag
	entity.Tags = make([]entities.Tag, 0, len(dto.Tags))
	for _, tagDto := range dto.Tags {
		entity.Tags = append(entity.Tags, *tagDto.ToEntity(0))
	}

	return &entity
}
