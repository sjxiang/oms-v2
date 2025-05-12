package adapters

import (
	"github.com/sjxiang/oms-v2/common/pb"
	"github.com/sjxiang/oms-v2/order/domain"
)


var defaultOrder = []*domain.Order{
	{
		ID: "20002",
		CustomerID: "10001",
		Status: "PENDING",
		PaymentLink: "fake-payment-link-1",
		Items: []*pb.Item{
			{
				Id: "fake-item-id-1",
				Name: "fake-product-id-1",
				Quantity: 1,
				PriceId: "fake-price-id-1",
			},
			{
				Id: "fake-item-id-2",
				Name: "fake-product-id-2",
				Quantity: 2,
				PriceId: "fake-price-id-2",
			},
		},
	},
}
