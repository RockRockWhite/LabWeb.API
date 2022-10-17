package dtos

import (
	"github.com/RockRockWhite/LabWeb.API/entities"
	"time"
)

// PaperAddDto 添加论文Dto
type PaperAddDto struct {
	Title       string              // 论文标题
	Abstract    string              // 论文简介 富文本
	Thumbnail   string              // 论文缩略图
	Link        string              // 论文链接
	Authors     string              // 论文作者
	Pdf         string              // 论文pdf
	Code        string              // 论文代码
	State       entities.PaperState // 论文状态 枚举
	PublishedAt time.Time           // 发布时间
	PublishedIn string              // 发表单位
}

// ToEntity 转换成Entity
func (dto *PaperAddDto) ToEntity(lastModifiedId uint) *entities.Paper {
	entity := entities.Paper{
		Title:          dto.Title,
		Abstract:       dto.Abstract,
		Thumbnail:      dto.Thumbnail,
		Link:           dto.Link,
		Pdf:            dto.Pdf,
		Code:           dto.Code,
		Authors:        dto.Authors,
		State:          dto.State,
		PublishedAt:    dto.PublishedAt,
		PublishedIn:    dto.PublishedIn,
		LastModifiedId: lastModifiedId,
	}

	return &entity
}
