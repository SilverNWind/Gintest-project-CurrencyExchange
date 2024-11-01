package router

import (
	"exchangeapp/controllers"
	middlewares "exchangeapp/middleswares"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	//配置cors中间件
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	// 路由组auth
	auth := r.Group("/api/auth")
	{
		// auth.POST("/login", func(ctx *gin.Context) {
		// 	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
		// 		"msg": "Login Success",
		// 	})
		// })
		// auth.POST("/register", func(ctx *gin.Context) {
		// 	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
		// 		"msg": "Register Success",
		// 	})
		// })

		auth.POST("/login", controllers.Login)
		auth.POST("/register", controllers.Register)
	}

	api := r.Group("/api")
	api.GET("/exchangeRates", controllers.GetExchangeRates)
	api.Use(middlewares.AuthMiddleWare())
	{
		api.POST("/exchangeRates", controllers.CreateExchangeRates)

		api.POST("/articles", controllers.CreateArticles)
		api.GET("/articles", controllers.GetArticles)
		api.GET("/articles/:id", controllers.GetArticlesByID)

		api.POST("/articles/:id/like", controllers.LikeArticle)
		api.GET("/articles/:id/like", controllers.GetArticleLikes)
	}
	// api.POST("/exchangeRates", middlewares.AuthMiddleWare(), controllers.CreateExchangeRates)

	return r

}
