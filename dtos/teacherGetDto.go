package dtos

import (
	"github.com/RockRockWhite/LabWeb.API/entities"
)

type TeacherGetDto struct {
	Id             uint   // 教师Id
	FullName       string // 教师姓名
	Avatar         string // 教师头像
	Title          string // 教师职称
	Email          string // 教师邮箱
	Phone          string // 教师手机号
	Intro          string // 教师简历 富文本
	RelatedLinks   string // 教师相关链接 富文本
	lastModifiedId uint   // 最后修改者Id
}

func ParseTeacherEntity(t *entities.Teacher) *TeacherGetDto {
	dto := TeacherGetDto{
		Id:             t.ID,
		FullName:       t.FullName,
		Avatar:         t.Avatar,
		Title:          t.Title,
		Email:          t.Email,
		Phone:          t.Phone,
		Intro:          t.Intro,
		RelatedLinks:   t.RelatedLinks,
		lastModifiedId: t.LastModifiedId,
	}

	return &dto
}
