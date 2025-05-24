package discovery

import "context"


type Registry interface {
	// 注册
	Register(ctx context.Context, instanceID, serviceName, grpcAddr string) error
	// 注销
	Deregister(ctx context.Context, instanceID, serviceName string) error
	// 发现
	Discover(ctx context.Context, serviceName string) ([]string, error)
	// 健康检查
	HealthCheck(instanceID, serviceName string) error
}
