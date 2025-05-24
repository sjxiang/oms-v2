package discovery

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"go.uber.org/zap"

	"github.com/sjxiang/oms-v2/common/discovery/consul"
)


func RegisterToConsul(ctx context.Context, log *zap.Logger, opts ...OptFunc) (func() error, error) {

	// 基础配置初始化
    o := defaultOpts()
    for _, fn := range opts {
        fn(&o)
    }

    // 必要参数校验
    if o.serviceName == "" || o.serviceID == "" || o.serviceAddr == "" {
        return nil, fmt.Errorf("缺少必要参数：serviceName, serviceID, serviceAddr")
    }

    // 初始化客户端
    registry, err := consul.New(o.consulHttpAddr)
    if err != nil {
        return nil, fmt.Errorf("consul客户端创建失败: %v", err)
    }

    // 注册服务
    if err := registry.Register(ctx, o.serviceID, o.serviceName, o.serviceAddr); err != nil {
        return nil, fmt.Errorf("服务注册失败: %v", err)
    }

	// 健康检查循环
	go func() {
		for {
			if err := registry.HealthCheck(o.serviceID, o.serviceName); err!= nil {
				log.Error("心跳包失败", zap.String("service", o.serviceName), zap.Error(err))
			}
			
			time.Sleep(1 * time.Second)
		}
	}()

	log.Info("服务注册成功",
	zap.String("service", o.serviceName),
	zap.String("address", o.serviceAddr))

	// 清理函数
	cleanup := func() error {
		return registry.Deregister(ctx, o.serviceID, o.serviceName)
	}

	return cleanup, nil
}


func GetServiceAddr(ctx context.Context, log *zap.Logger, opts ...OptFunc) (string, error) {
	// 基础配置初始化
	o := defaultOpts()
	for _, fn := range opts {
		fn(&o)
	}

	// 必要参数校验
	if o.serviceName == "" || o.serviceID == "" || o.serviceAddr == "" {
		return "", fmt.Errorf("缺少必要参数：serviceName, serviceID, serviceAddr")
	}

	// 初始化客户端
	registry, err := consul.New(o.consulHttpAddr)
	if err != nil {
		return "", fmt.Errorf("consul客户端创建失败: %v", err)
	}

	endpoints, err := registry.Discover(ctx, o.serviceName)
	if err != nil {
		return "", err
	}

	count := len(endpoints)
	if count == 0  {
		return "", fmt.Errorf("got empty %s addrs from consul", o.serviceName)
	}


	log.Info("检测到可用服务实例",
	zap.String("service", o.serviceName),
	zap.Int("instance_count", count),
	zap.Strings("endpoints", endpoints))

	/* 
	 	参数 n 是一个非负整数，表示生成的随机数的范围。
		如果 n=0，那么 rand.Intn(0) 会触发 panic，因为 0 不是一个有效的范围。
		如果 n>0，那么 rand.Intn(n) 会生成一个在 [0, n) 范围内的随机整数。
	 */

	idx := rand.Intn(count) 

	return endpoints[idx], nil
}