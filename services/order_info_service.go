package services

import (
	"context"

	"github.com/thangpham4/self-project/entities"
	"github.com/thangpham4/self-project/pkg/logger"
	"github.com/thangpham4/self-project/repo"
)

type OrderInfoService struct {
	orderRepo repo.OrderInfoRepo
	logger    logger.Logger
}

func NewOrderInfoService(
	orderRepo repo.OrderInfoRepo,
) *OrderInfoService {
	return &OrderInfoService{
		orderRepo: orderRepo,
		logger:    logger.Factory("OrderInfoService"),
	}
}

func (m *OrderInfoService) Create(ctx context.Context, order *entities.OrderInfo) (*entities.OrderInfoTransform, error) {
	return m.orderRepo.Create(ctx, order)
}

func (m *OrderInfoService) GetByID(ctx context.Context, id uint) (*entities.OrderInfoTransform, error) {
	return m.orderRepo.GetByID(ctx, id)
}

func (m *OrderInfoService) GetByCustomerID(ctx context.Context, customerID int32) ([]*entities.OrderInfoTransform, error) {
	return m.orderRepo.GetByCustomerID(ctx, customerID)
}

func (m *OrderInfoService) GetByProductID(ctx context.Context, productID int32) ([]*entities.OrderInfoTransform, error) {
	return m.orderRepo.GetByProductID(ctx, productID)
}
