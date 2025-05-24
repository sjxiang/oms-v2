package command

import (
	"context"
	"fmt"
	"errors"

	"github.com/sjxiang/oms-v2/common/pb"
	"github.com/sjxiang/oms-v2/order/app/query"
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
	Handle(ctx context.Context, cmd CreateOrder) (*CreateOrderResult, error)
}


func NewCreateOrderHandler(orderRepo domain.Repository, stockGrpc query.StockService, logger *zap.Logger) CreateOrderHandler {
	if orderRepo == nil {
		panic("order repo is nil")
	}

	baseHandler := createOrderHandler{
		orderRepo: orderRepo,
		stockGrpc: stockGrpc,
	}

	return ApplyQueryCreateOrderDecorators(baseHandler, logger)
}


type createOrderHandler struct {
	orderRepo domain.Repository
	stockGrpc query.StockService
}

func (h createOrderHandler) Handle(ctx context.Context, cmd CreateOrder) (result *CreateOrderResult, err error) {
// 	validItems, err := h.validate(ctx, cmd.Items)
// 	if err != nil {
// 		return nil, err
// 	}


	// // 1. 调用 `stock` gRPC 服务
	// if err := impl.stockGrpc.CheckIfItemsInStock(ctx, cmd.Items); err != nil {
	// 	return nil, err
	// }

	// stockResponse, err := impl.stockGrpc.GetItems(ctx, []string{"1"})
	// if err != nil {
	// 	return nil, err
	// }

	// o, err := impl.orderRepo.Create(ctx, &domain.Order{
	// 	CustomerID: cmd.CustomerID,
	// 	Items:      stockResponse,
	// })
	// if err != nil {
	// 	return nil, err
	// }

	// return &CreateOrderResult{
	// 	OrderID: o.ID,
	// }, nil

	panic("")
}


func (h createOrderHandler) validate(ctx context.Context, items []*pb.ItemWithQuantity) ([]*pb.Item, error) {
	if len(items) == 0 {
		return nil, errors.New("must have at least one item")
	}
	// 合并相同的 item id
	mergedItems := packItems(items)

	// 调用 stock service 检查库存是否充足
	if err := h.stockGrpc.CheckIfItemsInStock(ctx, mergedItems); err != nil {
		return nil, err
	}

	ids := make([]string, 0, len(mergedItems))
	for _, item := range mergedItems {
		ids = append(ids, item.Id)
	}

	// 调用 stock service 获取商品信息
	return h.stockGrpc.GetItems(ctx, ids)
}



func packItems(items []*pb.ItemWithQuantity) []*pb.ItemWithQuantity {
	merged := make(map[string]int32, 0)

	for _, item := range items {
		merged[item.Id] = item.Quantity
	}

	mergedItems := make([]*pb.ItemWithQuantity, 0, len(merged))

	for id, quantity := range merged {
		mergedItems = append(mergedItems, &pb.ItemWithQuantity{
			Id:       id,
			Quantity: quantity,
		})
	}

	return mergedItems
}


// ----------------------------
// 装饰器 logging
// ----------------------------

type createOrderHandlerLogging struct {
	logger *zap.Logger
	base   CreateOrderHandler
}

func (h createOrderHandlerLogging) Handle(ctx context.Context, cmd CreateOrder) (result *CreateOrderResult, err error) {
	
	h.logger.Debug("开始处理执行", 
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
			h.logger.Error("failed", zap.Error(err))
		} else {
			h.logger.Info("success")
		}
	}()

	return h.base.Handle(ctx, cmd)

}


// buildCreateOrderHandlerChain

func ApplyQueryCreateOrderDecorators(base CreateOrderHandler, logger *zap.Logger) CreateOrderHandler {
	return createOrderHandlerLogging{
		logger: logger,
		base:   base,
	}
}
