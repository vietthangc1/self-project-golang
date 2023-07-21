package services

import (
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
