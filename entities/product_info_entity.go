package entities

import "time"

type ProductInfo struct {
	ID          uint   `gorm:"autoIncrement" json:"id"`
	Name        string `json:"name"`
	Price       int32  `json:"price"`
	Category    string `json:"category"`
	SubCategory string `json:"sub_category"`
	SKU         string `json:"sku"`
	ImageUri    string `json:"image_uri"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	CreatedBy   string `json:"created_by"`
	UpdatedBy   string `json:"updated_by"`
}
