package entities

import (
	"gorm.io/gorm"
)

// Teacher 教师实体类
type Teacher struct {
	gorm.Model

	FullName       string `gorm:"not null"`     // 教师姓名
	Avatar         string `gorm:"not null"`     // 教师头像
	Title          string `gorm:"not null"`     // 教师职称
	Email          string `gorm:"default:null"` // 教师邮箱
	Phone          string `gorm:"default:null"` // 教师手机号
	Intro          string `gorm:"not null"`     // 教师简历 富文本
	RelatedLinks   string `gorm:"not null"`     // 教师相关链接 富文本
	LastModifiedId uint   `gorm:"default:null"` // 最后修改者Id
}

func (Teacher) TableName() string {
	return "t_teachers"
}
