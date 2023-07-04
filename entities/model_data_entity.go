package entities

type ModelDataItem struct {
	ProductID int32   `json:"product_id"`
	Rank      float32 `json:"rank"`
}

type ModelDataMaster struct {
	Key         string           `json:"key"`
	ProductRank []*ModelDataItem `json:"product_rank"`
}

type ModelInfo struct {
	ID            uint         `gorm:"autoIncrement" json:"id,omitempty"`
	Code          string       `json:"code"`
	ModelSourceID uint         `json:"model_source_id,omitempty"`
	Source        *ModelSource `json:"source,omitempty" gorm:"foreignKey:ModelSourceID;references:ID"`
}

type ModelSource struct {
	ID        uint   `gorm:"autoIncrement" json:"id"`
	SheetID   string `json:"sheet_id"`
	SheetName string `json:"sheet_name"`
}
