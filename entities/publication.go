package entities

import (
	"gorm.io/gorm"
	"time"
)

// Publication 论文实体类
type Publication struct {
	gorm.Model

	Title          string           // 论文标题
	Abstract       string           // 论文简介 富文本
	Content        string           // 论文内容 富文本
	Authors        string           // 论文作者
	State          PublicationState // 论文状态 枚举
	PublishedAt    time.Time        // 发布时间
	PublishedIn    string           // 发表单位
	LastModifiedId uint             `gorm:"default:null"` // 最后修改者Id

}

func (Publication) TableName() string {
	return "p_publications"
}

type PublicationState int

const (
	Publication_Public PublicationState = iota
	Publication_Private
)
