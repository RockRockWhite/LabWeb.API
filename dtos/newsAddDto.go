package dtos

import (
	"github.com/RockRockWhite/LabWeb.API/entities"
	"github.com/RockRockWhite/LabWeb.API/utils"
	"github.com/petersunbag/coven"
)

// NewsAddDto 添加新闻Dto
type NewsAddDto struct {
	Title          string // 新闻标题
	Content        string // 新闻内容
	LastModifiedId uint   // 最后修改者Id
}

func GetNewsAddDtoConverter() *coven.Converter {
	converter, err := coven.NewConverter(entities.News{}, NewsAddDto{})
	if err != nil {
		utils.GetLogger().Fatalln("failed to new converter of NewsAddDto.")
	}

	return converter
}
