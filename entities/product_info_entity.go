package entities

import "time"

type ProductInfo struct {
	ID          uint      `gorm:"autoIncrement" json:"id"`
	Name        string    `json:"name"`
	Price       int32     `json:"price"`
	Category    string    `json:"category"`
	SubCategory string    `json:"sub_category"`
	SKU         string    `json:"sku" gorm:"unique"`
	ImageURI    string    `json:"image_uri"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedBy   string    `json:"created_by"`
	UpdatedBy   string    `json:"updated_by"`
}
