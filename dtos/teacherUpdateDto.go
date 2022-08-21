package dtos

import (
	"github.com/RockRockWhite/LabWeb.API/entities"
)

// TeacherUpdateDto 修改教师Dto
type TeacherUpdateDto struct {
	FullName     string // 教师姓名
	Avatar       string // 教师头像
	Title        string // 教师职称
	Email        string // 教师邮箱
	Phone        string // 教师手机号
	Intro        string // 教师简历 富文本
	RelatedLinks string // 教师相关链接 富文本
}

// TeacherDtoFromEntity 从entity转换UpdateDto
func TeacherDtoFromEntity(t *entities.Teacher) *TeacherUpdateDto {
	return &TeacherUpdateDto{
		FullName:     t.FullName,
		Avatar:       t.Avatar,
		Title:        t.Title,
		Email:        t.Email,
		Phone:        t.Phone,
		Intro:        t.Intro,
		RelatedLinks: t.RelatedLinks,
	}
}

// ApplyUpdateToEntity 将Update应用到Entity
func (dto *TeacherUpdateDto) ApplyUpdateToEntity(entity *entities.Teacher, lastModifiedId uint) {
	entity.FullName = dto.FullName
	entity.Avatar = dto.Avatar
	entity.Title = dto.Title
	entity.Email = dto.Email
	entity.Phone = dto.Phone
	entity.Intro = dto.Intro
	entity.RelatedLinks = dto.RelatedLinks
	entity.LastModifiedId = lastModifiedId
}
