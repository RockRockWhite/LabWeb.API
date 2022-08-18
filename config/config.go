package config

import (
	"fmt"
	"github.com/spf13/viper"
)

// 加载配置文件

func Init(fileName string) {
	viper.SetConfigFile(fileName)

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
