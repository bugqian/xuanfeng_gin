package intranet

import (
	"github.com/gin-gonic/gin"
	"xuanfeng_gin/internal/api/middleware"
	ext "xuanfeng_gin/pkg/gin-ext"
)

func LoadRoute(router *gin.Engine) {

	router.Use(middleware.GinLogger(), middleware.GinRecovery(true))
	router.Use(middleware.Cors())
	api := router.Group("/api/app")
	user := api.Group("/user")
	{
		user.POST("/create", ext.H(userCreate)) // 新增用户
	}

}
