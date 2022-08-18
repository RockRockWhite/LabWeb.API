package dtos

import "github.com/RockRockWhite/LabWeb.API/entities"

// TagGetDto Get博文标签Dto
type TagGetDto struct {
	Id uint // 博文标签Id

	ArticleId uint   // 博文Id
	Name      string // 标签名称
}

func ParseTagEntity(tag *entities.Tag) *TagGetDto {
	return &TagGetDto{
		Id:        tag.ID,
		ArticleId: tag.ArticleId,
		Name:      tag.Name,
	}
}
