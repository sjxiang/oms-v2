package query

import (
	"context"

	"github.com/sjxiang/oms-v2/common/pb"
)


type StockService interface {
	CheckIfItemsInStock(ctx context.Context, items []*pb.ItemWithQuantity) error
	GetItems(ctx context.Context, itemIDs []string) ([]*pb.Item, error)
}
