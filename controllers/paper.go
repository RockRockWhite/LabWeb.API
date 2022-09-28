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
	"time"
)

var papersRepository *services.PapersRepository

func init() {
	papersRepository = services.GetPapersRepository()
}

// AddPaper 添加论文
func AddPaper(c *gin.Context) {
	var dto dtos.PaperAddDto

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
	if entity.PublishedAt.IsZero() {
		entity.PublishedAt = time.Now()
	}

	_, err := papersRepository.AddPaper(entity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.ErrorDto{
			Message:          err.Error(),
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	c.JSON(http.StatusCreated, dtos.ParsePaperEntity(entity))
}

// GetPaper 获得论文
func GetPaper(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	if !papersRepository.PaperExists(uint(id)) {
		c.JSON(http.StatusNotFound, dtos.ErrorDto{
			Message:          fmt.Sprintf("Paper %v not found!", id),
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	entity, err := papersRepository.GetPaper(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.ErrorDto{
			Message:          err.Error(),
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	// 转换为Dto
	c.JSON(http.StatusOK, dtos.ParsePaperEntity(entity))
}

// GetPapers 批量获得论文
func GetPapers(c *gin.Context) {
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

	entities, err := papersRepository.GetPapers(limit, (page-1)*limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.ErrorDto{
			Message:          err.Error(),
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}
	// 转换为Dto
	getDtos := make([]dtos.PaperGetDto, 0, len(entities))
	for _, entity := range entities {
		getDtos = append(getDtos, *dtos.ParsePaperEntity(&entity))
	}

	c.JSON(http.StatusOK, getDtos)
}

// PatchPaper 修改论文
func PatchPaper(c *gin.Context) {
	// 获得更新id
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if !papersRepository.PaperExists(uint(id)) {
		c.JSON(http.StatusNotFound, dtos.ErrorDto{
			Message:          fmt.Sprintf("Paper %v not found!", id),
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	entity, err := papersRepository.GetPaper(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.ErrorDto{
			Message:          err.Error(),
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	// 获得用户信息 判断用户是否对该博文具有修改权
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
	dto := dtos.PaperDtoFromEntity(entity)
	utils.ApplyJsonPatch(dto, patchJson)
	dto.ApplyUpdateToEntity(entity, claims.Id)

	// 更新数据库
	err = papersRepository.UpdatePaper(entity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.ErrorDto{
			Message:          err.Error(),
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	c.Status(http.StatusNoContent)
}

// DeletePaper 删除论文
func DeletePaper(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	if !papersRepository.PaperExists(uint(id)) {
		c.JSON(http.StatusNotFound, dtos.ErrorDto{
			Message:          fmt.Sprintf("Paper %v not found!", id),
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	err := papersRepository.DeletePaper(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.ErrorDto{
			Message:          err.Error(),
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}
	c.Status(http.StatusNoContent)
}

// CountPaper 获得论文数量
func CountPaper(c *gin.Context) {
	c.JSON(http.StatusOK, struct {
		Count int64
	}{
		Count: papersRepository.Count(),
	})
}
