package repo

import (
	"context"

	"github.com/thangpham4/self-project/entities"
)

type ProductInfoRepo interface {
	Get(ctx context.Context, id uint) (*entities.ProductInfo, error)
	Create(ctx context.Context, product *entities.ProductInfo) (*entities.ProductInfo, error)
}
