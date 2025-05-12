package ports

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sjxiang/oms-v2/order/app"
	"github.com/sjxiang/oms-v2/order/app/query"
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
func (h HTTPServer) GetOrder(c *gin.Context) {

	args := query.GetCustomerOrder{
		CustomerID: c.Param("customer_id"),   // 获取顾客编码
		OrderID: c.Param("order_id"),         // 获取订单编码
	} 

	o, err := h.app.Queries.GetCustomerOrderHandler.Handle(c, args)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success", 
		"data": o,
	})
}

// (POST /customer/{customer_id}/orders)
func (HTTPServer) CreateOrder(c *gin.Context) {
	customerID := c.Param("customer_id")   // 获取 customer_id
	
	c.JSON(http.StatusOK, gin.H{
		"customer_id": customerID,
	})
}
