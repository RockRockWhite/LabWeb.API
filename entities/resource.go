package entities

import (
	"gorm.io/gorm"
)

// Resource 资源实体类
type Resource struct {
	gorm.Model
	Link           string // 资源链接
	Description    string // 资源描述
	LastModifiedId uint   `gorm:"default:null"` // 最后修改者Id

}

func (Resource) TableName() string {
	return "r_resources"
}
