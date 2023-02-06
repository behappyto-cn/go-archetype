package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"go-archetype/api/demo"
	"go-archetype/common/constant"
	"go-archetype/common/router"
)

// 入口
func main() {
	// 地址注册
	engine, routerGroup := router.Register()

	// 示例
	demoApi.Demo(routerGroup)

	// 启动 Http 请求
	if err := engine.Run(fmt.Sprintf(":%d", constant.ServicePort)); err != nil {
		log.Fatalln("Start Server Fail!", err)
	}
}
