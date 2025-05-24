package main

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/gin-gonic/gin"
	grpc_tags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sjxiang/oms-v2/common/config"
	"github.com/sjxiang/oms-v2/common/discovery"
	"github.com/sjxiang/oms-v2/common/logger"
	"github.com/sjxiang/oms-v2/common/pb"
	"github.com/sjxiang/oms-v2/order/app"
	"github.com/sjxiang/oms-v2/order/ports"
	"github.com/sjxiang/oms-v2/order/service"
)

func init() {
	if err := config.NewViperConfig(); err != nil {
		panic(fmt.Errorf("初始化配置失败: %w", err))
	}
}


func main() {
	
	// 初始化配置
	cfg := struct{
		Salt            string
		LogLevel        string
		HttpAddr        string
		GrpcAddr        string
		ConsulHttpAddr  string
		ServiceName     string
	}{
		Salt:            viper.GetString("order.salt"),
		LogLevel:        viper.GetString("order.log-level"),
		HttpAddr:        viper.GetString("order.http-addr"),
		GrpcAddr:        viper.GetString("order.grpc-addr"),
		ConsulHttpAddr:  viper.GetString("order.consul-http-addr"),
		ServiceName:     viper.GetString("order.service-name"),

	}

	// 初始化日志
	log, err := logger.New(cfg.LogLevel, cfg.ServiceName)
	if err != nil {
		panic(err)
	}
	defer log.Sync()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 初始化应用
	application, cleanUp := service.NewApplication(ctx, log)
	defer cleanUp()

	// 服务注册
	deregisterFn, err := discovery.RegisterToConsul(ctx, log, 
		discovery.WithServiceID(cfg.Salt), 
		discovery.WithServiceName(cfg.ServiceName),
		discovery.WithServiceAddr(cfg.GrpcAddr), 
		discovery.WithConsulHttpAddr(cfg.ConsulHttpAddr),
	)
	if err != nil {
		panic(err)
	}
	defer func() {
		deregisterFn()
	}()

	// 润 1
	go runGrpcServer(cfg.GrpcAddr, log, application)
	
	// 润 2
	runHttpServer(cfg.HttpAddr, log, application)
}


func runHttpServer(httpAddr string, log *zap.Logger, application app.Application) {
	// 填充业务逻辑
	server, err := ports.NewHTTPServer(application)
	if err != nil {
		log.Fatal("cannot create server", zap.Error(err))
	}

	// 创建 Gin
	router := gin.New()

	// 注册路由
	ports.RegisterHandlersWithOptions(router, server, ports.GinServerOptions{
		BaseURL:     "/api/v1",
		Middlewares: []gin.HandlerFunc{gin.Logger(), gin.Recovery()},
	})

	// 启动服务
	err = router.Run(httpAddr)
	if err!= nil {
		log.Fatal("cannot start server", zap.Error(err))
	}
}


func runGrpcServer(grpcAddr string, log *zap.Logger, application app.Application) {
	// 填充业务逻辑
	server, err := ports.NewGrpcServer(application)
	if err != nil {
		log.Fatal("cannot create server", zap.Error(err))
	}

	// 创建 gRPC 服务器
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpc_tags.UnaryServerInterceptor(grpc_tags.WithFieldExtractor(grpc_tags.CodeGenRequestFieldExtractor)),
			withLoggingUnaryServerInterceptor(log),
		),
	)

	// 注册服务
	pb.RegisterOrderServiceServer(grpcServer, server)

	// 监听端口
	ln, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatal("cannot create listener", zap.Error(err))
	}

	// 打日志
	log.Info("start gRPC server at", zap.String("addr", ln.Addr().String()))
	
	// 启动服务
	err = grpcServer.Serve(ln)
	if err != nil {
		log.Fatal("cannot start gRPC server", zap.Error(err))
	}
}


// 日志拦截器
func withLoggingUnaryServerInterceptor(log *zap.Logger) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context, 
		req any, 
		info *grpc.UnaryServerInfo, 
		handler grpc.UnaryHandler) (resp any, err error) {
		
	 
		start := time.Now()
		resp, err = handler(ctx, req)
		elapsed := time.Since(start)
		
		statusCode := codes.Unknown
		if st, ok := status.FromError(err); ok {
			statusCode = st.Code()
		}
	
		if err != nil {
			log.Error("处理 grpc 请求失败 _(:3 」∠ )_", zap.Error(err))
		}
		
		log.Info("接收了一个 grpc 请求 😎", 
			zap.String("protocol", "grpc"), 
			zap.String("method", info.FullMethod), 
			zap.Int64("status_code", int64(statusCode)), 
			zap.String("status_text", statusCode.String()), 
			zap.Duration("duration", elapsed),  // 耗时
		)

		return resp, err
	}
}
