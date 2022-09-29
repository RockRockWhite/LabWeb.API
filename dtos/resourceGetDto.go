package dtos

import (
	"github.com/RockRockWhite/LabWeb.API/entities"
	"github.com/RockRockWhite/LabWeb.API/utils"
	"github.com/petersunbag/coven"
	"time"
)

type ResourceGetDto struct {
	Id             uint   // 资源Id
	Title          string // 资源标题
	Link           string // 资源链接
	Description    string // 资源描述
	LastModifiedId uint   // 最后修改者Id
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func GetResourceGetDtoConverter() *coven.Converter {
	option := &coven.StructOption{
		AliasFields: map[string]string{"Id": "ID"},
	}

	converter, err := coven.NewConverterOption(ResourceGetDto{}, entities.Resource{}, option)
	if err != nil {
		utils.GetLogger().Fatalln("failed to new converter of ResourceGetDto.")
	}

	return converter
}
