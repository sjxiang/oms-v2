package query

import (
	"context"

	"go.uber.org/zap"

	"github.com/sjxiang/oms-v2/common/pb"
	"github.com/sjxiang/oms-v2/stock/domain"
)


// func (w queryLoggingWrapper) Handle(ctx context.Context, cmd GetCustomerOrder) (result *domain.Order, err error) {
// 	w.logger.Debug("开始处理查询",
// 		zapcore.Field{
// 			Key: "query",
// 			Type: zapcore.StringType,
// 			String: "__Get_Customer_Order__",
// 		},
// 		zapcore.Field{
// 			Key: "query_body",
// 			Type: zapcore.StringType,
// 			String: fmt.Sprintf("%#v", cmd),
// 		},
// 	)
// 	defer func() {
// 		if err != nil {
// 			w.logger.Error("查询处理失败", zap.Error(err))
// 		} else {
// 			w.logger.Info("查询处理完成")
// 		}
// 	}()

// 	return w.base.Handle(ctx, cmd)
// }


type CheckIfItemsInStock struct {
    ItemIDs []*pb.ItemWithQuantity
}

type CheckIfItemsInStockHandler interface {
    Handle(ctx context.Context, query CheckIfItemsInStock) ([]*pb.Item, error)
}


type checkIfItemsInStockHandler struct {
    repo domain.Repository
}

func (h checkIfItemsInStockHandler) Handle(ctx context.Context, query CheckIfItemsInStock) ([]*pb.Item, error) {
    return nil, nil
}


func NewCheckIfItemsInStockHandler(repo domain.Repository, logger *zap.Logger) CheckIfItemsInStockHandler {
	if repo == nil {
		panic("stock repo is nil")
	}
	
	baseHandler := checkIfItemsInStockHandler{
		repo: repo,
	}

	return ApplyQueryCheckIfItemsInStockDecorators(baseHandler, logger)
}



// ----------------------------
// 装饰器 logging
// ----------------------------

type CheckIfItemsInStockHandlerLogging struct {
	logger *zap.Logger
	base   CheckIfItemsInStockHandler
}

func (h CheckIfItemsInStockHandlerLogging) Handle(ctx context.Context, query CheckIfItemsInStock) ([]*pb.Item, error) {
	return h.base.Handle(ctx, query)
} 


func ApplyQueryCheckIfItemsInStockDecorators(base CheckIfItemsInStockHandler, logger *zap.Logger) CheckIfItemsInStockHandler {
	return CheckIfItemsInStockHandlerLogging{
		logger: logger, 
		base:   base,
	}
}
