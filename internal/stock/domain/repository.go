package domain


import (
	"context"

	"github.com/sjxiang/oms-v2/common/pb"
)


type Repository interface {
	// CheckIfItemsInStock(ctx context.Context, req *pb.CheckIfItemsInStockRequest) (*pb.CheckIfItemsInStockResponse, error) 
	GetItems(ctx context.Context, ids []string) ([]*pb.Item, error) 
}