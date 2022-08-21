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

	router := gin.Default()

	// 配置中间件
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

	token := router.Group("/token")
	{
		token.POST("", controllers.CreateToken)
	}

	paper := router.Group("/papers")
	{
		paper.GET("/:id", controllers.GetPaper)
		paper.GET("/", middlewares.JwtAuth(middlewares.Role_All, nil), controllers.GetPapers)
		paper.POST("/", middlewares.JwtAuth(middlewares.Role_Admin, nil), controllers.GetPapers)
		paper.PATCH("/:id", middlewares.JwtAuth(middlewares.Role_Admin, nil), controllers.PatchPaper)
		paper.DELETE("/:id", middlewares.JwtAuth(middlewares.Role_Admin, nil), controllers.DeletePaper)
	}

	return router
}
