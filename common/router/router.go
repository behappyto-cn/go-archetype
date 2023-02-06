package router

import (
	"github.com/gin-gonic/gin"
	"go-archetype/common/constant"
	"go-archetype/common/interceptor"
)

// Register 初始化
func Register() (*gin.Engine, *gin.RouterGroup) {
	// 网络请求
	router := gin.New()
	// 配置 context-path
	routerGroup := router.Group(constant.ContextPath)
	// 拦截器
	routerGroup.Use(interceptor.HttpInterceptor())

	return router, routerGroup
}
