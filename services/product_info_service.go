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

const (
	productNoImageURL string = "https://t4.ftcdn.net/jpg/04/70/29/97/360_F_470299797_UD0eoVMMSUbHCcNJCdv2t8B2g1GVqYgs.jpg"
)

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
	productsArray, err := u.productRepo.GetMany(ctx, ids)
	if err != nil {
		u.logger.Error(err, "error in get many products")
		return nil, err
	}
	for _, product := range productsArray {
		if product.ImageURI == "" {
			product.ImageURI = productNoImageURL
		}
	}
	return productsArray, nil
}

func (u *ProductInfoService) GetAll(ctx context.Context) ([]*entities.ProductInfo, error) {
	u.logger.Info("running get all service")
	return u.productRepo.GetAll(ctx)
}
