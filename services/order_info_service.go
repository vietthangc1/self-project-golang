package services

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/thangpham4/self-project/config"
	"github.com/thangpham4/self-project/entities"
	"github.com/thangpham4/self-project/pkg/apix"
	"github.com/thangpham4/self-project/pkg/logger"
	"github.com/thangpham4/self-project/repo"
)

type OrderInfoService struct {
	orderRepo repo.OrderInfoRepo
	apiClient apix.APICaller
	logger    logger.Logger
}

func NewOrderInfoService(
	orderRepo repo.OrderInfoRepo,
	apiClient apix.APICaller,
) *OrderInfoService {
	return &OrderInfoService{
		orderRepo: orderRepo,
		apiClient: apiClient,
		logger:    logger.Factory("OrderInfoService"),
	}
}

func (m *OrderInfoService) Create(ctx context.Context, order *entities.OrderInfoTransform) (*entities.OrderInfoTransform, error) {
	return m.orderRepo.Create(ctx, order)
}

//nolint:govet
func (m *OrderInfoService) GetByID(ctx context.Context, id uint) (*entities.OrderInfoTransformProductInfo, error) {
	orderInfo, err := m.orderRepo.GetByID(ctx, id)
	if err != nil {
		m.logger.Error(err, "Error in getting order info", "order_id", id)
	}

	out := &entities.OrderInfoTransformProductInfo{
		ID:         orderInfo.ID,
		CustomerID: orderInfo.CustomerID,
		Products:   []entities.OrderProductInfo{},
		TotalValue: 0,
	}

	for _, product := range orderInfo.Products {
		productID := fmt.Sprintf("%d", product.ID)
		productURI := fmt.Sprintf("http://%s%s%s", config.ProductAPI, productsRoute, productID)

		resp, err := m.apiClient.Get(ctx, productURI, nil)
		if err != nil {
			m.logger.Error(err, "cannot get products info")
			return nil, err
		}

		var productsInfo []*entities.ProductInfo
		err = json.Unmarshal(resp, &productsInfo)
		if err != nil {
			m.logger.Error(err, "cannot get products info")
			return nil, err
		}
		value := productsInfo[0].Price * product.Quantity
		out.Products = append(out.Products, entities.OrderProductInfo{
			ProductInfo: *productsInfo[0],
			Quantity:    product.Quantity,
			Value:       value,
		})
		out.TotalValue += value
	}
	return out, err
}

func (m *OrderInfoService) GetByCustomerID(ctx context.Context, customerID int32) ([]*entities.OrderInfoTransform, error) {
	return m.orderRepo.GetByCustomerID(ctx, customerID)
}

func (m *OrderInfoService) GetByProductID(ctx context.Context, productID int32) ([]*entities.OrderInfoTransform, error) {
	return m.orderRepo.GetByProductID(ctx, productID)
}
