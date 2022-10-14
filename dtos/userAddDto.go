package dtos

import (
	"github.com/RockRockWhite/LabWeb.API/entities"
	"github.com/RockRockWhite/LabWeb.API/utils"
)

// UserAddDto 创建用户Dto
type UserAddDto struct {
	Username  string // 昵称
	Fullname  string `gorm:"default:null"` // 实验室成员全名
	Password  string // 密码
	Email     string // 邮箱
	Telephone string // 手机号码
	AvatarUrl string // 头像链接
}

// ToEntity 转换成Entity
func (dto *UserAddDto) ToEntity() *entities.User {
	salt := utils.GenerateSalt()
	passwordHash := utils.EncryptPasswordHash(dto.Password, salt)

	return &entities.User{
		Username:     dto.Username,
		Fullname:     dto.Fullname,
		PasswordHash: passwordHash,
		Salt:         salt,
		Email:        dto.Email,
		VerifyState:  false,
		Telephone:    dto.Telephone,
		IsAdmin:      false,
	}
}
