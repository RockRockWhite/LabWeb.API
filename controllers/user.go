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

var userRepository *services.UserRepository

// InitUserController 初始化用户Controller
func InitUserController() {
	userRepository = services.NewUserRepository(true)
}

// AddUser 添加用户
func AddUser(c *gin.Context) {
	var userDto dtos.UserAddDto

	if err := c.ShouldBind(&userDto); err != nil {
		c.JSON(http.StatusBadRequest, dtos.ErrorDto{
			Message:          "Bind Model Error",
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	// 保证用户名不重复
	if userRepository.UsernameExists(userDto.Username) {
		c.JSON(http.StatusBadRequest, dtos.ErrorDto{
			Message:          fmt.Sprintf("Username %v exists", userDto.Username),
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	entity := userDto.ToEntity()
	userRepository.AddUser(entity)

	c.JSON(http.StatusCreated, dtos.ParseUserEntity(entity))
}

// GetUser 获得用户
func GetUser(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	if !userRepository.UserExists(uint(id)) {
		c.JSON(http.StatusNotFound, dtos.ErrorDto{
			Message:          fmt.Sprintf("User id %v not found!", id),
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	user := userRepository.GetUser(uint(id))
	// 转换为Dto
	c.JSON(http.StatusOK, dtos.ParseUserEntity(user))
}

// PatchUser 修改用户
func PatchUser(c *gin.Context) {
	// 获得更新id
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if !userRepository.UserExists(uint(id)) {
		c.JSON(http.StatusNotFound, dtos.ErrorDto{
			Message:          fmt.Sprintf("User id %v not found!", id),
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}
	user := userRepository.GetUser(uint(id))

	// 获得用户信息 判断用户是否对该博文具有修改权
	// 修改权: 改博文为用户所有 或 该用户是管理员
	claims := c.MustGet("claims").(*utils.JwtClaims)

	if user.ID != claims.Id && !claims.IsAdmin {
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
	dto := dtos.UserUpdateDtoFromEntity(user)
	utils.ApplyJsonPatch(dto, patchJson)
	dto.ApplyUpdateToEntity(user)

	// 保证用户名不重复
	if userRepository.UsernameExists(user.Username) {
		c.JSON(http.StatusBadRequest, dtos.ErrorDto{
			Message:          fmt.Sprintf("Username %v exists", user.Username),
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	// 更新数据库
	userRepository.UpdateUser(user)

	c.Status(http.StatusNoContent)
}

// PutUser 替换用户
func PutUser(c *gin.Context) {
	var updateDto dtos.UserUpdateDto

	if err := c.ShouldBind(&updateDto); err != nil {
		c.JSON(http.StatusBadRequest, dtos.ErrorDto{
			Message:          "Bind Model Error",
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	// 保证用户名不重复
	if userRepository.UsernameExists(updateDto.Username) {
		c.JSON(http.StatusBadRequest, dtos.ErrorDto{
			Message:          fmt.Sprintf("Username %v exists", updateDto.Username),
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	// 获得更新id
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	// 获得用户信息
	claims := c.MustGet("claims").(*utils.JwtClaims)

	// 处理用户不存在
	if !userRepository.UserExists(uint(id)) {
		entity := entities.User{}
		entity.ID = uint(id)
		updateDto.ApplyUpdateToEntity(&entity)
		userRepository.AddUser(&entity)

		c.JSON(http.StatusCreated, dtos.ParseUserEntity(&entity))
		return
	}

	// 处理用户存在
	entity := userRepository.GetUser(uint(id))
	// 判断是否有修改权
	if entity.ID != claims.Id && !claims.IsAdmin {
		c.JSON(http.StatusForbidden, dtos.ErrorDto{
			Message:          "Permission denied for changing this resource!",
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	updateDto.ApplyUpdateToEntity(entity)
	userRepository.UpdateUser(entity)

	c.Status(http.StatusNoContent)
}

// DeleteUser 删除用户
func DeleteUser(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	if !userRepository.UserExists(uint(id)) {
		c.JSON(http.StatusNotFound, dtos.ErrorDto{
			Message:          fmt.Sprintf("User id %v not found!", id),
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	user := userRepository.GetUser(uint(id))
	// 获得用户信息 判断用户是否对该博文具有修改权
	// 修改权: 改博文为用户所有 或 该用户是管理员
	claims := c.MustGet("claims").(*utils.JwtClaims)

	if user.ID != claims.Id && !claims.IsAdmin {
		c.JSON(http.StatusForbidden, dtos.ErrorDto{
			Message:          "Permission denied for changing this resource!",
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	userRepository.DeleteUser(uint(id))
	c.Status(http.StatusNoContent)
}
