package main

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"github.com/sjxiang/oms-v2/common/api"
	"github.com/sjxiang/oms-v2/common/conf"
	"github.com/sjxiang/oms-v2/common/logger"
	"github.com/sjxiang/oms-v2/common/pb"
	"github.com/sjxiang/oms-v2/order/ports"
	"github.com/sjxiang/oms-v2/order/service"
)

func init() {
	if err := conf.NewViperConfig(); err != nil {
		// 初始化配置失败
		panic(fmt.Errorf("初始化配置失败: %w", err))
	}

	fmt.Println("config", viper.GetStringMap("order"))
}


func main() {
	serverName := viper.GetString("order.service-name")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger, err := logger.NewLogger()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	application := service.NewApplication(ctx, logger)

	// 启动 grpc 服务
	go api.RunGrpcServer(serverName, func(server *grpc.Server) {
		srv := ports.NewGrpcServer(application)
		pb.RegisterOrderServiceServer(server, srv)
	})

	// 启动 http 服务
	api.RunHTTPServer(serverName, func(router *gin.Engine) {	
		srv := ports.NewHTTPServer(application)	
		ports.RegisterHandlersWithOptions(router, srv, ports.GinServerOptions{
			BaseURL:     "/api/v1",
			Middlewares: nil,
		})
	})

}

