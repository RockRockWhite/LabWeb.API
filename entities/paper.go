package entities

import (
	"gorm.io/gorm"
	"time"
)

// Paper 论文实体类
type Paper struct {
	gorm.Model

	Title          string     // 论文标题
	Abstract       string     // 论文简介 富文本
	Content        string     // 论文内容 富文本
	Authors        string     // 论文作者
	State          PaperState // 论文状态 枚举
	PublishedAt    time.Time  `gorm:"default:null"` // 发布时间
	PublishedIn    string     `gorm:"default:null"` // 发表单位
	LastModifiedId uint       `gorm:"default:null"` // 最后修改者Id

}

func (Paper) TableName() string {
	return "p_papers"
}

type PaperState int

const (
	PaperState_Public PaperState = iota
	PaperState_Private
)
