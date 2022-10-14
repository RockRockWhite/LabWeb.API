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

var usersRepository *services.UsersRepository

func init() {
	usersRepository = services.GetUsersRepository()
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
	if usersRepository.UsernameExists(userDto.Username) {
		c.JSON(http.StatusBadRequest, dtos.ErrorDto{
			Message:          fmt.Sprintf("Username %v exists", userDto.Username),
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	entity := userDto.ToEntity()
	_, err := usersRepository.AddUser(entity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, dtos.ParseUserEntity(entity))
}

// GetUser 获得用户
func GetUser(c *gin.Context) {
	username := c.Param("username")
	if !usersRepository.UsernameExists(username) {
		c.JSON(http.StatusNotFound, dtos.ErrorDto{
			Message:          fmt.Sprintf("User %v not found.", username),
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	user, err := usersRepository.GetUserByName(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	// 转换为Dto
	c.JSON(http.StatusOK, dtos.ParseUserEntity(user))
}

// GetUsers 批量获得用户
func GetUsers(c *gin.Context) {
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
	filter := services.UserFilter(c.DefaultQuery("filter", "all"))

	if limitQueryErr != nil {
		c.JSON(http.StatusBadRequest, dtos.ErrorDto{
			Message:          "Incorrect query field limit.",
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	entities, err := usersRepository.GetUsers(limit, (page-1)*limit, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.ErrorDto{
			Message:          err.Error(),
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}
	// 转换为Dto
	getDtos := make([]dtos.UserGetDto, 0, len(entities))
	for _, entity := range entities {
		getDtos = append(getDtos, *dtos.ParseUserEntity(&entity))
	}

	c.JSON(http.StatusOK, getDtos)
}

// PatchUser 修改用户
func PatchUser(c *gin.Context) {
	// 获得更新id
	username := c.Param("username")
	if !usersRepository.UsernameExists(username) {
		c.JSON(http.StatusNotFound, dtos.ErrorDto{
			Message:          fmt.Sprintf("User %v not found!", username),
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	user, err := usersRepository.GetUserByName(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
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
	if usersRepository.UsernameExists(user.Username) {
		c.JSON(http.StatusBadRequest, dtos.ErrorDto{
			Message:          fmt.Sprintf("Username %v exists", user.Username),
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	// 更新数据库
	err = usersRepository.UpdateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusNoContent)
}

// PatchPassword 修改用户密码
func PatchPassword(c *gin.Context) {
	// 获得更新id
	username := c.Param("username")
	if !usersRepository.UsernameExists(username) {
		c.JSON(http.StatusNotFound, dtos.ErrorDto{
			Message:          fmt.Sprintf("User %v not found!", username),
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	var passwordDto dtos.UserPasswordDto

	if err := c.ShouldBind(&passwordDto); err != nil {
		c.JSON(http.StatusBadRequest, dtos.ErrorDto{
			Message:          "Bind Model Error",
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	user, err := usersRepository.GetUserByName(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	// 验证密码
	if !utils.ValifyPasswordHash(passwordDto.OldPassword, user.Salt, user.PasswordHash) {
		c.JSON(http.StatusForbidden, dtos.ErrorDto{
			Message:          "Password Error",
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	// 修改密码
	user.PasswordHash = utils.EncryptPasswordHash(passwordDto.Password, user.Salt)

	// 更新数据库
	err = usersRepository.UpdateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusNoContent)
}

// DeleteUser 删除用户
func DeleteUser(c *gin.Context) {
	username := c.Param("username")
	if !usersRepository.UsernameExists(username) {
		c.JSON(http.StatusNotFound, dtos.ErrorDto{
			Message:          fmt.Sprintf("User %v not found!", username),
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	err := usersRepository.DeleteUserByName(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.Status(http.StatusNoContent)
}

// CountUser 获得用户数量
func CountUser(c *gin.Context) {
	filter := services.UserFilter(c.DefaultQuery("filter", "all"))

	c.JSON(http.StatusOK, struct {
		Count int64
	}{
		Count: usersRepository.Count(filter),
	})
}
