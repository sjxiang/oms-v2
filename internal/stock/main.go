package main

import (
	"context"
	"fmt"

	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"github.com/sjxiang/oms-v2/common/api"
	"github.com/sjxiang/oms-v2/common/conf"
	"github.com/sjxiang/oms-v2/common/pb"
	"github.com/sjxiang/oms-v2/stock/ports"
	"github.com/sjxiang/oms-v2/stock/service"
)

func init() {
	if err := conf.NewViperConfig(); err != nil {
		panic(fmt.Errorf("初始化配置失败: %w", err))
	}

	fmt.Println("config", viper.GetStringMap("stock"))
}


func main() {
	serviceName := viper.GetString("stock.service-name")
	serviceType := viper.GetString("stock.server-to-run")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	application := service.NewApplication(ctx)

	switch serviceType {
	case "grpc":
		// 启动 grpc 服务
		api.RunGrpcServer(serviceName , func(server *grpc.Server) {
			svc := ports.NewGrpcServer(application)
			pb.RegisterStockServiceServer(server, svc)
		})
	case "http":
		panic("unimplemented")
	default:
		panic("unexpected server type")
	}
}
