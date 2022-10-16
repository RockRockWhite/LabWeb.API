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

var todoRepository *services.TodosRepository

func init() {
	todoRepository = services.GetTodoRepository()
}

// AddTodos 添加Todo项
func AddTodos(c *gin.Context) {
	var addDto dtos.TodoAddDto

	if err := c.ShouldBind(&addDto); err != nil {
		c.JSON(http.StatusBadRequest, dtos.ErrorDto{
			Message:          "Bind Model Error",
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	// 获得用户信息
	claims := c.MustGet("claims").(*utils.JwtClaims)

	var entity entities.Todo
	dtos.GetTodoAddDtoConverter().Convert(&entity, &addDto)
	entity.UserId = claims.Id

	_, err := todoRepository.AddTodo(&entity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.ErrorDto{
			Message:          err.Error(),
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	var getDto dtos.TodoGetDto
	err = dtos.GetTodoGetDtoConverter().Convert(&getDto, &entity)
	c.JSON(http.StatusCreated, getDto)
}

// GetTodo 获得todo
func GetTodo(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	if !todoRepository.TodoExists(uint(id)) {
		c.JSON(http.StatusNotFound, dtos.ErrorDto{
			Message:          fmt.Sprintf("Todo %v not found!", id),
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	entity, err := todoRepository.GetTodo(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.ErrorDto{
			Message:          err.Error(),
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	// 转换为Dto
	var getDto dtos.TodoGetDto
	dtos.GetTodoGetDtoConverter().Convert(&getDto, entity)
	c.JSON(http.StatusOK, getDto)
}

// GetTodos 批量获得todo列表
func GetTodos(c *gin.Context) {
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
	filter := c.DefaultQuery("filter", "")

	entities, err := todoRepository.GetTodosList(limit, (page-1)*limit, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.ErrorDto{
			Message:          err.Error(),
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}
	// 转换为Dto
	getDtos := make([]dtos.TodoGetDto, 0, len(entities))
	for _, entity := range entities {
		var each dtos.TodoGetDto
		dtos.GetTodoGetDtoConverter().Convert(&each, &entity)

		user, err := usersRepository.GetUserById(entity.UserId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, dtos.ErrorDto{
				Message:          err.Error(),
				DocumentationUrl: viper.GetString("Document.Url"),
			})
			return
		}

		each.Username = user.Username
		each.Fullname = user.Fullname
		getDtos = append(getDtos, each)
	}

	c.JSON(http.StatusOK, getDtos)
}

// PatchTodo 修改todo
func PatchTodo(c *gin.Context) {
	// 获得更新id
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if !todoRepository.TodoExists(uint(id)) {
		c.JSON(http.StatusNotFound, dtos.ErrorDto{
			Message:          fmt.Sprintf("Todo %v not found!", id),
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	entity, err := todoRepository.GetTodo(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.ErrorDto{
			Message:          err.Error(),
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
	var updateDto dtos.TodoUpdateDto
	dtos.GetTodoUpdateDtoEntityConverter().Convert(&updateDto, entity)
	utils.ApplyJsonPatch(&updateDto, patchJson)
	dtos.GetTodoEntityUpdateDtoConverter().Convert(entity, &updateDto)

	// 更新数据库
	err = todoRepository.UpdateTodo(entity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.ErrorDto{
			Message:          err.Error(),
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	c.Status(http.StatusNoContent)
}

// DeleteTodo 删除todo
func DeleteTodo(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	if !todoRepository.TodoExists(uint(id)) {
		c.JSON(http.StatusNotFound, dtos.ErrorDto{
			Message:          fmt.Sprintf("Todo %v not found!", id),
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	err := todoRepository.DeleteTodo(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.ErrorDto{
			Message:          err.Error(),
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}
	c.Status(http.StatusNoContent)
}

// CountTodos 获得todo数量
func CountTodos(c *gin.Context) {
	filter := c.DefaultQuery("filter", "")
	c.JSON(http.StatusOK, struct {
		Count int64
	}{
		Count: todoRepository.Count(filter),
	})
}

// CountTodosSelf 获得自己创建的todo数量
func CountTodosSelf(c *gin.Context) {
	// 获得用户信息
	claims := c.MustGet("claims").(*utils.JwtClaims)

	c.JSON(http.StatusOK, struct {
		Count int64
	}{
		Count: todoRepository.Count(strconv.Itoa(int(claims.Id))),
	})
}

// GetTodosSelf 批量获得自己创建的todo列表
func GetTodosSelf(c *gin.Context) {
	// 获得page limit
	page, pageQueryErr := strconv.Atoi(c.DefaultQuery("page", "1"))
	if pageQueryErr != nil {
		c.JSON(http.StatusBadRequest, dtos.ErrorDto{
			Message:          "Incorrect query field page",
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}
	// 获得用户信息
	claims := c.MustGet("claims").(*utils.JwtClaims)

	limit, limitQueryErr := strconv.Atoi(c.DefaultQuery("limit", "20"))
	if limitQueryErr != nil {
		c.JSON(http.StatusBadRequest, dtos.ErrorDto{
			Message:          "Incorrect query field limit.",
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	entities, err := todoRepository.GetTodosList(limit, (page-1)*limit, strconv.Itoa(int(claims.Id)))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.ErrorDto{
			Message:          err.Error(),
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}
	// 转换为Dto
	getDtos := make([]dtos.TodoGetDto, 0, len(entities))
	for _, entity := range entities {
		var each dtos.TodoGetDto
		dtos.GetTodoGetDtoConverter().Convert(&each, &entity)
		getDtos = append(getDtos, each)
	}

	c.JSON(http.StatusOK, getDtos)
}
