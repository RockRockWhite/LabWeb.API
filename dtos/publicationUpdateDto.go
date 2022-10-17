package dtos

import (
	"github.com/RockRockWhite/LabWeb.API/entities"
	"time"
)

// PaperUpdateDto 修改论文Dto
type PaperUpdateDto struct {
	Title       string              // 论文标题
	Abstract    string              // 论文简介 富文本
	Thumbnail   string              // 论文缩略图
	Link        string              // 论文链接
	Pdf         string              // 论文pdf
	Code        string              // 论文代码
	Authors     string              // 论文作者
	State       entities.PaperState // 论文状态 枚举
	PublishedAt time.Time           // 发布时间
	PublishedIn string              // 发表单位
}

// PaperDtoFromEntity 从entity转换UpdateDto
func PaperDtoFromEntity(p *entities.Paper) *PaperUpdateDto {
	return &PaperUpdateDto{
		Title:       p.Title,
		Abstract:    p.Abstract,
		Thumbnail:   p.Thumbnail,
		Link:        p.Link,
		Pdf:         p.Pdf,
		Code:        p.Code,
		Authors:     p.Authors,
		State:       p.State,
		PublishedAt: p.PublishedAt,
		PublishedIn: p.PublishedIn,
	}
}

// ApplyUpdateToEntity 将Update应用到Entity
func (dto *PaperUpdateDto) ApplyUpdateToEntity(entity *entities.Paper, lastModifiedId uint) {
	entity.Title = dto.Title
	entity.Abstract = dto.Abstract
	entity.Thumbnail = dto.Thumbnail
	entity.Link = dto.Link
	entity.Authors = dto.Authors
	entity.State = dto.State
	entity.PublishedAt = dto.PublishedAt
	entity.PublishedIn = dto.PublishedIn
	entity.LastModifiedId = lastModifiedId
}
