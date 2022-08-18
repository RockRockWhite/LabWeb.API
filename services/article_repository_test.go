package services

import (
	"fmt"
	"github.com/RockRockWhite/LabWeb.API/config"
	"github.com/RockRockWhite/LabWeb.API/entities"
	"testing"
)

func TestArticleRepository_AddArticle(t *testing.T) {
	// 初始化配置文件
	config.Init("../config/config.yml")

	repository := NewArticleRepository(true)
	repository.AddArticle(&entities.Article{
		UserId:  1,
		Title:   "测试博文",
		Content: "# hello world",
		Views:   0,
		Tags: []entities.Tag{{
			ArticleId: 0,
			Name:      "测试Tag",
		}},
		Comments: nil,
		Stars:    nil,
	})
}

func TestArticleRepository_GetArticle(t *testing.T) {
	// 初始化配置文件
	config.Init("../config/config.yml")

	repository := NewArticleRepository(true)

	article := repository.GetArticle(1)

	fmt.Printf("%+v", article)
}

func TestArticleRepository_GetArticles(t *testing.T) {
	// 初始化配置文件
	config.Init("../config/config.yml")

	repository := NewArticleRepository(true)

	articles := repository.GetArticles()

	fmt.Printf("%+v", articles)
}

func TestArticleRepository_AddStar(t *testing.T) {
	// 初始化配置文件
	config.Init("../config/config.yml")

	repository := NewArticleRepository(true)

	if _, err := repository.AddStar(&entities.Star{UserId: 1, ArticleId: 1}); err != nil {
		t.Fatalf("%v", err)
	}
}

func TestArticleRepository_GetStar(t *testing.T) {
	// 初始化配置文件
	config.Init("../config/config.yml")

	repository := NewArticleRepository(true)

	star, _ := repository.GetStar(1)

	fmt.Printf("%+v", star)
}

func TestArticleRepository_GetStars(t *testing.T) {
	// 初始化配置文件
	config.Init("../config/config.yml")

	repository := NewArticleRepository(true)
	stars := repository.GetStars(1)

	fmt.Printf("%+v", stars)
}

func TestArticleRepository_DeleteStar(t *testing.T) {
	// 初始化配置文件
	config.Init("../config/config.yml")

	repository := NewArticleRepository(true)
	repository.DeleteStar(1)

	if repository.StarExists(1) {
		t.Fatalf("delete fatalF")
	}
}

func TestArticleRepository_AddTag(t *testing.T) {
	// 初始化配置文件
	config.Init("../config/config.yml")

	repository := NewArticleRepository(true)
	repository.AddTag(&entities.Tag{ArticleId: 1, Name: "测试Tag"})
}

func TestArticleRepository_GetTag(t *testing.T) {
	// 初始化配置文件
	config.Init("../config/config.yml")

	repository := NewArticleRepository(true)
	tag, _ := repository.GetTag(1)

	fmt.Printf("%+v", tag)
}

func TestArticleRepository_GetTags(t *testing.T) {
	// 初始化配置文件
	config.Init("../config/config.yml")

	repository := NewArticleRepository(true)
	tags := repository.GetTags(1)

	fmt.Printf("%+v", tags)
}

func TestArticleRepository_UpdateTag(t *testing.T) {
	// 初始化配置文件
	config.Init("../config/config.yml")

	repository := NewArticleRepository(true)
	tag, _ := repository.GetTag(1)
	tag.Name = "被我修改过了"

	repository.UpdateTag(tag)
}

func TestArticleRepository_DeleteTag(t *testing.T) {
	// 初始化配置文件
	config.Init("../config/config.yml")

	repository := NewArticleRepository(true)
	repository.DeleteTag(1)
}

func TestArticleRepository_AddComment(t *testing.T) {
	// 初始化配置文件
	config.Init("../config/config.yml")

	repository := NewArticleRepository(true)
	repository.AddComment(&entities.Comment{
		UserId:    1,
		Content:   "这是一条评论测试F",
		ArticleId: 1,
		ParentId:  0,
	})
}

func TestArticleRepository_GetComment(t *testing.T) {
	// 初始化配置文件
	config.Init("../config/config.yml")

	repository := NewArticleRepository(true)

	comment, _ := repository.GetComment(1)

	fmt.Printf("%+v", comment)
}

func TestArticleRepository_GetComments(t *testing.T) {
	// 初始化配置文件
	config.Init("../config/config.yml")

	repository := NewArticleRepository(true)

	comments := repository.GetComments(1)

	fmt.Printf("%+v", comments)
}

func TestArticleRepository_UpdateComment(t *testing.T) {
	// 初始化配置文件
	config.Init("../config/config.yml")

	repository := NewArticleRepository(true)
	comment, _ := repository.GetComment(1)
	comment.Content = "这个被我change了"

	repository.UpdateComment(comment)
}

func TestArticleRepository_DeleteComment(t *testing.T) {
	// 初始化配置文件
	config.Init("../config/config.yml")

	repository := NewArticleRepository(true)
	repository.DeleteComment(1)
}
