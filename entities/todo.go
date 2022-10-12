package entities

import (
	"gorm.io/gorm"
	"time"
)

// Todo Todo表实体类
type Todo struct {
	gorm.Model
	Title       string    `gorm:"not null"` // Todo表标题
	Description string    `gorm:"not null"` // Todo描述
	From        time.Time `gorm:"not null"` // 开始时间
	To          time.Time `gorm:"not null"` // 完成时间
	Finished    bool      // 完成状态
	UserId      uint      `gorm:"not null"` // 创建用户Id
}

func (Todo) TableName() string {
	return "t_todos"
}
