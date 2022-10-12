package routers

import (
	"fmt"
	"github.com/RockRockWhite/LabWeb.API/controllers"
	"github.com/RockRockWhite/LabWeb.API/dtos"
	"github.com/RockRockWhite/LabWeb.API/middlewares"
	"github.com/RockRockWhite/LabWeb.API/services"
	"github.com/RockRockWhite/LabWeb.API/utils"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"strconv"
)

func InitApiRouter() *gin.Engine {
	// 初始化Controllers
	router := gin.Default()

	// 配置中间件
	router.Use(middlewares.Cors, middlewares.Logger)

	user := router.Group("/users")
	{
		user.GET("/:username", controllers.GetUser)
		user.POST("", controllers.AddUser)
		user.PATCH(
			"/:username",
			middlewares.JwtAuth(middlewares.Role_Admin|middlewares.Role_Cond, func(c *gin.Context) bool {
				username := c.Param("username")
				claims := c.MustGet("claims").(*utils.JwtClaims)
				return username == claims.Username
			}), controllers.PatchUser)
		user.DELETE("/:username",
			middlewares.JwtAuth(middlewares.Role_Admin|middlewares.Role_Cond, func(c *gin.Context) bool {
				username := c.Param("username")
				claims := c.MustGet("claims").(*utils.JwtClaims)
				return username == claims.Username
			}), controllers.DeleteUser)
	}

	token := router.Group("/tokens")
	{
		token.POST("", controllers.CreateToken)
	}

	paper := router.Group("/papers")
	{
		paper.GET("/:id", controllers.GetPaper)
		paper.GET("/count", controllers.CountPaper)
		paper.GET("", middlewares.JwtAuth(middlewares.Role_Admin, nil), controllers.GetPapers)
		paper.GET("/public", controllers.GetPapersPublic)
		paper.GET("/private", middlewares.JwtAuth(middlewares.Role_Admin, nil), controllers.GetPapersPrivate)
		paper.POST("", middlewares.JwtAuth(middlewares.Role_Admin, nil), controllers.AddPaper)
		paper.PATCH("/:id", middlewares.JwtAuth(middlewares.Role_Admin, nil), controllers.PatchPaper)
		paper.DELETE("/:id", middlewares.JwtAuth(middlewares.Role_Admin, nil), controllers.DeletePaper)
	}

	teacher := router.Group("/teachers")
	{
		teacher.GET("/:id", controllers.GetTeacher)
		teacher.GET("/count", controllers.CountTeachers)
		teacher.GET("", controllers.GetTeachers)
		teacher.POST("", middlewares.JwtAuth(middlewares.Role_Admin, nil), controllers.AddTeacher)
		teacher.PATCH("/:id", middlewares.JwtAuth(middlewares.Role_Admin, nil), controllers.PatchTeacher)
		teacher.DELETE("/:id", middlewares.JwtAuth(middlewares.Role_Admin, nil), controllers.DeleteTeacher)
	}

	news := router.Group("/news")
	{
		news.GET("/:id", controllers.GetNews)
		news.GET("/count", controllers.CountNews)
		news.GET("", controllers.GetNewsList)
		news.POST("", middlewares.JwtAuth(middlewares.Role_Admin, nil), controllers.AddNews)
		news.PATCH("/:id", middlewares.JwtAuth(middlewares.Role_Admin, nil), controllers.PatchNews)
		news.DELETE("/:id", middlewares.JwtAuth(middlewares.Role_Admin, nil), controllers.DeleteNews)
	}

	configs := router.Group("/configs")
	{
		configs.GET("/:key", controllers.GetConfig)
		configs.PUT("/:key", middlewares.JwtAuth(middlewares.Role_Admin, nil), controllers.PutConfig)
	}

	resources := router.Group("/resources")
	{
		resources.GET("/:id", controllers.GetResource)
		resources.GET("/count", controllers.CountResources)
		resources.GET("", controllers.GetResourcesList)
		resources.POST("", middlewares.JwtAuth(middlewares.Role_Admin, nil), controllers.AddResources)
		resources.PATCH("/:id", middlewares.JwtAuth(middlewares.Role_Admin, nil), controllers.PatchResources)
		resources.DELETE("/:id", middlewares.JwtAuth(middlewares.Role_Admin, nil), controllers.DeleteResources)
	}

	todos := router.Group("/todos")
	{
		todos.GET("/:id", middlewares.JwtAuth(middlewares.Role_Admin|middlewares.Role_Cond, func(c *gin.Context) bool {
			id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

			if !services.GetTodoRepository().TodoExists(uint(id)) {
				c.JSON(http.StatusNotFound, dtos.ErrorDto{
					Message:          fmt.Sprintf("Todo %v not found!", id),
					DocumentationUrl: viper.GetString("Document.Url"),
				})
				return false
			}
			entity, _ := services.GetTodoRepository().GetTodo(uint(id))
			claims := c.MustGet("claims").(*utils.JwtClaims)
			return entity.UserId == claims.Id
		}), controllers.GetTodo)
		todos.GET("/count", middlewares.JwtAuth(middlewares.Role_Admin, nil), controllers.CountTodos)
		todos.GET("", middlewares.JwtAuth(middlewares.Role_Admin, nil), controllers.GetTodos)
		todos.POST("", middlewares.JwtAuth(middlewares.Role_All, nil), controllers.AddTodos)
		todos.PATCH("/:id", middlewares.JwtAuth(middlewares.Role_Admin|middlewares.Role_Cond, func(c *gin.Context) bool {
			id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

			if !services.GetTodoRepository().TodoExists(uint(id)) {
				c.JSON(http.StatusNotFound, dtos.ErrorDto{
					Message:          fmt.Sprintf("Todo %v not found!", id),
					DocumentationUrl: viper.GetString("Document.Url"),
				})
				return false
			}
			entity, _ := services.GetTodoRepository().GetTodo(uint(id))
			claims := c.MustGet("claims").(*utils.JwtClaims)
			return entity.UserId == claims.Id
		}), controllers.PatchTodo)
		todos.DELETE("/:id", middlewares.JwtAuth(middlewares.Role_Admin|middlewares.Role_Cond, func(c *gin.Context) bool {
			id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

			if !services.GetTodoRepository().TodoExists(uint(id)) {
				c.JSON(http.StatusNotFound, dtos.ErrorDto{
					Message:          fmt.Sprintf("Todo %v not found!", id),
					DocumentationUrl: viper.GetString("Document.Url"),
				})
				return false
			}
			entity, _ := services.GetTodoRepository().GetTodo(uint(id))
			claims := c.MustGet("claims").(*utils.JwtClaims)
			return entity.UserId == claims.Id
		}), controllers.DeleteTodo)

		todos.GET("/self/count", middlewares.JwtAuth(middlewares.Role_All, nil), controllers.CountTodos)
		todos.GET("/self", middlewares.JwtAuth(middlewares.Role_All, nil), controllers.GetTodos)

	}

	return router
}
