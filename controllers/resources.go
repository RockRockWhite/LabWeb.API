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

var resourcesRepository *services.ResourcesRepository

func init() {
	resourcesRepository = services.GetResourcesRepository()
}

// AddResources 添加资源
func AddResources(c *gin.Context) {
	var addDto dtos.ResourceAddDto

	if err := c.ShouldBind(&addDto); err != nil {
		c.JSON(http.StatusBadRequest, dtos.ErrorDto{
			Message:          "Bind Model Error",
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	// 获得用户信息
	claims := c.MustGet("claims").(*utils.JwtClaims)

	var entity entities.Resource
	dtos.GetResourceAddDtoConverter().Convert(&entity, &addDto)
	entity.LastModifiedId = claims.Id

	_, err := resourcesRepository.AddResource(&entity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.ErrorDto{
			Message:          err.Error(),
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	var getDto dtos.ResourceGetDto
	err = dtos.GetResourceGetDtoConverter().Convert(&getDto, &entity)
	c.JSON(http.StatusCreated, getDto)
}

// GetResource 获得资源
func GetResource(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	if !resourcesRepository.ResourceExists(uint(id)) {
		c.JSON(http.StatusNotFound, dtos.ErrorDto{
			Message:          fmt.Sprintf("Resources %v not found!", id),
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	entity, err := resourcesRepository.GetResource(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.ErrorDto{
			Message:          err.Error(),
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	// 转换为Dto
	var getDto dtos.ResourceGetDto
	dtos.GetResourceGetDtoConverter().Convert(&getDto, entity)
	c.JSON(http.StatusOK, getDto)
}

// GetResourcesList 批量获得资源
func GetResourcesList(c *gin.Context) {
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

	entities, err := resourcesRepository.GetResources(limit, (page-1)*limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.ErrorDto{
			Message:          err.Error(),
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}
	// 转换为Dto
	getDtos := make([]dtos.ResourceGetDto, 0, len(entities))
	for _, entity := range entities {
		var each dtos.ResourceGetDto
		dtos.GetResourceGetDtoConverter().Convert(&each, &entity)
		getDtos = append(getDtos, each)
	}

	c.JSON(http.StatusOK, getDtos)
}

// PatchResources 修改资源
func PatchResources(c *gin.Context) {
	// 获得更新id
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if !resourcesRepository.ResourceExists(uint(id)) {
		c.JSON(http.StatusNotFound, dtos.ErrorDto{
			Message:          fmt.Sprintf("Resources %v not found!", id),
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	entity, err := resourcesRepository.GetResource(uint(id))
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
	var updateDto dtos.ResourceUpdateDto
	dtos.GetResourceUpdateDtoEntityConverter().Convert(&updateDto, entity)
	utils.ApplyJsonPatch(&updateDto, patchJson)
	dtos.GetResourceEntityUpdateDtoConverter().Convert(entity, &updateDto)
	entity.LastModifiedId = claims.Id

	// 更新数据库
	err = resourcesRepository.UpdateResource(entity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.ErrorDto{
			Message:          err.Error(),
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	c.Status(http.StatusNoContent)
}

// DeleteResources 删除资源
func DeleteResources(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	if !resourcesRepository.ResourceExists(uint(id)) {
		c.JSON(http.StatusNotFound, dtos.ErrorDto{
			Message:          fmt.Sprintf("Resources %v not found!", id),
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	err := resourcesRepository.DeleteResource(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.ErrorDto{
			Message:          err.Error(),
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}
	c.Status(http.StatusNoContent)
}

// CountResources 获得资源数量
func CountResources(c *gin.Context) {
	c.JSON(http.StatusOK, struct {
		Count int64
	}{
		Count: resourcesRepository.Count(),
	})
}
