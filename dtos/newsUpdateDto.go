package dtos

import (
	"github.com/RockRockWhite/LabWeb.API/entities"
	"github.com/RockRockWhite/LabWeb.API/utils"
	"github.com/petersunbag/coven"
)

// NewsUpdateDto 修改新闻Dto
type NewsUpdateDto struct {
	Title   string // 新闻标题
	Content string // 新闻内容
}

func GetNewsUpdateDtoEntityConverter() *coven.Converter {
	option := &coven.StructOption{
		AliasFields: map[string]string{"Id": "ID"},
	}

	converter, err := coven.NewConverterOption(NewsUpdateDto{}, entities.News{}, option)
	if err != nil {
		utils.GetLogger().Fatalln("failed to new converter of NewsUpdateDto.")
	}

	return converter
}

func GetNewsEntityUpdateDtoConverter() *coven.Converter {
	option := &coven.StructOption{
		AliasFields: map[string]string{"ID": "Id"},
	}

	converter, err := coven.NewConverterOption(entities.News{}, NewsUpdateDto{}, option)
	if err != nil {
		utils.GetLogger().Fatalln("failed to new converter of NewsUpdateDto.")
	}

	return converter
}
