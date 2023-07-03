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
	ID     uint        `gorm:"autoIncrement" json:"id"`
	Code   string      `json:"code"`
	Source ModelSource `json:"source"`
}

type ModelSource struct {
	SheetID   string `json:"sheet_id"`
	SheetName string `json:"sheet_name"`
}
