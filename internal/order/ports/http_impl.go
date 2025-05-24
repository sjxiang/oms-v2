package ports

import (
	"net/http"

	"github.com/gin-gonic/gin"
	
	"github.com/sjxiang/oms-v2/common/pb"
	"github.com/sjxiang/oms-v2/order/app"
	"github.com/sjxiang/oms-v2/order/app/command"
	"github.com/sjxiang/oms-v2/order/app/query"
)

// ensure that we've conformed to the `ServerInterface` with a compile-time check
var _ ServerInterface = (*HTTPServer)(nil)

type HTTPServer struct{
	app app.Application
}

func NewHTTPServer(app app.Application) (HTTPServer, error) {
	return HTTPServer{app: app}, nil 
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

	// 参数校验, 略


	o, err := h.app.Queries.GetCustomerOrderHandler.Handle(c, query.GetCustomerOrder{
		CustomerID: c.Param("customer_id"),   // 获取顾客 ID
		OrderID:    c.Param("order_id"),      // 获取订单 ID
	})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success", 
		"data": gin.H{
			"order": o,
		},
	})
}

// (POST /customer/{customer_id}/orders)
func (h HTTPServer) CreateOrder(c *gin.Context) {
	
	var req CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	o, err := h.app.Commands.CreateOrderHandler.Handle(c, command.CreateOrder{
		CustomerID: req.CustomerId,
		Items:      packItems(req.Items),
	})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    gin.H{
			"order_id":    o.OrderID,
			"customer_id": req.CustomerId,
		},
	})
}

type ItemWithQuantity struct {
	Id       string `json:"id,omitempty"`
	Quantity int32  `json:"quantity,omitempty"`
}

type CreateOrderRequest struct {
	CustomerId string              `json:"customer_id,omitempty"`
	Items      []*ItemWithQuantity `json:"items,omitempty"`
}


func packItem(item *ItemWithQuantity) *pb.ItemWithQuantity {
	if item == nil {
		return nil
	}

	return &pb.ItemWithQuantity{
		Id:       item.Id,
		Quantity: item.Quantity,
	}
}

func packItems(items []*ItemWithQuantity) []*pb.ItemWithQuantity {
	res := make([]*pb.ItemWithQuantity, 0)

	for _, item := range items {
		if i := packItem(item); i != nil {
			res = append(res, i)		
		}
	}

	return res
}