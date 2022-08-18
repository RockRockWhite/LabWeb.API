package dtos

import "github.com/RockRockWhite/LabWeb.API/entities"

// ArticleUpdateDto 添加博文Dto
type ArticleUpdateDto struct {
	Title   string // 博文标题
	Content string // 博文内容
}

// ArticleUpdateDtoFromEntity 从entity转换UpdateDto
func ArticleUpdateDtoFromEntity(article *entities.Article) *ArticleUpdateDto {
	return &ArticleUpdateDto{
		Title:   article.Title,
		Content: article.Content,
	}
}

// ApplyUpdateToEntity 将Update应用到Entity
func (dto *ArticleUpdateDto) ApplyUpdateToEntity(entity *entities.Article) {
	entity.Title = dto.Title
	entity.Content = dto.Content
}
