package client


import (
	"context"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/sjxiang/oms-v2/common/pb"
)


func NewStockGrpcClient(ctx context.Context) (pb.StockServiceClient, func() error, error) {
	
	addr := viper.GetString("stock.grpc-addr")

	opts, err := buildGrpcDialOpts()
	if err != nil {
		return nil, nil, err
	}
	
	conn, err := grpc.NewClient(addr, opts...)
	if err!= nil {
		return nil, nil, err
	}

	closeFn := func() error {
		return conn.Close()
	}

	return pb.NewStockServiceClient(conn), closeFn, nil
}


func buildGrpcDialOpts() ([]grpc.DialOption, error) {
	return []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}, nil
}