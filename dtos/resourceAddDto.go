package dtos

import (
	"github.com/RockRockWhite/LabWeb.API/entities"
	"github.com/RockRockWhite/LabWeb.API/utils"
	"github.com/petersunbag/coven"
)

// ResourceAddDto 添加资源Dto
type ResourceAddDto struct {
	Title       string // 资源标题
	Link        string // 资源链接
	Description string // 资源描述
}

func GetResourceAddDtoConverter() *coven.Converter {
	converter, err := coven.NewConverter(entities.Resource{}, ResourceAddDto{})
	if err != nil {
		utils.GetLogger().Fatalln("failed to new converter of ResourceAddDto.")
	}

	return converter
}
