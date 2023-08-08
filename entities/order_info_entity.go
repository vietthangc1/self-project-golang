package entities

import (
	"encoding/json"

	"github.com/thangpham4/self-project/pkg/logger"
)

type OrderInfo struct {
	ID         uint   `gorm:"autoIncrement" json:"id"`
	Products   string `json:"products"`
	CustomerID int32  `json:"customer_id"`
}

type OrderInfoTransform struct {
	ID         uint           `gorm:"autoIncrement" json:"id"`
	Products   []OrderProduct `json:"products"`
	CustomerID int32          `json:"customer_id"`
}

type OrderProduct struct {
	ID       uint  `gorm:"autoIncrement" json:"id"`
	Quantity int32 `json:"quantity"`
}

type OrderInfoTransformProductInfo struct {
	ID         uint               `gorm:"autoIncrement" json:"id"`
	Products   []OrderProductInfo `json:"products"`
	CustomerID int32              `json:"customer_id"`
	TotalValue int32              `json:"total_value"`
}

type OrderProductInfo struct {
	ProductInfo ProductInfo `json:"product_info"`
	Quantity    int32       `json:"quantity"`
	Value       int32       `json:"value"`
}

func (o *OrderInfo) Transform() *OrderInfoTransform {
	products := []OrderProduct{}
	productsString := o.Products

	err := json.Unmarshal([]byte(productsString), &products)
	if err != nil {
		logger.Error(err, "error in unmarshal")
		return &OrderInfoTransform{
			ID:         o.ID,
			CustomerID: o.CustomerID,
		}
	}

	return &OrderInfoTransform{
		ID:         o.ID,
		CustomerID: o.CustomerID,
		Products:   products,
	}
}

func (o *OrderInfoTransform) ReTransform() *OrderInfo {
	products := o.Products
	out := &OrderInfo{
		ID:         o.ID,
		CustomerID: o.CustomerID,
	}

	productsStr, err := json.Marshal(products)
	if err != nil {
		return out
	}

	out.Products = string(productsStr)
	return out
}
