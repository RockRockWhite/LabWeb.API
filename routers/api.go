package routers

import (
	"github.com/RockRockWhite/LabWeb.API/controllers"
	"github.com/RockRockWhite/LabWeb.API/middlewares"
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
		blog.POST("/", middlewares.JwtAuth(false), controllers.AddArticle)
		blog.PUT("/:id", middlewares.JwtAuth(false), controllers.PutArticle)
		blog.PATCH("/:id", middlewares.JwtAuth(false), controllers.PatchArticle)
		blog.DELETE("/:id", middlewares.JwtAuth(false), controllers.DeleteArticle)
	}

	user := router.Group("/users")
	{
		user.GET("/:username", controllers.GetUser)
		user.POST("/", controllers.AddUser)
		user.PATCH("/:username", middlewares.JwtAuth(false), controllers.PatchUser)
		user.DELETE("/:username", middlewares.JwtAuth(false), controllers.DeleteUser)
	}

	token := router.Group("/token")
	{
		token.POST("", controllers.CreateToken)
	}

	paper := router.Group("/papers")
	{
		paper.GET("/:id", controllers.GetPaper)
		paper.GET("/", middlewares.JwtAuth(false), controllers.GetPapers)
		paper.POST("/", middlewares.JwtAuth(true), controllers.GetPapers)
		paper.PATCH("/:id", middlewares.JwtAuth(true), controllers.PatchPaper)
		paper.DELETE("/:id", middlewares.JwtAuth(true), controllers.DeletePaper)
	}

	return router
}
