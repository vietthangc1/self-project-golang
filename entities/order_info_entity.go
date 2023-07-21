package entities

import (
	"strconv"
	"strings"
)

type OrderInfo struct {
	ID         uint    `gorm:"autoIncrement" json:"id"`
	TotalValue float32 `json:"total_value"`
	ProductIDs string  `json:"product_ids"`
	CustomerID int32   `json:"customer_id"`
}

type OrderInfoTransform struct {
	ID         uint    `gorm:"autoIncrement" json:"id"`
	TotalValue float32 `json:"total_value"`
	ProductIDs []int32 `json:"product_ids"`
	CustomerID int32   `json:"customer_id"`
}

type OrderProduct struct {
	ID          uint        `gorm:"autoIncrement" json:"id"`
	ProductID   uint        `json:"product_id"`
	ProductInfo ProductInfo `json:"product_info"`
	Quantity    int32       `json:"quantity"`
	Value       int32       `json:"value"`
}

func (o *OrderInfo) Transform() *OrderInfoTransform {
	productIDs := []int32{}
	productIDString := o.ProductIDs
	if productIDString == "" || productIDString == "[]" {
		productIDs = []int32{}
	}
	modelIDArr := strings.Split(productIDString[1:len(productIDString)-1], ",")
	for _, modelID := range modelIDArr {
		modelIDStr := strings.TrimSpace(modelID)
		modelID, err := strconv.ParseInt(modelIDStr, 10, 32)
		if err != nil {
			continue
		}
		productIDs = append(productIDs, int32(modelID))
	}
	return &OrderInfoTransform{
		ID:         o.ID,
		TotalValue: o.TotalValue,
		CustomerID: o.CustomerID,
		ProductIDs: productIDs,
	}
}
