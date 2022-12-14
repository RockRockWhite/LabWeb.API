package entities

import (
	"gorm.io/gorm"
)

// User 用户实体类
type User struct {
	gorm.Model
	Username     string `gorm:"unique"`       // 实验室成员昵称 需要保证唯一
	Fullname     string `gorm:"default:null"` // 实验室成员全名
	PasswordHash string // 密码
	Salt         string // 密码盐值
	Email        string `gorm:"default:null"` // 邮箱
	VerifyState  bool   // 邮箱验证状态
	Telephone    string `gorm:"default:null"` // 手机号码
	IsAdmin      bool   // 是否管理员
	AvatarUrl    string `gorm:"default:null"` // 头像链接
	Config       string `gorm:"default:null"` // 用户配置文件
}

func (User) TableName() string {
	return "u_user"
}
