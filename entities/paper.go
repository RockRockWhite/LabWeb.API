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
	Authors        string     // 论文作者
	Thumbnail      string     // 论文缩略图
	Link           string     // 论文链接
	Pdf            string     // 论文pdf
	Code           string     // 论文代码
	State          PaperState // 论文状态 枚举
	PublishedAt    time.Time  // 发布时间
	PublishedIn    string     `gorm:"default:null"` // 发表单位
	LastModifiedId uint       `gorm:"default:null"` // 最后修改者Id

}

func (Paper) TableName() string {
	return "p_papers"
}

type PaperState int

const (
	PaperState_Private PaperState = iota
	PaperState_Public
)
