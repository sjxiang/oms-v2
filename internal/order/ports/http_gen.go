package ports

import (
	"github.com/gin-gonic/gin"
)


// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /ping)
	GetPing(c *gin.Context)

	// (GET /customer/{customer_id}/orders/{order_id})
	GetOrder(c *gin.Context) 

	// (POST /customer/{customer_id}/orders)
	CreateOrder(c *gin.Context) 

}

// ServerInterfaceWrapper 讲上下文转换为参数
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
}

type MiddlewareFunc func(c *gin.Context)

// GetPing operation middleware
func (siw *ServerInterfaceWrapper) GetPing(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetPing(c)
}

func (siw *ServerInterfaceWrapper) GetOrder(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetOrder(c)
}

func (siw *ServerInterfaceWrapper) CreateOrder(c *gin.Context)  {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.CreateOrder(c)
}


// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
}


// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router gin.IRouter, si ServerInterface, options GinServerOptions) {
	
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
	}

	// 注册路由
	router.GET(options.BaseURL+"/ping", wrapper.GetPing)
	router.GET(options.BaseURL+"/customer/:customer_id/orders/:order_id", wrapper.GetOrder)
	router.POST(options.BaseURL+"/customer/:customer_id/orders", wrapper.CreateOrder)
}