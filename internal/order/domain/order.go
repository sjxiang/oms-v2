package domain

import "github.com/sjxiang/oms-v2/common/pb"


type Order struct {
	ID          string
	CustomerID  string
	Status      string
	PaymentLink string
	Items       []*pb.Item
}