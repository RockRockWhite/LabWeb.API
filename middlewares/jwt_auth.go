package middlewares

import (
	"github.com/RockRockWhite/LabWeb.API/dtos"
	"github.com/RockRockWhite/LabWeb.API/utils"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"strings"
)

// JwtAuth JwtToken验证中间件
func JwtAuth(roleFlag Role, condFunc func(c *gin.Context) bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" || strings.Fields(token)[0] != "Bearer" {
			// 没有传token参数
			c.JSON(http.StatusUnauthorized, dtos.ErrorDto{
				Message:          "Requires bearer token in filed {Authorization}.",
				DocumentationUrl: viper.GetString("Document.Url"),
			})

			c.Abort()
			return
		}
		token = strings.Fields(token)[1]

		claims, err := utils.ParseJwtToken(token)

		if err != nil {
			c.JSON(http.StatusUnauthorized, dtos.ErrorDto{
				Message:          "Token expired or the other error occurred",
				DocumentationUrl: viper.GetString("Document.Url"),
			})

			c.Abort()
			return
		}

		// Claims写入上下文
		c.Set("claims", claims)

		access := false
		if roleFlag&Role_All != 0 { // 所有人通行
			access = true
		} else if roleFlag&Role_Admin != 0 && claims.IsAdmin { // 管理员通行
			access = true
		} else if roleFlag&Role_Cond != 0 && condFunc != nil && condFunc(c) { // 符合条件用户通行
			access = true
		}

		if !access {
			c.JSON(http.StatusForbidden, dtos.ErrorDto{
				Message:          "The token cannot access this resource.",
				DocumentationUrl: viper.GetString("Document.Url"),
			})

			c.Abort()
			return
		}
	}
}

type Role int

const (
	Role_Admin Role = 1 // 管理员
	Role_Cond  Role = 2 // 满足条件的用户
	Role_All   Role = 4 // 所有用户
)
