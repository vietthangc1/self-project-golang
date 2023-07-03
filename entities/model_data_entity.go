package entities

type ModelDataItem struct {
	ProductID int32
	Rank      float32
}

type ModelDataMaster struct {
	Key         string
	ProductRank []*ModelDataItem
}

type ModelInfo struct {
	ID            uint        `gorm:"autoIncrement" json:"id"`
	Code          string      `json:"code"`
	ModelSourceID uint        `json:"model_source_id"`
	Source        *ModelSource `json:"source" gorm:"foreignKey:ModelSourceID;references:ID"`
}

type ModelSource struct {
	ID        uint   `gorm:"autoIncrement" json:"id"`
	SheetID   string `json:"sheet_id"`
	SheetName string `json:"sheet_name"`
}
