package utils

import (
	"github.com/RockRockWhite/LabWeb.API/config"
	"github.com/go-redis/redis/v9"
)

var _rdb *redis.Client

func init() {
	_rdb = redis.NewClient(&redis.Options{
		Addr:     config.GetString("Redis.Addr"),
		Password: config.GetString("Redis.Password"),
		DB:       config.GetInt("Redis.DB"),
	})
}

func GetReidsClient() *redis.Client {
	return _rdb
}
