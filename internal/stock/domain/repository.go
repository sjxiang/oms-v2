package domain


import (
	"context"

	"github.com/sjxiang/oms-v2/common/pb"
)


type Repository interface {
	GetItems(ctx context.Context, ids []string) ([]*pb.Item, error) 
}