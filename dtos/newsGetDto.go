package dtos

import (
	"github.com/RockRockWhite/LabWeb.API/entities"
)

type NewsGetDto struct {
	Id             uint   // 新闻Id
	Title          string // 新闻标题
	Content        string // 新闻内容
	LastModifiedId uint   // 最后修改者Id
}

func ParseNewsEntity(n *entities.News) *NewsGetDto {
	dto := NewsGetDto{
		Id:             n.ID,
		Title:          n.Title,
		Content:        n.Content,
		LastModifiedId: n.LastModifiedId,
	}

	return &dto
}
