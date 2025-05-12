package service

import (
	"context"
	
	"go.uber.org/zap"

	"github.com/sjxiang/oms-v2/order/adapters"
	"github.com/sjxiang/oms-v2/order/app"
	"github.com/sjxiang/oms-v2/order/app/query"
)


func NewApplication(ctx context.Context, logger *zap.Logger) app.Application {

	orderRepo := adapters.NewMemoryOrderRepository(logger)

	return app.Application{
		Commands: app.Commands{

		},
		Queries: app.Queries{
			GetCustomerOrderHandler: query.NewCustomerOrderHandler(orderRepo, logger),
		},
	}
}