package command

import (
	"context"
	"fmt"

	"github.com/sjxiang/oms-v2/common/pb"
	"github.com/sjxiang/oms-v2/order/domain"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type CreateOrder struct {
	CustomerID string
	Items      []*pb.ItemWithQuantity  // 前端传递, 商品编码和商品数量
}

type CreateOrderResult struct {
	OrderID string
}

type CreateOrderHandler interface {
	Handle(ctx context.Context, cmd CreateOrder) (result *CreateOrderResult, err error)
}

type CreateOrderHandlerImpl struct {
	orderRepo domain.Repository
}

func (impl CreateOrderHandlerImpl) Handle(ctx context.Context, cmd CreateOrder) (result *CreateOrderResult, err error) {
	
	// 1. 调用 `stock` gRPC 服务
	var stockResponse []*pb.Item
		
	for _, item := range cmd.Items {
		stockResponse = append(stockResponse, &pb.Item{
			Id:       item.Id,
			Quantity: item.Quantity,
		})

	}

	o, err := impl.orderRepo.Create(ctx, &domain.Order{
		CustomerID: cmd.CustomerID,
		Items:      stockResponse,
	})
	if err != nil {
		return nil, err
	}

	return &CreateOrderResult{
		OrderID: o.ID,
	}, nil
}

func NewCreateOrderHandler(orderRepo domain.Repository,  logger *zap.Logger) CreateOrderHandler {
	return applyCreateOrderCommandWrapper(
		CreateOrderHandlerImpl{
			orderRepo: orderRepo,
		}, logger)
}


// 集成
func applyCreateOrderCommandWrapper(handler CreateOrderHandler, logger *zap.Logger) CreateOrderHandler {
	return createOrderCommandLoggingWrapper{
		logger: logger,
		base:   handler,
	}
}


// 套娃组件
type createOrderCommandLoggingWrapper struct {
	logger *zap.Logger
	base   CreateOrderHandler
}

func (w createOrderCommandLoggingWrapper) Handle(ctx context.Context, cmd CreateOrder) (result *CreateOrderResult, err error) {
	w.logger.Debug("开始处理执行", 
		zapcore.Field{
			Key: "command",
			Type: zapcore.StringType,
			String: "__Create_Order__",
		},
		zapcore.Field{
			Key: "command_body",
			Type: zapcore.StringType,
			String: fmt.Sprintf("%#v", cmd),
		},
	)
	defer func() {
		if err != nil {
			w.logger.Error("处理创建订单命令失败", zap.Error(err))
		} else {
			w.logger.Info("处理创建订单命令成功")
		}
	}()

	return w.base.Handle(ctx, cmd)
}
