package entities

type BlockInfo struct {
	ID          uint   `gorm:"autoIncrement" json:"id"`
	Code        string `json:"code" gorm:"unique"`
	Description string `json:"description"`
	ModelIDs    string `json:"model_ids"`
}
