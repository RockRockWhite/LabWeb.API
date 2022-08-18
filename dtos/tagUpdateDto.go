package dtos

import (
	"github.com/RockRockWhite/LabWeb.API/entities"
)

// TagUpdateDto 博文标签Dto
type TagUpdateDto struct {
	Name string // 标签名称
}

// TagUpdateDtoFromEntity 从entity转换UpdateDto
func TagUpdateDtoFromEntity(tag *entities.Tag) *TagUpdateDto {
	return &TagUpdateDto{Name: tag.Name}
}

// ApplyUpdateToEntity 将Update应用到Entity
func (dto *TagUpdateDto) ApplyUpdateToEntity(entity *entities.Tag) {
	entity.Name = dto.Name
}
