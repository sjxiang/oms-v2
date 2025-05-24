package ports

import (
	"github.com/gin-gonic/gin"
)

// 约束



// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /ping)
	GetPing(c *gin.Context)

	// (GET /customer/{customer_id}/orders/{order_id})
	GetOrder(c *gin.Context) 

	// (POST /customer/{customer_id}/orders)
	CreateOrder(c *gin.Context) 

}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []gin.HandlerFunc
}


// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router gin.IRouter, si ServerInterface, opts GinServerOptions) {
	
	// 创建路由分组, 统一应用中间件
	group := router.Group(opts.BaseURL)
	group.Use(opts.Middlewares...)  // 统一应用中间件

	// 注册路由（直接指向接口方法）
	group.GET("/ping", si.GetPing)
	group.GET("/customer/:customer_id/orders/:order_id", si.GetOrder)
	group.POST("/customer/:customer_id/orders", si.CreateOrder)
}
