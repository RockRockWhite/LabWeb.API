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

	blog := router.Group("/article")
	{
		blog.GET("/:id", controllers.GetArticle)
		blog.GET("/", controllers.GetArticles)
		blog.POST("/", middlewares.JwtAuth(), controllers.AddArticle)
		blog.PUT("/:id", middlewares.JwtAuth(), controllers.PutArticle)
		blog.PATCH("/:id", middlewares.JwtAuth(), controllers.PatchArticle)
		blog.DELETE("/:id", middlewares.JwtAuth(), controllers.DeleteArticle)
	}

	user := router.Group("/user")
	{
		user.GET("/:id", controllers.GetUser)
		user.POST("/", controllers.AddUser)
		user.PUT("/:id", middlewares.JwtAuth(), controllers.PutUser)
		user.PATCH("/:id", middlewares.JwtAuth(), controllers.PatchUser)
		user.DELETE("/:id", middlewares.JwtAuth(), controllers.DeleteUser)
	}

	token := router.Group("/token")
	{
		token.POST("", controllers.CreateToken)
	}

	return router
}
