package consul

import (
	"context"
	"fmt"
	"net"
	"strconv"

	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
)

func (r *Registry) Register(ctx context.Context, instanceID, serviceName, addr string) error {
	
	host, portStr, err := net.SplitHostPort(addr)
	if err != nil {
		return fmt.Errorf("invalid host:port %v", err)
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return fmt.Errorf("invalid port: %v", err)
	}

	agentServiceCheck := &api.AgentServiceCheck{
		CheckID:       instanceID,
		TLSSkipVerify: false,
		TTL:           "5s",                    // 健康检查间隔（默认 5s）
		Timeout:       "5s",                    // 检查超时时间（默认 5s）
		DeregisterCriticalServiceAfter: "10s",  // 服务不可用后注销时间（默认 10s）
	}
	
	return r.client.Agent().ServiceRegister(&api.AgentServiceRegistration{
		ID:       instanceID,
		Name:     serviceName,
		Address:  host,
		Port:     port,
		Check:    agentServiceCheck,
	})
}


func (r *Registry) Deregister(ctx context.Context, instanceID, serviceName string) error {
	
	r.log.Info("服务实例已从Consul注销", 
		zap.String("instance-id", instanceID),  
		zap.String("service", serviceName), 
	)

	return r.client.Agent().CheckDeregister(instanceID)
}

func (r *Registry) Discover(ctx context.Context, serviceName string) ([]string, error) {
	
	entries, _,  err := r.client.Health().Service(serviceName, "", true, nil)
	if err != nil {
		return nil, err
	}

	var endpoints []string
	for _, entry := range entries {
		endpoints = append(endpoints, fmt.Sprintf("%s:%d", entry.Service.Address, entry.Service.Port))
	}

	return endpoints, nil
}

func (r *Registry) HealthCheck(instanceID, serviceName string) error {
	return r.client.Agent().UpdateTTL(instanceID, "online", api.HealthPassing)
}



