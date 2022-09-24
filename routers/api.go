package routers

import (
	"github.com/RockRockWhite/LabWeb.API/controllers"
	"github.com/RockRockWhite/LabWeb.API/middlewares"
	"github.com/RockRockWhite/LabWeb.API/utils"
	"github.com/gin-gonic/gin"
)

func InitApiRouter() *gin.Engine {
	// 初始化Controllers
	controllers.InitArticleController()
	controllers.InitUserController()
	controllers.InitPaperController()
	controllers.InitTeacherController()
	controllers.InitNewsController()

	router := gin.Default()

	// 配置中间件
	router.Use(middlewares.Cors)
	router.Use(middlewares.Logger())

	blog := router.Group("/articles")
	{
		blog.GET("/:id", controllers.GetArticle)
		blog.GET("/", controllers.GetArticles)
		//blog.POST("/", middlewares.JwtAuth(), controllers.AddArticle)
		//blog.PUT("/:id", middlewares.JwtAuth(), controllers.PutArticle)
		//blog.PATCH("/:id", middlewares.JwtAuth(), controllers.PatchArticle)
		//blog.DELETE("/:id", middlewares.JwtAuth(), controllers.DeleteArticle)
	}

	user := router.Group("/users")
	{
		user.GET("/:username", controllers.GetUser)
		user.POST("/", controllers.AddUser)
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
		paper.GET("/", middlewares.JwtAuth(middlewares.Role_Admin, nil), controllers.GetPapers)
		paper.POST("/", middlewares.JwtAuth(middlewares.Role_Admin, nil), controllers.AddPaper)
		paper.PATCH("/:id", middlewares.JwtAuth(middlewares.Role_Admin, nil), controllers.PatchPaper)
		paper.DELETE("/:id", middlewares.JwtAuth(middlewares.Role_Admin, nil), controllers.DeletePaper)
	}

	teacher := router.Group("/teachers")
	{
		teacher.GET("/:id", controllers.GetTeacher)
		teacher.GET("/", middlewares.JwtAuth(middlewares.Role_Admin, nil), controllers.GetTeachers)
		teacher.POST("/", middlewares.JwtAuth(middlewares.Role_Admin, nil), controllers.AddTeacher)
		teacher.PATCH("/:id", middlewares.JwtAuth(middlewares.Role_Admin, nil), controllers.PatchTeacher)
		teacher.DELETE("/:id", middlewares.JwtAuth(middlewares.Role_Admin, nil), controllers.DeleteTeacher)
	}

	news := router.Group("/news")
	{
		news.GET("/:id", controllers.GetNews)
		news.GET("/", middlewares.JwtAuth(middlewares.Role_Admin, nil), controllers.GetNewsList)
		news.POST("/", middlewares.JwtAuth(middlewares.Role_Admin, nil), controllers.AddNews)
		news.PATCH("/:id", middlewares.JwtAuth(middlewares.Role_Admin, nil), controllers.PatchNews)
		news.DELETE("/:id", middlewares.JwtAuth(middlewares.Role_Admin, nil), controllers.DeleteNews)
	}

	return router
}
