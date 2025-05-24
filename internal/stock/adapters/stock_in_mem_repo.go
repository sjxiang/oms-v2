package adapters

import (
	"context"
	"sync"

	"github.com/sjxiang/oms-v2/common/pb"
	"github.com/sjxiang/oms-v2/stock/domain"
)


type MemoryStockRepository struct {
	lock  *sync.RWMutex
	store map[string]*pb.Item
}


var fakeData = map[string]*pb.Item{
	"商品A": {
		Id:       "编号 1",
		Name:     "玉溪",
		Quantity: 100,
		PriceId:  "33",
	},
	"商品B": {
		Id:       "编号 2",
		Name:     "茅台",
		Quantity: 100,
		PriceId:  "1200",
	},
}


func NewMemoryStockRepository() *MemoryStockRepository {
	
	return &MemoryStockRepository{
		lock:  &sync.RWMutex{},
		store: fakeData,
	}
}


func (m *MemoryStockRepository) GetItems(ctx context.Context, ids []string) ([]*pb.Item, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	var (
		res []*pb.Item
		missingItemIDs []string
	)

	for _, id := range ids {
		if item, exist := m.store[id]; exist {
			res = append(res, item)
		} else {
			missingItemIDs = append(missingItemIDs, id)
		}
	}

	if len(res) == len(ids) {
		return res, nil
	}

	return nil, domain.NotFoundError{
		MissingItemIDs: missingItemIDs,
	}
}


