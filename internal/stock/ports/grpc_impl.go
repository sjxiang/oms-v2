package ports

import (
	"context"

	"github.com/sjxiang/oms-v2/common/pb"
	"github.com/sjxiang/oms-v2/stock/app"
)

type GrpcServer struct {
	app app.Application
}

func NewGrpcServer(app app.Application) *GrpcServer {
	return &GrpcServer{app: app}
}

func (s *GrpcServer) CheckIfItemsInStock(ctx context.Context, req *pb.CheckIfItemsInStockRequest) (*pb.CheckIfItemsInStockResponse, error) {
	return &pb.CheckIfItemsInStockResponse{
		
	}, nil
}

func (s *GrpcServer) GetItems(ctx context.Context, req *pb.GetItemsRequest) (*pb.GetItemsResponse, error) {
	return &pb.GetItemsResponse{
		
	}, nil 
}