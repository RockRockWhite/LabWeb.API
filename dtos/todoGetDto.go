package dtos

import (
	"github.com/RockRockWhite/LabWeb.API/entities"
	"github.com/RockRockWhite/LabWeb.API/utils"
	"github.com/petersunbag/coven"
	"time"
)

type TodoGetDto struct {
	Id          uint      // 新闻Id
	Title       string    `gorm:"not null"` // Todo表标题
	Description string    `gorm:"not null"` // Todo描述
	From        time.Time `gorm:"not null"` // 开始时间
	To          time.Time `gorm:"not null"` // 完成时间
	Finished    bool      // 完成状态
	UserId      uint      `gorm:"not null"` // 创建用户Id
	Username    string    // 昵称
	Fullname    string    // 实验室成员全名
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func GetTodoGetDtoConverter() *coven.Converter {
	option := &coven.StructOption{
		AliasFields: map[string]string{"Id": "ID"},
	}

	converter, err := coven.NewConverterOption(TodoGetDto{}, entities.Todo{}, option)
	if err != nil {
		utils.GetLogger().Fatalln("failed to new converter of TodoGetDto.")
	}

	return converter
}
