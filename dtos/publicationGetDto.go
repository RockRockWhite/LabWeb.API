package dtos

import (
	"github.com/RockRockWhite/LabWeb.API/entities"
	"time"
)

type PublicationGetDto struct {
	Title          string                    // 论文标题
	Abstract       string                    // 论文简介 富文本
	Content        string                    // 论文内容 富文本
	Authors        string                    // 论文作者
	State          entities.PublicationState // 论文状态 枚举
	PublishedAt    time.Time                 // 发布时间
	PublishedIn    string                    // 发表单位
	LastModifiedId uint                      // 最后修改者Id
}

func ParsePublicationEntity(p *entities.Publication) *PublicationGetDto {
	dto := PublicationGetDto{
		Title:       p.Title,
		Abstract:    p.Abstract,
		Content:     p.Content,
		Authors:     p.Authors,
		State:       p.State,
		PublishedAt: p.PublishedAt,
		PublishedIn: p.PublishedIn,
	}

	return &dto
}
