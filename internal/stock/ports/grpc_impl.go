package ports

import (
	"context"

	"go.uber.org/zap"

	"github.com/sjxiang/oms-v2/common/pb"
	"github.com/sjxiang/oms-v2/stock/app"
	"github.com/sjxiang/oms-v2/stock/app/query"
)

type GrpcServer struct {
	pb.UnimplementedStockServiceServer

	app    app.Application
	logger *zap.Logger
}

func NewGrpcServer(app app.Application, logger *zap.Logger) (*GrpcServer, error) {
	return &GrpcServer{app: app, logger: logger}, nil
}


func (s *GrpcServer) CheckIfItemsInStock(ctx context.Context, req *pb.CheckIfItemsInStockRequest) (*pb.CheckIfItemsInStockResponse, error) {

	items, err := s.app.Queries.CheckIfItemsInStock.Handle(ctx, query.CheckIfItemsInStock{ItemIDs: req.Items})
	if err != nil {
		return nil, err
	}

	return &pb.CheckIfItemsInStockResponse{InStock: true, Items: items}, nil
}


func (s *GrpcServer) GetItems(ctx context.Context, req *pb.GetItemsRequest) (*pb.GetItemsResponse, error) {
	
	items, err := s.app.Queries.GetItems.Handle(ctx, query.GetItems{ItemIDs: req.ItemIds})
	if err != nil {
		return nil, err
	} 
	
	return &pb.GetItemsResponse{Items: items}, nil 
}