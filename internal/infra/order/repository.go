package order

import (
	"context"
	
	orderDom "application-design-test-master/internal/domain/order"
)

type Repository struct {
}

func NewRepository() *Repository {
	return &Repository{}
}

var orders []orderDom.Order

func (r *Repository) GetAll(_ context.Context) []orderDom.Order {
	return orders
}

func (r *Repository) Create(_ context.Context, order orderDom.Order) bool {
	orders = append(orders, order)
	return true
}
