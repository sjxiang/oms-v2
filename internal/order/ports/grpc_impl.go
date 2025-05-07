package ports

import (
	"context"

	emptypb "google.golang.org/protobuf/types/known/emptypb"

	"github.com/sjxiang/oms-v2/common/pb"
	"github.com/sjxiang/oms-v2/order/app"
)

type GrpcServer struct {
	app app.Application
}

func NewGrpcServer(app app.Application) *GrpcServer {
	return &GrpcServer{app: app}
}

func (s *GrpcServer) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (resp *emptypb.Empty, err error) {
	return &emptypb.Empty{}, nil
}

func (s *GrpcServer) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (resp *pb.Order, err error) {
	return &pb.Order{}, nil
}

func (s *GrpcServer) UpdateOrder(ctx context.Context, req *pb.Order) (resp *emptypb.Empty, err error) {
	return &emptypb.Empty{}, nil
}