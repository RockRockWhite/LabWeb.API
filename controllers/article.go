package controllers

import (
	"fmt"
	"github.com/RockRockWhite/LabWeb.API/dtos"
	"github.com/RockRockWhite/LabWeb.API/entities"
	"github.com/RockRockWhite/LabWeb.API/services"
	"github.com/RockRockWhite/LabWeb.API/utils"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"strconv"
)

var articleRepository *services.ArticleRepository

// InitArticleController 初始化博文Controller
func InitArticleController() {
	articleRepository = services.NewArticleRepository(true)
}

// AddArticle 添加博文
func AddArticle(c *gin.Context) {
	var articleDto dtos.ArticleAddDto

	if err := c.ShouldBind(&articleDto); err != nil {
		c.JSON(http.StatusBadRequest, dtos.ErrorDto{
			Message:          "Bind Model Error",
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	// 获得用户信息
	claims := c.MustGet("claims").(*utils.JwtClaims)

	entity := articleDto.ToEntity(claims.Id)
	articleRepository.AddArticle(entity)

	c.JSON(http.StatusCreated, dtos.ParseArticleEntity(entity))
}

// GetArticle 获得博文
func GetArticle(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	if !articleRepository.ArticleExists(uint(id)) {
		c.JSON(http.StatusNotFound, dtos.ErrorDto{
			Message:          fmt.Sprintf("Article %v not found!", id),
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	article := articleRepository.GetArticle(uint(id))
	// 转换为Dto
	c.JSON(http.StatusOK, dtos.ParseArticleEntity(article))
}

// GetArticles 批量获得博文
func GetArticles(c *gin.Context) {
	// 获得page limit
	page, pageQueryErr := strconv.Atoi(c.DefaultQuery("page", "1"))
	if pageQueryErr != nil {
		c.JSON(http.StatusBadRequest, dtos.ErrorDto{
			Message:          "Incorrect query field page",
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}
	limit, limitQueryErr := strconv.Atoi(c.DefaultQuery("limit", "20"))
	if limitQueryErr != nil {
		c.JSON(http.StatusBadRequest, dtos.ErrorDto{
			Message:          "Incorrect query field limit",
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	articles := articleRepository.GetArticles(limit, (page-1)*limit)

	// 转换为Dto
	articleDtos := make([]dtos.ArticleGetDto, 0, len(articles))
	for _, article := range articles {
		articleDtos = append(articleDtos, *dtos.ParseArticleEntity(&article))
	}

	c.JSON(http.StatusOK, articleDtos)
}

// PatchArticle 修改文章
func PatchArticle(c *gin.Context) {
	// 获得更新id
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if !articleRepository.ArticleExists(uint(id)) {
		c.JSON(http.StatusNotFound, dtos.ErrorDto{
			Message:          fmt.Sprintf("Article %v not found!", id),
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}
	article := articleRepository.GetArticle(uint(id))

	// 获得用户信息 判断用户是否对该博文具有修改权
	// 修改权: 改博文为用户所有 或 该用户是管理员
	claims := c.MustGet("claims").(*utils.JwtClaims)

	if article.UserId != claims.Id && !claims.IsAdmin {
		c.JSON(http.StatusForbidden, dtos.ErrorDto{
			Message:          "Permission denied for changing this resource!",
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	// 获得patchJson
	patchJson, getRawDataErr := c.GetRawData()
	if getRawDataErr != nil {
		c.JSON(http.StatusBadRequest, dtos.ErrorDto{
			Message:          "Bind Model Error",
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	// 应用patch
	dto := dtos.ArticleUpdateDtoFromEntity(article)
	utils.ApplyJsonPatch(dto, patchJson)
	dto.ApplyUpdateToEntity(article)

	// 更新数据库
	articleRepository.UpdateArticle(article)

	c.Status(http.StatusNoContent)
}

// PutArticle 替换文章
func PutArticle(c *gin.Context) {
	var articleDto dtos.ArticleUpdateDto

	if err := c.ShouldBind(&articleDto); err != nil {
		c.JSON(http.StatusBadRequest, dtos.ErrorDto{
			Message:          "Bind Model Error",
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	// 获得更新id
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	// 获得用户信息
	claims := c.MustGet("claims").(*utils.JwtClaims)

	// 处理文章不存在
	if !articleRepository.ArticleExists(uint(id)) {
		entity := entities.Article{UserId: claims.Id}
		entity.ID = uint(id)
		articleDto.ApplyUpdateToEntity(&entity)
		articleRepository.AddArticle(&entity)

		c.JSON(http.StatusCreated, dtos.ParseArticleEntity(&entity))
		return
	}

	// 处理文章存在
	entity := articleRepository.GetArticle(uint(id))
	// 判断是否有修改权
	if entity.UserId != claims.Id && !claims.IsAdmin {
		c.JSON(http.StatusForbidden, dtos.ErrorDto{
			Message:          "Permission denied for changing this resource!",
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	articleDto.ApplyUpdateToEntity(entity)
	articleRepository.UpdateArticle(entity)

	c.Status(http.StatusNoContent)
}

// DeleteArticle 删除文章
func DeleteArticle(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	if !articleRepository.ArticleExists(uint(id)) {
		c.JSON(http.StatusNotFound, dtos.ErrorDto{
			Message:          fmt.Sprintf("Article %v not found!", id),
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}
	article := articleRepository.GetArticle(uint(id))

	// 获得用户信息 判断用户是否对该博文具有修改权
	// 修改权: 改博文为用户所有 或 该用户是管理员
	claims := c.MustGet("claims").(*utils.JwtClaims)

	if article.UserId != claims.Id && !claims.IsAdmin {
		c.JSON(http.StatusForbidden, dtos.ErrorDto{
			Message:          "Permission denied for changing this resource!",
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	articleRepository.DeleteArticle(uint(id))
	c.Status(http.StatusNoContent)
}
