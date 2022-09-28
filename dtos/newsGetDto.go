package dtos

import (
	"github.com/RockRockWhite/LabWeb.API/entities"
	"github.com/RockRockWhite/LabWeb.API/utils"
	"github.com/petersunbag/coven"
)

type NewsGetDto struct {
	Id             uint   // 新闻Id
	Title          string // 新闻标题
	Content        string // 新闻内容
	LastModifiedId uint   // 最后修改者Id
}

func GetNewsGetDtoConverter() *coven.Converter {

	option := &coven.StructOption{
		AliasFields: map[string]string{"Id": "ID"},
	}

	converter, err := coven.NewConverterOption(NewsGetDto{}, entities.News{}, option)
	if err != nil {
		utils.GetLogger().Fatalln("failed to new converter of NewsGetDto.")
	}

	return converter
}
