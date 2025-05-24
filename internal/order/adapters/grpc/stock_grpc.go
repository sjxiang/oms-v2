package grpc

import (
	"context"

	"go.uber.org/zap"

	"github.com/sjxiang/oms-v2/common/pb"
)


type StockGRPC struct {
	client pb.StockServiceClient
	log    *zap.Logger
}

func NewStockGrpc(client pb.StockServiceClient, log *zap.Logger) *StockGRPC {
	return &StockGRPC{
		client: client,
		log:    log,
	}
}

func (impl StockGRPC) CheckIfItemsInStock(ctx context.Context, items []*pb.ItemWithQuantity) error {
	_, err := impl.client.CheckIfItemsInStock(ctx, &pb.CheckIfItemsInStockRequest{
		Items: items,
	})

	return err 
}

func (impl StockGRPC) GetItems(ctx context.Context, itemIDs []string) ([]*pb.Item, error) {
	resp, err := impl.client.GetItems(ctx, &pb.GetItemsRequest{
		ItemIds: itemIDs,
	})
	if err != nil {
		return nil, err
	}

	return resp.Items, nil
}


