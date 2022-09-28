package dtos

import (
	"github.com/RockRockWhite/LabWeb.API/entities"
	"time"
)

type PaperGetDto struct {
	Id             uint                // 论文Id
	Title          string              // 论文标题
	Abstract       string              // 论文简介 富文本
	Thumbnail      string              // 论文缩略图
	Link           string              // 论文链接
	Authors        string              // 论文作者
	State          entities.PaperState // 论文状态 枚举
	PublishedAt    time.Time           // 发布时间
	PublishedIn    string              // 发表单位
	LastModifiedId uint                // 最后修改者Id
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func ParsePaperEntity(p *entities.Paper) *PaperGetDto {
	dto := PaperGetDto{
		Id:          p.ID,
		Title:       p.Title,
		Abstract:    p.Abstract,
		Thumbnail:   p.Thumbnail,
		Link:        p.Link,
		Authors:     p.Authors,
		State:       p.State,
		PublishedAt: p.PublishedAt,
		PublishedIn: p.PublishedIn,
	}

	return &dto
}
