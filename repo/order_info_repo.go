package repo

import (
	"context"

	"github.com/thangpham4/self-project/entities"
)

type OrderInfoRepo interface {
	Create(ctx context.Context, order *entities.OrderInfoTransform) (*entities.OrderInfoTransform, error)
	GetByID(ctx context.Context, id uint) (*entities.OrderInfoTransform, error)
	GetByCustomerID(ctx context.Context, customerID int32) ([]*entities.OrderInfoTransform, error)
	GetByProductID(ctx context.Context, productID int32) ([]*entities.OrderInfoTransform, error)
}
