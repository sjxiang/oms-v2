package api

import (
	"context"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sjxiang/oms-v2/common/xlog"
)

// 日志拦截器
func InterceptorLogger() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context, 
		req any, 
		info *grpc.UnaryServerInfo, 
		handler grpc.UnaryHandler,
		) (resp any, err error) {
			
		startTime := time.Now()
		result, err := handler(ctx, req)
		duration := time.Since(startTime)
	
		statusCode := codes.Unknown
		if st, ok := status.FromError(err); ok {
			statusCode = st.Code()
		}
	
		if err != nil {
			xlog.Error("err", zap.Error(err))
		}
		
		xlog.Info("received a gRPC request", zapcore.Field{
			Key: "protocol",
			Type: zapcore.StringType,
			String: "grpc",
		}, zapcore.Field{
			Key: "method",
			Type: zapcore.StringType,
			String: info.FullMethod,
		}, zapcore.Field{
			Key: "status_code",
			Type: zapcore.Int64Type,
			Integer: int64(statusCode),
		}, zapcore.Field{
			Key: "status_text",
			Type: zapcore.StringType,
			String: statusCode.String(),
		}, zapcore.Field{
			Key: "duration",
			Type: zapcore.DurationType,
			Integer: int64(duration),
		})
		
		return result, err
	}
}
