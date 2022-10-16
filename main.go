package main

import (
	"github.com/RockRockWhite/LabWeb.API/config"
	"github.com/RockRockWhite/LabWeb.API/dtos"
	"github.com/RockRockWhite/LabWeb.API/routers"
	"github.com/RockRockWhite/LabWeb.API/services"
	"github.com/RockRockWhite/LabWeb.API/utils"
)

func main() {
	// 初始化并运行路由
	router := routers.InitApiRouter()
	// 创建管理员用户
	InitAdmin()
	// 运行
	err := router.Run(config.GetString("HttpServer.Port"))

	if err != nil {
		utils.GetLogger().Fatal(err.Error())
	}
}

func InitAdmin() {
	username := config.GetString("Admin.Username")
	password := config.GetString("Admin.Password")

	if !services.GetUsersRepository().UsernameExists(username) {
		userDto := dtos.UserAddDto{
			Username: username,
			Password: password,
		}
		entity := userDto.ToEntity()
		entity.IsAdmin = true

		_, err := services.GetUsersRepository().AddUser(entity)
		if err != nil {
			utils.GetLogger().Fatal(err.Error())
		}
	}
}
