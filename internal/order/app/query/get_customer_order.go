package query

import (
	"context"
	"fmt"


	"github.com/sjxiang/oms-v2/order/domain"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)


type GetCustomerOrder struct {
	CustomerID string
	OrderID string
}

type GetCustomerOrderHandler interface {
	Handle(ctx context.Context, query GetCustomerOrder) (result *domain.Order, err error)
}

type GetCustomerOrderHandlerImpl struct {
	orderRepo domain.Repository
}

func (impl GetCustomerOrderHandlerImpl) Handle(ctx context.Context, query GetCustomerOrder) (result *domain.Order, err error) {
	return impl.orderRepo.Get(ctx, query.OrderID, query.CustomerID)
}

func NewCustomerOrderHandler(orderRepo domain.Repository, logger *zap.Logger) GetCustomerOrderHandler {
	return applyQueryWrapper(
		GetCustomerOrderHandlerImpl{
			orderRepo: orderRepo,
		}, logger)
}


// 集成套娃
func applyQueryWrapper(handler GetCustomerOrderHandler, logger *zap.Logger) GetCustomerOrderHandler {
	return queryLoggingWrapper{logger: logger, base: handler}
}


// 套娃组件
type queryLoggingWrapper struct {
	logger *zap.Logger
	base   GetCustomerOrderHandler
}

func (q queryLoggingWrapper) Handle(ctx context.Context, cmd GetCustomerOrder) (result *domain.Order, err error) {
	q.logger.Debug("开始处理查询", 
		zapcore.Field{
			Key: "query",
			Type: zapcore.StringType,
			String: "__Get_Customer_Order__",
		},
		zapcore.Field{
			Key: "query_body",
			Type: zapcore.StringType,
			String: fmt.Sprintf("%#v", cmd),
		},
	)
	defer func() {
		if err != nil {
			q.logger.Error("查询处理失败", zap.Error(err))
		} else {
			q.logger.Info("查询处理完成")
		}
	}()

	return q.base.Handle(ctx, cmd)
}
