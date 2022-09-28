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

var newsRepository *services.NewsRepository

func init() {
	newsRepository = services.GetNewsRepository()
}

// AddNews 添加新闻
func AddNews(c *gin.Context) {
	var addDto dtos.NewsAddDto

	if err := c.ShouldBind(&addDto); err != nil {
		c.JSON(http.StatusBadRequest, dtos.ErrorDto{
			Message:          "Bind Model Error",
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	// 获得用户信息
	claims := c.MustGet("claims").(*utils.JwtClaims)

	var entity entities.News
	dtos.GetNewsAddDtoConverter().Convert(&entity, &addDto)
	entity.LastModifiedId = claims.Id

	_, err := newsRepository.AddNews(&entity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.ErrorDto{
			Message:          err.Error(),
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	var getDto dtos.NewsGetDto
	err = dtos.GetNewsGetDtoConverter().Convert(&getDto, &entity)
	c.JSON(http.StatusCreated, getDto)
}

// GetNews 获得新闻
func GetNews(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	if !newsRepository.NewsExists(uint(id)) {
		c.JSON(http.StatusNotFound, dtos.ErrorDto{
			Message:          fmt.Sprintf("News %v not found!", id),
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	entity, err := newsRepository.GetNews(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.ErrorDto{
			Message:          err.Error(),
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	// 转换为Dto
	var getDto dtos.NewsGetDto
	dtos.GetNewsGetDtoConverter().Convert(&getDto, entity)
	c.JSON(http.StatusOK, getDto)
}

// GetNewsList 批量获得新闻
func GetNewsList(c *gin.Context) {
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
			Message:          "Incorrect query field limit.",
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	entities, err := newsRepository.GetNewsList(limit, (page-1)*limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.ErrorDto{
			Message:          err.Error(),
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}
	// 转换为Dto
	getDtos := make([]dtos.NewsGetDto, 0, len(entities))
	for _, entity := range entities {
		var each dtos.NewsGetDto
		dtos.GetNewsGetDtoConverter().Convert(&each, &entity)
		getDtos = append(getDtos, each)
	}

	c.JSON(http.StatusOK, getDtos)
}

// PatchNews 修改新闻
func PatchNews(c *gin.Context) {
	// 获得更新id
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if !newsRepository.NewsExists(uint(id)) {
		c.JSON(http.StatusNotFound, dtos.ErrorDto{
			Message:          fmt.Sprintf("News %v not found!", id),
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	entity, err := newsRepository.GetNews(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.ErrorDto{
			Message:          err.Error(),
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	claims := c.MustGet("claims").(*utils.JwtClaims)

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
	var updateDto dtos.NewsUpdateDto
	dtos.GetNewsUpdateDtoEntityConverter().Convert(&updateDto, entity)
	utils.ApplyJsonPatch(&updateDto, patchJson)
	dtos.GetNewsEntityUpdateDtoConverter().Convert(entity, &updateDto)
	entity.LastModifiedId = claims.Id

	// 更新数据库
	err = newsRepository.UpdateNews(entity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.ErrorDto{
			Message:          err.Error(),
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	c.Status(http.StatusNoContent)
}

// DeleteNews 删除新闻
func DeleteNews(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	if !newsRepository.NewsExists(uint(id)) {
		c.JSON(http.StatusNotFound, dtos.ErrorDto{
			Message:          fmt.Sprintf("News %v not found!", id),
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	err := newsRepository.DeleteNews(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.ErrorDto{
			Message:          err.Error(),
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}
	c.Status(http.StatusNoContent)
}
