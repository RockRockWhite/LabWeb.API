package dtos

import (
	"github.com/RockRockWhite/LabWeb.API/entities"
	"github.com/RockRockWhite/LabWeb.API/utils"
	"github.com/petersunbag/coven"
	"time"
)

// TodoAddDto 添加新闻Dto
type TodoAddDto struct {
	Title       string    `gorm:"not null"` // Todo表标题
	Description string    `gorm:"not null"` // Todo描述
	From        time.Time `gorm:"not null"` // 开始时间
	To          time.Time `gorm:"not null"` // 完成时间
}

func GetTodoAddDtoConverter() *coven.Converter {
	converter, err := coven.NewConverter(entities.Todo{}, TodoAddDto{})
	if err != nil {
		utils.GetLogger().Fatalln("failed to new converter of TodoAddDto.")
	}

	return converter
}
