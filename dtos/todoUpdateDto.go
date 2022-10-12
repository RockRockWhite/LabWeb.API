package dtos

import (
	"github.com/RockRockWhite/LabWeb.API/entities"
	"github.com/RockRockWhite/LabWeb.API/utils"
	"github.com/petersunbag/coven"
	"time"
)

// TodoUpdateDto 修改todo Dto
type TodoUpdateDto struct {
	Title       string    `gorm:"not null"` // Todo表标题
	Description string    `gorm:"not null"` // Todo描述
	From        time.Time `gorm:"not null"` // 开始时间
	To          time.Time `gorm:"not null"` // 完成时间
	Finished    bool      // 完成状态
}

func GetTodoUpdateDtoEntityConverter() *coven.Converter {
	option := &coven.StructOption{
		AliasFields: map[string]string{"Id": "ID"},
	}

	converter, err := coven.NewConverterOption(TodoUpdateDto{}, entities.Todo{}, option)
	if err != nil {
		utils.GetLogger().Fatalln("failed to new converter of TodoUpdateDto.")
	}

	return converter
}

func GetTodoEntityUpdateDtoConverter() *coven.Converter {
	option := &coven.StructOption{
		AliasFields: map[string]string{"ID": "Id"},
	}

	converter, err := coven.NewConverterOption(entities.Todo{}, TodoUpdateDto{}, option)
	if err != nil {
		utils.GetLogger().Fatalln("failed to new converter of TodoUpdateDto.")
	}

	return converter
}
