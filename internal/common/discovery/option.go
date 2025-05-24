package discovery


import (
	"fmt"
	"math/rand"
	"time"
)


type Opts struct {
	// Consul 连接配置
	consulHttpAddr string   // Consul 服务地址 (e.g. "localhost:8500")

	// 服务注册配置
	serviceName    string   // 服务名 (必填)
	serviceID      string   // 服务唯一 ID
	serviceAddr    string   // 服务监听地址 (e.g. ":8080")	
}


type OptFunc func(*Opts) 

func defaultOpts() Opts {
	return Opts{
		consulHttpAddr: "localhost:8500",
	}
}


func WithConsulHttpAddr(consulHttpAddr string) OptFunc {
	return func(opts *Opts) {
		opts.consulHttpAddr = consulHttpAddr
	}
}

func WithServiceName(serviceName string) OptFunc {
	return func(opts *Opts) {
		opts.serviceName = serviceName
	}
}

func WithServiceAddr(serviceAddr string) OptFunc {
	return func(opts *Opts) {
		opts.serviceAddr = serviceAddr
	}
}


func WithServiceID(salt string) OptFunc {
	return func(opts *Opts) {
		opts.serviceID = fmt.Sprintf("%s-%s-%d", opts.serviceName, salt, randomNum())
	}
}


func randomNum() int {
	return rand.New(rand.NewSource(time.Now().Unix())).Int()
}
 
