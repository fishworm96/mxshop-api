package initialize

import (
	"fmt"

	"mxshop-api/user-web/global"
	"mxshop-api/user-web/proto"

	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func InitSrvConn() {
	// 从注册中心获取到用户服务的信息
	cfg := api.DefaultConfig()
	consulInfo := global.ServerConfig.ConsulInfo
	cfg.Address = fmt.Sprintf("%s:%d", consulInfo.Host, consulInfo.Port)

	userSrvHost := ""
	userSrvPort := 0
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	// data, err := client.Agent().ServicesWithFilter(fmt.Sprintf("Service == \"%s\"", global.ServerConfig.UserSrvInfo.Name))
	data, err := client.Agent().ServicesWithFilter(fmt.Sprintf(`Service == "%s"`, global.ServerConfig.UserSrvInfo.Name))
	if err != nil {
		panic(err)
	}
	for _, value := range data {
		userSrvHost = value.Address
		userSrvPort = value.Port
		break
	}
	if userSrvHost == "" {
		zap.S().Fatal("[InitSrvConn] 连接 [用户服务失败]")
		return
	}

	// 拨号连接用户grpc服务器
	userConn, err := grpc.Dial(fmt.Sprintf("%s:%d", userSrvHost, userSrvPort), grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("[GetUserList] 连接 [用户服务失败]", "msg", err.Error())
	}
	// 生成grcp的client并调用接口
	userSrvClient := proto.NewUserClient(userConn)
	global.UserSrvClient = userSrvClient
}
