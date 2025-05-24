package query

import (
	"context"

	"go.uber.org/zap"

	"github.com/sjxiang/oms-v2/common/pb"
	"github.com/sjxiang/oms-v2/stock/domain"
)


type GetItems struct {
    ItemIDs []string
}

type GetItemsHandler interface {
    Handle(ctx context.Context, query GetItems) ([]*pb.Item, error)
}


type getItemsHandler struct {
    repo domain.Repository
}

func (h getItemsHandler) Handle(ctx context.Context, query GetItems) ([]*pb.Item, error) {
    return nil, nil
}


func NewGetItemsHandler(repo domain.Repository, logger *zap.Logger) GetItemsHandler {
	if repo == nil {
		panic("stock repo is nil")
	}
	
	baseHandler := getItemsHandler{
		repo: repo,
	}
	
	return ApplyQueryGetItemsDecorators(baseHandler, logger)
}



// ----------------------------
// 装饰器 logging
// ----------------------------

type GetItemsHandlerLogging struct {
	logger *zap.Logger
	base   GetItemsHandler
}

func (h GetItemsHandlerLogging) Handle(ctx context.Context, query GetItems) ([]*pb.Item, error) {
	return h.base.Handle(ctx, query)
}

func ApplyQueryGetItemsDecorators(base GetItemsHandler, logger *zap.Logger) GetItemsHandler {
	return GetItemsHandlerLogging{
		logger: logger,
		base:   base,
	}
}


