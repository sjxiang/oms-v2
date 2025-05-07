package api

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)


// 运行 http 服务
func RunHTTPServer(serviceName string, wrapper func(router *gin.Engine)) {
	addr := viper.Sub(serviceName).GetString("http-addr")
	if addr == "" {
		panic("empty http address")
	}

	RunHTTPServerOnAddr(addr, wrapper)
}


func RunHTTPServerOnAddr(addr string, wrapper func(router *gin.Engine)) {
	router := gin.New()

	setMiddleware(router)	
	wrapper(router)

	if err := router.Run(addr); err != nil {
		panic(err)
	}
}


func setMiddleware(router *gin.Engine) {
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
}