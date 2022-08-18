package services

import (
	"fmt"
	"github.com/RockRockWhite/LabWeb.API/entities"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ArticleRepository struct {
	db *gorm.DB
}

// NewArticleRepository 创建新博客Repository
func NewArticleRepository(autoMigrate bool) *ArticleRepository {
	Host := viper.GetString("DataBase.Host")
	Port := viper.GetString("DataBase.Port")
	Username := viper.GetString("DataBase.Username")
	Password := viper.GetString("DataBase.Password")
	DBName := viper.GetString("DataBase.DBName")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", Username, Password, Host, Port, DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("Fatal error open database:%s %s \n", dsn, err))
	}

	// 完成Article迁移
	if autoMigrate {
		if err := db.AutoMigrate(&entities.Article{}); err != nil {
			panic(fmt.Errorf("Fatal migrate database %s : %s \n", "Article", err))
		}

		if err := db.AutoMigrate(&entities.Star{}); err != nil {
			panic(fmt.Errorf("Fatal migrate database %s : %s \n", "Star", err))
		}

		if err := db.AutoMigrate(&entities.Tag{}); err != nil {
			panic(fmt.Errorf("Fatal migrate database %s : %s \n", "Tag", err))
		}

		if err := db.AutoMigrate(&entities.Comment{}); err != nil {
			panic(fmt.Errorf("Fatal migrate database %s : %s \n", "Comment", err))
		}
	}

	repository := ArticleRepository{db}
	return &repository
}

// GetArticle  从id获得博文
func (repository *ArticleRepository) GetArticle(id uint) *entities.Article {
	if !repository.ArticleExists(id) {
		panic(fmt.Errorf("article id %v not exists", id))
	}
	var article entities.Article
	if result := repository.db.First(&article, id); result.Error != nil {
		panic(fmt.Errorf("failed to get article id %v : %s", id, result.Error))
	}

	// 获得子成员
	article.Stars = repository.GetStars(id)
	article.Tags = repository.GetTags(id)
	article.Comments = repository.GetComments(id)

	return &article
}

func (repository *ArticleRepository) GetArticles(limit int, offset int) []entities.Article {
	var articles []entities.Article
	if result := repository.db.Limit(limit).Offset(offset).Find(&articles); result.Error != nil {
		panic(fmt.Errorf("failed to get articles: %s", result.Error))
	}

	// 获得各个子成员
	for i, _ := range articles {
		articles[i].Stars = repository.GetStars(articles[i].ID)
		articles[i].Tags = repository.GetTags(articles[i].ID)
		articles[i].Comments = repository.GetComments(articles[i].ID)
	}

	return articles
}

// AddArticle 创建博文
func (repository *ArticleRepository) AddArticle(article *entities.Article) uint {
	if result := repository.db.Create(article); result.Error != nil {
		panic(fmt.Errorf("failed to add article %+v : %s", article, result.Error))
	}

	// 添加Tags
	for i, _ := range article.Tags {
		article.Tags[i].ArticleId = article.ID
		repository.AddTag(&article.Tags[i])
	}

	return article.ID
}

// UpdateArticle 更新博文
func (repository *ArticleRepository) UpdateArticle(article *entities.Article) {
	// 更新子项
	for i, _ := range article.Tags {
		article.Tags[i].ArticleId = article.ID
		repository.UpdateTag(&article.Tags[i])
	}
	for i, _ := range article.Comments {
		article.Comments[i].ArticleId = article.ID
		repository.UpdateComment(&article.Comments[i])
	}
	for i, _ := range article.Stars {
		article.Stars[i].ArticleId = article.ID
		repository.UpdateStar(&article.Stars[i])
	}

	if result := repository.db.Save(&article); result.Error != nil {
		panic(fmt.Errorf("failed to update article %+v : %s", article, result.Error))
	}
}

// DeleteArticle 删除博文
func (repository *ArticleRepository) DeleteArticle(id uint) {
	if !repository.ArticleExists(id) {
		panic(fmt.Errorf("article id %v not exists", id))
	}

	// 删除各个子成员
	article := repository.GetArticle(id)
	for i, _ := range article.Stars {
		repository.DeleteStar(article.Stars[i].ID)
	}
	for i, _ := range article.Tags {
		repository.DeleteTag(article.Tags[i].ID)
	}
	for i, _ := range article.Comments {
		repository.DeleteComment(article.Comments[i].ID)
	}

	if result := repository.db.Delete(&entities.Article{}, id); result.Error != nil {
		panic(fmt.Errorf("failed to delete article id %v : %s", id, result.Error))
	}
}

// ArticleExists 判断该id是否存在
func (repository *ArticleRepository) ArticleExists(id uint) bool {
	var article entities.Article
	result := repository.db.First(&article, id)

	return result.RowsAffected >= 1
}

// GetStar 获得单个点赞
func (repository *ArticleRepository) GetStar(id uint) (*entities.Star, error) {
	if !repository.StarExists(id) {
		panic(fmt.Errorf("star id %v not exists", id))
	}

	var star entities.Star
	result := repository.db.First(&star, id)

	return &star, result.Error
}

// GetStars 获得博文的所有点赞
func (repository *ArticleRepository) GetStars(articleId uint) []entities.Star {
	if !repository.ArticleExists(articleId) {
		panic(fmt.Errorf("article id %v not exists", articleId))
	}

	var stars []entities.Star
	if result := repository.db.Where(&entities.Star{ArticleId: articleId}).Find(&stars); result.Error != nil {
		panic(fmt.Errorf("failed to get stars articleId %v : %s", articleId, result.Error))
	}

	return stars
}

