package service

import (
	"context"

	"github.com/sjxiang/oms-v2/stock/adapters"
	"github.com/sjxiang/oms-v2/stock/app"
	"github.com/sjxiang/oms-v2/stock/app/query"
	"go.uber.org/zap"
)


func NewApplication(ctx context.Context, logger *zap.Logger) app.Application {
	
	stockRepo := adapters.NewMemoryStockRepository()

	return app.Application{
		Commands: app.Commands{},
		Queries: app.Queries{
			CheckIfItemsInStock: query.NewCheckIfItemsInStockHandler(stockRepo, logger),
			GetItems:            query.NewGetItemsHandler(stockRepo, logger),
		},
	}
}