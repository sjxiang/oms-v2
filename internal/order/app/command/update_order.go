package command

import (
	"context"
	"fmt"

	"github.com/sjxiang/oms-v2/order/domain"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)


type UpdateOrder struct {
	Order    *domain.Order
	UpdateFn func(context.Context, *domain.Order) (*domain.Order, error)
}


type UpdateOrderResult struct {
	OrderID string
}

type UpdateOrderHandler interface {
	Handle(ctx context.Context, cmd UpdateOrder) (result *UpdateOrder, err error)
}


func NewUpdateOrderHandler(orderRepo domain.Repository, logger *zap.Logger) UpdateOrderHandler {
	if orderRepo == nil {
		panic("order repo is nil")
	}

	baseHandler := updateOrderHandler{
		orderRepo: orderRepo,
	}
	
	return ApplyQueryUpdateOrderDecorators(baseHandler, logger)
}



type updateOrderHandler struct {
	orderRepo domain.Repository
}

func (h updateOrderHandler) Handle(ctx context.Context, cmd UpdateOrder) (result *UpdateOrder, err error) {
	return nil, h.orderRepo.Update(ctx, cmd.Order, cmd.UpdateFn)
}


// ----------------------------
// 装饰器 logging
// ----------------------------

type updateOrderHandlerLogging struct {
	logger *zap.Logger
	base   UpdateOrderHandler
}

func (h updateOrderHandlerLogging) Handle(ctx context.Context, cmd UpdateOrder) (result *UpdateOrder, err error) {

	h.logger.Debug("开始处理执行", 
		zapcore.Field{
			Key: "command",
			Type: zapcore.StringType,
			String: "__Update_Order__",
		},
		zapcore.Field{
			Key: "command_body",
			Type: zapcore.StringType,
			String: fmt.Sprintf("%#v", cmd),
		},
	)

	defer func() {
		if err != nil {
			h.logger.Error("failed", zap.Error(err))
		} else {
			h.logger.Info("success")
		}
	}()

	return h.base.Handle(ctx, cmd)
}


func ApplyQueryUpdateOrderDecorators(base UpdateOrderHandler, logger *zap.Logger) UpdateOrderHandler {
	return updateOrderHandlerLogging{
		logger: logger,
		base:   base,
	}
}
