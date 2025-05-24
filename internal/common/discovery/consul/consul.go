package consul


import (
	"fmt"
	"sync"

	"go.uber.org/zap"
	"github.com/hashicorp/consul/api"
)

var (
	once         sync.Once
	consulClient *Registry
	initErr      error
	initialAddr  string // 保存首次初始化的地址
)

// Registry 封装 Consul 客户端
type Registry struct {
	client *api.Client
	log    *zap.Logger
}

// New 创建或返回已初始化的Consul客户端单例
// 所有调用必须使用相同的consulHttpAddr，否则返回错误
func New(consulHttpAddr string) (*Registry, error) {
	once.Do(func() {
		cfg := api.DefaultConfig()
		cfg.Address = consulHttpAddr

		client, err := api.NewClient(cfg)
		if err != nil {
			initErr = fmt.Errorf("consul初始化失败: %w", err)
			return
		}

		consulClient = &Registry{client: client}
		initialAddr = consulHttpAddr // 记录首次成功初始化的地址
	})

	// 优先返回初始化错误
	if initErr != nil {
		return nil, initErr
	}

	// 检查后续调用地址是否一致
	if consulHttpAddr != initialAddr {
		return nil, fmt.Errorf("客户端已用地址[%s]初始化，拒绝使用新地址[%s]", initialAddr, consulHttpAddr)
	}

	return consulClient, nil
}