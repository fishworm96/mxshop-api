package main

import (
	"fmt"
	"mxshop-api/user-web/global"
	"mxshop-api/user-web/initialize"

	"go.uber.org/zap"
)

func main() {
	port := 8021

	// 初始化全局logger
	global.InitLogger()
	// 初始化routers
	Router := initialize.Routers()

	zap.S().Debugf("启动服务器,端口: %d", port)
	if err := Router.Run(fmt.Sprintf(":%d", port)); err != nil {
		zap.S().Panic("启动失败", err.Error())
	}
}
