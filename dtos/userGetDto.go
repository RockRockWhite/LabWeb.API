package dtos

import (
	"github.com/RockRockWhite/LabWeb.API/entities"
)

// UserGetDto 获得用户Dto
type UserGetDto struct {
	Id          uint   // 用户Id
	Username    string // 昵称
	Fullname    string // 实验室成员全名
	Email       string // 邮箱
	VerifyState bool   // 邮箱验证状态
	Telephone   string // 手机号码
	IsAdmin     bool   // 是否管理员
	AvatarUrl   string // 头像链接
}

// ParseUserEntity 将entity转换为GetDto
func ParseUserEntity(user *entities.User) *UserGetDto {
	return &UserGetDto{
		Id:          user.ID,
		Username:    user.Username,
		Fullname:    user.Fullname,
		Email:       user.Email,
		VerifyState: user.VerifyState,
		Telephone:   user.Telephone,
		IsAdmin:     user.IsAdmin,
		AvatarUrl:   user.AvatarUrl,
	}
}
