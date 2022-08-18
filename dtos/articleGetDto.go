package dtos

import (
	"github.com/RockRockWhite/LabWeb.API/entities"
	"time"
)

type ArticleGetDto struct {
	Id        uint      // 博文Id
	CreatedAt time.Time // 博文创建时间
	UpdatedAt time.Time // 博文修改时间

	Titile  string // 博文标题
	Content string // 博文内容
	Views   uint   // 博文浏览量

	Tags     []TagGetDto     // 博文标签
	Comments []CommentGetDto // 博文评论
	Stars    []StarGetDto    // 博文点赞
}

func ParseArticleEntity(article *entities.Article) *ArticleGetDto {
	dto := ArticleGetDto{
		Id:        article.ID,
		CreatedAt: article.CreatedAt,
		UpdatedAt: article.UpdatedAt,
		Titile:    article.Title,
		Content:   article.Content,
		Views:     article.Views,
		Tags:      nil,
		Comments:  nil,
		Stars:     nil,
	}
	// 转换Tags
	dto.Tags = make([]TagGetDto, 0, len(article.Tags))
	for _, tag := range article.Tags {
		dto.Tags = append(dto.Tags, *ParseTagEntity(&tag))
	}

	// 转换Comments
	dto.Comments = make([]CommentGetDto, 0, len(article.Comments))
	for _, comment := range article.Comments {
		dto.Comments = append(dto.Comments, *ParseCommentEntity(&comment))
	}

	// 转换Stars
	dto.Stars = make([]StarGetDto, 0, len(article.Stars))
	for _, star := range article.Stars {
		dto.Stars = append(dto.Stars, *ParseStarEntity(&star))
	}

	return &dto
}
