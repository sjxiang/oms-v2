package query

import (
	"context"
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/sjxiang/oms-v2/order/domain"
)


type GetCustomerOrder struct {
	CustomerID string
	OrderID string
}

type GetCustomerOrderHandler interface {
	Handle(ctx context.Context, query GetCustomerOrder) (*domain.Order, error)
}

func NewGetCustomerOrderHandler(orderRepo domain.Repository, logger *zap.Logger) GetCustomerOrderHandler {
	if orderRepo == nil {
		panic("order repo is nil")
	}

	baseHandler := getCustomerOrderHandler{
		orderRepo: orderRepo,
	}

	return ApplyQueryGetCustomerOrderDecorators(baseHandler, logger)
}


type getCustomerOrderHandler struct {
	orderRepo domain.Repository
}

func (h getCustomerOrderHandler) Handle(ctx context.Context, query GetCustomerOrder) (*domain.Order, error) {
	return h.orderRepo.Get(ctx, query.OrderID, query.CustomerID)
}

// ----------------------------
// 装饰器 logging
// ----------------------------

type getCustomerOrderHandlerLogging struct {
	logger *zap.Logger
	base   GetCustomerOrderHandler
}

func (h getCustomerOrderHandlerLogging) Handle(ctx context.Context, query GetCustomerOrder) (result *domain.Order, err error) {
	
	h.logger.Debug("开始处理查询", 
		zapcore.Field{
			Key: "query",
			Type: zapcore.StringType,
			String: "__Get_Customer_Order__",
		},
		zapcore.Field{
			Key: "query_body",
			Type: zapcore.StringType,
			String: fmt.Sprintf("%#v", query),
		},
	)

	defer func() {
		if err != nil {
			h.logger.Error("failed", zap.Error(err))
		} else {
			h.logger.Info("success")
		}
	}()

	return h.base.Handle(ctx, query)
} 


func ApplyQueryGetCustomerOrderDecorators(base GetCustomerOrderHandler, logger *zap.Logger) GetCustomerOrderHandler {
	return getCustomerOrderHandlerLogging{
		logger: logger,
		base:   base,
	}
}
