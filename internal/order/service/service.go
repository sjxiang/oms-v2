package service

import (
	"context"

	"go.uber.org/zap"

	"github.com/sjxiang/oms-v2/common/client"
	"github.com/sjxiang/oms-v2/order/adapters"
	"github.com/sjxiang/oms-v2/order/adapters/grpc"
	"github.com/sjxiang/oms-v2/order/app"
	"github.com/sjxiang/oms-v2/order/app/command"
	"github.com/sjxiang/oms-v2/order/app/query"
)


func NewApplication(ctx context.Context, log *zap.Logger) (app.Application, func()) {

	stockClient, closeStockClientFn, err := client.NewStockGrpcClient(ctx)
	if err != nil {
		panic(err)
	}

	stockGrpc := grpc.NewStockGrpc(stockClient, log)

	return newApplication(ctx, stockGrpc, log), func() {
		closeStockClientFn()
	}
}

func newApplication(ctx context.Context, stockGrpc query.StockService, log *zap.Logger) app.Application {
	orderRepo := adapters.NewMemoryOrderRepository(log)
	
	return app.Application{
		Commands: app.Commands{
			CreateOrderHandler: command.NewCreateOrderHandler(orderRepo, stockGrpc, log),
			UpdateOrderHandler: command.NewUpdateOrderHandler(orderRepo, log),
		},
		Queries: app.Queries{
			GetCustomerOrderHandler: query.NewGetCustomerOrderHandler(orderRepo, log),
		},
	}
}