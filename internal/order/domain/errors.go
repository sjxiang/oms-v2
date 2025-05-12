package domain

import "fmt"


type NotFoundError struct {
	OrderID string
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("order with ID %s not found", e.OrderID)	
}