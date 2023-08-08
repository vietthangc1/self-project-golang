package mysql

import (
	"context"
	"fmt"

	"github.com/thangpham4/self-project/entities"
	"github.com/thangpham4/self-project/repo"
	"gorm.io/gorm"
)

var (
	_ repo.OrderInfoRepo = &OrderMysql{}
)

type OrderMysql struct {
	db *gorm.DB
}

func NewOrderMysql(
	db *gorm.DB,
) *OrderMysql {
	return &OrderMysql{
		db: db,
	}
}

func (m *OrderMysql) Create(ctx context.Context, order *entities.OrderInfoTransform) (*entities.OrderInfoTransform, error) {
	orderInput := order.ReTransform()
	err := m.db.WithContext(ctx).Create(orderInput).Error
	if err != nil {
		return nil, err
	}
	return orderInput.Transform(), nil
}

func (m *OrderMysql) GetByID(ctx context.Context, id uint) (*entities.OrderInfoTransform, error) {
	var orderResult = &entities.OrderInfo{
		ID: id,
	}
	err := m.db.WithContext(ctx).First(orderResult).Error
	if err != nil {
		return nil, err
	}
	return orderResult.Transform(), nil
}

func (m *OrderMysql) GetByCustomerID(ctx context.Context, customerID int32) ([]*entities.OrderInfoTransform, error) {
	ordersResult := []*entities.OrderInfo{}
	err := m.db.WithContext(ctx).Where("customer_id = ", customerID).Find(ordersResult).Error
	if err != nil {
		return nil, err
	}
	ordersResultTransform := []*entities.OrderInfoTransform{}
	for _, order := range ordersResult {
		ordersResultTransform = append(ordersResultTransform, order.Transform())
	}
	return ordersResultTransform, nil
}

func (m *OrderMysql) GetByProductID(ctx context.Context, productID int32) ([]*entities.OrderInfoTransform, error) {
	ordersResult := []*entities.OrderInfo{}
	err := m.db.WithContext(ctx).Where("product_ids like ", fmt.Sprintf("%d,", productID)).Find(ordersResult).Error
	if err != nil {
		return nil, err
	}
	ordersResultTransform := []*entities.OrderInfoTransform{}
	for _, order := range ordersResult {
		ordersResultTransform = append(ordersResultTransform, order.Transform())
	}
	return ordersResultTransform, nil
}
