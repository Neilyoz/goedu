package router

import (
	"github.com/gin-gonic/gin"
	"goedu/controller"
	"goedu/middleware"
)

func Router(engine *gin.Engine) *gin.Engine {
	apiRouter := engine.Group("api")
	{
		apiRouter.GET("/", controller.Index)

		apiRouter.POST("/register", controller.Register)   // 用户注册
		apiRouter.POST("/login", controller.Login)         // 用户登录
		apiRouter.GET("/article", controller.ArticleIndex) // 文章列表

		// 加载权限中间件
		apiRouter.Use(middleware.Auth())
		apiRouter.POST("/article/create", controller.ArticleCreate)
		apiRouter.POST("/article/edit", controller.ArticleEdit)
		apiRouter.GET("/article/:id", controller.ArticleDetail)
	}

	return engine
}
