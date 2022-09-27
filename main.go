package main

import (
	"github.com/RockRockWhite/LabWeb.API/config"
	"github.com/RockRockWhite/LabWeb.API/routers"
)

func main() {
	// 初始化并运行路由
	router := routers.InitApiRouter()

	_ = router.Run(config.GetString("HttpServer.Port"))
}
