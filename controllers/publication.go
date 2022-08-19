package controllers

import (
	"fmt"
	"github.com/RockRockWhite/LabWeb.API/dtos"
	"github.com/RockRockWhite/LabWeb.API/services"
	"github.com/RockRockWhite/LabWeb.API/utils"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"strconv"
)

var publicationRepository *services.PublicationRepository

// InitPublicationController 初始化Controller
func InitPublicationController() {
	publicationRepository = services.NewPublicationRepository(true)
}

// AddPublication 添加论文
func AddPublication(c *gin.Context) {
	var dto dtos.PublicationAddDto

	if err := c.ShouldBind(&dto); err != nil {
		c.JSON(http.StatusBadRequest, dtos.ErrorDto{
			Message:          "Bind Model Error",
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	// 获得用户信息
	claims := c.MustGet("claims").(*utils.JwtClaims)

	entity := dto.ToEntity(claims.Id)
	_, err := publicationRepository.AddPublication(entity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.ErrorDto{
			Message:          err.Error(),
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	c.JSON(http.StatusCreated, dtos.ParsePublicationEntity(entity))
}

// GetPublication 获得论文
func GetPublication(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	if !publicationRepository.PublicationExists(uint(id)) {
		c.JSON(http.StatusNotFound, dtos.ErrorDto{
			Message:          fmt.Sprintf("Publication %v not found!", id),
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	entity, err := publicationRepository.GetPublication(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.ErrorDto{
			Message:          err.Error(),
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	// 转换为Dto
	c.JSON(http.StatusOK, dtos.ParsePublicationEntity(entity))
}

// GetPublications 批量获得论文
func GetPublications(c *gin.Context) {
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

	entities, err := publicationRepository.GetPublications(limit, (page-1)*limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.ErrorDto{
			Message:          err.Error(),
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}
	// 转换为Dto
	getDtos := make([]dtos.PublicationGetDto, 0, len(entities))
	for _, entity := range entities {
		getDtos = append(getDtos, *dtos.ParsePublicationEntity(&entity))
	}

	c.JSON(http.StatusOK, getDtos)
}

// PatchPublication 修改论文
func PatchPublication(c *gin.Context) {
	// 获得更新id
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if !articleRepository.ArticleExists(uint(id)) {
		c.JSON(http.StatusNotFound, dtos.ErrorDto{
			Message:          fmt.Sprintf("Publication %v not found!", id),
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	entity, err := publicationRepository.GetPublication(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.ErrorDto{
			Message:          err.Error(),
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	// 获得用户信息 判断用户是否对该博文具有修改权 TODO
	claims := c.MustGet("claims").(*utils.JwtClaims)
	if !claims.IsAdmin {
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
	dto := dtos.PublicationDtoFromEntity(entity)
	utils.ApplyJsonPatch(dto, patchJson)
	dto.ApplyUpdateToEntity(entity, claims.Id)

	// 更新数据库
	err = publicationRepository.UpdatePublication(entity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.ErrorDto{
			Message:          err.Error(),
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	c.Status(http.StatusNoContent)
}

// DeletePublication 删除论文
func DeletePublication(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	if !publicationRepository.PublicationExists(uint(id)) {
		c.JSON(http.StatusNotFound, dtos.ErrorDto{
			Message:          fmt.Sprintf("Publication %v not found!", id),
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	err := publicationRepository.DeletePublication(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.ErrorDto{
			Message:          err.Error(),
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}
	c.Status(http.StatusNoContent)
}
