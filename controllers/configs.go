package controllers

import (
	"context"
	"fmt"
	"github.com/RockRockWhite/LabWeb.API/config"
	"github.com/RockRockWhite/LabWeb.API/dtos"
	"github.com/RockRockWhite/LabWeb.API/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PutConfig(c *gin.Context) {
	key := c.Param("key")
	dto := struct {
		Value string
	}{}
	if err := c.ShouldBind(&dto); err != nil {
		c.JSON(http.StatusBadRequest, dtos.ErrorDto{
			Message:          "Bind Model Error",
			DocumentationUrl: config.GetString("Document.Url"),
		})
		// todo: log
		return
	}
	if err := utils.GetReidsClient().Set(context.Background(), key, dto.Value, 0).Err(); err != nil {
		c.JSON(http.StatusInternalServerError, dtos.ErrorDto{
			Message:          fmt.Sprintf("Set config %s Error: %s", key, err.Error()),
			DocumentationUrl: config.GetString("Document.Url"),
		})
		return
	}

	c.JSON(http.StatusCreated, struct {
		Key   string
		Value string
	}{
		Key:   key,
		Value: dto.Value,
	})

}

func GetConfig(c *gin.Context) {
	key := c.Param("key")

	value, err := utils.GetReidsClient().Get(context.Background(), key).Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.ErrorDto{
			Message:          fmt.Sprintf("get config %s Error: %s", key, err.Error()),
			DocumentationUrl: config.GetString("Document.Url"),
		})
		return
	}

	c.JSON(http.StatusCreated, struct {
		Key   string
		Value string
	}{
		Key:   key,
		Value: value,
	})
}
