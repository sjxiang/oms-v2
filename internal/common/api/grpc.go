package api

import (
	"net"

	"github.com/sjxiang/oms-v2/common/xlog"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	grpc_tags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
)

// 启动 grpc 服务
func RunGrpcServer(serviceName string, registerServer func(server *grpc.Server)) {
	addr := viper.Sub(serviceName).GetString("grpc-addr")
	if addr == "" {
		// TODO: Warning log
		addr = viper.GetString("fallback-grpc-addr")
	}
	RunGrpcServerOnAddr(addr, registerServer)
}



func RunGrpcServerOnAddr(addr string, registerServer func(server *grpc.Server)) {
	
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpc_tags.UnaryServerInterceptor(grpc_tags.WithFieldExtractor(grpc_tags.CodeGenRequestFieldExtractor)),
			// 日志拦截器
			InterceptorLogger(),
		),
	)

	// 注册服务
	registerServer(grpcServer)

	// 监听端口
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		xlog.Fatal("cann't create listener", zap.Error(err))
	}

	xlog.Info("start gRPC server at", zap.String("addr", ln.Addr().String()))
	
	// 启动服务
	if err := grpcServer.Serve(ln); err != nil {
		xlog.Fatal("gRPC server failed to serve", zap.Error(err))
	}
}