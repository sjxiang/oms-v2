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


type UpdateOrderHandlerImpl struct {
	orderRepo domain.Repository
}

func (impl UpdateOrderHandlerImpl) Handle(ctx context.Context, cmd UpdateOrder) (result *UpdateOrder, err error) {
	err = impl.orderRepo.Update(ctx, cmd.Order, cmd.UpdateFn)
	if err != nil {
		return nil, err
	}

	return nil, nil 
}

func NewUpdateOrderHandler(orderRepo domain.Repository, logger *zap.Logger) UpdateOrderHandler {
	return applyUpdateOrderCommandWrapper(UpdateOrderHandlerImpl{
		orderRepo: orderRepo,
	}, logger)
}

func applyUpdateOrderCommandWrapper(handler UpdateOrderHandler, logger *zap.Logger) UpdateOrderHandler {
	return UpdateOrderCommandLoggingWrapper{logger: logger, base: handler}
}

type UpdateOrderCommandLoggingWrapper struct {
	logger *zap.Logger
	base   UpdateOrderHandler
}

func (w UpdateOrderCommandLoggingWrapper) Handle(ctx context.Context, cmd UpdateOrder) (result *UpdateOrder, err error) {
	w.logger.Debug("开始处理更新订单命令",
		zapcore.Field{
			Key: "command",
			Type: zapcore.StringType,
			String: "__Update_Order__",
		},
		zapcore.Field{
			Key: "command_body",
			Type: zapcore.StringType,
			String: fmt.Sprintf("%+v", cmd),
		},
	)

	defer func() {
		if err != nil {
			w.logger.Error("处理更新订单命令失败", zap.Error(err))
		} else {
			w.logger.Info("处理更新订单命令成功")
		}
	}()

	return w.base.Handle(ctx, cmd)
}