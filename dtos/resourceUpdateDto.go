package dtos

import (
	"github.com/RockRockWhite/LabWeb.API/entities"
	"github.com/RockRockWhite/LabWeb.API/utils"
	"github.com/petersunbag/coven"
)

// ResourceUpdateDto 修改资源Dto
type ResourceUpdateDto struct {
	Link        string // 资源链接
	Description string // 资源描述
}

func GetResourceUpdateDtoEntityConverter() *coven.Converter {
	option := &coven.StructOption{
		AliasFields: map[string]string{"Id": "ID"},
	}

	converter, err := coven.NewConverterOption(ResourceUpdateDto{}, entities.Resource{}, option)
	if err != nil {
		utils.GetLogger().Fatalln("failed to new converter of ResourceUpdateDto.")
	}

	return converter
}

func GetResourceEntityUpdateDtoConverter() *coven.Converter {
	option := &coven.StructOption{
		AliasFields: map[string]string{"ID": "Id"},
	}

	converter, err := coven.NewConverterOption(entities.Resource{}, ResourceUpdateDto{}, option)
	if err != nil {
		utils.GetLogger().Fatalln("failed to new converter of ResourceUpdateDto.")
	}

	return converter
}
