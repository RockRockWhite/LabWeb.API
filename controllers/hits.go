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

func PutHit(c *gin.Context) {
	hits, err := utils.GetReidsClient().Incr(context.Background(), "hits").Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.ErrorDto{
			Message:          fmt.Sprintf("incr hits Error: %s", err.Error()),
			DocumentationUrl: config.GetString("Document.Url"),
		})
		return
	}

	c.JSON(http.StatusCreated, struct {
		Hits uint64
	}{
		Hits: uint64(hits),
	})
}
