package domain


import "context"


type Repository interface {
	Create(ctx context.Context, o *Order) (*Order, error)
	Get(ctx context.Context, id, customerID string) (*Order, error)
	Update(ctx context.Context, o *Order, updateFn func(context.Context, *Order) (*Order, error)) error
}