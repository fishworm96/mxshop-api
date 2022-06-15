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

	/* 
		1.S()可以获取一个全局的sugar，可以让我们自己设置一个全局的logger
		2.日志是分级别的，debug，info，warn，error，fetal
		3.S函数和L函数很有用，提供了一个全局的安全访问logger的途径
	*/

	zap.S().Debugf("启动服务器,端口: %d", port)
	if err := Router.Run(fmt.Sprintf(":%d", port)); err != nil {
		zap.S().Panic("启动失败", err.Error())
	}
}
