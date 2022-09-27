package main

import (
	"github.com/RockRockWhite/LabWeb.API/config"
	"github.com/RockRockWhite/LabWeb.API/routers"
	"github.com/RockRockWhite/LabWeb.API/utils"
)

func main() {
	// 初始化并运行路由
	router := routers.InitApiRouter()
	err := router.Run(config.GetString("HttpServer.Port"))

	if err != nil {
		utils.GetLogger().Fatal(err.Error())
	}
}
