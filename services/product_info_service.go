package services

import (
	"context"

	"github.com/thangpham4/self-project/entities"
	"github.com/thangpham4/self-project/pkg/logger"
	"github.com/thangpham4/self-project/repo"
)

type ProductInfoService struct {
	productRepo repo.ProductInfoRepo
	logger      logger.Logger
}

func NewProductInfoService(
	productRepo repo.ProductInfoRepo,
) *ProductInfoService {
	return &ProductInfoService{
		productRepo: productRepo,
		logger:      logger.Factory("ProductInfoService"),
	}
}

func (u *ProductInfoService) Create(ctx context.Context, product *entities.ProductInfo) (*entities.ProductInfo, error) {
	return u.productRepo.Create(ctx, product)
}

func (u *ProductInfoService) Get(ctx context.Context, id uint) (*entities.ProductInfo, error) {
	return u.productRepo.Get(ctx, id)
}

func (u *ProductInfoService) GetMany(ctx context.Context, ids []uint) ([]*entities.ProductInfo, error) {
	return u.productRepo.GetMany(ctx, ids)
}

func (u *ProductInfoService) GetAll(ctx context.Context) ([]*entities.ProductInfo, error) {
	return u.productRepo.GetAll(ctx)
}
