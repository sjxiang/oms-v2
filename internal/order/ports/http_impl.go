package ports

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sjxiang/oms-v2/order/app"
)

// ensure that we've conformed to the `ServerInterface` with a compile-time check
var _ ServerInterface = (*HTTPServer)(nil)

type HTTPServer struct{
	app app.Application
}

func NewHTTPServer(app app.Application) HTTPServer {
	return HTTPServer{app: app}
}


type Pong struct {
	Ping string `json:"ping"`
}

// (GET /ping)
func (HTTPServer) GetPing(ctx *gin.Context) {
	resp := Pong{
		Ping: "pong",
	}

	ctx.JSON(http.StatusOK, resp)
}


// (GET /customer/{customer_id}/orders/{order_id})
func (HTTPServer) GetOrder(c *gin.Context) {
	customerID := c.Param("customer_id")   // 获取 customer_id
    orderID := c.Param("order_id")         // 获取 order_id

	c.JSON(http.StatusOK, gin.H{
		"customer_id": customerID,
		"order_id":    orderID,
	})
}

// (POST /customer/{customer_id}/orders)
func (HTTPServer) CreateOrder(c *gin.Context) {
	customerID := c.Param("customer_id")   // 获取 customer_id
	
	c.JSON(http.StatusOK, gin.H{
		"customer_id": customerID,
	})
}
