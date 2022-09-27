package services

import (
	"fmt"
	"github.com/RockRockWhite/LabWeb.API/config"
	"github.com/RockRockWhite/LabWeb.API/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var _db *gorm.DB

func init() {
	host := config.GetString("DataBase.Host")
	port := config.GetString("DataBase.Port")
	username := config.GetString("DataBase.Username")
	password := config.GetString("DataBase.Password")
	dbname := config.GetString("DataBase.DBName")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, dbname)
	var err error
	_db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		utils.GetLogger().Fatal("Fatal error open database:%s %s \n", dsn, err)
	} else {
		utils.GetLogger().Printf("Opened database:%s %s \n", dsn, err)
	}
}

func getDB() *gorm.DB {
	return _db
}
