package dtos

import (
	"github.com/RockRockWhite/LabWeb.API/entities"
)

// TeacherAddDto 添加教师Dto
type TeacherAddDto struct {
	FullName     string // 教师姓名
	Avatar       string // 教师头像
	Title        string // 教师职称
	Email        string // 教师邮箱
	Phone        string // 教师手机号
	Intro        string // 教师简历 富文本
	RelatedLinks string // 教师相关链接 富文本
}

// ToEntity 转换成Entity
func (dto *TeacherAddDto) ToEntity(lastModifiedId uint) *entities.Teacher {
	entity := entities.Teacher{
		FullName:       dto.FullName,
		Avatar:         dto.Avatar,
		Title:          dto.Title,
		Email:          dto.Email,
		Phone:          dto.Phone,
		Intro:          dto.Intro,
		RelatedLinks:   dto.RelatedLinks,
		LastModifiedId: lastModifiedId,
	}

	return &entity
}
