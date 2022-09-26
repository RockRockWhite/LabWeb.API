package main

import (
	"github.com/RockRockWhite/LabWeb.API/config"
	"github.com/RockRockWhite/LabWeb.API/routers"
	"github.com/RockRockWhite/LabWeb.API/utils"
	"github.com/sirupsen/logrus"
)

func main() {
	// 初始化logger
	utils.InitLogger(config.GetString("Logger.LogFile"), logrus.DebugLevel, "2006-01-02 15:04:05")
	utils.Logger().Infof("| [service] | ***** Service started ***** |")
	defer utils.Logger().Infof("| [service] | ***** Service stoped ***** |")

	// 初始化并运行路由
	router := routers.InitApiRouter()

	_ = router.Run(config.GetString("HttpServer.Port"))
}
