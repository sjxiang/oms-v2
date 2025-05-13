package adapters

import (
	"context"
	"strconv"
	"sync"
	"time"

	"go.uber.org/zap"

	"github.com/sjxiang/oms-v2/order/domain"
)

type MemoryOrderRepository struct {
	lock   *sync.RWMutex
	store  []*domain.Order
	logger *zap.Logger
}

func NewMemoryOrderRepository(logger *zap.Logger) *MemoryOrderRepository {
	
	return &MemoryOrderRepository{
		lock:   &sync.RWMutex{},
		store:  make([]*domain.Order, 0),
		logger: logger,
	}
}


func (m *MemoryOrderRepository) Create(ctx context.Context, o *domain.Order) (*domain.Order, error) {
	m.lock.Lock()
	defer m.lock.Unlock()

	newOrder := &domain.Order{
		ID:          strconv.FormatInt(time.Now().Unix(), 10),
		CustomerID:  o.CustomerID,
		Status:      o.Status,
		PaymentLink: o.PaymentLink,
		Items:       o.Items,
	}
	
	m.store = append(m.store, newOrder)

	m.logger.Debug("__inmemory_repo_create_order__", zap.Any("before", o), zap.Any("after", newOrder))

	return newOrder, nil
}

func (m *MemoryOrderRepository) Get(ctx context.Context, id, customerID string) (*domain.Order, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	for _, o := range m.store {
		if o.ID == id && o.CustomerID == customerID {
	
			m.logger.Debug("__inmemory_repo_get_order__", zap.String("id", id), zap.String("customer_id", customerID), zap.Any("order", o))
			return o, nil
		}
	}

	return nil, domain.NotFoundError{
		OrderID: id,
	}
}

func (m *MemoryOrderRepository) Update(ctx context.Context, o *domain.Order, 
		updateFn func(context.Context, *domain.Order) (*domain.Order, error)) error {

	m.lock.RLock()
	defer m.lock.RUnlock()


	found := false
	for i, order := range m.store {
		
		if order.ID == o.ID && order.CustomerID == o.CustomerID {
			found = true

			updatedOrder, err := updateFn(ctx, order)
			if err != nil {
				return err
			}

			m.logger.Debug("__inmemory_repo_update_order__", zap.String("id", o.ID), zap.String("customer_id", o.CustomerID), zap.Any("order", updatedOrder))
			m.store[i] = updatedOrder
		}
	}

	if !found {
		return domain.NotFoundError{
			OrderID: o.ID,
		}
	}
	
	return nil
}