// AddStar 添加一条点赞记录
func (repository *ArticleRepository) AddStar(star *entities.Star) (uint, error) {
	if !repository.ArticleExists(star.ArticleId) {
		panic(fmt.Errorf("article id %v not exists", star.ArticleId))
	}

	result := repository.db.Create(star)

	return star.ID, result.Error
}

// UpdateStar 更新一条点赞记录
func (repository *ArticleRepository) UpdateStar(star *entities.Star) {
	if !repository.ArticleExists(star.ArticleId) {
		panic(fmt.Errorf("article id %v not exists", star.ArticleId))
	}

	if result := repository.db.Save(star); result.Error != nil {
		panic(fmt.Errorf("failed to update star %v : %s", star, result.Error))
	}
}

// DeleteStar 删除一条点赞记录
func (repository *ArticleRepository) DeleteStar(id uint) {
	if !repository.StarExists(id) {
		panic(fmt.Errorf("star id %v not exists", id))
	}

	var star entities.Star
	if result := repository.db.Delete(&star, id); result.Error != nil {
		panic(fmt.Errorf("failed to delete star id %v : %s", id, result.Error))
	}
}

// StarExists 判断该点赞是否存在
func (repository *ArticleRepository) StarExists(id uint) bool {
	var star entities.Star
	result := repository.db.First(&star, id)

	return result.RowsAffected >= 1
}

// GetTag 获得单个F的Tag
func (repository *ArticleRepository) GetTag(id uint) (*entities.Tag, error) {
	if !repository.TagExists(id) {
		panic(fmt.Errorf("tag id %v not exists", id))
	}

	var tag entities.Tag
	result := repository.db.First(&tag, id)

	return &tag, result.Error
}

// GetTags 获得博客的所有Tag
func (repository *ArticleRepository) GetTags(articleId uint) []entities.Tag {
	if !repository.ArticleExists(articleId) {
		panic(fmt.Errorf("article id %v not exists", articleId))
	}

	var tags []entities.Tag
	if result := repository.db.Where(&entities.Tag{ArticleId: articleId}).Find(&tags); result.Error != nil {
		panic(fmt.Errorf("failed to get tags articleId %v : %s", articleId, result.Error))
	}

	return tags
}

// AddTag 增加一条Tag
func (repository *ArticleRepository) AddTag(tag *entities.Tag) uint {
	if !repository.ArticleExists(tag.ArticleId) {
		panic(fmt.Errorf("article id %v not exists", tag.ArticleId))
	}

	if result := repository.db.Create(tag); result.Error != nil {
		panic(fmt.Errorf("failed to add tag %+v : %s", tag, result.Error))
	}

	return tag.ID
}

// UpdateTag 更新Tag
func (repository *ArticleRepository) UpdateTag(tag *entities.Tag) {
	if !repository.ArticleExists(tag.ArticleId) {
		panic(fmt.Errorf("article id %v not exists", tag.ArticleId))
	}

	if result := repository.db.Save(tag); result.Error != nil {
		panic(fmt.Errorf("failed to update tag %+v : %s", tag, result.Error))
	}
}

// DeleteTag 删除Tag
func (repository *ArticleRepository) DeleteTag(id uint) {
	if !repository.TagExists(id) {
		panic(fmt.Errorf("tag id %v not exists", id))
	}

	var tag entities.Tag
	if result := repository.db.Delete(&tag, id); result.Error != nil {
		panic(fmt.Errorf("failed to delete tag id %v : %s", id, result.Error))
	}
}

// TagExists Tag是否存在
func (repository *ArticleRepository) TagExists(id uint) bool {
	var tag entities.Tag
	result := repository.db.First(&tag, id)

	return result.RowsAffected >= 1
}

// GetComment 获得评论
func (repository *ArticleRepository) GetComment(id uint) (*entities.Comment, error) {
	if !repository.CommentExists(id) {
		panic(fmt.Errorf("comment id %v not exists", id))
	}

	var comment entities.Comment
	result := repository.db.First(&comment, id)

	return &comment, result.Error
}

// GetComments 获得评论
func (repository *ArticleRepository) GetComments(articleId uint) []entities.Comment {
	if !repository.ArticleExists(articleId) {
		panic(fmt.Errorf("article id %v not exists", articleId))
	}

	var articles []entities.Comment
	if result := repository.db.Where(&entities.Comment{ArticleId: articleId}).Find(&articles); result.Error != nil {
		panic(fmt.Errorf("failed to get comments articleId %v : %s", articleId, result.Error))
	}

	return articles
}

// AddComment 添加一条评论
func (repository *ArticleRepository) AddComment(comment *entities.Comment) (uint, error) {
	if !repository.ArticleExists(comment.ArticleId) {
		panic(fmt.Errorf("article id %v not exists", comment.ArticleId))
	}

	result := repository.db.Create(comment)

	return comment.ID, result.Error
}

// UpdateComment 更新一条评论
func (repository *ArticleRepository) UpdateComment(comment *entities.Comment) {
	if !repository.ArticleExists(comment.ArticleId) {
		panic(fmt.Errorf("article id %v not exists", comment.ArticleId))
	}

	if result := repository.db.Save(comment); result.Error != nil {
		panic(fmt.Errorf("failed to update comment %v : %s", comment, result.Error))
	}
}

// DeleteComment 删除一条评论
func (repository *ArticleRepository) DeleteComment(id uint) {
	if !repository.CommentExists(id) {
		panic(fmt.Errorf("comment id %v not exists", id))
	}

	var comment entities.Comment
	if result := repository.db.Delete(&comment, id); result.Error != nil {
		panic(fmt.Errorf("failed to delete comment id %v : %s", id, result.Error))
	}
}

// CommentExists 判断评论否存在
func (repository *ArticleRepository) CommentExists(id uint) bool {
	var comment entities.Comment
	result := repository.db.First(&comment, id)

	return result.RowsAffected >= 1
}
