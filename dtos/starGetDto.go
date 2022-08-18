package dtos

import "github.com/RockRockWhite/LabWeb.API/entities"

// StarGetDto 博文点赞GetDto
type StarGetDto struct {
	Id uint // 博文点赞Id

	UserId    uint // 点赞者Id
	ArticleId uint // 点赞博文Id
}

// ParseStarEntity 将entity转换为GetDto
func ParseStarEntity(star *entities.Star) *StarGetDto {
	return &StarGetDto{
		Id:        star.ID,
		UserId:    star.UserId,
		ArticleId: star.ArticleId,
	}
}
