package entities

type CustomerInfo struct {
	ID   uint   `gorm:"autoIncrement" json:"id"`
	Email string `json:"email"`
}
